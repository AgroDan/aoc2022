package broken

/*
	The age-old problem with the humble engineer
	over-engineering something that doesn't

	friggen

	work
*/

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

/*
 * DEFINE CONSTANTS
 */

const ( // start position/end position
	Start = iota
	End
)

const ( // directions
	Up = iota
	Down
	Left
	Right
)

const Alphabet string = "abcdefghijklmnopqrstuvwxyz"

/*
 * DEFINE TERRAIN
 */

type Position struct {
	X int
	Y int
}

func GetDirection(direction int, pos Position) Position {
	// this function returns a position based on the provided direction.
	// NOTE: THIS DOES NOT CHECK IF WE HIT A BORDER!
	p := Position{}
	switch direction {
	case 0:
		// up or north
		p.X = pos.X
		p.Y = pos.Y - 1
	case 1:
		// down or south
		p.X = pos.X
		p.Y = pos.Y + 1
	case 2:
		// left or west
		p.X = pos.X - 1
		p.Y = pos.Y
	case 3:
		// right or east
		p.X = pos.X + 1
		p.Y = pos.Y
	}
	return p
}

type HillMap [][]string

func (h HillMap) ValidPosition(pos Position) bool {
	// returns true if the position provided
	// is actually on the map.

	// LEFT OFF HERE. WENT TO GET DRUNK.
	// THIS CODE IS LITERALLY DRIVING ME TO DRINK
	if pos.X < 0 {
		return false
	}
	if pos.Y < 0 {
		return false
	}
	if pos.X > h.Width() {
		return false
	}
	if pos.Y > h.Height() {
		return false
	}
	return true
}

// a space saver, maybe
func (h HillMap) Width() int {
	// returns the last index on the X Axis
	return len(h[0]) - 1
}
func (h HillMap) Height() int {
	return len(h) - 1
}

func (h HillMap) FindPoint(point int) Position {
	var letter string = ""
	if point == 0 {
		// Start
		letter = "S"
	} else if point == 1 {
		// End
		letter = "E"
	}
	for i := 0; i < len(h); i++ {
		// y is first
		for j := 0; j < len(h[i]); j++ {
			// this is x
			if h[i][j] == letter {
				p := Position{
					X: j,
					Y: i,
				}
				return p
			}
		}
	}

	// if you got here, not found!
	p := Position{
		X: -1,
		Y: -1,
	}
	return p
}

func (h HillMap) GetLetter(pos Position) (string, error) {
	// returns the letter of any arbitrary position on the map
	if !h.ValidPosition(pos) {
		return "", errors.New("invalid position")
	}
	return h[pos.Y][pos.X], nil
}

func Compare(a, b Position) bool {
	if a.X == b.X && a.Y == b.Y {
		return true
	} else {
		return false
	}
}

/*
 * DEFINE EXPLORER
 */

type Explorer struct {
	pos         Position
	breadcrumbs []Position
	h           HillMap
}

func (ex *Explorer) HaveWeBeenThere(pos Position) bool {
	// returns a True or False if we have previously been to this position.
	for _, v := range ex.breadcrumbs {
		if Compare(pos, v) {
			return true
		}
	}
	return false
}

func (ex *Explorer) CanWeMove(pos Position) bool {
	letter, err := ex.h.GetLetter(pos)
	if err != nil {
		// invalid direction, so no we can't move
		return false
	}
	if letter == "E" && ex.GetCurrentLetter() == "z" {
		return true
	} else if letter == "E" {
		return false
	}
	if letter == "S" {
		return false
	}

	// have we been there?
	if ex.HaveWeBeenThere(pos) {
		return false
	}

	currLetter := ex.GetCurrentLetter()

	letterNum := strings.Index(Alphabet, letter)
	currLetterNum := strings.Index(Alphabet, currLetter)

	if currLetter == letter {
		// the letter in this direction is equal
		// to the current direction

		// now let's consider less than -1 to be
		// not worth it
		// if currLetterNum-letterNum > 1 {
		// 	return false
		// }

		return true
	}

	if letterNum-currLetterNum == 1 {
		// the difference is one greater, which is valid
		return true
	}
	// otherwise, no sir
	return false
}

func (ex *Explorer) GetCurrentLetter() string {
	// returns the letter at the current position
	// on the map.
	return ex.h[ex.pos.Y][ex.pos.X] // this MAY need to be string-ified
}

