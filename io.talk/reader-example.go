package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("io.slide")
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 30)
	n, err := f.Read(buf)
	fmt.Printf("read %d bytes (%q)\n", n, buf[0:n])
	fmt.Printf("error %v\n", err)
}
