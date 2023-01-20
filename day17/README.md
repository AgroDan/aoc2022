# Dan's Introduction

Oh. My. God. This one kicked the ever-loving crap out of me. From getting the first part done to Part 2 being an exercise in math and pattern recognition...I was just mentally relieved in getting the gold stars for this one. I'm sure this is another case of over-engineering the whole thing, but I'm happy with how it turned out regardless. I even tried to optimize it a bit as much as I could given the amount of processing and memory this allocates, so this was a nice win for me I think in terms of learning Go. Here's how I figured everything out and what I learned along the way.

## Part 1

I approached it by making pretty much everything an object that can be worked on. I decided to start with the `Cavern` object which was the 7-object-wide cavern that all the rocks fell through, and it had a definitive "ground" as well as an an infinitely-growing ceiling that would appear as needed to help allocate properly. This would come in handy whenever a new block spawned.

Each Rock was essentially two kinds of objects. A Rock object which held the container to build a rock from including all of the "offsets" of each rock extension, which would be used to create each shape. So for example, to make a "plus sign" shaped rock, the offsets were:

```
(X: 0, Y: 0) <-- bottom piece
(X: 0, Y: 1) <-- center piece
(X: 1, Y: 1) <-- right-most piece
(X: -1, Y: 1) <-- left-most piece
(X: 0, Y: 2) <-- top piece

Looks like:
 #
###
 #
```

Considering these are offsets to create the shape, it was easy to build from these values. Next, there was the `TransitionRock` object, which took in the `Rock` object shape to determine how to _move_ the rock. In a sense, this `TransitionalRock` object was in charge of where the rock physically was inside the Cavern as it was falling. It would apply the proper calculations to each rock for each direction of movement it would go. If it was moved right, it would add `1` to each `X` value in the transitional rock. To move left, subtract `1` from the `X` value. Similarly, if it went down, it would subtract `1` from each `Y` value.

Before it would move though, it would check to see if any part of the rock would go beyond a wall limit or touch another rock. If it checked out that the rock was capable of moving in that direction without overlapping onto another object, it would update the values in the `TransitionalRock` object and cycle again. It would continue this cycle of traversing the `jets` loop until it could _not_ move downwards, meaning it would come to a rest. At this point, it would write `#`'s in the Cavern object where the rock came to rest and continue on. Based on this, I was able to determine the height of the tower by simply counting upwards and checking for the first line that _didn't_ contain any rock objects, and this would be the height. There might be a better way to handle that, but whatever, it worked. The challenge was to see how high the rocks will stack after `2022` rocks fall.

## Part 2

Oh boy, this was the ridiculous part. The challenge was to determine how tall the rocks will stack to if a bazillion rocks fall. Obviously not a bazillion, but basically a number so high that it would take an unreasonable amount of time to calculate as was done in Part 1. Specifically, the number is `1_000_000_000_000` (one trillion). To accomplish this, you had to look for a shortcut. And this shortcut can be applied to similar cryptographic algorithms as well...since all you're looking for is a pattern. If you can find the pattern, you can extrapolate. And to accomplish this, I needed to add a few additional aspects of my code to look for them. Enter the `Metrics.go` file, which contains all the functions I used to look for those patterns.

So the basic premise was that the jets pattern is cyclical. Once you reach the end of the jets list, you start over from the beginning. Just like the rocks, once the square rock drops, you cycle back to the horizontal line block in order. So _eventually_ you will cycle through the jets and rocks pattern to the point where they will repeat. Based on that, considering the "keyspace" is limited (a 7-unit-wide cavern, 5 types of rocks and a finite limit to the jets pattern), then eventually given a large enough cycle the pattern should repeat. So if we can find out _when_ that pattern repeats and how high it was, we can multiply enough times to get the height at one trillion rocks.

Most of this was done by hand, but obtaining the metrics helped determine the pattern. The idea was thus: Generate a large-enough rock pile and drop something like `10_000` rocks. At every single rock that is dropped, create a map object consisting of the `RockType` that was dropped and the `JetIndex` of what jet we are in. Of that hash, save two values: how many rocks have dropped so far, and how high the rock formation is. Once you cycle through all the rocks, check back at the map you created of everything and obtain the deltas of each value. Every time that `JetIndex` and `RockType` combination fell, how much _higher_ was the rock formation, and how many _more_ rocks have fallen?

From those values, the ones I care about are the deltas that are **the same**. Because if I can determine that the delta of rocks fallen at the same `JetIndex` and `RockType` is 20 and 30 (for example), and the last time they overlapped was at 250 and 400 respectively, I know that after 250 rocks fall, _every 20 rocks that fall afterwards will add 30 to the height_.

So based on this, here is what I deteremined.

