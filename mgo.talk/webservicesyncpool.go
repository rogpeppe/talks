package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatalf("cannot dial mongoDB: %v", err)
	}
	db := session.DB("mongodbtalk")
	collection := db.C("people")
	log.Println("starting service on port 55667")
	err = http.ListenAndServe(":55667", NewStatusHandler(collection))
	if err != nil {
		log.Fatal(err)
	}
}

type person struct {
	Name       string `bson:"_id"`
	Status     string
	StatusTime time.Time
}

type statusHandler struct {
	mux        *http.ServeMux
	collection *mgo.Collection
}

func newStatusHandler() *statusHandler {
	h := &statusHandler{
		mux: http.NewServeMux(),
	}
	h.mux.HandleFunc("/latest", h.serveLatest)
	h.mux.HandleFunc("/status/", h.serveStatus)
	return h
}

// POOL OMIT
var statusHandlerPool = sync.Pool{
	New: func() interface{} {
		return newStatusHandler()
	},
}

func getStatusHandler() *statusHandler {
	return statusHandlerPool.Get().(*statusHandler)
}

func (h *statusHandler) Close() {
	h.collection.Database.Session.Close()
	h.collection = nil
	statusHandlerPool.Put(h)
}

// ENDPOOL OMIT

func NewStatusHandler(collection *mgo.Collection) http.Handler {
	handler := func(w http.ResponseWriter, req *http.Request) {
		h := getStatusHandler()
		defer h.Close()

		newSession := collection.Database.Session.Copy()
		h.collection = collection.With(newSession)

		h.mux.ServeHTTP(w, req)
	}
	return http.HandlerFunc(handler)
}

func (h *statusHandler) serveLatest(w http.ResponseWriter, req *http.Request) {
	var p person
	if err := h.collection.Find(nil).Sort("-statustime").One(&p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s is %s at %v\n", p.Name, p.Status, p.StatusTime)
}

func (h *statusHandler) serveStatus(w http.ResponseWriter, req *http.Request) {
	name := strings.TrimPrefix(req.URL.Path, "/status/")
	switch req.Method {
	case "PUT":
		h.servePutStatus(w, req, name)
	case "GET":
		h.serveGetStatus(w, req, name)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *statusHandler) servePutStatus(w http.ResponseWriter, req *http.Request, name string) {
	status, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := h.collection.Upsert(bson.D{{"_id", name}}, bson.D{{
		"$set", bson.D{{
			"status", strings.TrimSpace(string(status)),
		}, {
			"statustime", time.Now(),
		}},
	}}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Printf("inserted %q", name)
}

func (h *statusHandler) serveGetStatus(w http.ResponseWriter, req *http.Request, name string) {
	var p person
	err := h.collection.Find(bson.D{{"_id", name}}).One(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s is %s at %v\n", p.Name, p.Status, p.StatusTime)
}
