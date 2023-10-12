## Example 1:

The "Fan-Out" pattern, which is used to distribute data from one source channel to multiple worker goroutines. Here's how it works:

1. The `FanOut` function takes a source channel (`src`) and the number of worker goroutines (`n`) to distribute the data to.

2. It creates an array of destination channels (`dests`) and launches `n` goroutines, each with its own destination channel. These destination channels are used to distribute the data from the source channel to the workers.

3. Inside each worker goroutine, a `for` loop reads data from the source channel and forwards it to the associated destination channel. This allows the data to be distributed concurrently to all worker goroutines.

4. The `main` function starts by creating a source channel using the `source` function, which populates the source channel with integers from 1 to 10.

5. Then, the `FanOut` function is called with the source channel and the desired number of worker goroutines (in this case, 5). The `FanOut` function creates the destination channels and distributes the data to the workers.

6. A `sync.WaitGroup` (`wg`) is used to ensure that the `main` function waits for all workers to finish processing.

7. For each destination channel created by `FanOut`, a worker goroutine is started. These worker goroutines read data from their respective destination channels and print it.

8. The `main` function waits for all worker goroutines to complete using `wg.Wait()`. After all workers finish processing, the program exits.

In summary, the "Fan-Out" pattern efficiently distributes data from a single source to multiple worker goroutines, allowing for concurrent processing and workload distribution. It is especially useful in scenarios where you need to parallelize tasks across multiple workers.