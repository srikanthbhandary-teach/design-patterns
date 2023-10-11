package main

import (
	"design-patterns/structural"
	"log"
	"net/http"
	"time"
)

func main() {
	interval := 5 * time.Second
	maxRequestCount := 5
	RateLimiter := structural.NewRateLimiter(interval, maxRequestCount)
	server := structural.Server{}
	proxy := structural.NewRateLimiterProxy(RateLimiter, &server)
	log.Println(http.ListenAndServe(":8080", proxy))
}
