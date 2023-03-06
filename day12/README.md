# NOTICE - this isn't the answer

This challenge messed me up so much that I had to re-do the entire thing. Only after starting over from scratch (see the `day12_2` directory) did I notice what I was doing wrong: **I was considering the starting position to be one lower than the lowest elevation square.** This caused my score to be off by 2, always. I did the same thing on the second iteration of this code and damn near wanted to give up.

Regardless, however, I am keeping this code for two reasons. One, I learned a great deal about proper structuring in Go. There were a lot of neato tricks I learned from this challenge that I want to keep around so I can refer back to. Specifically Go's weird way of dealing with `enum` data types. Which is to say, *not really that intuitive*. Still though, it did the job so no complaints other than the aforementioned I suppose.

Two, this is the saddest/craziest/most extreme example I can give of **COMPLETELY** over-engineering an end goal. This was originally planned to work with a random number generator, where this code would simply try every single path it could by randomly making decisions until it eventually made it to the end enough times, and it would simply print the smallest amount of steps the explorer would take over the last time it printed. This surprisingly actually worked with the test input with much smaller data to work with. It was incredibly accurate and did the job well. But the actual challenge input would be exponentially longer and in fact never actually made it to the end in a reasonable amount of time. I eventually scrapped this technique.

Eventually I discovered [Dijkstra's Algorithm](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm) that explained the most efficient way of handling this. I then jerry-rigged the original "random choice" code to work with a queue using the aforementioned algorithm, and no matter what always came up off by two in my puzzle input, even though it would work with the test input. What's worse, it even worked with *other people's puzzle input* but not mine! This took an extremely long time to troubleshoot until eventually I just re-read the puzzle instructions and discovered the small (yet major) discrepency.

I am leaving my failures here for all to see so that you can mock me as penance for my own stupidity.

# Future Dan's Edit

This is not Dijkstra's Algorithm. This is in fact a Breadth-First-Search algorithm. My fault. I apparently never used Dijkstra's algorithm. Not sure why I thought I did. Most likely confused with reviewing things on youtube.