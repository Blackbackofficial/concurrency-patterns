## Example 1: Coordinated Termination

The "Done Channel Pattern" takes center stage in this example to elegantly coordinate the termination of multiple goroutines while receiving notifications of their completion.

1. A `done` channel makes its debut, poised to receive notifications of each goroutine's work termination. The channel's buffer size aligns with `numWorkers`, mirroring the number of goroutines under observation.

2. Within a loop `for i := 0; i < numWorkers; i++`, a chorus of goroutines springs to life. Each of them graciously summons the `worker(id, done)` function. These diligent workers carry out their tasks, and once they've reached the grand finale, they dispatch a `true` signal to the `done` channel. This signals the triumphant conclusion of their efforts.

3. Yet another goroutine emerges, assigned the noble task of keeping vigil over its peers with a steadfast `wg.Wait()`. After every single worker has taken their bow, this guardian closes the curtains on the `done` channel, gently pulling it into the realms of finality with `close(done)`.

4. In the audience, the main goroutine patiently awaits each performer's exit by partaking in a ceremonial dance `for range done`. Whenever a triumphant `true` proclamation is voiced through the `done` channel, the main stage acknowledges this herald with a resounding message, signifying the completion of yet another virtuoso's work.

To encapsulate, the "Done Channel Pattern" steps into the limelight, deftly orchestrating the synchronized execution of multiple goroutines. The `done` channel orchestrates the delivery of notifications for each goroutine's final bow, and the main maestro gracefully listens for the last note, ensuring a harmonious conclusion to their collective performance.