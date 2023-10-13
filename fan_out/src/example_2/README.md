## Example 2: The Dance with Graceful Stops

In this rendition of the "Fan-Out" pattern, the lead dancer, represented by the source channel (`src`), gracefully directs its moves to a troupe of worker goroutines. The context (`ctx`) adds a new layer of choreography, allowing a graceful exit. Here's the dance:

1. **The Choreographer**: The `FanOutWithCancel` function takes center stage, alongside the source channel (`src`), the chosen number of worker goroutines (`n`), and the graceful context (`ctx`). The task at hand is to distribute the lead dancer's moves and provide an exit strategy.

2. **Preparation on Stage**: The `FanOutWithCancel` function creates an array of destination channels (`dests`) and launches the same number of worker goroutines. Each worker is assigned its own destination channel, all in harmonious coordination.

3. **In the Spotlight**: Inside each worker's routine, a `for` loop is the spotlight where they await the lead dancer's moves. These dedicated workers read data from the source channel and pass it to their designated destination channel. They also keep a vigilant eye on the graceful exit signal from the context (`ctx`).

4. **Opening Act**: The show begins with the `main` function setting the stage. It creates the source channel with the help of the `source` function, filling it with integers ranging from 1 to 10.

5. **Fan-Out and Graceful Dance**: The `FanOutWithCancel` function orchestrates the dance, gracefully distributing the data from the source channel to a group of worker goroutines (in this case, 5), each with its own destination channel. A `sync.WaitGroup` (`wg`) makes sure everything is in perfect harmony.

6. **Grand Ensemble**: Each destination channel prepared by `FanOutWithCancel` gets its own dedicated dancer, represented by a worker goroutine. These dancers read data from their assigned destination channels and dance with grace.

7. **Simulated Work**: After some graceful moves, the `main` function introduces a 1-second pause with `time.Sleep(1 * time.Second)` to simulate some work.

8. **The Final Bow**: The context is gracefully canceled using `cancel()`, setting the stage for a graceful exit.

9. **Curtains Fall with Applause**: The `main` function waits for all worker goroutines to complete their performance with `wg.Wait()`. After a round of applause, the program exits gracefully.

In summary, this performance beautifully demonstrates the "Fan-Out" pattern with context cancellation. Data is gracefully distributed from a source channel to multiple worker goroutines for parallel processing, with the added elegance of graceful exits, ensuring a performance with both precision and poise.