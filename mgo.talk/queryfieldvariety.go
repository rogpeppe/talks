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

	type gtCondition struct {
		Value string `bson:"$gt"`
	}
	type gtQuery struct {
		Id gtCondition `bson:"_id"`
	}
	for i, q := range []interface{}{
		bson.M{"_id": bson.M{"$gt": "Alice"}},     // HL
		bson.D{{"_id", bson.D{{"$gt", "Alice"}}}}, // HL
		gtQuery{gtCondition{"Alice"}},             // HL
	} {
		var p person
		iter := collection.Find(q).Iter()
		for iter.Next(&p) {
			fmt.Printf("query %d: %s is %s at %v\n", i, p.Name, p.Status, p.StatusTime)
		}
		if err := iter.Err(); err != nil {
			log.Fatalf("iteration error: %v", err)
		}
	}
	// ENDLOOP OMIT
}
