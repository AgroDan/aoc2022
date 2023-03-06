# Dan's Introduction

I am super proud of this one as well. Mostly because I didn't need to go check my algorithm against someone else. I assessed the situation, decided on either Breadth-First-Search or Depth-First-Search (Went with BFS), then just let it run. Granted this didn't churn through it terribly quickly, but at least it wasn't obscenely long. I was able to run the algorithm 3 times for a total of `30m, 49s`. Not too bad. I just kicked it off and went to lunch, came back to an answer. Cool!

The biggest challenge here wasn't _that_ difficult in terms of comprehension, but it had an incredible amount of moving parts to contend with. Immediately I thought, "oh okay this is another map I have to ingest and keep track of." And that was the plan to begin with, just keep a 2-dimensional array of the current map and its state. But towards the end of development after I started creating the pathfinding algorithm, I realized that I needed to not only keep track of all the directions I was heading in, _but also_ keep track of _all the blizzards_. When I started pushing state values onto the queue it got so complex and memory-intensive (keeping a copy of the map in memory for each choice must have used a _ton_ of memory) that I had to just...start over. I didn't go completely tabula rasa here, but I did move the `Elf.go` and `Valley.go` files into the `broken` directory so I had some way to reference them. I scrapped the idea of keeping track of the Elf as a struct object and also scrapped the idea of keeping a copy of the map, but rather kept a copy of _every single Blizzard's coordinates_, which took up a significantly smaller amount of memory. When I started referencing positions in terms of numbers and borders on the map after ingestion, things got a lot more efficient.

Just goes to show you, the Sunk Cost Fallacy is my worst enemy. Sometimes I just have to blow away the code and start again to make it more efficient, rather than jerry-rig the right way into the old, less efficient way. Turns out that in blowing away my old code and re-writing it, I discovered a bug that probably would have taken me a looooong time to discover had I improvised the solution into my old code. Y'know, I really need to acquire a rubber duck.

## Part 1

Ignoring the first part in which my code didn't work properly, my way of handling this was first ingesting the size of the map by determining its length and width. Since the map was just a big rectangle with a border (and an entrance and exit in the upper left and lower right respectively), all I really needed was 3 coordinates: The bottom right coordinate (didn't need upper left because technically upper left was always `row: 0, col: 0`), the Entrance, and Exit coordinates. Everything in the first row was a wall except for the entrance, everything in the last row was a wall except for the exit, and for every other row, the `0`th and last column were walls. So since blizzards could essentially wrap around, I had to made 2 functions that had to deal with that: `IsBorder()` and `Wrap()`. `IsBorder()` would accept a coordinate as a parameter and return a `true` or `false` answer whether or not the coordinate was a wall, and the `Wrap()` function would accept a coordinate as a paramter and return the opposite side coordinate as if it wrapped around to the other side.

I then created the `MoveOne()` function which simply incremented the map by one minute, meaning it would move every single Blizzard one block in its specific direction, and wrap around if necessary.

Additionally, I created two types of functions that obtained the surrounding coordinates. If the item in question was a `Blizzard{}`, then it would use `GetSurroundingForBlizzard()`, which would incorporate the concept of wrapping around when obtaining all 4 surrounding coordinates. It turns out I never needed that since it's a blizzard that only travels in one direction, but I was in the zone, man. And second, I created `GetSurroundingForElf()`, which would get surrounding coordinates for N, S, E and W directions, but would exclude a direction if that was a wall. 

Finally, I created the extremely important function of `GetValidDirections()`, which would poll all 4 directions from the Elf's current location INCLUDING the current coordinate the elf was in and determine if a Blizzard would occupy that space on the next move, a job performed by the `Blizzard.Peek()` function. It simply checked every single blizzard to see if the coordinate of its next movement was the same as the space. If not, that coordinate was added to the return value of `GetValidDirections()`, if so it was skipped.

Now with that out of the way, I wanted to add _one more_ method to the mix, which was the "Poor man's hash" function. I could have simply hashed every single coordinate of blizzard with something like an MD5 hash, but that was way too many calculations given the amount of calculations already being performed, so I simply created a string of every numerical coordinate of every blizzard on the map. I did this because _eventually_ every single blizzard will repeat the pattern, and the way I saw it is that if the Elf is ever in a previously visited location where all the blizzards were in the _exact same location as well_, then I can throw out that branch because I will learn nothing new from it. This was my answer to the `visited` list in breadth-first-search algorithms, since the map was always changing, I needed some way to determine a place that I can consider actually visited. 

