package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("URL is Down", err)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s -> Status Code: %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1] // https://www.google.com
			file += ".txt"

			fmt.Printf("Writing response body to %s\n", file)

			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	wg.Done()
}

func main() {
	urls := []string{"https://www.golang.org", "https://www.google.com", "https://www.medium.com"}

	// 1.
	// Create a new instance of sync.WaitGroup (we’ll call it symply wg)
	// This WaitGroup is used to wait for all the goroutines that have been launched to finish.
	var wg sync.WaitGroup

	// Call wg.Add(n) method before attempting to launch the go routine.
	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
		fmt.Println(strings.Repeat("#", 20))
	}
	fmt.Println("Number of Goroutines", runtime.NumGoroutine())
	wg.Wait()
}
