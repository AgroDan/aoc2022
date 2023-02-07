# Dan's Introduction

Wow. I'm both super proud of this and super ashamed. Proud that I was able to figure this out myself without getting help from better developers. Ashamed that the code came out as ugly and "spaghetti-code" as I could possibly make it. Seriously, I started referring to items as things like `m[m[thisMonkey].PrimaryMonkey].Name`, which is just dreadful. Still though, the code is fast and accurate. I had to do a bit of re-learning basic algrebra but it was worth it. I managed to get this using my own home-grown algorithm with a little shade of Depth-First-Search. There's something about creating your own recursive algorithm that just feels really good, y'know?

## Part 1

Not so bad. As usual, I wrote up the concept of creating objects for pretty much everything I could, but in this case the objects were `Monkeys`. Each `Monkey` had a bunch of datatypes associated with it, but it all seemed to be pretty useful. Basically if the monkey's job was to perform some operation based on the result of 2 other monkeys, I would store all of that. The primary and secondary monkey (so I had some recollection of order when performing the operation), and the operation associated with it, where `+` and `-` were addition and subtraction, and `/` and `*` were division and multiplication. If the monkey's job was to shout out a number, then the number would be stored in `Monkey.Result` and the `Monkey.JobStatus` would be set to `true`.

With that taken care of, all I needed to do was find the `root` Monkey and perform a recursive function which would continue to dive into the subroutines associated with the chosen monkey and it would spit out the end result of all of the operations associated with the entire monkey hierarchy. I don't blame you if this is confusing, it made more sense to me in my head.

## Part 2

Now the operations have flipped a bit. Basically I needed to ignore the operation associated with the `root` monkey and change the operation to `=`, which stated that the two numbers provided are equal. To make it weirder, the `humn` monkey is actually me! I have to ignore the value given to _that_ "monkey" and insert my own value to make sure that the `root` monkey's equality statement evaluates to `true`. To accomplish that, I had to kinda work backwards. Basically I had to code a function that determined which "branch" lead to the `humn` object, and took the numeric resulting value of the alternate fork and used that as a known value. From there I used basic algebra to determine the unknowns, drilling down all the way to the `humn` object.

This is a lot harder to explain than I thought, and even harder to code in an "easy" way.

## What I Learned

I didn't really use any new techniques here. I think probably the best takeaway from this is writing a recursive function that actually worked insanely fast. Score.

# --- Day 21: Monkey Math ---

The monkeys are back! You're worried they're going to try to steal your stuff again, but it seems like they're just holding their ground and making various monkey noises at you.

Eventually, one of the elephants realizes you don't speak monkey and comes over to interpret. As it turns out, they overheard you talking about trying to find the grove; they can show you a shortcut if you answer their riddle.

Each monkey is given a job: either to yell a specific number or to yell the result of a math operation. All of the number-yelling monkeys know their number from the start; however, the math operation monkeys need to wait for two other monkeys to yell a number, and those two other monkeys might also be waiting on other monkeys.

Your job is to work out the number the monkey named root will yell before the monkeys figure it out themselves.

For example:

```
root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32
```

Each line contains the name of a monkey, a colon, and then the job of that monkey:

```
    A lone number means the monkey's job is simply to yell that number.
    A job like aaaa + bbbb means the monkey waits for monkeys aaaa and bbbb to yell each of their numbers; the monkey then yells the sum of those two numbers.
    aaaa - bbbb means the monkey yells aaaa's number minus bbbb's number.
    Job aaaa * bbbb will yell aaaa's number multiplied by bbbb's number.
    Job aaaa / bbbb will yell aaaa's number divided by bbbb's number.
```

So, in the above example, monkey drzm has to wait for monkeys hmdt and zczc to yell their numbers. Fortunately, both hmdt and zczc have jobs that involve simply yelling a single number, so they do this immediately: 32 and 2. Monkey drzm can then yell its number by finding 32 minus 2: 30.

Then, monkey sjmn has one of its numbers (30, from monkey drzm), and already has its other number, 5, from dbpl. This allows it to yell its own number by finding 30 multiplied by 5: 150.

This process continues until root yells a number: 152.

However, your actual situation involves considerably more monkeys. What number will the monkey named root yell?

Your puzzle answer was `51928383302238`.

# --- Part Two ---

Due to some kind of monkey-elephant-human mistranslation, you seem to have misunderstood a few key details about the riddle.

First, you got the wrong job for the monkey named root; specifically, you got the wrong math operation. The correct operation for monkey root should be =, which means that it still listens for two numbers (from the same two monkeys as before), but now checks that the two numbers match.

Second, you got the wrong monkey for the job starting with humn:. It isn't a monkey - it's you. Actually, you got the job wrong, too: you need to figure out what number you need to yell so that root's equality check passes. (The number that appears after humn: in your input is now irrelevant.)

In the above example, the number you need to yell to pass root's equality test is 301. (This causes root to get the same number, 150, from both of its monkeys.)

What number do you yell to pass root's equality test?

Your puzzle answer was `3305669217840`.

Both parts of this puzzle are complete! They provide two gold stars: **

At this point, you should return to your Advent calendar and try another puzzle.

If you still want to see it, you can get your puzzle input.