func (ex *Explorer) GetPossibleDirections() ([]int, error) {
	// This function will check all four directions and
	// check if we can move there.
	// var n, s, e, w Position
	var PossibleDirections = []int{
		Up,
		Down,
		Left,
		Right,
	}

	var direction []int

	for _, dir := range PossibleDirections {
		// Get the position first
		dirPos := GetDirection(dir, ex.pos)

		// Now, is it valid? (meaning at a border or not)
		if ex.h.ValidPosition(dirPos) {
			// yes, it's valid. Now can we move there?
			if ex.CanWeMove(dirPos) {
				// we can? Then this is a possible direction.
				direction = append(direction, dir)
			}
		}
	}

	if len(direction) > 0 {
		return direction, nil
	} else {
		return direction, errors.New("no possible direction to move")
	}
}

func (ex *Explorer) Move(pos Position) error {
	// This ONLY checks if we are moving to a valid coordinate.
	if ex.h.ValidPosition(pos) {
		ex.breadcrumbs = append(ex.breadcrumbs, ex.pos)
		ex.pos = pos
		return nil
	}
	return errors.New("invalid map coordinate entered")
}

func (ex *Explorer) AreWeThereYet() bool {
	// simply returns a true or false if we finished the map
	if ex.GetCurrentLetter() == "E" {
		return true
	} else {
		return false
	}
}

func parseMap(lines []string) HillMap {
	// This will parse the data line by line
	h := make([][]string, 0)
	for _, v := range lines {
		lineSlice := make([]string, 0)
		for _, c := range v {
			lineSlice = append(lineSlice, string(c))
		}
		h = append(h, lineSlice)
	}
	return h
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
	hm := parseMap(lines)

	// fmt.Println(hm)

	// start := h.FindPoint(Start)
	// end := h.FindPoint(End)
	// fmt.Printf("Start: +%v, End: +%v\n", start, end)

	// for i := 0; i < ((h.Height()+1)*(h.Width()+1)); i++ {

	// }

	lowest := ((hm.Height() + 1) * (hm.Width() + 1))

	// for i := 0; i < 1000; i++ {
	for {
		startingPoint := hm.FindPoint(Start)
		// fmt.Printf("Starting point X:%d Y:%d\n", startingPoint.X, startingPoint.Y)
		agr0 := Explorer{
			pos: startingPoint,
			h:   hm,
		}
		for {
			if agr0.AreWeThereYet() {
				if len(agr0.breadcrumbs) < lowest {
					lowest = len(agr0.breadcrumbs)
				}
				fmt.Printf("Current Quickest Route: %d\r", lowest)
				// pp.Print(agr0.breadcrumbs)

			}
			directions, err := agr0.GetPossibleDirections()
			if err != nil {
				// we are stuck, break and try again
				// fmt.Printf("Stuck at X:%d, Y:%d\n", agr0.pos.X, agr0.pos.Y)
				break
			}
			randDir := rand.Intn(len(directions))
			agr0.Move(GetDirection(directions[randDir], agr0.pos))
		}
	}

	// practice!
	// startingPoint := hm.FindPoint(Start)
	// fmt.Printf("Starting point X:%d Y:%d\n", startingPoint.X, startingPoint.Y)
	// agr0 := Explorer{
	// 	pos: startingPoint,
	// 	h:   hm,
	// }

	// for {
	// 	fmt.Printf("Current Pos: X:%d, Y:%d\n", agr0.pos.X, agr0.pos.Y)
	// 	if agr0.AreWeThereYet() {
	// 		fmt.Printf("Got there in %d moves!\n", len(agr0.breadcrumbs))
	// 	}
	// 	directions, err := agr0.GetPossibleDirections()
	// 	if err != nil {
	// 		fmt.Printf("Stuck at X:%d, Y:%d\n", agr0.pos.X, agr0.pos.Y)
	// 		break
	// 	}
	// 	randDir := rand.Intn(len(directions))
	// 	switch directions[randDir] {
	// 	case 0:
	// 		fmt.Printf("Going ^\n")
	// 	case 1:
	// 		fmt.Printf("Going v\n")
	// 	case 2:
	// 		fmt.Printf("Going <\n")
	// 	case 3:
	// 		fmt.Printf("Going >\n")
	// 	}
	// 	agr0.Move(GetDirection(directions[randDir], agr0.pos))
	// }
	// fmt.Println()
}
