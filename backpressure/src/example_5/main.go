package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create a goroutine pool with a limit of 3 active goroutines.
	pool := NewWorkerPool(3)

	// Create tasks for execution.
	tasks := []Task{
		{"Task 1", sync.Mutex{}, 2},
		{"Task 2", sync.Mutex{}, 3},
		{"Task 3", sync.Mutex{}, 1},
		{"Task 4", sync.Mutex{}, 4},
		{"Task 5", sync.Mutex{}, 2},
	}

	// Start executing tasks in the pool.
	for _, task := range tasks {
		pool.Submit(task)
	}

	// Wait for all tasks to complete.
	pool.Wait()

	fmt.Println("All tasks are completed.")
}

// Task represents a task to be executed.
type Task struct {
	Name     string
	mu       sync.Mutex // Mutex for synchronization
	Duration time.Duration
}

// Execute performs the task.
func (t *Task) Execute() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.Duration = 3
	fmt.Printf("%s is running...\n", t.Name)
	time.Sleep(t.Duration * time.Second)
	fmt.Printf("%s is done.\n", t.Name)
}

// WorkerPool represents a goroutine pool for task execution.
type WorkerPool struct {
	workers    int
	semaphore  chan struct{}
	taskQueue  chan Task
	wg         sync.WaitGroup
	stopSignal chan struct{}
}

// NewWorkerPool creates a new goroutine pool with the specified number of active workers.
func NewWorkerPool(workers int) *WorkerPool {
	pool := &WorkerPool{
		workers:    workers,
		semaphore:  make(chan struct{}, workers),
		taskQueue:  make(chan Task, 100), // Task queue size
		stopSignal: make(chan struct{}),
	}

	// Start active goroutines for task execution.
	for i := 0; i < workers; i++ {
		go pool.worker()
	}

	// Wait for all goroutines to finish on exit.
	go func() {
		pool.wg.Wait()
		close(pool.taskQueue)
		close(pool.stopSignal)
	}()

	return pool
}

// Submit sends a task to the pool for execution.
func (p *WorkerPool) Submit(task Task) {
	p.taskQueue <- task
}

// Wait waits for all tasks to complete in the pool.
func (p *WorkerPool) Wait() {
	<-p.stopSignal
}

// worker represents an active goroutine that executes tasks.
func (p *WorkerPool) worker() {
	p.wg.Add(1)
	defer p.wg.Done()

	for {
		select {
		case task, ok := <-p.taskQueue:
			if !ok {
				// Task channel is closed, terminate the goroutine.
				return
			}

			// Acquire a semaphore to limit the number of active goroutines.
			p.semaphore <- struct{}{}

			// Execute the task.
			task.Execute()

			// Release the semaphore after task execution.
			<-p.semaphore

		case <-p.stopSignal:
			// Received a termination signal, terminate the goroutine.
			return
		}
	}
}
