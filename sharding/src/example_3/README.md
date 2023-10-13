## Example 3:

In this example, the "Sharding" pattern is applied to efficiently distribute and manage data among multiple shards. Here's a detailed explanation of how it works:

1. **Data Structure**: The `Data` struct is used to represent key-value pairs, where each pair consists of a `Key` (a string) and a `Value` (an integer).

2. **Main Function**:
   - An array of `Data` elements is created, representing the data that needs to be distributed across multiple shards.
   - A slice of maps, `shards`, is created to simulate the shards. In this example, three shards are used.
   - A `sync.WaitGroup` (`wg`) is employed to ensure that all goroutines complete their work before the program proceeds.

3. **Goroutines for Data Distribution**:
   - For each `Data` element in the `data` slice, a goroutine is launched. Each goroutine's responsibility is to distribute a piece of data to one of the three shards.
   - The `hashCode` function is used to calculate the shard index for a given key. In this demonstration, a simple hash function is used, but in real-world applications, a more complex hashing method would be preferable. The shard index is determined by the result of `hashCode(d.Key) % 3`, ensuring even distribution among the three shards.
   - Subsequently, the goroutine adds the data to the appropriate shard.

4. **Waiting for Goroutines**: In the `main` function, `wg.Wait()` is utilized to ensure that all goroutines finish their work before proceeding.

5. **Outputting Data from Shards**:
   - After all data is successfully distributed, the `main` function proceeds to output data from all shards.
   - It systematically iterates through the `shards` slice, printing both the keys and values stored within each shard.

The "Sharding" pattern demonstrated in this code effectively showcases the distribution and management of data among multiple shards, allowing for parallel processing and efficient data organization. It proves to be particularly valuable in scenarios where data needs to be divided into smaller partitions to enhance scalability and support parallelism.