## Example 3:
This code also demonstrates the "Fan-In" pattern for merging data from multiple channels into a single channel. Here's how it works:

1. The `mergeData` function takes a `done` channel and multiple data channels (`datas...`) as arguments. The `done` channel is used to signal the termination of the merging process, while the `datas` channels contain data to be merged.

2. Inside the `mergeData` function, a `sync.WaitGroup` (`wg`) is created to keep track of the goroutines launched for each data channel. Another channel, `merged`, is created to store the merged data.

3. For each data channel in the `datas` slice, a goroutine is launched. These goroutines continuously select data from the input channels and send it to the `merged` channel. They also listen for signals from the `done` channel to stop processing and return.

4. The `merged` channel is closed once all the goroutines have finished processing their data. A separate goroutine waits for the completion of the goroutines using `wg.Wait()` and then closes the `merged` channel using `close(merged)`.

5. In the `main` function, a `done` channel is created to signal the termination of the merging process. This channel is deferred for closure to ensure it's closed when the program exits.

6. Two data channels (`data1` and `data2`) are created and populated with integer values using separate goroutines.

7. The `mergeData` function is called with the `done` channel and the two data channels as arguments. This function combines the data from `data1` and `data2` into the `merged` channel.

8. In the `main` function, a loop reads and prints values from the `merged` channel using `for val := range merged`. This loop continues until all data has been merged and processed.

In summary, this code demonstrates the "Fan-In" pattern to merge data from multiple channels into a single channel efficiently. The use of the `done` channel allows for graceful termination of the merging process, and the code efficiently combines data from different sources asynchronously.