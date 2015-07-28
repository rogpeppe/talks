package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

	q := bson.M{"_id": bson.M{"$gt": "Alice"}} // HL
	iter := collection.Find(q).Iter()
	var p person
	for iter.Next(&p) {
		fmt.Printf("%s is %s at %v\n", p.Name, p.Status, p.StatusTime)
	}
	if err := iter.Err(); err != nil {
		log.Fatalf("iteration error: %v", err)
	}
}
