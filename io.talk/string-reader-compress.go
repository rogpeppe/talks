package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	r := &StringReader{
		S: "Some data that we're going to read.\n",
	}

	outFile, err := os.Create("/tmp/text.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	w := gzip.NewWriter(outFile)

	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("compressed to %v\n", outFile.Name())
}

// StringReader implements io.Reader for a string.
type StringReader struct {
	// S holds the string that's being read.
	S string
}

// Read implements io.Reader by reading from r.S
// and updating it to hold the unread data.
func (r *StringReader) Read(buf []byte) (int, error) {
	if len(r.S) == 0 {
		return 0, io.EOF
	}
	n := copy(buf, r.S)
	r.S = r.S[n:]
	return n, nil
}

