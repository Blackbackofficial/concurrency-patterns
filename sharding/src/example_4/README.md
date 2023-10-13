## Example 4:

In this example, the "Sharding" pattern is applied to effectively balance and distribute requests among multiple shards. Here's a breakdown of how it works:

1. **Data Structure**: The `Request` struct is used to represent individual requests, each having an `ID` (an integer) and `Data` (a string).

2. **Main Function**:
   - An array of `Request` elements, `requests`, is created to represent the requests that need to be balanced and distributed across shards.
   - A slice of slices of `Request`, `shards`, is created to simulate the shards. In this example, three shards are employed.
   - A `sync.WaitGroup` (`wg`) is utilized to ensure that all goroutines finish their tasks before the program proceeds.

3. **Goroutines for Request Distribution**:
   - For each `Request` in the `requests` array, a goroutine is initiated. Each goroutine is responsible for distributing a request to one of the three shards.
   - The `loadBalance` function is employed to determine the shard index for a given request based on its `ID`. It uses a simple modulus operation to ensure requests are evenly distributed among the three shards.
   - Subsequently, the goroutine appends the request to the appropriate shard.

4. **Waiting for Goroutines**: In the `main` function, a call to `wg.Wait()` ensures that all goroutines complete their work before proceeding.

5. **Outputting Requests from Shards**:
   - After all requests have been evenly distributed, the `main` function proceeds to output the requests stored within each shard.
   - It systematically iterates through the `shards` slice, printing both the request IDs and their corresponding data within each shard.

The "Sharding" pattern showcased in this code is a powerful technique for the effective distribution of requests among multiple shards, facilitating load balancing and efficient request management. This pattern proves particularly valuable in scenarios where incoming requests need to be uniformly distributed across various partitions or shards to enhance overall system performance and scalability.