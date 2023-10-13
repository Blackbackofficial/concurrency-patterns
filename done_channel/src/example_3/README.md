## Example 3: Graceful Timeout Handling

In this entrancing act, we introduce a variation of the "Done Channel Pattern" that incorporates a timeout mechanism, ensuring that all participants respect their roles and preventing potential resource leaks.

Here's how this pattern performs:

1. Our play begins with the creation of a `result` channel, a vessel to receive the results of the mystical computations or operations performed by a dedicated goroutine.

2. A `done` channel, a signal channel marked by the distinctive `struct{}`, makes its appearance. This channel serves as a signal to gently beckon the goroutine to conclude its performance if a timeout should occur.

3. The grand launch commences with the rise of a new goroutine, symbolizing an asynchronous task. This dedicated performer seeks to send the fruits of its labor to the `result` channel. However, the performance is orchestrated within a `select` statement, an ensemble featuring two cases:
   - The first case elegantly presents the completed result to the `result` channel, signifying the conclusion of the task.
   - The second case listens intently for a signal from the `done` channel. If the `done` channel is discreetly closed, it prompts an immediate and graceful exit, preventing the task from extending its performance indefinitely.

4. The main stage is graced with another `select` statement, patiently awaiting one of two possible acts:
   - A result gracefully bestowed by the `result` channel is processed and celebrated.
   - If a duration of one second passes in solemn silence (as indicated by `time.After(time.Second)`), a perfectly choreographed timeout act is performed.

5. After the conclusion of either act—receiving a result or experiencing a timeout—a final flourish emerges. The code discreetly closes the `done` channel, signaling to the asynchronous performer to conclude its performance gracefully.

This pattern artfully allows for the concurrent execution of tasks while ensuring the graceful handling of timeouts. It is an invaluable addition to the repertoire when dealing with scenarios where timely responses are essential, and the show must go on even in the face of delayed or absent responses.