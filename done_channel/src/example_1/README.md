## Example 1:

The "Done channel pattern" is applied here to coordinate the termination of multiple goroutines and receive notifications of their completion.

1. A `done` channel is created, which will be used to receive notifications of the termination of each goroutine's work. The `done` channel has a buffer with a size equal to `numWorkers`, representing the number of goroutines to wait for.

2. A goroutine is created using a `for i := 0; i < numWorkers; i++` loop. In each goroutine, the `worker(id, done)` function is called, which performs some work and, upon completion, sends the value `true` to the `done` channel. This signals the completion of the goroutine's work.

3. Another goroutine is created to wait for the termination of all goroutines using `wg.Wait()`. After all goroutines have finished, this goroutine closes the `done` channel using `close(done)`.

4. The main goroutine waits for notifications of each goroutine's completion using `for range done`. As soon as the value `true` is sent to the `done` channel, the main goroutine receives this notification and prints a message indicating the completion of the goroutine's work.

In summary, the Done Channel pattern is used for synchronization and control of the execution of multiple goroutines. The `done` channel is employed to send notifications of each goroutine's completion, and the main goroutine waits to receive all these notifications, allowing for precise determination of when all the goroutines have finished.