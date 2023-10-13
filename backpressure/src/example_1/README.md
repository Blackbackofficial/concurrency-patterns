## Example 1:

This code uses the "backpressure" pattern using the `PressureGauge` structure. Here's how it works:

1. `PressureGauge` is a structure representing the "backpressure" control mechanism. It contains a channel `ch`, which has a limited bandwidth specified when creating an instance of the structure.

2. `New(limit int)` is a function that initializes `PressureGauge'. It fills the channel `ch` with the specified limit representing the maximum number of simultaneous tasks.

3. `Do(fn func()) error` is the `PressureGauge` method, which is used to perform the task taking into account "backpressure". It tries to extract the value from the `ch` channel. If there is bandwidth available, it executes the task (passed as a function `fn`) and then returns the bandwidth by adding the value back to the channel `ch'. If there is no bandwidth, it returns the error "out of capacity".

4. In the `runServer(pg *PressureGauge) function`PressureGauge` is used to control the number of simultaneous requests to the server. If there is available bandwidth (which is limited to 3 in this case), the server performs an "expensive operation" (the `expensiveOperation` function) and sends a response. If there is no available bandwidth, the server returns the HTTP status "429 Too Many Requests".

5. In the `main()` function, several goroutins are launched that make HTTP requests to the server. Using `PressureGauge`, they are limited to three simultaneous requests, which demonstrates the use of the "backpressure" pattern to control the flow of data and prevent server overload.

This code illustrates the use of "backpressure" to manage the load on the server and effectively manage resources, preventing system overload.