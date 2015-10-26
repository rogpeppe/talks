package httpfetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetURLAsString makes a GET request to the
// given URL and returns the result as a string.
func GetURLAsString(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("GET failed: %v", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("cannot read body: %v", err)
	}
	return string(data), nil
}
