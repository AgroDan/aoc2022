package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ParseBluePrint(line string) Blueprint {
	words := strings.Split(line, " ")
	index := words[1]
	index = strings.TrimSuffix(index, ":")

	indexInt, err := strconv.Atoi(index)
	if err != nil {
		panic(fmt.Sprintf("Error in conversion: %s\n", err))
	}

	// TODO: Find out how to make this code not look like dogwater

	oreRobotCost, err := strconv.Atoi(words[6])
	if err != nil {
		panic(fmt.Sprintf("Error in conversion: %s\n", err))
	}

	clayRobotCost, err := strconv.Atoi(words[12])
	if err != nil {
		panic(fmt.Sprintf("Error in conversion: %s\n", err))
	}
	obsRobotCostOre, err := strconv.Atoi(words[18])
	if err != nil {
		panic(fmt.Sprintf("Error in conversion: %s\n", err))
	}
	obsRobotCostClay, err := strconv.Atoi(words[21])
	if err != nil {
		panic(fmt.Sprintf("Error in conversion: %s\n", err))
	}
	geoRobotCostOre, err := strconv.Atoi(words[27])
	if err != nil {
		panic(fmt.Sprintf("Error in conversion: %s\n", err))
	}
	geoRobotCostObs, err := strconv.Atoi(words[30])
	if err != nil {
		panic(fmt.Sprintf("Error in conversion: %s\n", err))
	}

	OreRobot, err := NewRobot(oreRobotCost, 0, 0, "ore")
	if err != nil {
		panic(fmt.Sprintf("Error in robot creation: %s\n", err))
	}
	ClayRobot, err := NewRobot(clayRobotCost, 0, 0, "clay")
	if err != nil {
		panic(fmt.Sprintf("Error in robot creation: %s\n", err))
	}
	ObsidianRobot, err := NewRobot(obsRobotCostOre, obsRobotCostClay, 0, "obsidian")
	if err != nil {
		panic(fmt.Sprintf("Error in robot creation: %s\n", err))
	}
	GeodeRobot, err := NewRobot(geoRobotCostOre, 0, geoRobotCostObs, "geode")

	return Blueprint{
		Index:    indexInt,
		Ore:      OreRobot,
		Clay:     ClayRobot,
		Obsidian: ObsidianRobot,
		Geode:    GeodeRobot,
	}
}

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")
	timePtr := flag.Int("t", 24, "Amount of minutes to cycle through per blueprint")
	whichPtr := flag.Int("b", 0, "Which blueprint to work with. If 0, all of them")
	part2Ptr := flag.Bool("p", false, "Execute part 2 instead of part 1")
	debugPtr := flag.Bool("debug", false, "Print debug statements")

	flag.Parse()
	readFile, err := os.Open(*filePtr)

	if err != nil {
		fmt.Println("Fatal:", err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	// Insert code here

	if !*part2Ptr {
		if *whichPtr == 0 {
			qualityLevels := 0
			for _, line := range lines {
				b := ParseBluePrint(line)
				if *debugPtr {
					fmt.Printf("%s\n", b)
				}
				geodeAmt := TestBlueprint(b, *timePtr)
				if *debugPtr {
					fmt.Printf("Found geodes: %d\n", geodeAmt)
				}
				qualityLevels += (geodeAmt * b.Index)
			}
			fmt.Printf("Total Quality: %d\n", qualityLevels)
		} else {
			b := ParseBluePrint(lines[*whichPtr-1])
			if *debugPtr {
				fmt.Printf("%s", b)
			}

			geodeAmt := TestBlueprint(b, *timePtr)

			fmt.Printf("Total Amount of Geodes found: %#v\n", geodeAmt)
		}
	} else {
		// just get the first 3 lines
		maxGeodes := make([]int, 0)
		for _, line := range lines[:3] {
			b := ParseBluePrint(line)
			if *debugPtr {
				fmt.Printf("%s\n", b)
			}
			maxGeodes = append(maxGeodes, TestBlueprint(b, *timePtr))
			if *debugPtr {
				fmt.Printf("Max Geodes for Blueprint %d: %d\n", b.Index, maxGeodes[len(maxGeodes)-1])
			}
		}
		fmt.Printf("3 parts multiplied: %d\n", maxGeodes[0]*maxGeodes[1]*maxGeodes[2])
	}

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
