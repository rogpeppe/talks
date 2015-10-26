package main

import (
	"fmt"
	"io"
)

func main() {
	r := &StringReader{
		S: "Some data that we're going to read.\n",
	}

	buf := make([]byte, 10)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			fmt.Printf("read %d bytes (%q)\n", n, buf[0:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read error: %v\n", err)
			break
		}
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

