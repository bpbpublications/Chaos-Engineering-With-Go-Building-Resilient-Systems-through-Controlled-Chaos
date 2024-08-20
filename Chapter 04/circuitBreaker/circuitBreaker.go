package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

func main() {
	// Configure Hystrix
	hystrix.ConfigureCommand("api_request", hystrix.CommandConfig{
		Timeout:                1000, // Timeout in milliseconds
		MaxConcurrentRequests:  100,  // Maximum concurrent requests
		RequestVolumeThreshold: 10,   // Minimum number of requests before circuit breaker action
		SleepWindow:            5000, // Time in milliseconds to wait before testing the circuit breaker
		ErrorPercentThreshold:  50,   // Percentage of requests that should be allowed to fail before opening the circuit
	})

	// Define the handler for the API request
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		// Execute the API request within the circuit breaker
		err := hystrix.Do("api_request", func() error {
			// Simulate the API request
			// For demonstration, we'll simulate an error on every third request
			if time.Now().Second()%3 == 0 {
				// Simulate an error response
				return fmt.Errorf("API request failed")
			}
			// Simulate a successful response
			fmt.Println("API request executed successfully")
			return nil
		}, func(err error) error {
			// Fallback function in case of circuit breaker trip
			fmt.Println("Circuit breaker trip:", err.Error())
			return err
		})

		// Handle errors returned by the circuit breaker
		if err != nil {
			http.Error(w, "Service unavailable, kindly try again later !!!", http.StatusServiceUnavailable)
			return
		}

		// Send success response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API request successful"))
	})

	// Start the HTTP server
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
