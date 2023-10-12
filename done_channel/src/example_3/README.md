## Example 3:

This variation introduces a timeout mechanism to handle cases where a goroutine might not produce a result in a timely manner, preventing potential resource leaks.

Here's how this pattern works:

1. A `result` channel of type `interface{}` is created to receive the result of some computation or operation performed by a goroutine.

2. A `done` channel of type `struct{}` (an empty struct, also known as a signal channel) is created. This channel will be used to signal the goroutine to stop if a timeout occurs.

3. A new goroutine is launched, which represents some asynchronous task. This goroutine attempts to send a result to the `result` channel. However, it does so within a `select` statement, which has two cases:
    - The first case attempts to send the result to the `result` channel when the operation is completed.
    - The second case listens for a signal from the `done` channel. If the `done` channel is closed, it immediately aborts the operation, preventing it from blocking indefinitely.

4. In the main part of the code, there is another `select` statement. This one waits for one of two events:
    - If a result is received from the `result` channel, it is processed.
    - If a timeout of one second elapses (as specified by `time.After(time.Second)`), it handles the timeout scenario.

5. After either a result is received or a timeout occurs, the code closes the `done` channel to signal to the asynchronous goroutine that it should clean up and exit.

This pattern allows for the concurrent execution of a task with the ability to handle timeouts effectively. If the asynchronous task takes too long to produce a result, the timeout mechanism ensures that resources are not tied up indefinitely. It's a useful pattern when dealing with scenarios where timely responses are essential, and you want to ensure that the system can gracefully handle situations where responses are delayed or never arrive.