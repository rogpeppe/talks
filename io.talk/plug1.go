package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("io.slide")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The presentation slide data is:\n")
	os.Stdout.Write(data)
}
