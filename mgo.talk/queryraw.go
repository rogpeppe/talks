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

	var p bson.Raw // HL
	if err := collection.Find(nil).One(&p); err != nil {
		log.Fatalf("cannot get person: %v", err)
	}
	fmt.Printf("%q\n", p.Data)
}
