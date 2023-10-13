package main

import (
	"github.com/julienschmidt/httprouter"
	"golang.org/x/time/rate"
	"net/http"
)

func main() {
	router := httprouter.New()
	limiter := rate.NewLimiter(2, 1) // Limit to 2 requests per second.

	router.GET("/api/resource", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if limiter.Allow() {
			// Request Processing
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Data"))
		} else {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests"))
		}
	})

	http.ListenAndServe(":8080", router)
}
