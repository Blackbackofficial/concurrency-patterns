## The "Done channel pattern"

The provided examples demonstrate various scenarios where synchronization and coordination of goroutines are essential. Here's a summary of each example:

### Example 1:
- Demonstrates the "Done channel pattern" for coordinating the termination of multiple goroutines.
- A `done` channel is used for notifications of goroutines' completion.
- The main goroutine waits for notifications from each goroutine to determine when they have all finished.

### Example 2:
- Shows the use of the "Done channel pattern" to coordinate and wait for the completion of multiple HTTP requests.
- A `done` channel with a buffer is used to collect results from the goroutines.
- The main goroutine processes results as they become available.

### Example 3:
- Introduces a timeout mechanism to handle potential delays or timeouts in goroutines.
- A `done` channel is used to signal the asynchronous task to stop if a timeout occurs.
- This pattern allows for the handling of timeouts in concurrent operations.

### Example 4:
- Demonstrates the use of the `context.Context` with a cancellation mechanism to achieve similar synchronization.
- A `context` is created to monitor the operation's progress.
- The main goroutine uses a `select` statement to wait for results or a timeout, and it cancels the operation explicitly.

These examples showcase different aspects of concurrent programming in Go, emphasizing the importance of coordination, synchronization, and handling potential issues like timeouts.