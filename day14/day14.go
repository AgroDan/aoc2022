package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Sprintf("Could not read file: %s", err))
	}

	// myCave := NewCave(60, 20)
	// PrintCave(&myCave)
	rockCollection := NewRockCollection(dat)
	part1Cave := NewCave(rockCollection)

	// now drop the sand

	fmt.Printf("Part 1:\n\n")

	PourSand(&part1Cave, 500, 0)

	PrintCave(&part1Cave)

	fmt.Printf("Amount of sand units: %d\n\n", CountSand(&part1Cave))

	// now do part 2
	fmt.Printf("Part 2:\n\n")
	part2Cave := NewCavePart2(rockCollection)

	PourSandPart2(&part2Cave, 500, 0)
	PrintCave(&part2Cave)

	fmt.Printf("Amount of sand units: %d\n", CountSand(&part2Cave))
}
