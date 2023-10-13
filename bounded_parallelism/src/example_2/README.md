## Example 2:

Web-scraping to collect data from multiple websites.

Suppose you have a task to collect information from several websites. You want to scan each site in parallel to speed up data collection, but you don't want to overload network resources or site servers. In this case, limiting concurrency can be useful.

You could use the "Limited Concurrency" pattern to limit the number of concurrent HTTP requests to sites. For example, by setting a limit of 5 concurrent requests, you ensure that your scraper does not send more than 5 requests at the same time. This will allow:

Collect data from multiple sites at the same time, speeding up the process.
Prevent website servers from overloading, as you don't send too many requests at the same time.
Use resources efficiently, since too many parallel goroutines are not created, which can cause excessive memory consumption.