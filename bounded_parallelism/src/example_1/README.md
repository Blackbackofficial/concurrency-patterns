## Example 1:

The "Bounded Parallelism" pattern, which limits the number of concurrent operations to a specific level, ensuring that the system does not become overloaded. Here's an explanation of how it works:

1. **Process Function**:
    - The `Process` function takes two parameters: `concurrency` and `data`.
    - `concurrency` specifies the maximum number of concurrent operations that can be executed at the same time.
    - `data` is a slice of integers that represents the data to be processed.

2. **Semaphore (sem)**:
    - A channel called `sem` of type `struct{}` is created with a capacity of `concurrency`. This channel acts as a semaphore, allowing only a limited number of goroutines to execute concurrently.

3. **WaitGroup (wg)**:
    - A `sync.WaitGroup` called `wg` is created to wait for all goroutines to finish their work.

4. **Processing Data Concurrently**:
    - A loop iterates through the `data` slice, and for each element `d`, it acquires a semaphore slot by sending a value into the `sem` channel (`sem <- struct{}{}`). This operation blocks if the channel is full, effectively limiting the number of concurrent goroutines to `concurrency`.
    - A goroutine is started for each element in `data`. Inside the goroutine, the `doProcess` function is called to process the data.
    - The `doProcess` function simulates some work by sleeping for one second.
    - After the processing is complete, the goroutine releases the semaphore slot by receiving a value from the `sem` channel (`<-sem`), freeing up capacity for another goroutine.

5. **Waiting for Goroutines to Finish**:
    - The `wg.Wait()` call waits for all goroutines to finish their work. It ensures that the program doesn't exit until all processing is complete.

6. **Main Function**:
    - In the `main` function, data is retrieved using the `getData` function, which creates a slice of integers from 0 to 9.
    - The `Process` function is called with a concurrency limit of 3, which means that at most three data processing operations will occur concurrently.

7. **Output**:
    - The `doProcess` function simulates processing each piece of data by sleeping for one second. As each operation completes, it prints a message indicating the completion of processing.
