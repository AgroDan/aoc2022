# Dan's Introduction

Oh my poor brain. This was BY FAR the hardest challenge yet. Sadly I had to hard-code some of the deduction, but honestly even with the hard-coded data this challenge kept me from many hours of sleep. It took me a very long time to get right and to be perfectly honest there came a time where I considered very strongly about hanging up the towel. But I slept on it, had a hearty breakfast and a strong cup of coffee, rolled up my sleeves and went back to finding out my issue. I HATE the debugging process when your code works perfectly fine against the example data, but the challenge data is off. And a challenge like this is painfully unforgiving. If you get one thing wrong your answer will be wildly off by a huge degree. And that exact scenario plagued me for days until I had a Eureka moment. I cheered loudly when I got the message saying that I got the right answer. It was like climbing Mount Everest.

## Part 1

Part 1 was fairly easy. It involved a lot of steps sure, but building the foundation wasn't all that challenging. Pretty standard, there are 3 tiles to concern yourself with. A blank space (which for some reason I converted to `+` signs, and I can't remember why I bothered but when I realized how dumb that was I was too deep into the code that I didn't bother changing it back), an octothorpe `#` which symbolized a wall, and a `.` which symbolized a position on the map. The blank spaces meant presumably the ether, and by entering one blank space you would re-appear to the opposing side's first `.` character. You had to track the direction you were facing, and you were given a string of characters that looked something like: `10R5L5R10L4R5L5`, which had to be parsed. Numbers indicated how many steps forward your character had to take, and `R` and `L` meant to turn your character 90 degrees right and left, respectively. My logic was to parse the map, parse the directions, and create a `Monkey{}` object which symbolized my character. The `Monkey{}` object held the current X/Y coordinates of the monkey, as well as the direction the monkey was facing, symbolized as a numeric value using Go's `iota` command, to create a kind of `enum` datatype. I placed the character on the board and let 'er rip. Then at the end I calculated the final passcode using the X and Y coordinates (remembering to add 1 to both, because the makers of the advent of code count starting at `1` like some sort of weirdo) and the direction I was facing. Got this one on the first try.

## Part 2

This...is where it got...relentless. For a lack of a better word. Turns out that the map I was given can actually be folded into a cube. The actual map consisted of a flattened cube in which each side was 50x50. I just needed to figure out what each side of the cube was. THEN, the wrap-around rules became much different. Since it was now a cube, if your character goes over the edge, it pops up on the adjacent side of the cube instead of the opposing side of the flattened map. It was now up to me to determine what side is adjacent to what, AS WELL AS the rules for what direction I would be facing now on the flattened map AND the actual mapping of where I would end up on that side of the cube! This was a freakin' nightmare to imagine how to accomplish, but after days of getting involved with it I somehow managed. I'll try to explain...

This was the part of the example where I had to work with hard-coded data. I've seen other peoples' examples of using DFS to get each side of the cube automatically and determine the neighbors as well, but this would have set me back a LOT longer trying to figure out, and frankly I just wanted to see if I could get this done proper before I give myself the option of going back into the code and making it more...automatic. Regardless, I added the following additions based off of the map I was given:

I created the `Sextant{}` object which held the coordinates of the top-left corner of each side of the cube, and used numbers to explain what each side's adjacent neighbor is. Remember that I count starting at `0`, so side `1` == side `0`. In the below example, side 1 starts at `X: 50, Y: 0`, and above it is side 6, below is side 3, to the left is side 4 and to the right is side 2.

```go
c := [6]Sextant{
    // side 1, remember sides start at 0 here to add to the confusion of course
    {loc: Coord{X: 50, Y: 0},
        up:    5,
        down:  2,
        left:  3,
        right: 1,
    },
    // side 2
    {loc: Coord{X: 100, Y: 0},
        up:    5,
        down:  2,
        left:  0,
        right: 4,
    },
    // side 3
    {
        loc:   Coord{X: 50, Y: 50},
        up:    0,
        down:  4,
        left:  3,
        right: 1,
    },
    // side 4
    {
        loc:   Coord{X: 0, Y: 100},
        up:    2,
        down:  5,
        left:  0,
        right: 4,
    },
    // side 5
    {
        loc:   Coord{X: 50, Y: 100},
        up:    2,
        down:  5,
        left:  3,
        right: 1,
    },
    // side 6
    {
        loc:   Coord{X: 0, Y: 150},
        up:    3,
        down:  1,
        left:  0,
        right: 4,
    },
}
```

Finally, in the `MoveOnePartTwo()` function (which is a function that moves the monkey one step forwards, and returns `true` if the monkey was able to move forward, `false` if the monkey hits a wall), I added this matrix:

