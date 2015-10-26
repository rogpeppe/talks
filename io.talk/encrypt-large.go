package main

import (
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
	"text/template"
)

var bottleTemplate = template.Must(template.New("").Parse(`
{{range $i := .CountChan}}
There are {{$i}} bottles on the wall
{{end}}
`))

type params struct {
	CountChan chan int
}

func main() {
	// Create the output file.
	outFile, err := os.Create("/tmp/bottles.crypt")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	hash, err := generate(outFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("checksum of encrypted data: %x\n", hash.Sum(nil))
}

func generate(outFile io.Writer) (hash.Hash, error) {
	// Base64 encode to the output file; close when done to
	// flush any partial data remaining.
	b64w := base64.NewEncoder(base64.StdEncoding, outFile)
	defer b64w.Close()

	// Create the hasher for the checksum.
	hasherw := sha1.New()

	// We want to base64-encode and hash... and we can!
	b64AndHashw := io.MultiWriter(b64w, hasherw)

	// Encrypt to the hasher and to the file.
	encryptw := newEncrypter(b64AndHashw, "my secret key")

	// Compress to the encrypter; close when done to flush
	// any remaining data.
	gzipw := gzip.NewWriter(encryptw)
	defer gzipw.Close()

	// TEMPLATE OMIT
	bottleChan := make(chan int)
	go func() {
		for i := 100 * 1000; i >= 0; i-- {
			bottleChan <- i
		}
		close(bottleChan)
	}()
	err := bottleTemplate.Execute(gzipw, params{
		CountChan: bottleChan,
	})
	if err != nil {
		return nil, err
	}
	return hasherw, nil
}

func newEncrypter(w io.Writer, key string) io.Writer {
	// Create 32-byte (256 bit) key selects AES-256
	hashedKey := sha256.Sum256([]byte(key))

	// Use AES in output feedback mode.
	block, err := aes.NewCipher(hashedKey[:])
	if err != nil {
		panic(err) // Can only happen if key is wrong size.
	}
	// If the key is unique for each ciphertext, then it's ok to use a zero IV.
	iv := make([]byte, aes.BlockSize)
	return &cipher.StreamWriter{
		S: cipher.NewOFB(block, iv),
		W: w,
	}
}
