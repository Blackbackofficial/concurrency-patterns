## Example:

The "Fan-Out" pattern with context cancellation, which allows distributing data from a single source channel to multiple worker goroutines and gracefully stopping the distribution process. Here's how it works:

1. The `FanOutWithCancel` function takes a source channel (`src`), the number of worker goroutines (`n`), and a context (`ctx`) for cancellation.

2. It creates an array of destination channels (`dests`) and launches `n` goroutines, each with its own destination channel. These destination channels are used to distribute the data from the source channel to the workers.

3. Inside each worker goroutine, there is a `for` loop that reads data from the source channel and forwards it to the associated destination channel. The worker goroutines also listen for signals from the context to gracefully exit the loop and terminate.

4. The `main` function starts by creating a source channel using the `source` function, which populates the source channel with integers from 1 to 10.

5. Then, the `FanOutWithCancel` function is called with the source channel, the desired number of worker goroutines (in this case, 5), and a context for cancellation.

6. A `sync.WaitGroup` (`wg`) is used to ensure that the `main` function waits for all workers to finish processing.

7. For each destination channel created by `FanOutWithCancel`, a worker goroutine is started. These worker goroutines read data from their respective destination channels and print it.

8. To simulate some work, the `main` function introduces a 1-second delay using `time.Sleep(1 * time.Second`.

9. After the simulated work, the context is canceled with `cancel()`, which triggers the graceful termination of worker goroutines.

10. The `main` function waits for all worker goroutines to complete using `wg.Wait()`. After all workers finish processing and the context is canceled, the program exits gracefully.

In summary, the "Fan-Out" pattern with context cancellation efficiently distributes data from a single source to multiple worker goroutines, allowing for concurrent processing, and provides a means for graceful termination using context cancellation.
