package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {
	r := StringReader("Some data that we're going to read.\n")

	outFile, err := os.Create("/tmp/text.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	w := gzip.NewWriter(outFile)
	defer w.Close()

	_, err = io.Copy(w, &r)
	if err != nil {
		log.Fatal(err)
	}
}

// StringReader implements io.Reader for a string.
type StringReader string

// Read implements io.Reader by reading from the string and updating it
// so that it holds the unread part of the string.
func (r *StringReader) Read(buf []byte) (int, error) {
	if len(*r) == 0 {
		return 0, io.EOF
	}
	n := copy(buf, *r)
	*r = (*r)[n:]
	return n, nil
}
