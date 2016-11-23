package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"path"
	"log"
	"sync"
	
	"golang.org/x/net/context"
	"gopkg.in/macaroon-bakery.v2-unstable/bakery"
	"gopkg.in/macaroon-bakery.v2-unstable/httpbakery"
	"github.com/juju/idmclient"
)

func main() {
	content := NewContent()
	authorizer := bakery.ACLAuthorizer{
		AllowPublic: true,
		GetACL: func(_ context.Context, op bakery.Op) ([]string, error) {
			return content.GetACL(op.Entity, op.Action), nil
		},
	}
	identity, err := idmclient.New(idmclient.NewParams{
		BaseURL: "https://api.jujucharms.com/identity",
		Client: httpbakery.NewClient(),
	})
	if err != nil {
		log.Fatal(err)
	}
	handler := NewContentHandler(content)
	err = http.ListenAndServe(":61234", authHandler(handler, authorizer, idmclientShim{identity}))
	log.Fatal(err)
}

func authHandler(h http.Handler, authorizer bakery.Authorizer, identity bakery.IdentityClient) http.Handler {
	key, _ := bakery.GenerateKey()		// TODO check error!
	b := bakery.New(bakery.BakeryParams{
		Key: key,
		Authorizer: authorizer,
		Identity: identity,
	})
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ops := opsForRequest(req)
		macaroons := httpbakery.RequestMacaroons(req)
		_, err := b.Checker.Auth(macaroons...).Allow(req.Context(), ops...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, req)
	})
}

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
		h.content.Put(req.URL.Path, string(data), []string{bakery.Everyone})
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
	acls map[string][]string
}

func NewContent() *Content {
	return &Content{
		files: make(map[string]string),
	}
}

func (c *Content) GetACL(name, action string) []string {
	c.mu.Lock()
	defer c.mu.Unlock()
	if acl, ok := c.acls[name]; ok {
		return acl
	}
	if action != "PUT" {
		return nil
	}
	if acl, ok := c.acls[path.Dir(name)]; ok {
		return acl
	}
	return nil
}

func (c *Content) PutACL(name, action string, acl []string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.acls[name]; !ok {
		return ErrNotFound
	}
	c.acls[name] = acl
	return nil
}

func (c *Content) Put(name string, data string, acl []string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.files[name] = data
	c.acls[name] = acl
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
		action := "GET"
		if i == len(paths) - 1 {
			action = req.Method
		}
		ops[i] = bakery.Op{
			Entity: path,
			Action: action,
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

type idmClientShim struct {
	c idmclient.IdentityClient
}

func (c idmClientShim) DeclaredIdentity(attrs map[string]string) (auth.Identity, error) {
	return c.DeclaredIdentity(attrs)
}

func (c idmClientShim) IdentityFromContext(context.Context) (bakery.Identity, []checkers.Caveat, error) {
	return nil, c.c.IdentityCaveats(), nil
}
