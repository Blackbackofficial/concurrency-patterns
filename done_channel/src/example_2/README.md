## Example 2:

The "Done channel pattern" is used for coordinating and waiting for the completion of multiple HTTP requests.

Here's how the pattern is applied:

1. A `done` channel of type `chan RequestResult` is created with a buffer size equal to the number of URLs in the `urls` slice. This channel will be used to send the results of the HTTP requests.

2. A loop iterates through the `urls` slice. For each URL, a new goroutine is launched by calling `go processRequest(url, done)`. Each goroutine sends the result of the HTTP request to the `done` channel upon completion.

3. The main goroutine waits for the completion of all HTTP requests using a loop that iterates over the number of URLs. It receives the results from the `done` channel using `<-done`. This allows the main goroutine to process the results as they become available.

4. For each received result, the code checks whether the request was successful (`result.Success`). If the request was successful, it prints a message indicating success along with the URL. If the request failed, it prints an error message with the URL and the specific error that occurred.

5. Finally, after processing all the results, the `done` channel is closed using `close(done)`.

This pattern allows for concurrent execution of multiple HTTP requests and provides a way to collect and process the results as they arrive, making it a useful approach for handling parallel tasks.
