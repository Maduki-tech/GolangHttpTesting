package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func sendRequest(url string, data []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Response Status:", resp.Status)
	defer resp.Body.Close()

}

func main() {
	url := flag.String("url", "http://localhost:8080", "URL to send requests to")
	dataSize := flag.Int("data-size", 1024, "Size of data to send in bytes")
	numberOfRequests := flag.Int("requests", 10, "Number of requests to send")

	flag.Parse()

	data := make([]byte, *dataSize)
	for i := range data {
		data[i] = 'a'
	}

	wg := sync.WaitGroup{}

	start := time.Now()

	for i := 0; i < *numberOfRequests; i++ {
		wg.Add(1)
		go sendRequest(*url, data, &wg)
	}

	wg.Wait()

	fmt.Printf("It Took %v seconds\n", time.Since(start))
}
