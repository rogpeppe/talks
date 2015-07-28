package main

import (
	"github.com/kr/pretty"
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

	var p person
	if err := collection.Find(nil).One(&p); err != nil { // HL
		log.Fatalf("cannot get one person: %v", err)
	}
	pretty.Println(p)
}
