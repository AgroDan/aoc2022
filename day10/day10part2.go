package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func shouldWeDraw(X int, pos int) bool {
	// will return whether or not we draw a pixel
	xf := float64(X)
	posf := float64(pos)
	if math.Abs(posf-xf) <= 1 {
		return true
	} else {
		return false
	}
}

// need a "Cycle CRT instruction function" which draws or not on screen.

func parseInstruction(X, C *int, stillRunning *bool, col *int, inst string) {
	// note: addx instruction takes 2 cycles.
	//       noop takes 1 cycle, does nothing to X

	// read through to cycle 10! maybe even print out values for 20 cycles

	*C++ // increment the cycle regardless
	// Right here, determine if we should print.
	if *col == 40 {
		fmt.Printf("\n")
		*col = 0
	}
	if shouldWeDraw(*X, *col) {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
	*col++

	if *stillRunning {
		s := strings.Split(inst, " ")
		num, err := strconv.Atoi(s[1])
		if err != nil {
			panic("Could not read integer!")
		}
		*X += num
		*stillRunning = false
	} else if strings.HasPrefix(inst, "addx") {
		*stillRunning = true
	}

}

func main() {
	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	readFile.Close()

	// For this, we will set two variables. X for the x register, an C for the cycle count.
	var X int = 1
	var C int = 0
	var stillRunning bool = false
	var col int = 0
	for _, v := range lines {
		for {
			// fmt.Printf("Cycle %d: Reg X: %d, inst: %s\n", C, X, v)
			parseInstruction(&X, &C, &stillRunning, &col, v)
			if !stillRunning {
				break

			}
		}

		// if i >= 20 {
		// 	if counter == 40 {
		// 		fmt.Printf("Cycle %d: Signal Strength: %d\n", i, X*C)
		// 		counter = 0
		// 	}
		// 	counter++
		// }
		// if C >= 15 {
		// 	break
		// }
	}
}
