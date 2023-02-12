package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now()
	// filePtr := flag.String("f", "input", "Input file if not 'input'")
	debugPtr := flag.Bool("debug", false, "Runs through debug input")

	flag.Parse()
	var readFile *os.File
	var err error
	if *debugPtr {
		readFile, err = os.Open("testinput")
	} else {
		readFile, err = os.Open("input")
	}

	if err != nil {
		fmt.Println("Fatal:", err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	lines := make([]string, 0)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	// for fileScanner.Scan() {
	// 	lines = append(lines, fileScanner.Bytes())
	// }

	m := ReadMap(lines)

	// fmt.Printf("Contents of m: %#v\n", m)

	// plant the monkey at the starting position according to
	// the rules of the challenge
	X, Y := FindStart(&m)
	monke := NewMonkey(X, Y, right, &m)

	// Perform the traveling
	Travel(&m, monke)

	// print the map (may not be worthwhile using the challenge data
	// since it's so damn huge)
	// if *debugPtr {
	// 	PrintMonkeyMap(&m, monke)
	// }
	// PrintMonkeyMap(&m, monke)

	// remember, rows and cols start with 1 because the creator of this
	// challenge is a sadist
	endRow, endCol := monke.Y+1, monke.X+1
	password := (1000 * endRow) + (4 * endCol) + monke.facing

	fmt.Printf("Part 1: The password is %d\n", password)

	// now let's start again for step 2
	nX, nY := FindStart(&m)
	part2Monke := NewMonkey(nX, nY, right, &m)

	var newCube Cube
	if *debugPtr {
		newCube = BuildDebugCube()
	} else {
		newCube = BuildCube()
	}
	breadcrumbs := TravelPart2(&m, &newCube, part2Monke, *debugPtr)

	if *debugPtr {
		PrintMonkeyMap(&m, part2Monke, breadcrumbs)
	}
	// PrintMonkeyMap(&m, part2Monke, breadcrumbs)
	part2endRow, part2endCol := part2Monke.Y+1, part2Monke.X+1
	part2Password := (1000 * part2endRow) + (4 * part2endCol) + part2Monke.facing
	fmt.Printf("Part 2: The password is %d\n", part2Password)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
