//220901

package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.instagram.com/",
		"https://www.facebook.com/",
	}

	for _, url := range urls {
		//fmt.Println(url)
		hitURL(url)
	}
}

var errRequestFailed = errors.New("Request Failed")

func hitURL(url string) error {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
