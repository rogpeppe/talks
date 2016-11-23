package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {
	content := NewContent()
	handler := NewContentHandler(content)
	err := http.ListenAndServe(":61234", handler)
	log.Fatal(err)
}

var ErrNotFound = errors.New("not found")

func NewContentHandler(content *Content) *ContentHandler {
	return &ContentHandler{
		content: content,
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

func (c *Content) Get(name string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, ok := c.files[name]
	if !ok {
		return "", ErrNotFound
	}
	return data, nil
}
