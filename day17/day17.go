package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	Left = iota
	Right
	Down
)

func cycle(dir int, t *TransitionalRock) bool {
	// this function will cycle ONE direction, and THEN down.
	// if this ever returns false, then the transitional rock
	// has hit the floor (or as far down as it will go).
	switch dir {
	case Right:
		// don't care if we hit a wall. it should correct itself
		_ = t.Move(Right)
	case Left:
		_ = t.Move(Left)
	}

	// now move down. If we can move down, then it is still free-falling.
	// if we can't, then we hit the ground and the calling function is
	// responsible for writing the block.
	return t.Move(Down)
}

func main() {

	filePtr := flag.String("f", "input", "Input file if not 'input'")
	maxCyclesPtr := flag.Int("r", 2022, "Max rocks to iterate for")
	printPtr := flag.Bool("p", false, "Print map")
	metricPtr := flag.Bool("m", false, "Print metrics (lots of output)")

	flag.Parse()
	readFile, err := os.Open(*filePtr)

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	j, err := io.ReadAll(readFile)

	if err != nil {
		fmt.Println("Fatal:", err)
	}

	jets := make([]string, 0)
	for _, v := range j {
		jets = append(jets, fmt.Sprintf("%c", v))
	}

	// now loop through the jets
	counter := 0
	newRock := true         // start with a rock
	n := TransitionalRock{} // this will get overwritten
	rockCounter := 0        // +1 for each rock generated
	max := *maxCyclesPtr
	metrics := NewMetric()

	c := NewCavern() // new cave

	for {
		if newRock {
			// we need to generate a new rock. But before we do that, let's do some looking back.
			// first, find the highest rock so we know where to generate the new one

			highest, err := c.FindHighest()
			if err != nil {
				panic("Couldn't compute highest rock")
			}

			// take this rock type, the jet index, and return all the heights
			metrics.Add(rockCounter%5, counter%len(jets), highest, rockCounter)

			// keep this part here because it used the binary signature function
			//
			// // now let's look back only if it's in bounds
			// if highest > *lookBackPtr+*sigLengthPtr {
			// 	// fmt.Printf("Highest: %d, lookback+siglength: %d\n", highest, *lookBackPtr+*sigLengthPtr)
			// 	compareSig := c.GetSignature(*sigLengthPtr, highest-*sigLengthPtr)
			// 	// fmt.Printf("compare sig: %b\n", compareSig)
			// 	for i := highest - *lookBackPtr; i < highest; i++ {
			// 		if i+*sigLengthPtr > highest {
			// 			break
			// 		}
			// 		lineSig := c.GetSignature(*sigLengthPtr, i)
			// 		// fmt.Printf("Line sig: %b\n", lineSig)
			// 		if compareSig == lineSig {
			// 			fmt.Printf("Signature match at height: %d, prev at %d, delta: %d\n", highest, prev, highest-i)
			// 			prev = i
			// 		}
			// 	}
			// }

			thisRock := Rock{}
			switch rockCounter % 5 {
			case 0:
				thisRock = HorizontalLine()
			case 1:
				thisRock = PlusSign()
			case 2:
				thisRock = Corner()
			case 3:
				thisRock = VerticalLine()
			case 4:
				thisRock = Square()
			}
			newRock = false
			n = NewTransitionalRock(highest, thisRock, &c)
			rockCounter++
		}
		// infinite, will break normally
		if jets[counter%len(jets)] == "<" {
			if !cycle(Left, &n) {
				c.DrawRock(n)
				if rockCounter >= max {
					break
				}
				newRock = true
			}
		} else if jets[counter%len(jets)] == ">" {
			if !cycle(Right, &n) {
				c.DrawRock(n)
				if rockCounter >= max {
					break
				}
				newRock = true
			}
		}
		counter++
	}

	if *printPtr {
		c.PrintCavern()
	}
	fullHeight, err := c.FindHighest()
	if err != nil {
		fmt.Printf("could not find highest")
	}
	fmt.Printf("The height of these rocks is: %d\n", fullHeight)

	if *metricPtr {
		for k, v := range metrics.drops {
			dupes := FindDupes(metrics.GetDeltas(k))
			rockDupes := FindDupes(metrics.GetRockDeltas(k))
			if len(dupes) > 0 && len(rockDupes) > 0 {
				fmt.Printf("Duplicate Height Deltas: %s -> Deltas: %v\n", k, dupes)
				fmt.Printf("Rocks Dropped: %s -> Deltas %v\n", k, rockDupes)
				for _, w := range v {
					fmt.Printf("[r: %d, h: %d] ", w.rocksDropped, w.height)
				}
				fmt.Printf("\n\n")
			}
		}
	}
}
