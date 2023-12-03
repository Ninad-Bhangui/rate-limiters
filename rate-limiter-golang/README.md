# Rate limiters

## Token Bucket Algorithm
In Token bucket algorithm, we decide a bucketSize which contains number of tokens allowed. Tokens keeps filling tokens at a fixed rate. Everytime a request is sent to the rate limiter, it revokes tokens. If the bucket is empty, the rate limiter should reject the request.
I tried this in multiple iterations:

1. Based on refill rate I kept incrementing the counter every second in a seperate go routine.
2. After looking at other solutions online, I realised I can simply keep track of lastRefillTime, and everytime there's a request, I calculate how many tokens to fill based on `current time - lastRefillTime`. This is more performant than doing something every second.
3. I then integrated this with an actual API by creating middleware
4. Next step: This should work concurrently (use some mutex when incrementing/decrementing counter)

## Leaky Bucket Algorithm
This algorithm has two implementations.
1. Leaky Bucket as a meter
2. Leaky Bucket as a FIFO queue

The meter implementation is exactly similar in practice to Token bucket but just described differently (https://en.wikipedia.org/wiki/Leaky_bucket#As_a_meter). Instead of checking if bucket is full, you check if it's empty and instead of the bucket being filled at a fixed rate, here it leaks at a fixed rate.

The queue implementation, while different, I'm not sure if it's ideal for REST API rate limiting. If a request is queued, the client would simply wait for a few seconds till it's dequeued. Does not fit my mental model. (Need to read more)