```go
DirectionMatrix := [6][6]int{
    {-1, right, down, right, -1, right},
    {left, -1, left, -1, left, up},
    {up, up, -1, down, down, -1},
    {right, -1, right, -1, right, down},
    {-1, left, up, left, -1, left},
    {down, down, -1, up, up, -1},
}
```

Note that I set up `up`, `left`, `down` and `right` as enum data types set to 0, 1, 2, and 3. This matrix works like this -- if you are on side 3 and you are travelling to side 2, you sould access `DirectionMatrix[2][1]`, which returns `up`, meaning that if you are travelling from side 3 to 2, your new direction would be facing `up`. The `-1` representations were if I was checking the same side I was already on, or if I was trying to see what direction I would face on a side that was impossible to get to from where I was (without going through another side), basically the opposite side of the cube.

Of course, I had to figure all of this out by using visual aids! Not pictured: The three-dimensional cutout I made from printer paper of my map representation. I used this to determine the directions and adjacent sides.

Finally, I used a whiteboard to find out the behavior of where my character would appear going from one plane to another if the direction changes. I had to work out which would require an "inverse" representation to find out if I needed to move my character slightly differently, so for instance if I was moving `up` on one side, then on the next side I was moving `left`, I had to adjust my positioning to be the inverse of where I was relative to the adjacent size. It's hard to explain without drawing a diagram, and I'm not about to share my chicken-scratch diagrams on github. This allowed me to figure out the formula for where to place the charachter in relation to new sides. This took a majority of my time.

### Where I got hung up

Originally in the `DirectionMatrix` object I created, I represented a `-1` as three things: First if you were checking the same side for source and destination, Second if you were trying to see what direction you would face if you checked against two sides that were NOT adjacent (ie: impossible), and Third if you were NOT changing direction going from one side to another. This turned out to be my downfall, because for some reason I was able to use this logic on the example data and it worked just fine, however if I used it on the challenge data it failed miserably. Not a position anyone should ever be in. Once I added in ALL the directions you would have to be in regardless if it's the same direction or not, the code finally just...worked. What a feeling, man.

## Things I Learned

Well there's two basic things I learned in Go. First of all, you can do a 1:1 comparison between two struct types. I was pretty excited about that, since I'm not sure you can do that in Python without defining a `__eq__()` function to the object itself. Which is fine by the way, but it's nice that Go has this capability built in. For example:

```go
type Coord struct {
    X, Y int
}

// ... etc, func main() blah blah blah

firstCoord := Coord{
    X: 50,
    Y: 100,
}

secondCoord := Coord{
    X: 50,
    Y: 100,
}

if firstCoord == secondCoord {
    fmt.Println("Coords are equal!")
}

secondCoord.X = 500

if firstCoord != secondCoord {
    fmt.Println("Coords are no longer equal!")
}
```

The above conditions evaluate to true.