```
Based on my input and performing a Metrics Calculation, I can determine:

Rocks Fallen:  8622 with a Delta of 1730
      Height: 13175 with a Delta of 2647

So therefore I can predict that after (8622 + 1730 = 10_352) rocks fall, the height will
be (13_175 + 2647 = 15_822). I can verify that easily with this script.

Now, to get the one trillion height, it requires a little math.

First, subtract the amount of rocks fallen from one trillion:
    1_000_000_000_000
   -            8_622
   =  999_999_991_378

Now divide that by the Rock Delta, but only the floor, don't care about the remainder:
  999_999_9991_378 / 1730 = 578_034_677 (and change, but ignore the decimal)

We can get how close we are (and how much more we need to find out) by multiplying that
number by the Rock Delta again:

  578_034_677 * 1730 = 999_999_991_210

Subtract that from the amount of rocks fallen earlier to find out how much higher we
need to calculate later:

  999_999_991_378
 -999_999_991_210
 =            168

Great, hold onto that for later. Now we can find out how high the rock formation will
be at a height of 999_999_999_832:

  (578_034_677 * 2647) + 13175 = 1_530_057_803_194

Almost there...now I need to just find out how much higher the rock formation will grow
after 168 more rocks fall. Since I know how high it will be after 8_622 rocks fall, I
can just run the simulation again and see how high it will be after (8_622 + 168) rocks
fall.

<Run the simulation to drop 8790 rocks>
Height: 13434

Subtract that from the known height of 13_175...and we get 259. That is how much higher
the formation will be after 168 more rocks fall from 8622!

So therefore:
  1_530_057_803_194
 +              259
 =1_530_057_803_453

That's the answer!
```

