package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

var (
	// Define the rate limit: 1 requests per second with a burst of 1 requests.
	limiter = rate.NewLimiter(rate.Limit(1), 1)
)

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Wait for the limiter to allow the request.
	start := time.Now()
	ctx := context.Background()
	err := limiter.Wait(ctx)
	if err != nil {
		http.Error(w, "Too Many Requests to handle", http.StatusTooManyRequests)
		return
	}

	// Simulate processing the request.
	fmt.Fprintf(w, "Hello, Rate Limiting!\n")
	fmt.Printf("Request processed. Elapsed time: %s\n", time.Since(start))
}
