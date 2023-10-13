## Example 1: Harmonizing Data Streams

In this performance, the "Fan-In" pattern takes the stage, skillfully weaving together data from various channels into a harmonious symphony. Here's how this composition unfolds:

1. **The Ensemble Channel**: In the `FanIn` function, a channel named `dest` is created, serving as the main stage for the grand performance. This channel will unite data from multiple sources, much like a conductor unites the melodies of an orchestra.

2. **The Conductor's Baton**: A `sync.WaitGroup` named `wg` is introduced. It acts as the conductor's baton, ensuring that all musicians (goroutines) finish their performances before the grand finale.

3. **Musical Scores**: A `for` loop is initiated, like a musical score that guides each musician. Within this loop, a goroutine is launched for each data source (channel). These virtuoso goroutines read data from their respective sources (`src`) and gracefully deliver it to the `dest` channel. They repeat this act until their source is exhausted (`for v := range src`).

4. **Curtain Call**: For each source, another goroutine is called into action, conducting a final act of closing the source channel once it has performed its part. This ensures that no notes linger once the performance is complete (`defer close(src)`).

5. **The Grand Finale**: After all the virtuoso goroutines have been summoned to the stage, another goroutine takes the conductor's role. It monitors the musicians' performances, waiting for each one to finish their masterpiece, and only then does it gracefully close the grand `dest` channel (`close(dest)`).

6. **The Maestro's Baton Wave**: In the `main` function, the `sources` function orchestrates the creation of several sources, each a stream of data. Then, the `FanIn` function takes the podium, conducting the performance by passing these sources as the instruments in a symphony. The result is the `dest` channel where all the melodies are merged.

7. **The Awe-Inspiring Crescendo**: Finally, a `for` loop reads from the `dest` channel, like an eager audience absorbing the harmonious data. The data is then eloquently presented to the world.

In summary, the "Fan-In" pattern allows for the concurrent retrieval of data from multiple sources, such as channels, files, or network connections. These data streams are seamlessly woven together into a single channel, creating a magnificent composition for efficient data processing. Just like an orchestra, this pattern brings together the melodies of data sources into a symphony of information.