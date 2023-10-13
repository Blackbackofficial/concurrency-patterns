## Example 2:

The "Sharding" pattern is applied to create a sharded storage for distributed data management. Here's how it works:

1. **Shard**: This struct represents an individual shard, which is essentially a partition for storing data. Each shard is equipped with a read-write mutex (`sync.RWMutex`) to ensure thread-safe data access and a map (`map`) to store key-value pairs.

2. **ShardedStorage**: It's a collection of shards used for distributed data storage.

3. **NewShardedStorage (Creating a Sharded Storage)**: This function creates a new sharded storage with a specified number of shards (`numShards`). For each shard, a map is created for data storage, and the shards are stored within the `ShardedStorage`.

4. **Set (Adding Data)**: The `Set` method allows you to add key-value pairs to the sharded storage. It calculates the appropriate shard based on the key using the `shardIndex` function, locks the shard for writing, sets the value, and releases the lock.

5. **Get (Retrieving Data)**: The `Get` method retrieves the value associated with a key from the sharded storage. Similar to the `Set` method, it calculates the corresponding shard, locks it for reading, and retrieves the value while maintaining thread safety.

6. **shardIndex (Calculating Shard Index)**: This function calculates the index of the shard for a given key. In this example, it employs a simple hashing function for demonstration purposes, but in a real-world scenario, a more sophisticated hashing function should be used.

7. **hashString (Simple Hash Function)**: This function provides a simple hash calculation based on the characters in the input string.

8. **Main Function**: In the `main` function, a sharded storage with five shards is created. Key-value pairs are added to the sharded storage, and then, the data is retrieved and printed. The `shardIndex` function helps in determining which shard to use for each key, ensuring that the data is distributed across the shards.

The "Sharding" pattern here allows for the efficient distribution and management of data among multiple shards, ensuring that data can be accessed and modified in a parallel and thread-safe manner.