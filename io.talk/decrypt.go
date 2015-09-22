package main

import (
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	outFile, err := os.Create("/tmp/bottles.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	inFile, err := os.Open("/tmp/bottles.crypt")
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()
	if err := readEncryptedFile(outFile, inFile, "0a39169800daa19d19d9add1a772fff2b653310e"); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("successful decryption\n")
}

func readEncryptedFile(outFile io.Writer, inFile io.Reader, expectSum string) error {
	// Base64 decode the input file.
	b64r := base64.NewDecoder(base64.StdEncoding, inFile)

	// Write to the hasher while reading the decoded data.
	hasherw := sha1.New()
	teeRead := io.TeeReader(b64r, hasherw)

	// Decrypt the decoded data.
	decryptr := newDecrypter(teeRead, "my secret key")
	gzipr, err := gzip.NewReader(decryptr)
	if err != nil {
		return err
	}

	// Copy the decrypted data to the output file.
	_, err = io.Copy(outFile, gzipr)
	if err != nil {
		return err
	}

	// CHECKSUM OMIT
	// Check that everything arrived OK.
	if sum := fmt.Sprintf("%x", hasherw.Sum(nil)); sum != expectSum {
		return fmt.Errorf("checksum mismatch (got %s, want %s)", sum, expectSum)
	}
	return nil
}

func newDecrypter(r io.Reader, key string) io.Reader {
	// Create 32-byte (256 bit) key selects AES-256
	hashedKey := sha256.Sum256([]byte(key))

	// Use AES in output feedback mode.
	block, err := aes.NewCipher(hashedKey[:])
	if err != nil {
		panic(err) // Can only happen if key is wrong size.
	}
	// If the key is unique for each ciphertext, then it's ok to use a zero IV.
	iv := make([]byte, aes.BlockSize)

	return &cipher.StreamReader{
		S: cipher.NewOFB(block, iv),
		R: r,
	}
}
