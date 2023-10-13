## Example 3:

The "Sharding" pattern is also applied to distribute and manage data among multiple shards. Here's how it works:

1. **Data Structure**: The `Data` struct is used to represent key-value pairs, where `Key` is a string and `Value` is an integer.

2. **Main Function**:
    - An array of `Data` elements is created, representing the data that needs to be distributed across shards.
    - A slice of maps (`shards`) is created to simulate the shards, in this case, three shards.
    - A `sync.WaitGroup` (`wg`) is used to ensure that all goroutines finish their work.

3. **Goroutines for Data Distribution**:
    - For each `Data` element in the `data` slice, a goroutine is started. Each goroutine is responsible for distributing a piece of data to one of the three shards.
    - The `hashCode` function is used to calculate the shard index for a given key. It uses a simple hash function for this demonstration, but in real applications, a more complex hashing method should be employed. The shard index is determined by the result of `hashCode(d.Key) % 3`, ensuring even distribution among the three shards.
    - The goroutine then adds the data to the appropriate shard.

4. **Waiting for Goroutines**: The `main` function waits for all goroutines to finish their work using the `wg.Wait()` statement.

5. **Outputting Data from Shards**:
    - After all data is distributed, the `main` function proceeds to output data from all shards.
    - It iterates through the `shards` slice and prints the keys and values within each shard.

The "Sharding" pattern in this code demonstrates the distribution of data among multiple shards, allowing for parallel processing and efficient data management. It's particularly useful in scenarios where data needs to be divided into smaller partitions for scalability and parallelism.