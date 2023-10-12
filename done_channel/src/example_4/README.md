## Example 4:

This code demonstrates a scenario where the "Done channel pattern" is not directly applied, but the `context.Context` with a cancellation mechanism is used to achieve a similar goal.

Here's how this code works:

1. The `request` function simulates an asynchronous operation that takes approximately 5 seconds to complete. This operation is monitored for cancellation through the provided `context.Context`.

2. In the `main` function, a `context` is created using `context.WithCancel(context.Background())`, allowing you to cancel the context when needed. The `defer cancel()` statement is used to ensure that the context is canceled, but it's not relied upon for timely cancellation. Instead, explicit cancellation is performed later in the code.

3. A buffered channel named `result` is created to receive the result of the `request` operation.

4. A goroutine is launched to execute the `request` function. The goroutine uses the provided context `ctx` to perform the operation. If the operation completes successfully, it sends the result to the `result` channel. If an error occurs or the context is canceled, it prints an error message but does not block indefinitely.

5. The main part of the code uses a `select` statement to wait for one of the following conditions:
    - If a result is received from the `result` channel, it prints the result.
    - If one second elapses (as specified by `time.After(time.Second)`), it prints a "request timeout" message.

6. After waiting for either the result or the timeout, the `cancel()` function is explicitly called to cancel the context. This is important to ensure timely cancellation and cleanup.

7. Finally, a small delay with `time.Sleep` is introduced to allow time for the sub-goroutine to print a "cancel" message. This is not directly related to the "Done channel pattern" but serves as a waiting mechanism.

In this code, the `context.Context` is used for signaling and canceling the operation, similar to how a done channel would be used in the "Done channel pattern." The primary difference is that the cancellation is managed via the `context` itself, allowing for clean and timely cancellation of the operation, which is essential in many concurrent scenarios.