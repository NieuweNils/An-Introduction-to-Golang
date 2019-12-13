package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(channel chan Page, url string) {
	fmt.Println("Getting", url)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	channel <- Page{URL: url, Size: len(body)}
}

func main() {
	pages := make(chan Page)
	urls := []string{"https://example.com/",
		"https://golang.org/",
		"https://golang.org/doc",
		"https://www.facebook.com",
		"https://www.linkit.nl/"}

	// mini-DDOS facebook
	for i := 0; i < 200; i++ {
		urls = append(urls, "https://www.facebook.com/")
	}

	for _, url := range urls {
		go responseSize(pages, url)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Println("Response size of", page.URL, ":", page.Size)
	}

}

// mini-DDOS linkit
// for i := 0; i < 200; i++ {
// 	urls = append(urls, "https://www.linkit.nl/")
// }
