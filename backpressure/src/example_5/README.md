## Example 5:

A worker pool is a common concurrency pattern used to manage a limited number of worker goroutines that can process tasks concurrently.

In this code:

1. `Task` struct represents individual tasks that can be executed. Each task has a name and a duration.

2. `WorkerPool` struct represents the worker pool. It contains a specified number of worker goroutines (`workers`), a channel (`semaphore`) to limit the number of concurrently active workers, a channel (`taskQueue`) for queuing tasks, a `sync.WaitGroup` (`wg`) to track active workers, and a stop signal channel (`stopSignal`) for stopping the worker pool.

3. `NewWorkerPool` function creates a new worker pool with a specified number of workers. It starts goroutines for task execution and waits for their completion when the program exits.

4. `Submit` method is used to send tasks to the worker pool for execution.

5. `Wait` method waits for all tasks to complete.

6. `worker` function represents an active goroutine that executes tasks. It acquires a semaphore to limit the number of active goroutines, executes the task, and releases the semaphore.

The code demonstrates how tasks can be executed concurrently using a worker pool with a limited number of workers. The worker pool ensures that only a controlled number of tasks are executed in parallel, preventing resource overload and efficiently managing task execution.