package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/rogpeppe/httpfetch"
)

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
