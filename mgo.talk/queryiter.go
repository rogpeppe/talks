package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
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

	iter := collection.Find(nil).Iter() // HL
	var p person
	for iter.Next(&p) {
		fmt.Printf("%s is %s at %v\n", p.Name, p.Status, p.StatusTime)
	}
	if err := iter.Err(); err != nil {
		log.Fatalf("iteration error: %v", err)
	}
}
