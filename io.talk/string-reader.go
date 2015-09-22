package main
import (
	"fmt"
	"io"
)

func main() {
	r := StringReader("Some data that we're going to read.\n")

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
