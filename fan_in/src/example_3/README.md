## Example 3: Uniting Streams

In this exquisite showcase of the "Fan-In" pattern, data from multiple channels gracefully waltz together into a single harmonious channel. Let's explore the performance step by step:

1. **The Master of Ceremonies**: The `mergeData` function takes the stage, and with it, a `done` channel and a collection of data channels (`datas...`). The `done` channel serves as the conductor's baton, orchestrating the termination of this merging symphony, while the `datas` channels are the individual instrumentalists contributing their data.

2. **The Symphony Orchestra**: Within the grand concert hall of the `mergeData` function, a `sync.WaitGroup` (`wg`) stands as the conductor, keeping time for the orchestra of goroutines. The `merged` channel is the grand stage where the melodies from all instrumentalists are harmoniously united.

3. **Virtuosos in Action**: Each data channel (`datas...`) has its own virtuoso goroutine. These virtuosos deftly select notes from their respective channels and perform a beautiful concerto by sending them to the `merged` channel. They also keep a keen ear for signals from the `done` channel, ready to exit when the conductor's baton is lowered.

4. **The Grand Finale**: The concert hall bursts into applause as the `merged` channel is closed, signaling the end of the performance. A separate backstage conductor waits for the orchestra to conclude their virtuoso acts by listening for the echoes of `wg.Done()`. Once every musician has left the stage, the grand `merged` channel is officially closed (`close(merged)`).

5. **The Audience's Applause**: The audience, represented by the `main` function, arrives with a `done` channel in hand, ready to signal the final ovation. This channel, used to indicate the termination of the merging process, is gracefully deferred for closure.

6. **Instrumentalists in Harmony**: Two instrumentalists (`data1` and `data2`) step onto the stage and perform solos with integer values, sharing their melodies with the audience.

7. **A Symphony Begins**: The `mergeData` function is called to bring the instrumentalists together. It is instructed to follow the lead of the `done` channel and the data channels (`data1` and `data2`). The result is a harmonious composition in the `merged` channel, and the performance begins.

8. **A Captivated Audience**: The `main` function, embodying the audience, listens to the combined melodies from the `merged` channel, printing and savoring each note in a loop. The performance continues until the last note has been played and the final curtain falls.

In summary, this code exemplifies the "Fan-In" pattern, uniting data from various channels into one, allowing for efficient asynchronous data processing. The use of the `done` channel ensures a graceful conclusion to the symphony of data, creating a harmonious experience for all observers.