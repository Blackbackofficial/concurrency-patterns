## Example 4:

The "Sharding" pattern is applied to balance and distribute requests among multiple shards. Here's how it works:

1. **Data Structure**: The `Request` struct is used to represent individual requests, each with an `ID` (an integer) and `Data` (a string).

2. **Main Function**:
    - An array of `Request` elements, `requests`, is created to represent the requests that need to be balanced and distributed across shards.
    - A slice of slices of `Request`, `shards`, is created to simulate the shards. In this case, three shards are created.
    - A `sync.WaitGroup` (`wg`) is used to ensure that all goroutines finish their work before proceeding.

3. **Goroutines for Request Distribution**:
    - For each `Request` in the `requests` array, a goroutine is started. Each goroutine is responsible for distributing a request to one of the three shards.
    - The `loadBalance` function calculates the shard index for a given request based on its `ID`. It uses a simple modulus operation to evenly distribute requests among the three shards.
    - The goroutine then appends the request to the appropriate shard.

4. **Waiting for Goroutines**: The `main` function waits for all goroutines to finish their work using the `wg.Wait()` statement.

5. **Outputting Requests from Shards**:
    - After all requests are distributed, the `main` function proceeds to output the requests from all shards.
    - It iterates through the `shards` slice and prints the request IDs and data within each shard.

The "Sharding" pattern in this code demonstrates the distribution of requests among multiple shards, enabling load balancing and efficient request management. It's particularly useful in scenarios where incoming requests need to be evenly distributed across different partitions or shards for better performance and scalability.