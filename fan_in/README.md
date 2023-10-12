## The "Fan-In"
The "Fan-In" pattern allows combining data from multiple sources into one, enabling asynchronous reading and passing of data into a single channel. Key steps:

1. Create a destination channel.

2. Launch goroutines to read and pass data to the destination channel.

3. Use signaling channels for synchronization.

4. Close the destination channel after goroutines finish.

This is useful for efficient asynchronous processing of data from various sources.


### Example 1:
The "Fan-In" pattern is applied to combine data from various sources (channels) into one, allowing asynchronous reading of data from multiple sources and passing it to a single channel for further processing.

### Example 2:
This example also uses the "Fan-In" pattern to merge data from multiple channels into one. This enables efficient asynchronous handling of data from different sources.

### Example 3:
Here, the "Fan-In" pattern is similarly used to merge data from different channels, with the provision for graceful termination of the merging operation using an additional signal channel, `done`. This ensures efficient merging of data from diverse sources and a harmonious conclusion to the operation.

Each of these examples illustrates the application of the "Fan-In" pattern to combine data from different sources into a single channel, offering convenience and efficiency for asynchronous data processing from various origins.