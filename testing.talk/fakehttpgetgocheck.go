package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	gc "gopkg.in/check.v1"
	jujutesting "github.com/juju/testing"
)

func TestPackage(t *testing.T) {
	gc.TestingT(t)
}

var httpfetch =  struct {
	HTTPGet *func(string) (*http.Response, error)
	GetURLAsString func(string) (string, error)
}{
	GetURLAsString: GetURLAsString,
	HTTPGet: &httpGet,
}

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

var _ = gc.Suite(&suite{})

type suite struct{
	jujutesting.CleanupSuite
}

func (s *suite) TestGetURLAsStringHTTPGetError(c *gc.C) {
	s.PatchValue(httpfetch.HTTPGet, func(u string) (*http.Response, error) {
		return nil, errors.New("crash and burn")
	})
	got, err := httpfetch.GetURLAsString("http://0.1.2.3/")
	c.Assert(err, gc.ErrorMatches, "GET failed: crash and burn")
	c.Assert(got, gc.Equals, "")
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
	Name: "TestIt",
	F:    TestIt,
}}
