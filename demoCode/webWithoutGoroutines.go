package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(url string) Page {
	fmt.Println("Getting", url)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return Page{URL: url, Size: len(body)}
}

func main() {
	urls := []string{"https://example.com/",
		"https://golang.org/",
		"https://golang.org/doc",
		"https://facebook.com"}

	for i := 0; i < 100; i++ {
		urls = append(urls, "https://facebook.com")
	}
	pages := []Page{}
	for _, url := range urls {
		pages = append(pages, responseSize(url))
	}
	for _, page := range pages {
		fmt.Println(page)
	}
}
