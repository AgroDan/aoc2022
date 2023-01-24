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

func ParseLine(line string) Point {
	coords := strings.Split(line, ",")
	convX, err := strconv.Atoi(coords[0])
	convY, err := strconv.Atoi(coords[1])
	convZ, err := strconv.Atoi(coords[2])
	if err != nil {
		fmt.Println("Could not convert")
	}
	return Point{
		X: convX,
		Y: convY,
		Z: convZ,
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

	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	// Insert code here
	knownPoints := make([]Point, 0)

	for _, line := range lines {
		knownPoints = append(knownPoints, ParseLine(line))
	}

	highestDimension := 0
	for _, v := range knownPoints {
		if v.X > highestDimension {
			highestDimension = v.X
		}
		if v.Y > highestDimension {
			highestDimension = v.Y
		}
		if v.Z > highestDimension {
			highestDimension = v.Z
		}
	}

	s := NewSpace(highestDimension + 1)

	for _, v := range knownPoints {
		s.area[v.X][v.Y][v.Z] = '#'
	}

	// now let's just count everything.
	totalSides := 0
	for _, v := range knownPoints {
		totalSides += s.CheckPoint(v, false)
	}
	fmt.Printf("Total sides: %d\n", totalSides)

	// removed below because it only applies to test data
	// test := s.DetectBubble(Point{
	// 	X: 2,
	// 	Y: 2,
	// 	Z: 5,
	// })

	// fmt.Printf("If this works, then this should only return one cube of air: %v\n", test)

	fmt.Printf("Now let's find the bubbles. First, detect all the air spaces around lava...")
	allAirSpaces := make([]Point, 0)
	for _, v := range knownPoints {
		airSpaces := s.GetEmptyEdges(v)
		allAirSpaces = append(allAirSpaces, airSpaces...)
	}
	fmt.Printf("Done!\n")

	fmt.Printf("Now use Breadth-first-search to find any bubbles surrounded by lava...")
	GetBubbles := make([]Point, 0)
	for _, v := range allAirSpaces {
		bubble := s.DetectBubble(v)
		if len(bubble) > 0 {
			GetBubbles = append(GetBubbles, bubble...)
		}
	}
	fmt.Printf("Done!\n")

	// now De-dupe the bubbles...
	fmt.Printf("De-duping bubbles...")
	dedupeBubbles := RemoveDuplicatePoints(GetBubbles)
	fmt.Printf("Done!\n")

	fmt.Printf("How many bubbles total then? %d\n", len(dedupeBubbles))

	fmt.Printf("Obtaining surface area from inner air bubbles...")
	sumOfInnerSurfaceArea := 0
	for _, v := range dedupeBubbles {
		sumOfInnerSurfaceArea += s.CheckPoint(v, true) // look for lava surfaces
	}
	fmt.Printf("Done! Found total bubble SA to be %d!\n", sumOfInnerSurfaceArea)

	fmt.Printf("Therefore, total surface area minus bubbles is: %d\n", totalSides-sumOfInnerSurfaceArea)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
