package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var httpGet = http.Get

// GetURLAsString makes a GET request to the
// given URL and returns the result as a string.
func GetURLAsString(url string) (string, error) {
	resp, err := httpGet(url)
	if err != nil {
		return "", fmt.Errorf("GET failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GET returned unexpected status %q", resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("cannot read body: %v", err)
	}
	return string(data), nil
}

func TestGetURLAsStringNotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(http.NotFound))
	defer srv.Close()

	got, err := httpfetch.GetURLAsString(srv.URL)
	if err == nil {
		t.Fatalf("unexpected success with result %q", got)
	}
}

func main() {
	//	fmt.Println(testing.Benchmark(func(b *testing.B) {
	//		for i := 0; i < b.N; i++ {
	//			TestGetString(nil)
	//		}
	//	}))
	match := func(pat, str string) (bool, error) {
		return true, nil
	}
	m := testing.MainStart(match, tests, nil, nil)
	os.Exit(m.Run())
}

var tests = []testing.InternalTest{{
	Name: "TestGetURLAsStringNotFound",
	F:    TestGetURLAsStringNotFound,
}}
