package main

import (
	"context"
	"fmt"
	"time"
)

// request does not support context
func request(ctx context.Context) (any, error) {
	select {
	case <-time.After(5 * time.Second):
		return "Hello, World!", nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // or context.WithDeadline, or context.WithTimeout
	defer cancel()                                          // best practice, but don't rely on it for timely cancellation

	result := make(chan any)

	go func() {
		r, err := request(ctx)
		if err != nil {
			fmt.Println("request error: ", err)
			return
		}

		result <- r
	}()

	select {
	case r := <-result:
		fmt.Println("result: ", r)
	case <-time.After(time.Second):
		fmt.Println("request timeout")
	}

	cancel() // important! cancel now; don't defer

	time.Sleep(100 * time.Millisecond) // allow time for sub goroutine to print cancel msg
}