Finally I worked on the Breadth-First-Search function, where what worked for me was finally figuring out how to make a proper deep copy of the state object. With the idea of copying a _representation_ of the map rather than the map itself I was able to properly handle memory management better and without causing a tremendous memory leak. The first process took about 13 minutes or so to complete. Pretty happy with how it turned out.

## Part 2

Part 2 was an easy addition to part 1. This time I had to create new functions for the second part which took more arguments, specifically the `start` and `end` coordinates. Since they were no longer "baked" into the direction -- that is, no longer _just_ going from entrance to exit -- I had to re-create the functions to start from an arbitrary coordinate and end at another arbitrary coordinate. In addition, I had to _also_ return the `State{}` of the map and resulting blizzards when the last execution was run so I can continue on that pattern when I would ultimately have to run the function two more times, starting from the Exit back to the Entrance, and then back to the Exit again. It was consecutive, so I had to start from where the last one left off, so the blizzards would not be reset for each iteration.

This process took about 30 minutes total to run back and forth. I was happy to say that I got it on the first try. Well, I got it on the first _successful_ run-through. There was a ton of debugging going on to obtain this.

## What I Learned

Definitely learned some weird edge cases here. One thing I did was when working with an array of `rune` datatypes, printing the specific character will print exactly that. But if I want to print the ascii value of a number, I have to do some fancy footwork. Specifically I have to take the number and add it to the value of `48`, which in ascii is the string value of the number `0`, so if I wanted to print a `2` for example (specifically in this case when I wanted to print the actual map), I would add that value to `48` to make `50`, which is the ascii value of the number `2`. Was a silly little thing I learned because the ascii value of the `2` is an unprintable character.

I also learned that when using Go's `make()` function for a slice, if I specify the size of the slice I want to make I don't need to use the `append()` function. I can simply just reference the item within the slice and assign it to whatever.

I used the `copy()` builtin function to create a proper deep copy of a structure object. Something I'm sure I will use again in the future.

If I changed this to a Depth-First-Search algorithm, the program crashes. My guess is it never gets to the end before the memory completely fills up. Not worth it I guess.

# --- Day 24: Blizzard Basin ---

With everything replanted for next year (and with elephants and monkeys to tend the grove), you and the Elves leave for the extraction point.

Partway up the mountain that shields the grove is a flat, open area that serves as the extraction point. It's a bit of a climb, but nothing the expedition can't handle.

At least, that would normally be true; now that the mountain is covered in snow, things have become more difficult than the Elves are used to.

As the expedition reaches a valley that must be traversed to reach the extraction site, you find that strong, turbulent winds are pushing small blizzards of snow and sharp ice around the valley. It's a good thing everyone packed warm clothes! To make it across safely, you'll need to find a way to avoid them.

Fortunately, it's easy to see all of this from the entrance to the valley, so you make a map of the valley and the blizzards (your puzzle input). For example:

```
#.#####
#.....#
#>....#
#.....#
#...v.#
#.....#
#####.#
```

The walls of the valley are drawn as #; everything else is ground. Clear ground - where there is currently no blizzard - is drawn as .. Otherwise, blizzards are drawn with an arrow indicating their direction of motion: up (^), down (v), left (<), or right (>).

The above map includes two blizzards, one moving right (>) and one moving down (v). In one minute, each blizzard moves one position in the direction it is pointing:

```
#.#####
#.....#
#.>...#
#.....#
#.....#
#...v.#
#####.#
```

Due to conservation of blizzard energy, as a blizzard reaches the wall of the valley, a new blizzard forms on the opposite side of the valley moving in the same direction. After another minute, the bottom downward-moving blizzard has been replaced with a new downward-moving blizzard at the top of the valley instead:

```
#.#####
#...v.#
#..>..#
#.....#
#.....#
#.....#
#####.#
```

Because blizzards are made of tiny snowflakes, they pass right through each other. After another minute, both blizzards temporarily occupy the same position, marked 2:

