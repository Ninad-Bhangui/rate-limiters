# Rate limiters

## Token Bucket Algorithm
In Token bucket algorithm, we decide a bucketSize which contains number of tokens allowed. Tokens keeps filling tokens at a fixed rate. Everytime a request is sent to the rate limiter, it revokes tokens. If the bucket is empty, the rate limiter should reject the request.
I tried this in multiple iterations:

1. Based on refill rate I kept incrementing the counter every second in a seperate go routine.
2. After looking at other solutions online, I realised I can simply keep track of lastRefillTime, and everytime there's a request, I calculate how many tokens to fill based on `current time - lastRefillTime`. This is more performant than doing something every second.
3. Next step: Try integrating this with an actual API
4. Next Next step: This should work concurrently (use some mutex when incrementing/decrementing counter)