Second, Go's Modulo behaves...differently. At least for what I'm used to. Generally speaking Modulo is the remainder of division, allowing "clock" mathematics, essentially, whereby if I were to take any number and modulo it by 50, it would never return a greater number than 49. This is useful for all kinds of things. But Modulo in Go behaves much differently from Python, and after reading up on this apparently there is a specific and logical reason as to why, which evidently is more [technically correct](https://go.dev/ref/spec#Arithmetic_operators). 

FUTURE DAN EDIT: after I looked this up, apparently Go is more mathematically correct in that `%` is not so much the Euclidian Modulus, but rather the _remainder_. How about that, huh? Python uses the Euclidean Modulus.

The biggest example I can think of is:

Go's Modulo:
```go
fmt.Printf("Example: %d\n", -1 % 50)
```

Prints `-1`.

Python's Modulo:
```python
print(f"Example: {-1 % 50}")
```

Prints `49`, which is the desired outcome.

To fix this, I used this function:

```go
func BetterMod(x, y int) int {
	rem := x % y
	if rem < 0 {
		rem += y
	}
	return rem
}
```

Oh well, it is what it is.

# --- Day 22: Monkey Map ---

The monkeys take you on a surprisingly easy trail through the jungle. They're even going in roughly the right direction according to your handheld device's Grove Positioning System.

As you walk, the monkeys explain that the grove is protected by a force field. To pass through the force field, you have to enter a password; doing so involves tracing a specific path on a strangely-shaped board.

At least, you're pretty sure that's what you have to do; the elephants aren't exactly fluent in monkey.

The monkeys give you notes that they took when they last saw the password entered (your puzzle input).

For example:

```
        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5
```

The first half of the monkeys' notes is a map of the board. It is comprised of a set of open tiles (on which you can move, drawn .) and solid walls (tiles which you cannot enter, drawn #).

The second half is a description of the path you must follow. It consists of alternating numbers and letters:

    A number indicates the number of tiles to move in the direction you are facing. If you run into a wall, you stop moving forward and continue with the next instruction.
    A letter indicates whether to turn 90 degrees clockwise (R) or counterclockwise (L). Turning happens in-place; it does not change your current tile.

So, a path like 10R5 means "go forward 10 tiles, then turn clockwise 90 degrees, then go forward 5 tiles".

You begin the path in the leftmost open tile of the top row of tiles. Initially, you are facing to the right (from the perspective of how the map is drawn).

If a movement instruction would take you off of the map, you wrap around to the other side of the board. In other words, if your next tile is off of the board, you should instead look in the direction opposite of your current facing as far as you can until you find the opposite edge of the board, then reappear there.

For example, if you are at A and facing to the right, the tile in front of you is marked B; if you are at C and facing down, the tile in front of you is marked D:

```
        ...#
        .#..
        #...
        ....
...#.D.....#
........#...
B.#....#...A
.....C....#.
        ...#....
        .....#..
        .#......
        ......#.
```

It is possible for the next tile (after wrapping around) to be a wall; this still counts as there being a wall in front of you, and so movement stops before you actually wrap to the other side of the board.

By drawing the last facing you had with an arrow on each tile you visit, the full path taken by the above example looks like this:

```
        >>v#    
        .#v.    
        #.v.    
        ..v.    
...#...v..v#    
>>>v...>#.>>    
..#v...#....    
...>>>>v..#.    
        ...#....
        .....#..
        .#......
        ......#.
```

To finish providing the password to this strange input device, you need to determine numbers for your final row, column, and facing as your final position appears from the perspective of the original map. Rows start from 1 at the top and count downward; columns start from 1 at the left and count rightward. (In the above example, row 1, column 1 refers to the empty space with no tile on it in the top-left corner.) Facing is 0 for right (>), 1 for down (v), 2 for left (<), and 3 for up (^). The final password is the sum of 1000 times the row, 4 times the column, and the facing.

In the above example, the final row is 6, the final column is 8, and the final facing is 0. So, the final password is 1000 * 6 + 4 * 8 + 0: 6032.

Follow the path given in the monkeys' notes. What is the final password?

Your puzzle answer was `27492`.

# --- Part Two ---

As you reach the force field, you think you hear some Elves in the distance. Perhaps they've already arrived?

You approach the strange input device, but it isn't quite what the monkeys drew in their notes. Instead, you are met with a large cube; each of its six faces is a square of 50x50 tiles.

To be fair, the monkeys' map does have six 50x50 regions on it. If you were to carefully fold the map, you should be able to shape it into a cube!

In the example above, the six (smaller, 4x4) faces of the cube are:

```
        1111
        1111
        1111
        1111
222233334444
222233334444
222233334444
222233334444
        55556666
        55556666
        55556666
        55556666
```

You still start in the same position and with the same facing as before, but the wrapping rules are different. Now, if you would walk off the board, you instead proceed around the cube. From the perspective of the map, this can look a little strange. In the above example, if you are at A and move to the right, you would arrive at B facing down; if you are at C and move down, you would arrive at D facing up:

```
        ...#
        .#..
        #...
        ....
...#.......#
........#..A
..#....#....
.D........#.
        ...#..B.
        .....#..
        .#......
        ..C...#.
```

Walls still block your path, even if they are on a different face of the cube. If you are at E facing up, your movement is blocked by the wall marked by the arrow:

```
        ...#
        .#..
     -->#...
        ....
...#..E....#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.
```

Using the same method of drawing the last facing you had with an arrow on each tile you visit, the full path taken by the above example now looks like this:

```
        >>v#    
        .#v.    
        #.v.    
        ..v.    
...#..^...v#    
.>>>>>^.#.>>    
.^#....#....    
.^........#.    
        ...#..v.
        .....#v.
        .#v<<<<.
        ..v...#.
```

The final password is still calculated from your final position and facing from the perspective of the map. In this example, the final row is 5, the final column is 7, and the final facing is 3, so the final password is 1000 * 5 + 4 * 7 + 3 = 5031.

Fold the map into a cube, then follow the path given in the monkeys' notes. What is the final password?

Your puzzle answer was `78291`.

Both parts of this puzzle are complete! They provide two gold stars: **

At this point, you should return to your Advent calendar and try another puzzle.

If you still want to see it, you can get your puzzle input.