Another "nearly jumped out of my chair" moment when I got this. My poor cat. (Don't worry, he's fine. I didn't actually jump out of my chair)

## What I learned

In this one I learned how to set arguments which cleans up the code quite a bit. Now I can print the whole thing if I want to, or leave it out completely. It's actually pretty easy and I would argue even easier than python to work with command line flags. In addition to that, I also worked a little bit with variadic functions, though really I think I worked on it more on [Day 16](../day16/). Either way this was a pretty good learning experience.

There's also a time where I thought I would have to look for rows that are actually the same to find the patterns there. To accomplish that, I used kind of a "poor man's hash" algorithm, which honestly wasn't really a poor man's hash as it was a "super fast hash" by simply recording the binary representation of each or multiple rows. If a rock pattern existed on a row that looked like this:

```
|.##.#..|
```

Then the binary representation would be `0b0110100`, where each column is represented in the number. I could then combine multiple rows like this and use it to find a pattern relatively quickly. Ultimately I scrapped that idea when I went for the "same jet index/rock index" technique as listed above, but I kept the code here because I thought it was pretty cool and I can definitely use this technique elsewhere. Hashing functions are typically fast enough, but y'know what's faster? binary math. Don't need to cycle a bunch of times through a SHA256 algorithm per row when I can just flip a couple bits to accomplish the same thing.

# --- Day 17: Pyroclastic Flow ---

Your handheld device has located an alternative exit from the cave for you and the elephants. The ground is rumbling almost continuously now, but the strange valves bought you some time. It's definitely getting warmer in here, though.

The tunnels eventually open into a very tall, narrow chamber. Large, oddly-shaped rocks are falling into the chamber from above, presumably due to all the rumbling. If you can't work out where the rocks will fall next, you might be crushed!

The five types of rocks have the following peculiar shapes, where # is rock and . is empty space:

```
####

.#.
###
.#.

..#
..#
###

#
#
#
#

##
##
```

The rocks fall in the order shown above: first the - shape, then the + shape, and so on. Once the end of the list is reached, the same order repeats: the - shape falls first, sixth, 11th, 16th, etc.

The rocks don't spin, but they do get pushed around by jets of hot gas coming out of the walls themselves. A quick scan reveals the effect the jets of hot gas will have on the rocks as they fall (your puzzle input).

For example, suppose this was the jet pattern in your cave:

```
>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>
```

In jet patterns, < means a push to the left, while > means a push to the right. The pattern above means that the jets will push a falling rock right, then right, then right, then left, then left, then right, and so on. If the end of the list is reached, it repeats.

The tall, vertical chamber is exactly seven units wide. Each rock appears so that its left edge is two units away from the left wall and its bottom edge is three units above the highest rock in the room (or the floor, if there isn't one).

After a rock appears, it alternates between being pushed by a jet of hot gas one unit (in the direction indicated by the next symbol in the jet pattern) and then falling one unit down. If any movement would cause any part of the rock to move into the walls, floor, or a stopped rock, the movement instead does not occur. If a downward movement would have caused a falling rock to move into the floor or an already-fallen rock, the falling rock stops where it is (having landed on something) and a new rock immediately begins falling.

Drawing falling rocks with @ and stopped rocks with #, the jet pattern in the example above manifests as follows:

```
The first rock begins falling:
|..@@@@.|
|.......|
|.......|
|.......|
+-------+

Jet of gas pushes rock right:
|...@@@@|
|.......|
|.......|
|.......|
+-------+

Rock falls 1 unit:
|...@@@@|
|.......|
|.......|
+-------+

Jet of gas pushes rock right, but nothing happens:
|...@@@@|
|.......|
|.......|
+-------+

Rock falls 1 unit:
|...@@@@|
|.......|
+-------+

Jet of gas pushes rock right, but nothing happens:
|...@@@@|
|.......|
+-------+

Rock falls 1 unit:
|...@@@@|
+-------+

Jet of gas pushes rock left:
|..@@@@.|
+-------+

Rock falls 1 unit, causing it to come to rest:
|..####.|
+-------+

A new rock begins falling:
|...@...|
|..@@@..|
|...@...|
|.......|
|.......|
|.......|
|..####.|
+-------+

Jet of gas pushes rock left:
|..@....|
|.@@@...|
|..@....|
|.......|
|.......|
|.......|
|..####.|
+-------+

Rock falls 1 unit:
|..@....|
|.@@@...|
|..@....|
|.......|
|.......|
|..####.|
+-------+

Jet of gas pushes rock right:
|...@...|
|..@@@..|
|...@...|
|.......|
|.......|
|..####.|
+-------+

Rock falls 1 unit:
|...@...|
|..@@@..|
|...@...|
|.......|
|..####.|
+-------+

Jet of gas pushes rock left:
|..@....|
|.@@@...|
|..@....|
|.......|
|..####.|
+-------+

Rock falls 1 unit:
|..@....|
|.@@@...|
|..@....|
|..####.|
+-------+

Jet of gas pushes rock right:
|...@...|
|..@@@..|
|...@...|
|..####.|
+-------+

Rock falls 1 unit, causing it to come to rest:
|...#...|
|..###..|
|...#...|
|..####.|
+-------+

A new rock begins falling:
|....@..|
|....@..|
|..@@@..|
|.......|
|.......|
|.......|
|...#...|
|..###..|
|...#...|
|..####.|
+-------+
```

The moment each of the next few rocks begins falling, you would see this:

```
|..@....|
|..@....|
|..@....|
|..@....|
|.......|
|.......|
|.......|
|..#....|
|..#....|
|####...|
|..###..|
|...#...|
|..####.|
+-------+

|..@@...|
|..@@...|
|.......|
|.......|
|.......|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

|..@@@@.|
|.......|
|.......|
|.......|
|....##.|
|....##.|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

|...@...|
|..@@@..|
|...@...|
|.......|
|.......|
|.......|
|.####..|
|....##.|
|....##.|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

|....@..|
|....@..|
|..@@@..|
|.......|
|.......|
|.......|
|..#....|
|.###...|
|..#....|
|.####..|
|....##.|
|....##.|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

|..@....|
|..@....|
|..@....|
|..@....|
|.......|
|.......|
|.......|
|.....#.|
|.....#.|
|..####.|
|.###...|
|..#....|
|.####..|
|....##.|
|....##.|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

|..@@...|
|..@@...|
|.......|
|.......|
|.......|
|....#..|
|....#..|
|....##.|
|....##.|
|..####.|
|.###...|
|..#....|
|.####..|
|....##.|
|....##.|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+

|..@@@@.|
|.......|
|.......|
|.......|
|....#..|
|....#..|
|....##.|
|##..##.|
|######.|
|.###...|
|..#....|
|.####..|
|....##.|
|....##.|
|....#..|
|..#.#..|
|..#.#..|
|#####..|
|..###..|
|...#...|
|..####.|
+-------+
```

To prove to the elephants your simulation is accurate, they want to know how tall the tower will get after 2022 rocks have stopped (but before the 2023rd rock begins falling). In this example, the tower of rocks will be 3068 units tall.

How many units tall will the tower of rocks be after 2022 rocks have stopped falling?

Your puzzle answer was `3090`.

# --- Part Two ---

The elephants are not impressed by your simulation. They demand to know how tall the tower will be after 1000000000000 rocks have stopped! Only then will they feel confident enough to proceed through the cave.

In the example above, the tower would be 1514285714288 units tall!

How tall will the tower be after 1000000000000 rocks have stopped?

Your puzzle answer was `1530057803453`.

Both parts of this puzzle are complete! They provide two gold stars: **

At this point, you should return to your Advent calendar and try another puzzle.

If you still want to see it, you can get your puzzle input.