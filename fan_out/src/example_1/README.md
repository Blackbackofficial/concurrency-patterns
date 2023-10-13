## Example 1: Spreading the Load

In this graceful performance of the "Fan-Out" pattern, the lead dancer is the source channel (`src`) distributing its elegant moves to a troupe of worker goroutines. Here's the choreography:

1. **The Dance Director**: The `FanOut` function steps onto the stage, accompanied by the source channel (`src`) and a group of worker goroutines (`n`). The task at hand is to distribute the lead dancer's moves to each worker.

2. **Creating Dance Partners**: The `FanOut` function prepares the dance floor by crafting an array of destination channels (`dests`) and launching a corresponding number of worker goroutines. Each worker is given its own destination channel to ensure a smooth dance.

3. **The Choreographed Routine**: Within the spotlight of each worker's routine, a `for` loop is the stage where they eagerly await the lead dancer's moves. They read data from the source channel and gracefully pass it to their designated destination channel, all in perfect synchronization.

4. **The Grand Opening**: The show begins as the `main` function raises the curtains by creating the source channel with the help of the `source` function. This talented source channel is filled with integers ranging from 1 to 10, setting the stage for the grand performance.

5. **Fan-Out, Let the Dance Begin**: The `FanOut` function takes center stage, guiding the source channel and the chosen number of worker goroutines (in this case, 5). It orchestrates the distribution of data, creating destination channels for the dancers.

6. **Master of Ceremonies**: A `sync.WaitGroup` (`wg`) takes the role of the master of ceremonies, ensuring that the `main` function patiently waits for all dancers to complete their performance.

7. **The Dance Ensemble**: For each destination channel carefully crafted by `FanOut`, a dancer enters the spotlight. These dancers are represented by worker goroutines, each with the duty to read the data from their designated destination channel and perform their moves.

8. **A Standing Ovation**: The `main` function applauds the dancers, waiting for each of them to take their final bow with `wg.Wait()`. After every dancer has completed their dance, the curtains fall, and the show gracefully concludes.

In summary, this code elegantly demonstrates the "Fan-Out" pattern, where data from a single source channel is distributed efficiently to multiple worker goroutines for parallel processing. It's a dance of concurrency, a symphony of synchronization, and a performance of parallelism.