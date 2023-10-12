# Concurrency patterns

**10 popular concurrency patterns in Golang:**

1. **The Done channel pattern**: Used for signaling the completion of goroutines through a special "done" channel.

2. **The Fan-In pattern**: Allows you to combine data from multiple channels into one for collecting processing results.

3. **The Fan-Out pattern**: Distributes data to multiple goroutines for parallel processing.

4. **Sharding**: Used to distribute data among multiple goroutines or nodes in distributed systems.

5. **Bounded Parallelism**: Limits the number of concurrently executing goroutines to manage system load.

6. **Backpressure**: Helps control the flow of data in the system to avoid overloading handlers.

7. **The Or-pattern**: Combines results from multiple channels, completing when the first channel returns a result.

8. **The Tee pattern**: Copies input data into multiple output channels for parallel processing.

9. **The Bridge pattern**: Allows you to link multiple channels, passing data between them.

10. **The Pipeline pattern**: Creates sequential data processing, where each processing stage runs in a separate goroutine.