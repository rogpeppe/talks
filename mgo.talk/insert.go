package main

import (
	"gopkg.in/mgo.v2"

	"log"
	"time"
)

type person struct {
	Name       string `bson:"_id"`
	Status     string
	StatusTime time.Time
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatalf("cannot dial mongoDB: %v", err)
	}
	db := session.DB("mongodbtalk")
	collection := db.C("people")
	p := person{
		Name:       "Bob",
		Status:     "bored",
		StatusTime: time.Now(),
	}
	if err := collection.Insert(p); err != nil { // HL
		log.Fatalf("cannot insert document: %v", err)
	}
	log.Printf("added %s", p.Name)
}
