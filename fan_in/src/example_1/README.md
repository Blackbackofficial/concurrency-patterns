## Example 1:
The "Fan-In" pattern is applied, which is used to combine multiple channels into one, allowing you to read data from several sources and pass them into a single channel for further processing. Here's how it works:

1. In the `FanIn` function, a channel called `dest` is created, which will be returned as the result. This channel will be used to merge data from multiple sources.

2. Then, a variable `wg` of type `sync.WaitGroup` is created, which will be used to wait for the completion of all goroutines working with sources.

3. A `for` loop is initiated, iterating through the provided sources (channels). For each source, a separate goroutine is started. Each of these goroutines reads data from its respective source (`src`) and forwards it to the `dest` channel. This is done within the loop using `for v := range src`.

4. Additionally, for each source, a goroutine is launched to close the `src` channel after it has finished processing data. This is achieved with `defer close(src)`.

5. After all the goroutines responsible for reading from sources are created, another goroutine is initiated to monitor the completion of all these goroutines using `wg.Wait()`. Once all the goroutines have finished, this goroutine closes the `dest` channel using `close(dest)`.

6. In the `main` function, the `sources` function is called, which generates several sources and returns them as a slice. Then, the `FanIn` function is invoked, passing the created sources. The result of the `FanIn` function is the `dest` channel, where data from multiple sources is merged.

7. Subsequently, a `for` loop `for r := range dest` is used to read data from the combined `dest` channel and print it.

In summary, the "Fan-In" pattern allows for concurrent data retrieval from multiple sources (which can be channels, files, network connections, etc.) and combines the data into a single channel for convenient data processing.