## Example 2:
This example demonstrates the "Fan-In" pattern, which is used to merge data from multiple channels into a single channel. Here's how it works:

1. In the `mergeChannels` function, a channel named `merged` is created. This channel will be used to combine data from multiple input channels, referred to as `channels`.

2. A goroutine is created for each of the input channels (`channels`). These goroutines read data from their respective input channels (`ch1`, `ch2`, `ch3`) and send that data to the `merged` channel. Each goroutine waits for the completion of this operation using `defer wg.Done()`.

3. The `merged` channel is considered closed once data has been read from all the input channels. This closing operation is performed in an additional goroutine, which uses a `sync.WaitGroup` (`wg`) to track the completion of all goroutines. When all goroutines have completed and signified this with `wg.Done()`, this additional goroutine closes the `merged` channel using `close(merged)`.

4. In the `main` function, three input channels (`ch1`, `ch2`, and `ch3`) are created. Each of these channels is populated with numbers within the respective ranges: 1 to 5, 6 to 10, and 11 to 15.

5. After creating and populating the input channels, the `mergeChannels` function is called with these channels as arguments. The `mergeChannels` function combines the data from all input channels into the `merged` channel.

6. Then, in the `main` function, data is read from the `merged` channel using a `for val := range merged` loop. This loop allows for sequentially reading and printing the combined data from all input channels.

In summary, this code uses the "Fan-In" pattern to merge data from multiple input channels into a single channel, allowing you to efficiently work with data from different sources asynchronously.