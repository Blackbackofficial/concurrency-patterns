## Example 2: Coordinated HTTP Requests

In this second act, the "Done Channel Pattern" takes on the role of conductor, skillfully coordinating and awaiting the completion of multiple HTTP requests.

Here's how the pattern orchestrates this performance:

1. A `done` channel, meticulously crafted as `chan RequestResult`, takes center stage. Its buffer size is set equal to the number of URLs within the `urls` slice. This channel is poised to receive the results of the HTTP requests as they conclude.

2. As the overture begins, a loop elegantly glides through the `urls` slice, gracefully launching a new goroutine for each URL. These goroutines gracefully step forward with `go processRequest(url, done)`, each bearing the solemn duty of dispatching the HTTP request's result to the `done` channel upon its successful conclusion.

3. The main protagonist, our central goroutine, poised for the denouement, engages in a captivating dance. This dance, a loop that matches the number of URLs, performs a gentle reception `<-done` to collect the HTTP request results. As the results gracefully arrive, they are processed by the main protagonist in the order they are received.

4. For each received result, our code gracefully inspects whether the HTTP request was a triumph (`result.Success`). In the event of a resounding victory, it proclaims a message, a splendid testament to the success alongside the URL. However, if misfortune strikes and the request falters, a sorrowful elegy is performed, casting an error message with the URL in mourning, coupled with a poignant description of the error's nature.

5. With the final note played and the curtain falling, the "done" channel gracefully exits the stage, and the performance concludes. This is achieved through the command of `close(done)`.

The "Done Channel Pattern" in this production not only allows for the graceful and parallel execution of multiple HTTP requests but also provides a method for receiving, processing, and responding to the results as they unfold, culminating in a performance that harmonizes concurrent tasks with poise and efficiency.