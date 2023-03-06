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
	// ingest lines
	v := NewValley(lines)
	agr0 := v.Entrance
	v.PrintValley(agr0)

	// options := v.GetSurroundingForElf(agr0)
	// fmt.Printf("Surrounding: %#v\n", options)

	// fmt.Printf("hold onto yer butts...\n")
	// res := Traverse(agr0, v)
	// fmt.Printf("Total Minutes: %d\n", res)

	// part 2, which should also include the answer to part 1:

	// set up the beginning state
	startState := NewState(agr0, v)

	// marching to the exit for the first leg...
	fmt.Printf("Marching to the first exit...\n")
	min1, secondState := TraverseAndReturn(startState, v.Entrance, v.Exit)
	fmt.Printf("Completed in %d minutes.\n", min1)
	secondState.v.PrintValley(secondState.loc)

	// marching back to the entrance to pick up those snacks...
	fmt.Printf("Heading back to the entrance...\n")
	min2, thirdState := TraverseAndReturn(secondState, v.Exit, v.Entrance)
	fmt.Printf("Completed in %d minutes.\n", min2)

	// now back to the exit again.
	fmt.Printf("Now heading back to the exit with the snacks.\n")
	final, _ := TraverseAndReturn(thirdState, v.Entrance, v.Exit)
	fmt.Printf("Last leg completed in %d minutes\n", final)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
