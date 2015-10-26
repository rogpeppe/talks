package httpfetch_test

import (
	"net/http"
	"net/http/httptest"
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
