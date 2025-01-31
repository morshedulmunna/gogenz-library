package middleware

import (
	"log"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

// Rate limiter configuration
var rateLimiters = struct {
	sync.Mutex
	clients map[string]*rate.Limiter
}{clients: make(map[string]*rate.Limiter)}

// getLimiter returns a rate limiter for a given IP
func getLimiter(ip string) *rate.Limiter {

	rateLimiters.Lock()
	defer rateLimiters.Unlock()

	if limiter, exists := rateLimiters.clients[ip]; exists {
		return limiter
	}

	log.Println("Ip Address: ", ip)

	// Allow 10 requests per second with a burst of 5
	limiter := rate.NewLimiter(10, 5)
	rateLimiters.clients[ip] = limiter
	return limiter
}

// RateLimitMiddleware applies rate limiting to incoming requests
func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		limiter := getLimiter(ip)

		if !limiter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