```
#.#####
#.....#
#...2.#
#.....#
#.....#
#.....#
#####.#
```

After another minute, the situation resolves itself, giving each blizzard back its personal space:

```
#.#####
#.....#
#....>#
#...v.#
#.....#
#.....#
#####.#
```

Finally, after yet another minute, the rightward-facing blizzard on the right is replaced with a new one on the left facing the same direction:

```
#.#####
#.....#
#>....#
#.....#
#...v.#
#.....#
#####.#
```

This process repeats at least as long as you are observing it, but probably forever.

Here is a more complex example:

```
#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#
```

Your expedition begins in the only non-wall position in the top row and needs to reach the only non-wall position in the bottom row. On each minute, you can move up, down, left, or right, or you can wait in place. You and the blizzards act simultaneously, and you cannot share a position with a blizzard.

In the above example, the fastest way to reach your goal requires 18 steps. Drawing the position of the expedition as E, one way to achieve this is:

```
Initial state:
#E######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#

Minute 1, move down:
#.######
#E>3.<.#
#<..<<.#
#>2.22.#
#>v..^<#
######.#

Minute 2, move down:
#.######
#.2>2..#
#E^22^<#
#.>2.^>#
#.>..<.#
######.#

Minute 3, wait:
#.######
#<^<22.#
#E2<.2.#
#><2>..#
#..><..#
######.#

Minute 4, move up:
#.######
#E<..22#
#<<.<..#
#<2.>>.#
#.^22^.#
######.#

Minute 5, move right:
#.######
#2Ev.<>#
#<.<..<#
#.^>^22#
#.2..2.#
######.#

Minute 6, move right:
#.######
#>2E<.<#
#.2v^2<#
#>..>2>#
#<....>#
######.#

Minute 7, move down:
#.######
#.22^2.#
#<vE<2.#
#>>v<>.#
#>....<#
######.#

Minute 8, move left:
#.######
#.<>2^.#
#.E<<.<#
#.22..>#
#.2v^2.#
######.#

Minute 9, move up:
#.######
#<E2>>.#
#.<<.<.#
#>2>2^.#
#.v><^.#
######.#

Minute 10, move right:
#.######
#.2E.>2#
#<2v2^.#
#<>.>2.#
#..<>..#
######.#

Minute 11, wait:
#.######
#2^E^2>#
#<v<.^<#
#..2.>2#
#.<..>.#
######.#

Minute 12, move down:
#.######
#>>.<^<#
#.<E.<<#
#>v.><>#
#<^v^^>#
######.#

Minute 13, move down:
#.######
#.>3.<.#
#<..<<.#
#>2E22.#
#>v..^<#
######.#

Minute 14, move right:
#.######
#.2>2..#
#.^22^<#
#.>2E^>#
#.>..<.#
######.#

Minute 15, move right:
#.######
#<^<22.#
#.2<.2.#
#><2>E.#
#..><..#
######.#

Minute 16, move right:
#.######
#.<..22#
#<<.<..#
#<2.>>E#
#.^22^.#
######.#

Minute 17, move down:
#.######
#2.v.<>#
#<.<..<#
#.^>^22#
#.2..2E#
######.#

Minute 18, move down:
#.######
#>2.<.<#
#.2v^2<#
#>..>2>#
#<....>#
######E#
```

What is the fewest number of minutes required to avoid the blizzards and reach the goal?

Your puzzle answer was `297`.

# --- Part Two ---

As the expedition reaches the far side of the valley, one of the Elves looks especially dismayed:

He forgot his snacks at the entrance to the valley!

Since you're so good at dodging blizzards, the Elves humbly request that you go back for his snacks. From the same initial conditions, how quickly can you make it from the start to the goal, then back to the start, then back to the goal?

In the above example, the first trip to the goal takes 18 minutes, the trip back to the start takes 23 minutes, and the trip back to the goal again takes 13 minutes, for a total time of 54 minutes.

What is the fewest number of minutes required to reach the goal, go back to the start, then reach the goal again?

Your puzzle answer was `856`.

Both parts of this puzzle are complete! They provide two gold stars: **

At this point, you should return to your Advent calendar and try another puzzle.

If you still want to see it, you can get your puzzle input.