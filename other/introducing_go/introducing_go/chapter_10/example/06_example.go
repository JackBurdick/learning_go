package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

type HomePageSize struct {
	URL string
	Size int
}

func main() {
	urls := []string{
		"http://www.apple.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.microsoft.com",
	}

	results := make(chan HomePageSize)

	for _, url := range urls {
		// create channel for current url
		go func(url string) {
			// HTTP Get request
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			// read response
			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			// store response information
			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
			}
		}(url)
	}

	// loop through all results
	// print biggest web page size
	var biggest HomePageSize

		for range urls {
		result := <- results
		if result.Size > biggest.Size {
			biggest = result
		}
	}
	fmt.Println("The biggest home page:", biggest.URL)
}