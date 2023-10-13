## Example 1:

This example demonstrates the "Sharding" pattern, which efficiently distributes data among multiple shards, enabling parallel access and data management. Here's a detailed breakdown of how it works:

1. **Shard Structure**: The code introduces a `Shard` struct, which represents an individual shard. Each shard is equipped with read-write mutexes (`sync.RWMutex`) to ensure safe data access and a map (`map`) for data storage. Shards are used to partition data.

2. **ShardedMap**: The `ShardedMap` is essentially a slice of shards, with each shard having its mutex for secure data handling. In this example, the `ShardedMap` is parameterized with the data type `V`.

3. **NewShardedMap (Creating a Sharded Map)**: This function is responsible for creating and initializing a sharded map with a specified number of shards (`nShards`). For each shard, a map is created to store data.

4. **Get (Retrieving Data)**: The `Get` method is used to retrieve a value associated with a key from the sharded map. It selects the appropriate shard, locks it for reading (`RLock`), and returns the value. The mutex is released after reading.

5. **Set (Setting Data)**: The `Set` method is used to set a value for a given key in the sharded map. It selects the corresponding shard, locks it for writing (`Lock`), sets the value, and releases the mutex.

6. **Keys (Retrieving Keys)**: The `Keys` method gathers all keys from all shards. For safety, an additional mutex (`mu`) and a `sync.WaitGroup` (`wg`) are used to wait for the completion of key reading from all shards.

7. **shardIndex (Calculating Shard Index)**: This method calculates the shard index for a given key using SHA1 hashing.

8. **shard (Getting a Shard)**: The `shard` method returns the shard associated with the given key. It uses `shardIndex` to determine the shard's index.

9. **main (Main Function)**: In the `main` function, a sharded map with 5 shards is created. Then, three key-value pairs are added to it. The `Get` method is used to retrieve data for the keys "alpha," "beta," and "gamma." Finally, the keys available in the shards are printed using the `Keys` method.

The "Sharding" pattern in this code effectively manages data, distributes it among multiple shards, and provides parallel access to data in a multi-threaded environment. It's a valuable approach in scenarios where data needs to be divided into smaller partitions for scalability and parallel processing.