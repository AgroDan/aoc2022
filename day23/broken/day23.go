package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

func ParseMap(input [][]byte) GroveMap {

	g := NewGroveMap(len(input))
	for i, v := range input {
		for j, w := range v {
			g.p[i][j] = rune(w)
		}
	}
	return g
}

func BuildElves(g *GroveMap) ElfStack {
	e := make([]*Elf, 0)
	for i, row := range g.p {
		for j, col := range row {
			if col == '#' {
				n := NewElf(j, i)
				e = append(e, &n)
			}
		}
	}
	return e
}

func DebugCoord(c Coord, e *ElfStack, g *GroveMap) {
	// this function will dump as much info about the Elf
	// if there exists one at this coord at the time this
	// function is called.
	for _, v := range *e {
		if v.Loc.X == c.X && v.Loc.Y == c.Y {
			v.PrintElf()
			fmt.Printf("Is the Elf alone? %t\n", v.Alone(g))
		}
	}
}

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")

	flag.Parse()
	readFile, err := os.Open(*filePtr)

	if err != nil {
		fmt.Println("Fatal:", err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines [][]byte

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Bytes())
	}

	// Insert code here
	g := ParseMap(lines)
	e := BuildElves(&g)

	g.GrowMap(2) // grow the map to 10
	e.GrowMap(2) // update all the elves positions

	fmt.Println("Starting Map:")
	PrintMap(&g)

	// now determine out of the proposals who is actually moving.
	// if two or more elves propose the same coord, reset all actions.

	// now move the elves that are gonna move
	for i := 0; i < 10; i++ {
		// first half, propose directions
		for _, v := range e {
			v.ProposeDirection(&g)
		}

		// fmt.Printf("Before proposals:\n")
		// // debug
		// DebugCoord(Coord{Y: 4, X: 3}, &e, &g)

		// review the proposals...
		e.ReviewProposals()

		// // debug
		// fmt.Printf("After proposals:\n")
		// DebugCoord(Coord{Y: 4, X: 3}, &e, &g)

		// now move the elves
		for _, v := range e {
			if v.Move {
				g.MoveElf(v)
			}
		}

		// Do we grow?
		if g.ElvesAtBorder() {
			g.GrowMap(1)
			e.GrowMap(1)
		}

		fmt.Printf("\nStarting direction: %c\n", g.MoveQueue[0])
		fmt.Printf("Iteration %d, map:\n", i+1)
		PrintMap(&g)

		// now rotate the moveQueue
		g.Rotate()

		// and reset the elves
		e.ResetElves()
	}

	// fmt.Printf("%#v\n", g)

	fmt.Printf("Beginning to shrink map...\n")
	// start shinking the map
	for !g.ElvesAtBorder() {
		g.ShrinkMap()
		e.GrowMap(-1)
	}
	for !g.ElfAtEdge('N') {
		g.ShaveDirection('N', &e)
	}
	for !g.ElfAtEdge('S') {
		g.ShaveDirection('S', &e)
	}
	for !g.ElfAtEdge('E') {
		g.ShaveDirection('E', &e)
	}
	for !g.ElfAtEdge('W') {
		g.ShaveDirection('W', &e)
	}
	// fmt.Printf("Shaving left...\n")
	// g.ShaveDirection('W', &e)
	// fmt.Printf("Elf up: %t, Elf down: %t\n", g.ElfAtEdge('N'), g.ElfAtEdge('S'))
	// fmt.Printf("Elf Left: %t, Elf right: %t\n", g.ElfAtEdge('W'), g.ElfAtEdge('E'))

	fmt.Printf("\n\nFinal Map:\n")
	// Print what we have!
	PrintMap(&g)

	fmt.Printf("Based on the above calculations, the amount of empty ground tiles is: %d\n", g.CountEmpty())

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
