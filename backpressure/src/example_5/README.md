## Example 5: Efficient Task Execution with Worker Pool

A worker pool is a widely used concurrency pattern that efficiently manages a limited number of worker goroutines capable of processing tasks concurrently. In this code example:

1. **Task Structure** 📋:
    - The `Task` struct represents individual tasks that can be executed. Each task is defined by its name and a specified duration.

2. **WorkerPool Structure** 🏭:
    - The `WorkerPool` struct embodies the worker pool concept.
    - It includes attributes such as the number of worker goroutines (`workers`), a semaphore channel (`semaphore`) to control the concurrent workers, a task queue channel (`taskQueue`) for enqueuing tasks, a `sync.WaitGroup` (`wg`) to track active workers, and a stop signal channel (`stopSignal`) for gracefully stopping the worker pool.

3. **Creating the Worker Pool** 🚀:
    - The `NewWorkerPool` function is responsible for initializing a new worker pool with a specified number of worker goroutines.
    - It starts goroutines for task execution and ensures they are gracefully waited for when the program exits.

4. **Submitting Tasks** 📤:
    - The `Submit` method is used to send tasks to the worker pool for execution.

5. **Waiting for Task Completion** ⏳:
    - The `Wait` method allows the main program to wait for all tasks to complete execution.

6. **Worker Goroutine Functionality** 🧰:
    - The `worker` function represents an active goroutine within the worker pool. It carries out the execution of tasks.
    - To maintain control over the number of active goroutines, the `worker` function acquires a semaphore token before executing the task. After task execution, it releases the acquired semaphore token.

This code showcases the practical implementation of a worker pool, ensuring that a limited number of tasks are executed in parallel. By doing so, it effectively prevents resource overload, provides efficient management of task execution, and ensures smooth system operation.