package FUCKYOU

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

/*
 * HERE BE CONSTANTS
 */

const Alpha string = "abcdefghijklmnopqrstuvwxyz"

const (
	Start = iota
	End
)

const (
	North = iota
	South
	East
	West
)

/*
 * HERE BE STRUCTS
 */

type Queue struct {
	q []Explorer
}

type Explorer struct {
	pos      Coordinate
	distance int
}

type Coordinate struct {
	row int
	col int
}

type Breadcrumb struct {
	crumb []Coordinate
}

/*
 * HERE BE STRUCT FUNCTIONS
 */

func (q *Queue) Push(e Explorer) {
	// adds an explorer to the queue
	q.q = append(q.q, e)
}

func (q *Queue) Pop() (Explorer, error) {
	if len(q.q) == 0 {
		return Explorer{}, errors.New("empty queue")
	}
	ne := q.q[0]
	q.q = q.q[1:]
	return ne, nil
}

func NewExplorer(c Coordinate, d int) Explorer {
	return Explorer{
		pos:      c,
		distance: d,
	}
}

func (b *Breadcrumb) Add(c Coordinate) {
	// adds a breadcrumb to the list
	b.crumb = append(b.crumb, c)
}

func (b Breadcrumb) Exists(c Coordinate) bool {
	// checks to see if the supplied coordinate
	// is in the list of breadcrumbs.
	for _, v := range b.crumb {
		if v.row == c.row && v.col == c.col {
			return true
		}
	}
	return false
}

/*
 * HERE BE GENERIC FUNCTIONS
 */

func GetNumberValue(letter string) int {
	if letter == "S" {
		return 0
	} else if letter == "E" {
		return 25
	} else {
		return strings.Index(Alpha, letter)
	}
}

// func BuildMap(lines []string) [][]int {
// 	HillMap := make([][]int, 0)
// 	for _, i := range lines {
// 		col := make([]int, 0)
// 		for _, j := range i {
// 			numVal := GetNumberValue(string(j))
// 			col = append(col, numVal)
// 		}
// 		HillMap = append(HillMap, col)
// 	}
// 	return HillMap
// }

// going to take a less elegant approach for science
func BuildMap(lines []string) (HillMap [][]int, Start Coordinate, End Coordinate) {
	HillMap = make([][]int, 0)
	Start = Coordinate{}
	End = Coordinate{}
	for i_idx, i := range lines {
		col := make([]int, 0)
		for j_idx, j := range i {
			if j == 'S' {
				Start = Coordinate{
					col: j_idx,
					row: i_idx,
				}
			} else if j == 'E' {
				End = Coordinate{
					col: j_idx,
					row: i_idx,
				}
			}
			numVal := GetNumberValue(string(j))
			col = append(col, numVal)
		}
		HillMap = append(HillMap, col)
	}
	return HillMap, Start, End
}

func GetPoint(which int, HillMap [][]int) Coordinate {
	var find int
	coord := Coordinate{}
	switch which {
	case 0:
		find = GetNumberValue("a")
	case 1:
		find = GetNumberValue("z")
	default:
		coord.row = -1
		coord.col = -1
		return coord
	}

	for r_idx, i := range HillMap {
		for c_idx, j := range i {
			if j == find {
				coord.row = r_idx
				coord.col = c_idx
				return coord
			}
		}
	}

	coord.row, coord.col = -1, -1
	return coord
}

func ValidCoord(checkNum int, c Coordinate, b Breadcrumb, h [][]int) bool {
	// this will check if the provided coordinate has either
	// never been visited before, and is within the bounds of
	// the provided challenge. Which is to say no higher than
	// +1 than the current coordinate. The provided integer
	// checkNum is the value of the original coordinate, NOT
	// the provided coordinate. This only returns a true or
	// false based on whether or not the provided coordinate
	// is worth adding to the queue.
	if b.Exists(c) {
		// we've been here before. Can't go there.
		return false
	}

	targetNum := h[c.row][c.col]

	if targetNum <= (checkNum + 1) {
		return true
	} else {
		return false
	}
}

func GetDirection(direction int, c Coordinate, b Breadcrumb, h [][]int) (Coordinate, error) {
	// given a row and column as well as a direction, this will
	// return a row and column coordinate in that direction. An
	// error will be thrown if the direction is off the map.

	coord := Coordinate{}
	switch direction {
	case 0:
		// north
		coord.row = c.row + 1
		coord.col = c.col
	case 1:
		// south
		coord.row = c.row - 1
		coord.col = c.col
	case 2:
		// east
		coord.row = c.row
		coord.col = c.col + 1
	case 3:
		// west
		coord.row = c.row
		coord.col = c.col - 1
	default:
		return Coordinate{}, errors.New("invalid direction")
	}

	// now let's do some bounds checking
	if coord.row < 0 {
		return coord, errors.New("out of bounds, too far south")
	}
	if coord.col < 0 {
		// negative numbers are out of bounds
		return coord, errors.New("out of bounds, too far west")
	}
	if coord.row >= len(h) {
		return coord, errors.New("out of bounds, too far north")
	}
	if coord.col >= len(h[c.row]) {
		return coord, errors.New("out of bounds, too far east")
	}

	// now let's see if this coordinate is worth looking at
	checkNum := h[c.row][c.col]

	if !ValidCoord(checkNum, coord, b, h) {
		return coord, errors.New("invalid coordinate due to test constraints")
	}

	return coord, nil
}

func getLowestPath(h [][]int, Start, End Coordinate) int {
	// this will span out and find the lowest path
	// s := GetPoint(Start, h)

	// start the breadcrumbs
	b := Breadcrumb{}

	// get the end coords
	// end := GetPoint(End, h)

	// direction array
	dir := [4]int{North, South, East, West}

	// hello explorer!
	agr0 := NewExplorer(Start, 0)

	q := Queue{}

	q.Push(agr0)
	b.Add(Start)

	for {
		e, err := q.Pop()
		if err != nil {
			// empty queue, drop out
			break
		}

		if e.pos.col == End.col && e.pos.row == End.row {
			// this is it!
			return e.distance
		}

		for _, v := range dir {
			c, err := GetDirection(v, e.pos, b, h)
			if err != nil {
				// invalid dir, next!
				continue
			}

			// if you got here, valid direction. Create a new explorer and push it onto the queue.
			b.Add(c)
			newExp := NewExplorer(c, e.distance+1)
			q.Push(newExp)
		}
	}

	return 0 // if you got here, things didn't work
}

func main() {
	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	hm, Start, End := BuildMap(lines)
	// // testing...
	// for _, i := range hm {
	// 	for _, j := range i {
	// 		fmt.Printf("%02d ", j)
	// 	}
	// 	fmt.Printf("\n")
	// }

	// did we do it?
	result := getLowestPath(hm, Start, End)

	fmt.Printf("Current lowest path: %d\n", result)

	// row, col := GetPoint(Start, hm)
	// fmt.Printf("Starting points: %d %d\n", row, col)

	// row, col = GetPoint(End, hm)
	// fmt.Printf("Ending points: %d %d\n", row, col)

}
