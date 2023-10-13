## Example 1: Bounded Parallelism in Action

The "Bounded Parallelism" pattern is expertly demonstrated in this example, showcasing how it effectively limits the number of concurrent operations to prevent system overload. Let's delve into the details of how this pattern is implemented:

1. **Process Function**:
   - The `Process` function is the heart of this pattern, taking two vital parameters: `concurrency` and `data`.
   - `concurrency` is a crucial parameter, specifying the maximum number of concurrent operations that can be executed simultaneously.
   - `data` is an array of integers, representing the data that needs to be processed.

2. **Semaphore (sem)**:
   - The implementation begins with the creation of a channel called `sem` with a type of `struct{}` and a capacity of `concurrency`. This channel acts as a semaphore, allowing only a restricted number of goroutines to execute concurrently.

3. **WaitGroup (wg)**:
   - A `sync.WaitGroup` aptly named `wg` comes into play. It serves the purpose of waiting for all goroutines to finish their assigned work.

4. **Concurrent Data Processing**:
   - A loop gracefully traverses the `data` slice. For each element `d`, it tactfully acquires a semaphore slot by sending a value into the `sem` channel (`sem <- struct{}{}`). This operation is strategically designed to block if the channel is full, expertly limiting the number of concurrent goroutines to the specified `concurrency`.
   - A goroutine is judiciously launched for every element in `data`. Within the confines of each goroutine, the `doProcess` function is summoned to diligently process the data.
   - The `doProcess` function adeptly simulates work by inducing a one-second sleep.
   - Upon completion of the processing, the goroutine liberally releases the semaphore slot by deftly receiving a value from the `sem` channel (`<-sem`), thereby replenishing capacity for other goroutines.

5. **Waiting for Goroutines to Conclude**:
   - The harmonious call to `wg.Wait()` stands guard, ensuring that the program gracefully lingers until every single goroutine has diligently completed its work.

6. **Main Function**:
   - In the grand symphony of this code, data gracefully makes its entrance through the `getData` function. This function assembles an array of integers, spanning from 0 to 9.
   - The star of the show, the `Process` function, is brought into the limelight with a concurrency cap of 3, signifying that, at most, three data processing operations will proceed concurrently.

7. **Output**:
   - The crescendo of this performance is the `doProcess` function, artfully simulating the processing of each piece of data. With a one-second sleep, it gracefully marks the completion of each operation, elegantly signaling its conclusion through printed messages.