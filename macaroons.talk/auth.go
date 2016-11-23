package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"path"
	"sync"

	"gopkg.in/macaroon-bakery.v2-unstable/bakery"
)

func main() {
	http.ListenAndServe(":61234", authHandler(NewContentHandler()))
}

func authHandler(h http.Handler) http.Handler {
	key, _ := bakery.GenerateKey()		// TODO check error!
	b := bakery.New(bakery.BakeryParams{
		Key: key,
	})
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ops := opsForRequest(req)
		_, err := b.Checker.Auth().Allow(req.Context(), ops...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, req)
	})
}

func NewContentHandler() *ContentHandler {
	return &ContentHandler{
		content: NewContent(),
	}
}

type ContentHandler struct {
	content *Content
}
	

func (h *ContentHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "PUT":
		data, _ := ioutil.ReadAll(req.Body) // TODO check error!
		h.content.Put(req.URL.Path, string(data))
	case "GET":
		data, err := h.content.Get(req.URL.Path)
		if err != nil {
			if err == ErrNotFound {
				http.Error(w, "file not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		w.Write([]byte(data))
	default:
		http.Error(w, "only PUT and GET are allowed", http.StatusMethodNotAllowed)
	}
}

type Content struct {
	mu    sync.Mutex
	files map[string]string
}

func NewContent() *Content {
	return &Content{
		files: make(map[string]string),
	}
}

func (c *Content) Put(name string, data string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.files[name] = data
}

var ErrNotFound = errors.New("not found")

func (c *Content) Get(name string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, ok := c.files[name]
	if !ok {
		return "", ErrNotFound
	}
	return data, nil
}

// opsForRequest returns all the operations implied
// by the given request.
func opsForRequest(req *http.Request) []bakery.Op {
	paths := parents(req.URL.Path)
	ops := make([]bakery.Op, len(paths))
	for i, path := range paths {
		ops[i] = bakery.Op{
			Entity: path,
			Action: req.Method,
		}
	}
	return ops
}

// parents returns the given path and all its parents.
// For instance, given "/usr/rog/foo",
// it will return []string{"/usr/rog/foo", "/usr/rog", "/usr", "/"}
func parents(p string) []string {
	var all []string
	p = path.Clean(p)
	for {
		all = append(all, p)
		parent := path.Dir(p)
		if parent == p {
			break
		}
		p = parent
	}
	return all
}
