package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/macaroon.v2-unstable"
)

func main() {
	m, _ := macaroon.New(
		[]byte("root key"),
		[]byte("identifier"),
		"location",
		macaroon.LatestVersion,
	)
	m.AddFirstPartyCaveat("some caveat")
	data, _ := json.MarshalIndent(m, "", "\t")
	fmt.Printf("%s\n", data)
}
