# Dan's Introduction

This challenge was deceptively difficult! It seemed relatively easy conceptually, but after really working on it and trying to find out the answer on my own, I had to resort to looking at others' answers who were certainly much better than my approach. Still, this isn't a contest [for me at least], this is a learning experience. My intention here is just to learn Go, and I'm learning quite a bit doing this.

So there are two algorithms at work here: [Depth-First-Search](https://en.wikipedia.org/wiki/Depth-first_search), and the [Floyd-Warshall Algorithm](https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm). The former finds all permutations of possible routes until the time expires, and the latter finds the shortest distance between any two points on the entire map.

## Part 1 - My Logic
First, after parsing all valves into `Valve` objects, I used the `Floyd-Warshall Algorithm` to create a matrix/2-dimensional array that can be used to determine the shortest distance between two points. So after using said algorithm, I had the variable `matrix` which I can use like `matrix[from][to]` which would return a number, and this number was the shortest possible distance between the `from` and `to` points. I used this to find the most efficient paths.

Second, I used the Depth-First-Search algorithm to find all possible paths I can take before the time runs out or if I hit a wall where I can't go anywhere else. With these two functions, I created a new `Path` object which contained an unordered set of the valves I visited in the path, then calculated the total pressure output of each path and threw it into an array. I then simply chose the highest pressure output from the array as the answer to Part 1.

To make this faster, the only valves I cared about were the ones that had a positive value! Zero-value pressure valves (or as the challenge explained it, broken or stuck) were pointless to visit, so the only valves I cared about were >0. This meant that even though every vertex in the map was a valve, I could still walk through the valve room without doing anything and it would only take 1 minute to move from one valve to another. With a combination of DFS and Floyd Warshall, I could determine that it would take N steps to get from Valve X to Valve Y, and I could thus use those values to calculate the time it would take to travel to two different points.

## Part 2 - My Logic
Now that I had a capable way of finding out the highest possible output, it stands to reason that the most efficient way of calculating the highest possible pressure with a partner to open valves along with me is to have that partner _only open valves that I never open_. That seems relatively obvious since there's no point in re-opening a valve that's currently open. My first thought was to create some ridiculous step-by-step algorithm to run two processes concurrently and avoid places that have already been, but that was _severely_ over-complicating tthe issue. I'm only supposed to find the most efficient routes and return the highest possible pressure I can release.

So...why not just get the same list of all possible permutations and choose the two highest? That is...as long as those two possible permutations _never overlap at all_. So that means the first thing I need to find are path-pairs that have absolutely nothing in common with each other, meaning neither me nor the elephant ever cross paths, but in-so-doing we also release the most possible pressure.

So I simply took the list of all possible path permutations and looped through each possible path comparing _its_ visited valves with each path and returning _only the paths that were unique to each other_. Once I had all the unique pairs, I simply added the two total pressure values and chose the pair with the highest combined total This took nearly 2 minutes to complete, but eventually it returned the proper value. _Whew!_

## Things I Learned

I got a bit stronger in my knowledge of handling multi-dimensional arrays, as well as better utilized the unordered sets I learned from [Day 15](../day15/README.md). In addition to that, I discovered how to make multi-dimensional _maps_, which is a weird one if I'm being honest. But it allowed me to create the matrix object which allowed me to map something like `matrix["AZ"]["BD"]` which pointed to the lowest possible distance between the two points `AZ` and `BD`.

And finally I learned about the `time.Now()` function and `time.Since()` functions, which gives me a decent measure of how long my code took to process without setting a timer per se, but measuring the delta between when the timestamp was created and when the `Since()` function is called. Pretty neat!

I have to state that I am simply not smart enough to come up with these algorithms on my own, so these are somewhat plagiarized from better coders than myself. But again, not trying to compete or prove anything, this is generally for my own reference and practice.

# --- Day 16: Proboscidea Volcanium ---

The sensors have led you to the origin of the distress signal: yet another handheld device, just like the one the Elves gave you. However, you don't see any Elves around; instead, the device is surrounded by elephants! They must have gotten lost in these tunnels, and one of the elephants apparently figured out how to turn on the distress signal.

The ground rumbles again, much stronger this time. What kind of cave is this, exactly? You scan the cave with your handheld device; it reports mostly igneous rock, some ash, pockets of pressurized gas, magma... this isn't just a cave, it's a volcano!

You need to get the elephants out of here, quickly. Your device estimates that you have 30 minutes before the volcano erupts, so you don't have time to go back out the way you came in.

You scan the cave for other options and discover a network of pipes and pressure-release valves. You aren't sure how such a system got into a volcano, but you don't have time to complain; your device produces a report (your puzzle input) of each valve's flow rate if it were opened (in pressure per minute) and the tunnels you could use to move between the valves.

There's even a valve in the room you and the elephants are currently standing in labeled AA. You estimate it will take you one minute to open a single valve and one minute to follow any tunnel from one valve to another. What is the most pressure you could release?

For example, suppose you had the following scan output:

```
Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II
```

All of the valves begin closed. You start at valve AA, but it must be damaged or jammed or something: its flow rate is 0, so there's no point in opening it. However, you could spend one minute moving to valve BB and another minute opening it; doing so would release pressure during the remaining 28 minutes at a flow rate of 13, a total eventual pressure release of 28 * 13 = 364. Then, you could spend your third minute moving to valve CC and your fourth minute opening it, providing an additional 26 minutes of eventual pressure release at a flow rate of 2, or 52 total pressure released by valve CC.

Making your way through the tunnels like this, you could probably open many or all of the valves by the time 30 minutes have elapsed. However, you need to release as much pressure as possible, so you'll need to be methodical. Instead, consider this approach:

```
== Minute 1 ==
No valves are open.
You move to valve DD.

== Minute 2 ==
No valves are open.
You open valve DD.

== Minute 3 ==
Valve DD is open, releasing 20 pressure.
You move to valve CC.

== Minute 4 ==
Valve DD is open, releasing 20 pressure.
You move to valve BB.

== Minute 5 ==
Valve DD is open, releasing 20 pressure.
You open valve BB.

== Minute 6 ==
Valves BB and DD are open, releasing 33 pressure.
You move to valve AA.

== Minute 7 ==
Valves BB and DD are open, releasing 33 pressure.
You move to valve II.

== Minute 8 ==
Valves BB and DD are open, releasing 33 pressure.
You move to valve JJ.

== Minute 9 ==
Valves BB and DD are open, releasing 33 pressure.
You open valve JJ.

== Minute 10 ==
Valves BB, DD, and JJ are open, releasing 54 pressure.
You move to valve II.

== Minute 11 ==
Valves BB, DD, and JJ are open, releasing 54 pressure.
You move to valve AA.

== Minute 12 ==
Valves BB, DD, and JJ are open, releasing 54 pressure.
You move to valve DD.

== Minute 13 ==
Valves BB, DD, and JJ are open, releasing 54 pressure.
You move to valve EE.

== Minute 14 ==
Valves BB, DD, and JJ are open, releasing 54 pressure.
You move to valve FF.

== Minute 15 ==
Valves BB, DD, and JJ are open, releasing 54 pressure.
You move to valve GG.

== Minute 16 ==
Valves BB, DD, and JJ are open, releasing 54 pressure.
You move to valve HH.

== Minute 17 ==
Valves BB, DD, and JJ are open, releasing 54 pressure.
You open valve HH.

== Minute 18 ==
Valves BB, DD, HH, and JJ are open, releasing 76 pressure.
You move to valve GG.

== Minute 19 ==
Valves BB, DD, HH, and JJ are open, releasing 76 pressure.
You move to valve FF.

== Minute 20 ==
Valves BB, DD, HH, and JJ are open, releasing 76 pressure.
You move to valve EE.

== Minute 21 ==
Valves BB, DD, HH, and JJ are open, releasing 76 pressure.
You open valve EE.

== Minute 22 ==
Valves BB, DD, EE, HH, and JJ are open, releasing 79 pressure.
You move to valve DD.

== Minute 23 ==
Valves BB, DD, EE, HH, and JJ are open, releasing 79 pressure.
You move to valve CC.

== Minute 24 ==
Valves BB, DD, EE, HH, and JJ are open, releasing 79 pressure.
You open valve CC.

== Minute 25 ==
Valves BB, CC, DD, EE, HH, and JJ are open, releasing 81 pressure.

== Minute 26 ==
Valves BB, CC, DD, EE, HH, and JJ are open, releasing 81 pressure.

== Minute 27 ==
Valves BB, CC, DD, EE, HH, and JJ are open, releasing 81 pressure.

== Minute 28 ==
Valves BB, CC, DD, EE, HH, and JJ are open, releasing 81 pressure.

== Minute 29 ==
Valves BB, CC, DD, EE, HH, and JJ are open, releasing 81 pressure.

== Minute 30 ==
Valves BB, CC, DD, EE, HH, and JJ are open, releasing 81 pressure.
```

This approach lets you release the most pressure possible in 30 minutes with this valve layout, 1651.

Work out the steps to release the most pressure in 30 minutes. What is the most pressure you can release?

Your puzzle answer was `1595`.

# --- Part Two ---

You're worried that even with an optimal approach, the pressure released won't be enough. What if you got one of the elephants to help you?

It would take you 4 minutes to teach an elephant how to open the right valves in the right order, leaving you with only 26 minutes to actually execute your plan. Would having two of you working together be better, even if it means having less time? (Assume that you teach the elephant before opening any valves yourself, giving you both the same full 26 minutes.)

In the example above, you could teach the elephant to help you as follows:

```
== Minute 1 ==
No valves are open.
You move to valve II.
The elephant moves to valve DD.

== Minute 2 ==
No valves are open.
You move to valve JJ.
The elephant opens valve DD.

== Minute 3 ==
Valve DD is open, releasing 20 pressure.
You open valve JJ.
The elephant moves to valve EE.

== Minute 4 ==
Valves DD and JJ are open, releasing 41 pressure.
You move to valve II.
The elephant moves to valve FF.

== Minute 5 ==
Valves DD and JJ are open, releasing 41 pressure.
You move to valve AA.
The elephant moves to valve GG.

== Minute 6 ==
Valves DD and JJ are open, releasing 41 pressure.
You move to valve BB.
The elephant moves to valve HH.

== Minute 7 ==
Valves DD and JJ are open, releasing 41 pressure.
You open valve BB.
The elephant opens valve HH.

== Minute 8 ==
Valves BB, DD, HH, and JJ are open, releasing 76 pressure.
You move to valve CC.
The elephant moves to valve GG.

== Minute 9 ==
Valves BB, DD, HH, and JJ are open, releasing 76 pressure.
You open valve CC.
The elephant moves to valve FF.

== Minute 10 ==
Valves BB, CC, DD, HH, and JJ are open, releasing 78 pressure.
The elephant moves to valve EE.

== Minute 11 ==
Valves BB, CC, DD, HH, and JJ are open, releasing 78 pressure.
The elephant opens valve EE.

(At this point, all valves are open.)

== Minute 12 ==
Valves BB, CC, DD, EE, HH, and JJ are open, releasing 81 pressure.

...

== Minute 20 ==
Valves BB, CC, DD, EE, HH, and JJ are open, releasing 81 pressure.

...

== Minute 26 ==
Valves BB, CC, DD, EE, HH, and JJ are open, releasing 81 pressure.
```

With the elephant helping, after 26 minutes, the best you could do would release a total of 1707 pressure.

With you and an elephant working together for 26 minutes, what is the most pressure you could release?

Your puzzle answer was `2189`.

Both parts of this puzzle are complete! They provide two gold stars: **

At this point, you should return to your Advent calendar and try another puzzle.

If you still want to see it, you can get your puzzle input.