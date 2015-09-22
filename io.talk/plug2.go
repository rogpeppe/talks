package main

import (
	"crypto/sha1"
	"fmt"
	"log"
	"text/template"
)

var bottleTemplate = template.Must(template.New("").Parse(`
There are {{.Count}} bottles on the wall
`))

type params struct {
	Count int
}

func main() {
	bottleHash := sha1.New()

	err := bottleTemplate.Execute(bottleHash, params{
		Count: 99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("SHA1 hash: %x\n", bottleHash.Sum(nil))
}
