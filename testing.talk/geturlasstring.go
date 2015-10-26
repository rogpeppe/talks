package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/rogpeppe/httpfetch"
)

func TestGetURLAsStringSuccess(t *testing.T) {
	text := "hello, world\n"
	handler := func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(text))
	}
	srv := httptest.NewServer(http.HandlerFunc(handler))
	defer srv.Close()

	got, err := httpfetch.GetURLAsString(srv.URL)
	if err != nil {
		t.Fatalf("GetURLAsString error: %v", err)
	}
	if got != text {
		t.Fatalf("GetURLAsString returned invalid text; got %q want %q", got, text)
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
	Name: "TestGetURLAsStringSuccess",
	F:    TestGetURLAsStringSuccess,
}}
