package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type PressureGauge struct {
	ch chan struct{}
}

func New(limit int) *PressureGauge {
	ch := make(chan struct{}, limit)

	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}

	return &PressureGauge{ch: ch}
}

func (p *PressureGauge) Do(fn func()) error {
	select {
	case <-p.ch:
		fn()
		p.ch <- struct{}{}
		return nil
	default:
		return errors.New("out of capacity")
	}
}

func main() {
	pg := New(3)
	go runServer(pg)

	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()

			resp, err := http.Get("http://localhost:8080")
			if err != nil {
				fmt.Println("error: ", err.Error())
				return
			}

			defer resp.Body.Close()

			msg, _ := io.ReadAll(resp.Body)

			fmt.Println("response: ", string(msg))
		}()
	}

	wg.Wait()
}

func runServer(pg *PressureGauge) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Do(func() {
			w.Write([]byte(expensiveOperation()))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("too many requests"))
		}
	})
	http.ListenAndServe(":8080", nil)
}

func expensiveOperation() string {
	time.Sleep(2 * time.Second)
	return "done"
}
