//220901

package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	//var results map[string]string // expected to panic error
	//var results = map[string]string{} // would be work well
	var results = make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.instagram.com/",
		"https://www.facebook.com/",
	}
	//results["hello"] = "hello"
	for _, url := range urls {
		//fmt.Println(url)
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	//fmt.Println(results)
	for url, result := range results {
		fmt.Println(url, result)
	}
}

var errRequestFailed = errors.New("Request Failed")

func hitURL(url string) error {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}
