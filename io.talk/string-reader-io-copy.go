package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	r := &StringReader{
		S: "Some data that we're going to read.\n",
	}
	n, err := io.Copy(os.Stdout, r)
	if err != nil {
		fmt.Printf("error when copying: %v", err)
	} else {
		fmt.Printf("copied %d bytes\n", n)
	}
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

