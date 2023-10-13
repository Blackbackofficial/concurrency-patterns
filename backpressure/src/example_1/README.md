## Example 1: Managing Server Load with "Backpressure"

In this example, the "backpressure" pattern is implemented using the `PressureGauge` structure to control the load on a server. Here's how it works:

1. **PressureGauge Structure**:
    - The `PressureGauge` structure serves as the "backpressure" control mechanism.
    - It includes a channel, `ch`, which has a limited bandwidth specified when creating an instance of the structure.

2. **Initialization with Limit**:
    - The `New(limit int)` function initializes the `PressureGauge`. It pre-fills the `ch` channel with a specified limit, representing the maximum number of concurrent tasks allowed.

3. **Performing Tasks with "Backpressure"**:
    - The `Do(fn func()) error` method in the `PressureGauge` is used to execute tasks while considering "backpressure."
    - It attempts to extract a value from the `ch` channel. If bandwidth is available, the task (provided as a function `fn`) is executed. After task completion, the bandwidth is released by returning the value to the `ch` channel. If no bandwidth is available, it returns an error, indicating "out of capacity."

4. **Server Load Management**:
    - In the `runServer(pg *PressureGauge)` function, `PressureGauge` is utilized to manage the number of concurrent requests to the server.
    - If available bandwidth (limited to 3 in this case) exists, the server performs an "expensive operation" using the `expensiveOperation` function and sends a response.
    - In case there's no available bandwidth, the server responds with an HTTP status code of "429 Too Many Requests."

5. **Request Handling with "Backpressure"**:
    - In the `main()` function, several goroutines are launched to make HTTP requests to the server.
    - By leveraging the `PressureGauge`, these requests are restricted to a maximum of three simultaneous connections. This effectively showcases the application of the "backpressure" pattern to regulate data flow and prevent server overload.

This code exemplifies the practical use of "backpressure" in managing server load, ensuring resource efficiency, and averting system overload. It is a powerful pattern for controlling the flow of data and optimizing system performance.