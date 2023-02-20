package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

// func ParseMap(input [][]byte) GroveMap {

// 	g := NewGroveMap(len(input))
// 	for i, v := range input {
// 		for j, w := range v {
// 			g.p[i][j] = rune(w)
// 		}
// 	}
// 	return g
// }

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
	eList := NewElfStack()
	for r, line := range lines {
		for c, char := range line {
			if char == '#' {
				err := eList.Add(r, c)
				if err != nil {
					fmt.Printf("Something happened\n")
				}
			}
		}
	}

	// at this point we have all of our elves.
	// fmt.Printf("%#v\n", eList)

	fmt.Printf("Starting map:\n")
	eList.PrintElves()

	/*

		AT THIS POINT, we are going to go through a single step.

	*/

	// now let's set proposals.

	dirList := []string{"N", "S", "W", "E"}
	props := NewProposals()

	for i := 0; i < 10; i++ {
		fmt.Printf("\n")

		fmt.Printf("Iteration %d\nStarting dir: %s\n\n", i+1, dirList[0])

		for k := range eList.set {
			// for each elf
			// propose the direction
			row, col, err := ExtractCoords(k)
			if err != nil {
				fmt.Printf("Could not extract coords! %s\n", err)
				continue
			}

			// is elf alone?
			if eList.IsAlone(row, col) {
				// can't move, skip further checks
				// fmt.Printf("Elf is alone at %s", k)
				continue
			}

			// now set up the proposals
			// fmt.Printf("Dirlist: %+v\n", dirList)
			for _, v := range dirList { // the list of directions we check
				check := CheckDirection(v)
				isOccupied := false
				for _, checkDir := range check { // the 3 squares we check in this specific dir
					// for each direction
					// if there is anyone in either of these
					// directions,  break and check the next
					// direction
					diffR, diffC := GetDirection(checkDir)
					if eList.Exists(row+diffR, col+diffC) {
						// fmt.Printf("Found occupied\n")
						isOccupied = true
						continue
					}
				}

				if !isOccupied {
					// fmt.Printf("Proposed movement\n")
					// if not, then propose this direction
					diffR, diffC := GetDirection(v)
					newR := diffR + row
					newC := diffC + col
					props.DeclareProposal(newR, newC)
					eList.SubmitProposal(row, col, newR, newC)
					break
				}
			}
		}

		// fmt.Printf("%#v\n", eList)
		// fmt.Printf("Proposals: %#v\n", props)
		// fmt.Printf("All proposals: %#v\n", props)
		// now all of our proposals are submitted. Let's move if we can.
		for key, value := range eList.proposed {
			// let's see if this proposal is valid
			fromRow, fromCol, err := ExtractCoords(key)
			if err != nil {
				fmt.Printf("Could not extract FROM coords: %s\n", err)
			}

			toRow, toCol, err := ExtractCoords(value)
			if err != nil {
				fmt.Printf("Could not extract TO coords: %s\n", err)
			}

			if props.ValidProposal(toRow, toCol) {
				// if this is valid, then we move!
				// fmt.Printf("Valid proposal, will now move\n")
				err := eList.Move(fromRow, fromCol, toRow, toCol)
				if err != nil {
					fmt.Printf("Could not move properly: %s\n", err)
				}
			} else {
				// fmt.Printf("Invalid proposal! Row: %d, Col %d\n", toRow, toCol)
			}
		}

		// all proposals done, let's cycle
		eList.ClearProposed()
		dirList = RotateQueue(dirList)
		eList.PrintElves()
		props.ClearProposals()
	}
	// print second iter
	// fmt.Printf("%#v\n", eList)

	fmt.Printf("Total amount of empty space: %d\n", eList.GetEmptySpace())

	// Now to simulate how far we go, I'm going to just copy-paste the above with
	// a for loop checking to see when we all stop moving.

	part2eList := NewElfStack()
	for r, line := range lines {
		for c, char := range line {
			if char == '#' {
				err := part2eList.Add(r, c)
				if err != nil {
					fmt.Printf("Something happened\n")
				}
			}
		}
	}
	part2dirList := []string{"N", "S", "W", "E"}
	part2props := NewProposals()
	part2iter := 0

	for {
		// fmt.Printf("Iteration %d\nStarting dir: %s\n\n", i+1, dirList[0])
		part2iter++

		for k := range part2eList.set {
			// for each elf
			// propose the direction
			row, col, err := ExtractCoords(k)
			if err != nil {
				fmt.Printf("Could not extract coords! %s\n", err)
				continue
			}

			// is elf alone?
			if part2eList.IsAlone(row, col) {
				// can't move, skip further checks
				// fmt.Printf("Elf is alone at %s", k)
				continue
			}

			// now set up the proposals
			// fmt.Printf("Dirlist: %+v\n", dirList)
			for _, v := range part2dirList { // the list of directions we check
				check := CheckDirection(v)
				isOccupied := false
				for _, checkDir := range check { // the 3 squares we check in this specific dir
					// for each direction
					// if there is anyone in either of these
					// directions,  break and check the next
					// direction
					diffR, diffC := GetDirection(checkDir)
					if part2eList.Exists(row+diffR, col+diffC) {
						// fmt.Printf("Found occupied\n")
						isOccupied = true
						continue
					}
				}

				if !isOccupied {
					// fmt.Printf("Proposed movement\n")
					// if not, then propose this direction
					diffR, diffC := GetDirection(v)
					newR := diffR + row
					newC := diffC + col
					part2props.DeclareProposal(newR, newC)
					part2eList.SubmitProposal(row, col, newR, newC)
					break
				}
			}
		}

		// fmt.Printf("%#v\n", eList)
		// fmt.Printf("Proposals: %#v\n", props)
		// fmt.Printf("All proposals: %#v\n", props)
		// now all of our proposals are submitted. Let's move if we can.

		// lets see if this works
		if len(part2eList.proposed) <= 0 {
			break
		}
		for key, value := range part2eList.proposed {
			// let's see if this proposal is valid
			fromRow, fromCol, err := ExtractCoords(key)
			if err != nil {
				fmt.Printf("Could not extract FROM coords: %s\n", err)
			}

			toRow, toCol, err := ExtractCoords(value)
			if err != nil {
				fmt.Printf("Could not extract TO coords: %s\n", err)
			}

			if part2props.ValidProposal(toRow, toCol) {
				// if this is valid, then we move!
				// fmt.Printf("Valid proposal, will now move\n")
				err := part2eList.Move(fromRow, fromCol, toRow, toCol)
				if err != nil {
					fmt.Printf("Could not move properly: %s\n", err)
				}
			} else {
				// fmt.Printf("Invalid proposal! Row: %d, Col %d\n", toRow, toCol)
			}
		}

		// all proposals done, let's cycle
		part2eList.ClearProposed()
		part2dirList = RotateQueue(part2dirList)
		// part2eList.PrintElves()
		part2props.ClearProposals()
	}

	fmt.Printf("Finished iterating after %d cycles.\n", part2iter)
	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
