package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

/*
 * DEFINE CONSTANTS
 */

// behold, Golang's answer to enum datatypes!

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

const Alphabet string = "SabcdefghijklmnopqrstuvwxyzE"

/*
 * DEFINE BREADCRUMBS
 */

type Breadcrumbs struct {
	crumbs []Position
}

func (b Breadcrumbs) Exists(pos Position) bool {
	// checks to see if a breadcrumb already exists
	// in the set
	if len(b.crumbs) == 0 {
		return false
	}
	for _, v := range b.crumbs {
		if Compare(v, pos) {
			return true
		}
	}
	return false
}

func (b *Breadcrumbs) Add(pos Position) error {
	// all positions in the breadcrumb trail MUST BE UNIQUE!
	if b.Exists(pos) {
		return errors.New("breadcrumb already exists in set")
	}
	b.crumbs = append(b.crumbs, pos)
	return nil
}

/*
 * DEFINE TERRAIN
 */

type Position struct {
	X int
	Y int
}

func Compare(a, b Position) bool {
	if a.X == b.X && a.Y == b.Y {
		return true
	} else {
		return false
	}
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
	if pos.Y > h.Height() {
		return false
	}
	if pos.X > h.Width() {
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

/*
 * DEFINE EXPLORER
 */

type Explorer struct {
	pos      Position
	distance int
}

func NewExplorer(currPos Position, d int) Explorer {
	return Explorer{
		pos:      currPos,
		distance: d,
	}
}

// func (ex *Explorer) HaveWeBeenThere(pos Position) bool {
// 	// returns a True or False if we have previously been to this position.
// 	for _, v := range ex.breadcrumbs {
// 		if Compare(pos, v) {
// 			return true
// 		}
// 	}
// 	return false
// }

func (ex *Explorer) IncrementDistance() {
	// increments the distanct of the provided explorer by one
	ex.distance++
}

func (ex *Explorer) CanWeMove(pos Position, h *HillMap, b Breadcrumbs) bool {
	letter, err := h.GetLetter(pos)
	if err != nil {
		// invalid direction, so no we can't move
		return false
	}
	// if letter == "E" && ex.GetCurrentLetter(h) == "z" {
	// 	return true
	// } else if letter == "E" {
	// 	return false
	// }
	// if letter == "S" {
	// 	return false
	// }

	// have we been there?
	if b.Exists(pos) {
		fmt.Printf("Sorry, we've already been to location X:%d Y:%d!\n", pos.X, pos.Y)
		return false
	}

	currLetter := ex.GetCurrentLetter(h)

	letterNum := strings.Index(Alphabet, letter)
	currLetterNum := strings.Index(Alphabet, currLetter)

	fmt.Printf("Checking current letter: %s with next letter: %s\n", currLetter, letter)

	if currLetterNum >= letterNum {
		// potential letter is at least equal to the current letter
		fmt.Printf("Current letter %s is greater than or equal to %s!\n", currLetter, letter)
		return true
	}

	if letterNum-currLetterNum == 1 {
		// the difference is one greater, which is valid
		fmt.Printf("Current letter %s is 1 less than %s!\n", currLetter, letter)
		return true
	}
	// otherwise, no sir
	return false
}

func (ex *Explorer) GetCurrentLetter(h *HillMap) string {
	// returns the letter at the current position
	// on the map.
	return (*h)[ex.pos.Y][ex.pos.X] // this MAY need to be string-ified
}

func (ex *Explorer) GetPossibleDirections(h *HillMap, b Breadcrumbs) ([]int, error) {
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
		if h.ValidPosition(dirPos) {
			// yes, it's valid. Now can we move there?
			if ex.CanWeMove(dirPos, h, b) {
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

// this function is useless now because we're not going to "move" an
// explorer per se, but rather create a new explorer and increment the
// distance
//
// func (ex *Explorer) Move(pos Position) error {
// 	// This ONLY checks if we are moving to a valid coordinate.
// 	if ex.h.ValidPosition(pos) {
// 		ex.breadcrumbs = append(ex.breadcrumbs, ex.pos)
// 		ex.pos = pos
// 		return nil
// 	}
// 	return errors.New("invalid map coordinate entered")
// }

func (ex *Explorer) AreWeThereYet(h *HillMap) bool {
	// simply returns a true or false if we finished the map
	if ex.GetCurrentLetter(h) == "E" {
		return true
	} else {
		return false
	}
}

/*
 * DEFINE EXPLORER QUEUE
 */
// var ExplorerQueue []*Explorer
type ExplorerQueue struct {
	// This will be a FIFO queue
	queue []*Explorer
}

// func (e *ExplorerQueue) Pop() (Explorer, error) {
// 	// this one pops off the right
// 	if len(e.queue) == 0 {
// 		return Explorer{}, errors.New("zero length explorer queue")
// 	}

// 	index := len(e.queue) - 1
// 	var ne Explorer = *(e.queue)[index]
// 	e.queue = e.queue[:index]
// 	return ne, nil
// }

func (e *ExplorerQueue) Pop() (Explorer, error) {
	// this function will return the first entrant in the
	// explorerQueue and remove it from the slice.

	// first let's see if this is empty
	if len(e.queue) == 0 {
		return Explorer{}, errors.New("zero length explorer queue")
	}
	var ne Explorer = *e.queue[0] // next explorer
	e.queue = e.queue[1:]

	return ne, nil
}

func (e *ExplorerQueue) Push(new *Explorer) {
	// pushes onto the queue
	// NOTE! it is the caller's responsibility to
	// increment the Explorer distance
	e.queue = append(e.queue, new)
}

/*
 * GET SHORTEST ROUTE
 */

func GetShortestRoute(h *HillMap) int {
	// This should just put it all together.

	// distance Results
	var distanceResults []int

	startingPoint := h.FindPoint(Start)
	q := ExplorerQueue{}
	b := Breadcrumbs{}

	// Hello there, little explorer!
	agr0 := NewExplorer(startingPoint, 0)
	fmt.Printf("Starting Point: X:%d Y:%d\n", startingPoint.X, startingPoint.Y)

	// Now let's prime the queue!
	q.Push(&agr0)
	counter := 0
	for {
		// Iterate through the queue.
		myExplorer, err := q.Pop()
		if err != nil {
			// nothing left in the queue. Break?
			fmt.Printf("Nothing left in the queue!\n")
			break
		}
		// fmt.Printf("Queue size: %d\r", len(q.queue))

		if b.Exists(myExplorer.pos) {
			// we've been here already
			continue
		} else {
			// lay the breadcrumb
			b.Add(myExplorer.pos)
		}

		// did we make it to the destination?
		if myExplorer.AreWeThereYet(h) {
			// we're there, first to the post
			// return myExplorer.distance
			fmt.Printf("We made it after %d loops!\n", counter)
			distanceResults = append(distanceResults, myExplorer.distance)
			continue
		}

		directions, err := myExplorer.GetPossibleDirections(h, b)
		if err != nil {
			// no possible directions
			continue
		}
		for _, v := range directions {
			// given the directions, time to make some explorers

			// first get the position of the given direction
			dir := GetDirection(v, myExplorer.pos)

			// start up a new explorer with a new distance
			ne := NewExplorer(dir, myExplorer.distance+1)

			// // also, add to that explorer's breadcrumbs
			// ne.bc.Add(dir)

			// push this new explorer onto the queue
			q.Push(&ne)
		}
		counter++
	}
	// fmt.Printf("Queue: +%v, +%v, +%v\n", *q.queue[0], *q.queue[1], *q.queue[2])

	fmt.Printf("Looped %d times.\n", counter)
	fmt.Printf("Results: +%v\n", distanceResults)
	var lowest int
	if len(distanceResults) == 0 {
		return 0
	} else {
		lowest = distanceResults[0]
		for _, v := range distanceResults {
			if v < lowest {
				lowest = v
			}

		}
	}
	return lowest
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

	fmt.Printf("HillMap stats: Width: %d, Height: %d\n", hm.Width(), hm.Height())

	routeNum := GetShortestRoute(&hm)

	fmt.Printf("Please for the love of god let this work: %d\n", routeNum)

	// fmt.Println(hm)

	// start := h.FindPoint(Start)
	// end := h.FindPoint(End)
	// fmt.Printf("Start: +%v, End: +%v\n", start, end)

	// for i := 0; i < ((h.Height()+1)*(h.Width()+1)); i++ {

	// }

	// lowest := ((hm.Height() + 1) * (hm.Width() + 1))

	// // for i := 0; i < 1000; i++ {
	// for {
	// 	startingPoint := hm.FindPoint(Start)
	// 	// fmt.Printf("Starting point X:%d Y:%d\n", startingPoint.X, startingPoint.Y)
	// 	agr0 := Explorer{
	// 		pos: startingPoint,
	// 		h:   &hm,
	// 	}
	// 	for {
	// 		if agr0.AreWeThereYet() {
	// 			if len(agr0.breadcrumbs) < lowest {
	// 				lowest = len(agr0.breadcrumbs)
	// 			}
	// 			fmt.Printf("Current Quickest Route: %d\r", lowest)
	// 			// pp.Print(agr0.breadcrumbs)

	// 		}
	// 		directions, err := agr0.GetPossibleDirections()
	// 		if err != nil {
	// 			// we are stuck, break and try again
	// 			// fmt.Printf("Stuck at X:%d, Y:%d\n", agr0.pos.X, agr0.pos.Y)
	// 			break
	// 		}
	// 		randDir := rand.Intn(len(directions))
	// 		agr0.Move(GetDirection(directions[randDir], agr0.pos))
	// 	}
	// }

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
