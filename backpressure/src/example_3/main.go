package main

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/time/rate"
)

func main() {
	db, _ := sqlx.Open("postgres", "user=username dbname=mydb sslmode=disable")

	limiter := rate.NewLimiter(10, 1) // Limit to 10 requests per second.

	// ...

	var result string
	if limiter.Allow() {
		db.Get(&result, "SELECT data FROM sometable WHERE id=$1", 1)
	} else {
		// Processing too many requests
	}
}
