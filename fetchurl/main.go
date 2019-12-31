package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	Error    error
	Response *http.Response
}

func main() {

	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)

			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}
	done := make(chan interface{})
	defer func() {
		fmt.Println("closing done channel")
		close(done)
	}()

	go func() {
		urls := []string{"https://www.google.com", "https://badhost"}
		for result := range checkStatus(done, urls...) {
			if result.Error != nil {
				fmt.Printf("error: %v", result.Error)
				continue
			}
			fmt.Printf("Response: %v\n", result.Response.Status)
		}
	}()
	fmt.Println("Waiting for 5 seconds then terminates ongoing work")
	time.Sleep(5 * time.Second)
}
