package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		start := time.Now()

		data := make([]byte, r.ContentLength)
		r.Body.Read(data)

		// Simulate processing time
		// TODO: come up with some realistic processing
		processingTime := time.Duration(len(data)/1024) * time.Millisecond
		time.Sleep(processingTime)

		fmt.Printf("Processed %v bytes in %v\n", len(data), time.Since(start))
	}
}
