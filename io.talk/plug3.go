package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("io.slide")
	if err != nil {
		log.Fatal(err)
	}
	slideHash := sha1.New()
	_, err = io.Copy(slideHash, f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("SHA1 hash of slides: %x\n", slideHash.Sum(nil))
}
