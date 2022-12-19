package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInstruction(X, C *int, stillRunning *bool, inst string) {
	// note: addx instruction takes 2 cycles.
	//       noop takes 1 cycle, does nothing to X
	*C++ // increment the cycle regardless
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
	var C int = 1
	var stillRunning bool = false
	var counter int = 40
	var total int = 0
	for _, v := range lines {
		for {
			//fmt.Printf("Cycle %d: Reg X: %d, Signal Strength: %d, inst: %s\n", C, X, X*C, v)

			if C >= 20 {
				if counter == 40 {
					total += X * C
					fmt.Printf("Cycle %d: Reg X: %d, Signal Strength: %d\n", C, X, X*C)
					counter = 0
				}
				counter++
			}
			parseInstruction(&X, &C, &stillRunning, v)
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
	}
	fmt.Println("Total signal strengths:", total)
}
