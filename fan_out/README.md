## The "Fan-Out"

The "Fan-Out" pattern is a concurrency pattern used to distribute data from a single source to multiple worker goroutines. It involves creating a set of destination channels, launching worker goroutines, and distributing data concurrently to each worker. This pattern is useful for parallelizing tasks and improving processing efficiency in Go programs.

**Example 1: "Fan-Out" Pattern**
- Efficiently distributes data from one source channel to multiple worker goroutines.
- Creates destination channels for each worker.
- Demonstrates concurrent data processing.
- Utilizes a WaitGroup to synchronize worker completion.

**Example 2: "Fan-Out" Pattern with Context Cancellation**
- Extends "Fan-Out" with graceful termination using context cancellation.
- Worker goroutines respond to context signals for clean shutdown.
- Simulates work, cancels context, and waits for worker completion.

**Example 3: "Fan-Out" Pattern with `sync.WaitGroup`**
- Efficiently distributes data from a source channel to multiple worker goroutines.
- Uses a WaitGroup for worker synchronization.
- Demonstrates parallelized data processing.

These examples illustrate how to efficiently parallelize work by distributing data from a source to multiple worker goroutines.