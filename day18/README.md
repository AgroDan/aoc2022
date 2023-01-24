# Dan's Introduction

This one was fun! Not nearly as mind-bending as Day 17 and I actually was able to implement a previously-learned function to determine some real-world stuff! Specifically the "Breadth-First-Search" function. Today I learned that the only difference between DFS and BFS is that DFS uses a stack, and BFS uses a queue. OK!

## Part 1

The first part was pretty straightforward. In 3-Dimensional space, map out the surface area of the droplets of lava flying past at that moment in time. Since each block had 6 sides if they never touched another adjacent block, the steps I took was:

1. Create a 3-Dimensional array, and ingest every single known point as a specific block in the array.
2. For every single known point, determine the surface area by counting every single edge that _isn't_ touching another lava block
3. Add them all up

That was incredibly easy. It wasn't until the second half where the complexity began to ramp up.

## Part 2

Now apparently this is all wrong. We can't count just the surface area of a lava block against an air block, we need to _negate_ any block that may be inside of a bubble within a larger lava block! Basically if there is a 1-unit pocket of air that is surrounded on all 6 sides by lava, the answer to Part 1 will count the surface area within that air pocket, when what we _really_ need is just the surface area of the lava blob around it!

So that left me with quite a challenge. How do I determine bubbles? Especially if a bubble is bigger than just one cube? It took quite a bit of thinking (I'm just not as smart as some of those other speed programmers that kicked butt in this advent of code challenge), but eventually I came across the following logic:

1. For every point of lava, find all blocks around it that are "air" blocks.
2. For each air block, perform a Breadth-First-Search traversal outwards until either A) All blocks have been visited, or B) we traverse out to the edge of the graph, which would mean we are _not_ inside of a bubble. If we _are_ in a bubble, return every point we visited.
3. De-duplicate the list of "bubble zones" we acquired, since we are bound to have them. We only need unique bubble points.
4. Obtain the total surface area of all the lava edges of each of these bubbles
5. Subtract that surface area from the total surface area
6. Win!

## What I Learned

Honestly, not all that much. I mean, I did learn about the Breadth-First-Search algorithm as opposed to Depth-First-Search, but still it's not all that different. What set me back the most time was mostly my unwillingness to use a contrived 3-D set. I wanted as small a graph as possible. At first, I was tempted to use something like `var space [100][100][100]rune` and call it a day, but I wanted some way to dynamically grow the 3-D space as needed. At first I had this long convoluted algorithm to simply add to the space map in three dimensions, but with my pea-sized brain I simply could not wrap my head around the logic of doing that. Add to the X axis, then go back and revisit all the old items on the X axis and grow the Y axis and Z axis and...ugh, my brain meats are hurting just thinking back.

Ultimately I just rebuilt a new 3-D array in parallel, copied all the points in the map to the new map, and de-allocated the old map. I went from something to the tune of 80+ lines of code to 18. What a boneheaded move.

Other than that however, the rest was pretty smooth sailing. Again, I think I may have over-engineered this a bit and I can _probably_ make this code a bit more efficient, but hey whatever gets the job done. Besides, coding should be _fun!_

# --- Day 18: Boiling Boulders ---

You and the elephants finally reach fresh air. You've emerged near the base of a large volcano that seems to be actively erupting! Fortunately, the lava seems to be flowing away from you and toward the ocean.

Bits of lava are still being ejected toward you, so you're sheltering in the cavern exit a little longer. Outside the cave, you can see the lava landing in a pond and hear it loudly hissing as it solidifies.

Depending on the specific compounds in the lava and speed at which it cools, it might be forming obsidian! The cooling rate should be based on the surface area of the lava droplets, so you take a quick scan of a droplet as it flies past you (your puzzle input).

Because of how quickly the lava is moving, the scan isn't very good; its resolution is quite low and, as a result, it approximates the shape of the lava droplet with 1x1x1 cubes on a 3D grid, each given as its x,y,z position.

To approximate the surface area, count the number of sides of each cube that are not immediately connected to another cube. So, if your scan were only two adjacent cubes like 1,1,1 and 2,1,1, each cube would have a single side covered and five sides exposed, a total surface area of 10 sides.

Here's a larger example:

```
2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5
```

In the above example, after counting up all the sides that aren't connected to another cube, the total surface area is 64.

What is the surface area of your scanned lava droplet?

Your puzzle answer was `3364`.

# --- Part Two ---

Something seems off about your calculation. The cooling rate depends on exterior surface area, but your calculation also included the surface area of air pockets trapped in the lava droplet.

Instead, consider only cube sides that could be reached by the water and steam as the lava droplet tumbles into the pond. The steam will expand to reach as much as possible, completely displacing any air on the outside of the lava droplet but never expanding diagonally.

In the larger example above, exactly one cube of air is trapped within the lava droplet (at 2,2,5), so the exterior surface area of the lava droplet is 58.

What is the exterior surface area of your scanned lava droplet?

Your puzzle answer was `2006`.

Both parts of this puzzle are complete! They provide two gold stars: **

At this point, you should return to your Advent calendar and try another puzzle.

If you still want to see it, you can get your puzzle input.