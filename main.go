package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.instagram.com/",
		"https://www.naver.com/",
	}

	for _, url := range urls {
		// Goroutines 적용 전
		// result := "OK"
		// err := hitURL(url)
		// if err != nil {
		// 	result = "FAILED"
		// }
		// results[url] = result

		// Goroutines 적용 후
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}

	// for url, result := range results {
	// 	fmt.Println(url, result)
	// }

	// c := make(chan bool)
	// people := [2]string{"bewise", "sh"}
	// for _, person := range people {
	// 	go isSmart(person, c)
	// }
	// result := <-c
	// fmt.Println(result)
}

// func isSmart(person string, c chan bool) {
// 	time.Sleep(time.Second * 5)
// 	c <- true
// }

// Goroutines 적용 전
// func hitURL(url string) error {
// 	fmt.Println("Checking:", url)
// 	response, err := http.Get(url)
// 	if err != nil || response.StatusCode >= 400 {
// 		return errRequestFailed
// 	}
// 	return nil
// }

// Goroutins CPU 수에 따라 적용할 수 있는 방법이 다름
// Goroutines 적용 후
func hitURL(url string, c chan<- requestResult) {
	fmt.Println("Checking:", url)
	response, err := http.Get(url)
	status := "OK"
	if err != nil || response.StatusCode >= 400 {
		c <- requestResult{url: url, status: "FAILED"}
	} else {
		c <- requestResult{url: url, status: status}
	}
}
