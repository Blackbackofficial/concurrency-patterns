## Example 2:

This example demonstrates the "Sharding" pattern applied to create a sharded storage system for distributed data management. Here's a detailed breakdown of how it works:

1. **Shard Structure**: The code introduces a `Shard` struct, which represents an individual shard or partition for storing data. Each shard is equipped with a `sync.RWMutex` to ensure thread-safe data access and a `map` to store key-value pairs.

2. **ShardedStorage**: The `ShardedStorage` struct is designed as a collection of these shards, intended for distributed data storage.

3. **NewShardedStorage (Creating a Sharded Storage)**: This function is responsible for creating a new sharded storage system. It takes as input the desired number of shards (`numShards`). For each shard, a map is created to store data, and the shards are organized within the `ShardedStorage`.

4. **Set (Adding Data)**: The `Set` method allows users to add key-value pairs to the sharded storage. It calculates the appropriate shard for the data based on the key by utilizing the `shardIndex` function. It then locks the corresponding shard for writing, sets the value, and releases the lock, ensuring thread safety.

5. **Get (Retrieving Data)**: The `Get` method is used to retrieve the value associated with a key from the sharded storage. Similar to the `Set` method, it calculates the appropriate shard, locks it for reading using `RLock`, retrieves the value, and releases the lock, all while maintaining thread safety.

6. **shardIndex (Calculating Shard Index)**: This function calculates the index of the shard to be used for a given key. In this example, a simple hashing function is employed for demonstration purposes. In a production environment, a more sophisticated hashing method would be advisable.

7. **hashString (Simple Hash Function)**: The `hashString` function provides a basic hashing calculation based on the characters within the input string.

8. **Main Function**: In the `main` function, a sharded storage system with five shards is created. Key-value pairs are added to this sharded storage. Subsequently, the data is retrieved and printed. The `shardIndex` function plays a crucial role in determining which shard is used for each key, ensuring that the data is efficiently distributed across the shards.

The "Sharding" pattern, as illustrated in this code, enables the effective distribution and management of data across multiple shards. It ensures that data can be accessed and modified in a parallel and thread-safe manner, making it particularly valuable in scenarios where data needs to be divided into smaller partitions for improved scalability and parallel processing.