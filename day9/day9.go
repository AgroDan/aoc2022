package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type RopeEnd struct {
	x, y int
}

func newRopeEnd() RopeEnd {
	var r = RopeEnd{
		x: 0,
		y: 0,
	}
	return r
}

func (r *RopeEnd) MoveRope(direction rune) {
	// this function moves the rope end in ONE direction.
	// Meaning, given the direction, it will increase or
	// decrease the respective offset by one ONLY. This will
	// NOT consider diagonal, for that you should call this
	// function twice.
	switch direction {
	case 'U':
		// going up
		r.y++
	case 'D':
		// going down
		r.y--
	case 'L':
		// going left
		r.x--
	case 'R':
		// going right
		r.x++
	}
}

func DetermineTailDirection(h, t *RopeEnd) error {
	// This takes the head and tail's current position
	// and updates the tail's position.

	// first, are we occupying the same space?
	if h.x == t.x && h.y == t.y {
		return errors.New("Same")
	}
	// are we adjacent?
	// To determine if we are adjacent, we need to see if the xOffsets
	// AND the yOffsets are only 1 away!
	hfx := float64(h.x)
	hfy := float64(h.y)
	tfx := float64(t.x)
	tfy := float64(t.y)

	if math.Abs(hfx-tfx) == 1 && math.Abs(hfy-tfy) <= 1 {
		return errors.New("Adjacent")
	}
	if math.Abs(hfy-tfy) == 1 && math.Abs(hfx-tfx) <= 1 {
		return errors.New("Adjacent")
	}

	// Find out if we even have to go diagonal.
	// Are we in the same column OR row?
	if h.x == t.x || h.y == t.y {
		// we do not need to go diagonal
		if h.x > t.x+1 {
			// move 1 column right
			t.MoveRope('R')
		} else if h.x < t.x-1 {
			// move 1 column left
			t.MoveRope('L')
		} else if h.y > t.y+1 {
			// move 1 row up
			t.MoveRope('U')
		} else if h.y < t.y-1 {
			// move 1 row down
			t.MoveRope('D')
		} else {
			// this is redundant but whatever
			return errors.New("Same")
		}
	} else {
		// now we are moving diagonal. Check if up and right
		if h.x > t.x+1 {
			// we need to move right and...
			t.MoveRope('R')
			if h.y > t.y {
				t.MoveRope('U')
			} else if h.y < t.y {
				t.MoveRope('D')
			} else {
				return errors.New("Weird logic? Moved right but not up or down")
			}
		} else if h.x < t.x-1 {
			// we need to move left and...
			t.MoveRope('L')
			if h.y > t.y {
				t.MoveRope('U')
			} else if h.y < t.y {
				t.MoveRope('D')
			} else {
				return errors.New("Weird logic? Moved left but not up or down")
			}
		} else if h.y > t.y+1 {
			// we need to move up and...
			t.MoveRope('U')
			if h.x > t.x {
				t.MoveRope('R')
			} else if h.x < t.x {
				t.MoveRope('L')
			} else {
				return errors.New("Weird logic? Moved up but not right or left")
			}
		} else if h.y < t.y-1 {
			// we need to move down and...
			t.MoveRope('D')
			if h.x > t.x {
				t.MoveRope('R')
			} else if h.x < t.x {
				t.MoveRope('L')
			} else {
				return errors.New("Weird logic? Moved down but not right or left")
			}
		} else {
			return errors.New(fmt.Sprintf("Don't know what happened! h - x:%d y:%d, t - x:%d y:%d", h.x, h.y, t.x, t.y))
		}
	}
	return nil // we moved, no errors
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

	h := newRopeEnd()
	t := newRopeEnd()

	posCounter := make(map[string]int)
	stepCount := 0

	// count the first location
	pos := fmt.Sprintf("%s~%s", strconv.Itoa(t.x), strconv.Itoa(t.y))
	posCounter[pos]++

	for _, v := range lines {
		s := strings.Split(v, " ")
		dir := s[0]
		amt, err := strconv.Atoi(s[1])
		if err != nil {
			panic("Could not read integer!")
		}
		fmt.Printf("Direction: %s, Amount: %s\n", s[0], s[1])
		for i := 0; i < amt; i++ {

			r := []rune(dir)
			h.MoveRope(r[0])

			res := DetermineTailDirection(&h, &t)
			if res != nil {
				fmt.Printf("Step: %d) h - x:%d y:%d, t - x:%d y:%d - %s\n", stepCount, h.x, h.y, t.x, t.y, res)
			} else {
				fmt.Printf("Step: %d) h - x:%d y:%d, t - x:%d y:%d\n", stepCount, h.x, h.y, t.x, t.y)
			}
			// Count the new position
			pos = fmt.Sprintf("%s~%s", strconv.Itoa(t.x), strconv.Itoa(t.y))
			posCounter[pos]++

			stepCount++
		}
	}
	// var counter int = 0
	// for range posCounter {
	// 	counter++
	// }
	// fmt.Printf("posCounter: +%v\n", posCounter)
	fmt.Println("Possible positions:", len(posCounter))
}
