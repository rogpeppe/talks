package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

var bottleTemplate = template.Must(template.New("").Parse(`
There are {{.Count}} bottles on the wall
`))

type params struct {
	Count int
}

func main() {
	r, w := io.Pipe()
	go func() {
		err := bottleTemplate.Execute(w, params{
			Count: 99,
		})
		if err != nil {
			log.Fatal(err)
		}
		w.Close()
	}()
	data, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The template data is:\n")
	os.Stdout.Write(data)
}
