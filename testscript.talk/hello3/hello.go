package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var file = flag.String("f", "", "file to read message from")

func main() {
	os.Exit(main1())
}

func main1() int {
	flag.Parse()
	msg := "hello, world"
	if *file != "" {
		data, err := ioutil.ReadFile(*file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "cannot read file: ", err)
			return 1
		}
		msg = string(data)
	}
	fmt.Println(msg)
	return 0
}
