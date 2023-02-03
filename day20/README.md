# Dan's Introduction

I really enjoyed this one! I finally got a chance to use a data model prototype I learned back in college...a Linked List! Technically a Circular Bi-Directional Linked List. The benefits of a linked list is that the data doesn't really move within the memory, so there isn't a whole lot of calculations done on resizing the data. The negatives here is that it becomes fairly clumsy to refer to _specific items_ within the dataset, since really the only way you can refer to something like the fifth element in the list is to search for an anchor point (remember that it's circular, so there is not technical "start" point), and then hop 5 nodes after. However, re-ordering the list is trivial since all you're doing is changing the address locations of the `next` and `prev` values.

## Part 1

Not having made an actual linked list in...jeez, 20 years? I struggled a bit with the implementation. This was a special one though as it was bidirectional and circular. Normally a linked list is one direction and it's super easy to add and remove from it. But bidirectional offers a whole slew of new challenges that was a brain twister to figure out. If you want to insert a node inbetween two other nodes that are linked, first take the new node's `prev` and `next` variables and point them at the nodes we are inserting the node between. Then change the node behind it's `next` variable to the new node, and the node in front of it's `prev` variable to the new node. Also, if we're moving a node _from_ an old position, you have to update the old `prev` and `next` nodes variables.

...and if that's confusing, then yeah. It was confusing keeping track of it all. Eventually I figured it out. It wasn't a majority of the problem, but it took some time to figure out.

The next challenge I had to overcome was figuring out that if I am moving a node through the chain and the number provided was _larger than the chain_, it would ultimately wrap around and hit the original place it was in. I originally counted this position because I wasn't moving it through the chain, just finding out where to put it. Eventually I figured out that I couldn't count the same position I started, which was throwing off my numbers. Finally figured that out.

## Part 2

This part was a bit contrived. Basically do what we did before, just do it 10 times. Oh and also multiply the numbers in the provided file by some large integer. Because these numbers were so high, I had to find a way of doing this within a reasonable time. Enter the modulo.

The basic formula was each node was moved `number % (len(list) - 1)` places, which cut down the time to move everything considerably. It took me some time and effort to find out that formula, but I'm happy to know that I did it.

The code included executes in about 312ms. Nice!

## What I Learned

Linked lists can be confusing, but once you get them up and running they are pretty sweet. I think what I learned the most here is just a better handle on pointers and how they can be used to my benefit.

# --- Day 20: Grove Positioning System ---

It's finally time to meet back up with the Elves. When you try to contact them, however, you get no reply. Perhaps you're out of range?

You know they're headed to the grove where the star fruit grows, so if you can figure out where that is, you should be able to meet back up with them.

Fortunately, your handheld device has a file (your puzzle input) that contains the grove's coordinates! Unfortunately, the file is encrypted - just in case the device were to fall into the wrong hands.

Maybe you can decrypt it?

When you were still back at the camp, you overheard some Elves talking about coordinate file encryption. The main operation involved in decrypting the file is called mixing.

The encrypted file is a list of numbers. To mix the file, move each number forward or backward in the file a number of positions equal to the value of the number being moved. The list is circular, so moving a number off one end of the list wraps back around to the other end as if the ends were connected.

For example, to move the 1 in a sequence like 4, 5, 6, 1, 7, 8, 9, the 1 moves one position forward: 4, 5, 6, 7, 1, 8, 9. To move the -2 in a sequence like 4, -2, 5, 6, 7, 8, 9, the -2 moves two positions backward, wrapping around: 4, 5, 6, 7, 8, -2, 9.

The numbers should be moved in the order they originally appear in the encrypted file. Numbers moving around during the mixing process do not change the order in which the numbers are moved.

Consider this encrypted file:

```
1
2
-3
3
-2
0
4
```

Mixing this file proceeds as follows:

```
Initial arrangement:
1, 2, -3, 3, -2, 0, 4

1 moves between 2 and -3:
2, 1, -3, 3, -2, 0, 4

2 moves between -3 and 3:
1, -3, 2, 3, -2, 0, 4

-3 moves between -2 and 0:
1, 2, 3, -2, -3, 0, 4

3 moves between 0 and 4:
1, 2, -2, -3, 0, 3, 4

-2 moves between 4 and 1:
1, 2, -3, 0, 3, 4, -2

0 does not move:
1, 2, -3, 0, 3, 4, -2

4 moves between -3 and 0:
1, 2, -3, 4, 0, 3, -2
```

Then, the grove coordinates can be found by looking at the 1000th, 2000th, and 3000th numbers after the value 0, wrapping around the list as necessary. In the above example, the 1000th number after 0 is 4, the 2000th is -3, and the 3000th is 2; adding these together produces 3.

Mix your encrypted file exactly once. What is the sum of the three numbers that form the grove coordinates?

Your puzzle answer was `13967`.

# --- Part Two ---

The grove coordinate values seem nonsensical. While you ponder the mysteries of Elf encryption, you suddenly remember the rest of the decryption routine you overheard back at camp.

First, you need to apply the decryption key, 811589153. Multiply each number by the decryption key before you begin; this will produce the actual list of numbers to mix.

Second, you need to mix the list of numbers ten times. The order in which the numbers are mixed does not change during mixing; the numbers are still moved in the order they appeared in the original, pre-mixed list. (So, if -3 appears fourth in the original list of numbers to mix, -3 will be the fourth number to move during each round of mixing.)

Using the same example as above:

```
Initial arrangement:
811589153, 1623178306, -2434767459, 2434767459, -1623178306, 0, 3246356612

After 1 round of mixing:
0, -2434767459, 3246356612, -1623178306, 2434767459, 1623178306, 811589153

After 2 rounds of mixing:
0, 2434767459, 1623178306, 3246356612, -2434767459, -1623178306, 811589153

After 3 rounds of mixing:
0, 811589153, 2434767459, 3246356612, 1623178306, -1623178306, -2434767459

After 4 rounds of mixing:
0, 1623178306, -2434767459, 811589153, 2434767459, 3246356612, -1623178306

After 5 rounds of mixing:
0, 811589153, -1623178306, 1623178306, -2434767459, 3246356612, 2434767459

After 6 rounds of mixing:
0, 811589153, -1623178306, 3246356612, -2434767459, 1623178306, 2434767459

After 7 rounds of mixing:
0, -2434767459, 2434767459, 1623178306, -1623178306, 811589153, 3246356612

After 8 rounds of mixing:
0, 1623178306, 3246356612, 811589153, -2434767459, 2434767459, -1623178306

After 9 rounds of mixing:
0, 811589153, 1623178306, -2434767459, 3246356612, 2434767459, -1623178306

After 10 rounds of mixing:
0, -2434767459, 1623178306, 3246356612, -1623178306, 2434767459, 811589153
```

The grove coordinates can still be found in the same way. Here, the 1000th number after 0 is 811589153, the 2000th is 2434767459, and the 3000th is -1623178306; adding these together produces 1623178306.

Apply the decryption key and mix your encrypted file ten times. What is the sum of the three numbers that form the grove coordinates?

Your puzzle answer was `1790365671518`.

Both parts of this puzzle are complete! They provide two gold stars: **

At this point, you should return to your Advent calendar and try another puzzle.

If you still want to see it, you can get your puzzle input.