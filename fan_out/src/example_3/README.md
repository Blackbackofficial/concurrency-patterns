## Example 3: The Synchronized Ensemble

In this rendition of the "Fan-Out" pattern, a synchronized ensemble of worker goroutines receives instructions from the lead dancer (the source channel) and performs a graceful dance. Here's how they synchronize their steps:

1. **Lead Dancer and the Troupe**: The `FanOutWithWaitGroup` function sets the stage with a source channel (`src`) and assembles a troupe of worker goroutines. Each worker has a dedicated destination channel (`dest`) to receive the dance moves.

2. **Dance Preparation**: The function creates an array of destination channels (`dests`) and deploys worker goroutines, each with a designated destination channel. These channels facilitate the synchronized dance.

3. **In the Spotlight**: Inside each worker's routine, a `for` loop is the spotlight where they await the lead dancer's moves. These dancers read data from the source channel and then execute their graceful moves in parallel by forwarding the data to their respective destination channels. Closing the destination channels is their final bow.

4. **The Grand Stage**: In the `main` function, a grand stage is set. A source channel (`src`) is created to hold the integer values from 1 to 10.

5. **Preparing the Dance**: A dedicated goroutine is tasked with populating the source channel with these values. It also takes a graceful exit by deferring the closure of the source channel when the performance ends.

6. **Curtain Rises**: The `FanOutWithWaitGroup` function is called, and the grand ensemble of worker goroutines is given their destination channels. A `sync.WaitGroup` (`wg`) is assigned to oversee the synchronization.

7. **The Ensemble Takes the Stage**: Each destination channel, as assigned by `FanOutWithWaitGroup`, gets its own dedicated dancer. These workers read the data from their source channels and perform their dance steps by printing the worker's index and the data received.

8. **Final Bow and Applause**: After the ensemble's graceful performance, the program patiently waits for all workers to take their final bow with the `wg.Wait()` call.

In summary, this synchronized ensemble performs the "Fan-Out" pattern gracefully. Data is efficiently distributed from a source channel to multiple worker goroutines, and their synchronized dance makes the most of available resources.