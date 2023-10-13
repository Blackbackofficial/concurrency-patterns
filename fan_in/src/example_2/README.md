## Example 2: A Symphony of Streams

In this remarkable performance, the "Fan-In" pattern takes the stage once more, gracefully merging data from multiple channels into a magnificent composition. Let's take a closer look at how this symphony unfolds:

1. **The Maestro's Baton**: The `mergeChannels` function sets the stage with the creation of the `merged` channel. This channel serves as the conductor's baton, orchestrating the combination of melodies from various sources.

2. **Virtuoso Soloists**: For each input channel (`ch1`, `ch2`, and `ch3`), a virtuoso goroutine is summoned. Each of these virtuosos plays a unique melody, reading data from their respective input channels and eloquently delivering it to the grand `merged` channel. They exit the stage gracefully when their performance is complete, thanks to the `defer wg.Done()` encore.

3. **The Grand Finale**: The final act is a masterpiece. An additional goroutine ensures that the `merged` channel is closed only when all virtuosos have completed their performances. The `sync.WaitGroup` (`wg`) keeps a keen eye on each performer, and once they all bow out with `wg.Done()`, the conductor's baton gracefully closes the `merged` channel (`close(merged)`).

4. **A Multifaceted Overture**: In the `main` function, a trio of input channels (`ch1`, `ch2`, and `ch3`) are skillfully crafted. Each channel has its own range of numbers, ready to contribute to the symphony: 1 to 5, 6 to 10, and 11 to 15.

5. **The Harmonious Collaboration**: After preparing and populating the input channels, the `mergeChannels` function is called upon, and it harmoniously blends the melodies from the trio into the grand `merged` channel.

6. **The Enthralling Crescendo**: The `main` function takes a front-row seat as it listens to the grand composition, orchestrating a loop to read and appreciate the harmonious data from the `merged` channel. Each note, carefully combined from the individual input channels, is revealed in its full glory.

In summary, the "Fan-In" pattern orchestrates the merging of data from multiple input channels into a single harmonious channel, enabling efficient and asynchronous data processing. Just like a symphony composed of diverse musical instruments, this pattern brings together melodies from various sources into a harmonious composition for seamless data handling.