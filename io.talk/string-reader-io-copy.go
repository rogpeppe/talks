package main
import (
	"fmt"
	"io"
	"os"
)

func main() {
	r := StringReader("Some data that we're going to read.\n")
	n, err := io.Copy(os.Stdout, &r)
	if err != nil {
		fmt.Printf("error when copying: %v", err)
	} else {
		fmt.Printf("copied %d bytes\n", n)
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
