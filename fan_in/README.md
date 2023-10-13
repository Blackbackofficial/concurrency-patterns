## The "Fan-In" Pattern

The "Fan-In" pattern orchestrates the harmonious convergence of data from multiple sources into a single channel. It's like a symphony where various instruments join forces to create a unified melody. Here's how this pattern plays out:

1. **Create a Destination Channel**: Imagine it as the stage where all performers will ultimately share their artistry—a channel to receive the combined data.

2. **Launch Goroutines**: Just like musicians taking their positions, goroutines are orchestrated to read data from multiple sources and pass it to the destination channel asynchronously.

3. **Signaling for Synchronization**: The conductor uses signaling channels to ensure that the performance is synchronized, much like a well-choreographed dance.

4. **Curtain Call**: Once all performers have taken their final bow, the destination channel is gracefully closed, concluding the performance.

This pattern is a powerful tool for efficiently processing data from various sources asynchronously, much like the instruments of an orchestra coming together to create a masterpiece.

### Example 1: A Symphony of Data
In this act, the "Fan-In" pattern takes center stage, harmonizing data from diverse sources (channels). It showcases the beauty of asynchronous data handling, elegantly merging information from multiple origins into one channel for further processing.

### Example 2: The Ensemble Continues
In this sequel, the "Fan-In" pattern returns to unite data from various channels once again. The performance demonstrates the elegance of asynchronous data handling, allowing different sources to contribute their data efficiently.

### Example 3: A Grand Finale
The third installment maintains the tradition of the "Fan-In" pattern, merging data from diverse channels. What sets this act apart is the addition of a signal channel, "done," which serves as the conductor's final baton. It ensures that the merging operation concludes gracefully, harmonizing data from various sources until the last note.

Each of these examples showcases the application of the "Fan-In" pattern to unite data from different sources into a single channel, offering convenience and efficiency for asynchronous data processing from various origins. Just like a symphony, the "Fan-In" pattern orchestrates a melodious collaboration of data, creating a harmonious result.