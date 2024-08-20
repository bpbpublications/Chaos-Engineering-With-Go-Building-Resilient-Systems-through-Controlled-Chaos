package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gomodule/redigo/redis"
)

var redisPool *redis.Pool

func init() {
	redisPool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   100,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
}

func rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn := redisPool.Get()
		defer conn.Close()

		userId := r.RemoteAddr // Assuming IP as the user identifier
		window := 60           // 1 minute window
		limit := 10            // Maximum 10 requests per minute

		currentTimestamp := time.Now().Unix()
		startTimestamp := currentTimestamp - int64(window)

		// Remove expired entries from the sorted set
		_, err := conn.Do("ZREMRANGEBYSCORE", "rate_limiting:"+userId, "-inf", startTimestamp)
		if err != nil {
			log.Println("Error removing expired entries:", err)
		}

		// Count the number of entries within the time window
		count, err := redis.Int(conn.Do("ZCOUNT", "rate_limiting:"+userId, startTimestamp, currentTimestamp))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if count >= limit {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		// Increment request count
		_, err = conn.Do("ZADD", "rate_limiting:"+userId, currentTimestamp, currentTimestamp)
		if err != nil {
			log.Println("Error incrementing rate limit count:", err)
		}

		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	http.ListenAndServe(":8080", rateLimitMiddleware(mux))
}
