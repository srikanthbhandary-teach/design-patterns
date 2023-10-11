package structural

import (
	"net/http"
	"time"
)

type RateLimiter struct {
	Interval        time.Duration
	MaxRequestCount int
	Requests        map[string]int
	LastReset       map[string]time.Time
}

func NewRateLimiter(interval time.Duration, maxRequestCount int) *RateLimiter {
	return &RateLimiter{
		Interval:        interval,
		MaxRequestCount: maxRequestCount,
		Requests:        make(map[string]int),
		LastReset:       make(map[string]time.Time),
	}
}

func (r *RateLimiter) Allow(ip string) bool {
	if r.Requests[ip] >= r.MaxRequestCount {
		if time.Now().Sub(r.LastReset[ip]) >= r.Interval {
			r.Reset(ip)
		} else {
			return false
		}
	}
	r.Requests[ip]++
	return true
}

func (r *RateLimiter) Reset(ip string) {
	r.Requests[ip] = 0
	r.LastReset[ip] = time.Now()
}

type Server struct {
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type RateLimiterProxy struct {
	limiter *RateLimiter
	server  *Server
}

func NewRateLimiterProxy(limiter *RateLimiter, server *Server) *RateLimiterProxy {
	return &RateLimiterProxy{
		limiter: limiter,
		server:  server,
	}
}

func (r *RateLimiterProxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ip := req.RemoteAddr
	if !r.limiter.Allow(ip) {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte("Too many requests"))
		return
	}
	r.server.ServeHTTP(w, req)
}
