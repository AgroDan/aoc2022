package main

import (
	"errors"
	"fmt"
)

/*
	This file will hold all of the functions used by the sand.
	The sand operates under a basic principle. Given a starting
	position (default 500,0), it will operate under a series of
	"ticks", where each tick it will perform the following check:

	With current position of sand:
	1. Are we off the map?
		Yes: End loop
		No: Move to step 2
	2. Is there anything below it?
		No: fall by 1. Repeat step 2.
		Yes: move to step 3
	3. Is there anything below and to the left?
		Yes: move to step 4
		No: fall down left, repeat step 1
	4. Is there anything below and to the right?
		Yes: Rest here. No further movement.
		No: fall down left, repeat step 1
*/

type Sand struct {
	X, Y int
}

// constructor
func NewSand(x, y int) Sand {
	return Sand{
		X: x,
		Y: y,
	}
}

func (s *Sand) DropSand(c *Cave) (bool, error) {
	// This will drop the sand by one tick.
	// Returns true if the sand successfully moves.
	// Returns false if the sand rests.
	// throws an error if we are off the map.

	// what's below us

	// first, check to see if we're in the abyss.
	if s.Y == len(c.coords)-1 {
		return true, errors.New("Off the map, down into the abyss")
	}

	checkBlock := c.coords[s.Y+1][s.X]
	if checkBlock == "." {
		// air block
		s.Y++
		return true, nil
	}

	// otherwise, what's down left
	if s.X == 0 {
		// down left would be off the map. Throw error.
		return true, errors.New("Off the map, too far left")
	}

	checkBlock = c.coords[s.Y+1][s.X-1]
	if checkBlock == "." {
		// air block
		s.Y++
		s.X--
		return true, nil
	}

	// finally, what's down right
	if s.X == len(c.coords[s.Y])-1 {
		// down right would be off the map. Throw error.
		return true, errors.New("Off the map, too far right")
	}

	checkBlock = c.coords[s.Y+1][s.X+1]
	if checkBlock == "." {
		// air block
		s.Y++
		s.X++
		return true, nil
	}

	// otherwise, can't go any further. sand is at rest.
	return false, nil
}

func (s *Sand) DropSandPart2(c *Cave) bool {
	// This will drop the sand by one tick.
	// Returns true if the sand successfully moves.
	// Returns false if the sand rests.

	// what's below us? Not an abyss this time

	checkBlock := c.coords[s.Y+1][s.X]
	if checkBlock == "." {
		// air block
		s.Y++
		return true
	}

	// otherwise, what's down left
	if s.X == 0 {
		// we are all the way to the left and checking, so re-draw
		DrawLeft(c)
		// update the index of the sand!
		s.X++
	}

	checkBlock = c.coords[s.Y+1][s.X-1]
	if checkBlock == "." {
		// air block
		s.Y++
		s.X--
		return true
	}

	// finally, what's down right
	if s.X == len(c.coords[s.Y])-1 {
		// we are all the way to the right and checking, so re-draw
		DrawRight(c)
	}

	checkBlock = c.coords[s.Y+1][s.X+1]
	if checkBlock == "." {
		// air block
		s.Y++
		s.X++
		return true
	}

	// otherwise, can't go any further. sand is at rest.
	return false
}

func WriteSand(c *Cave, s Sand) error {
	// this writes the sand block in the position given.
	// Throws an error if the block is occupied.
	if c.coords[s.Y][s.X] != "." {
		return errors.New("cannot place sand in occupied space")
	}
	c.coords[s.Y][s.X] = "o"
	return nil
}

func PourSand(c *Cave, sandStartX, sandStartY int) {
	mapStartX, mapStartY := c.MapCoords(sandStartX, sandStartY)
	for {
		// keep on pouring
		s := NewSand(mapStartX, mapStartY)

		// edge case, if the beginning sand is NOT air
		if c.coords[mapStartY][mapStartX] != "." {
			fmt.Printf("Could not start sand!\n")
			break
		}

		// drop the sand
		for {
			sandState, err := s.DropSand(c)
			if err != nil {
				// can't drop sand anymore
				return
			}
			if !sandState {
				// sand comes to a rest
				WriteSand(c, s)
				break
			}
		}
	}
}

func PourSandPart2(c *Cave, sandStartX, sandStartY int) {
	for {
		// this will be in the for loop this time because the offsets
		// may change as the map gets re-drawn!
		mapStartX, mapStartY := c.MapCoords(sandStartX, sandStartY)
		// keep on pouring
		s := NewSand(mapStartX, mapStartY)

		// this is no longer the edge case, this is the end goal!
		if c.coords[mapStartY][mapStartX] != "." {
			// the sand has been plugged!
			// fmt.Printf("Could not start sand!\n")
			break
		}

		// drop the sand
		for {
			sandState := s.DropSandPart2(c)
			if !sandState {
				// sand comes to a rest
				WriteSand(c, s)
				break
			}
		}
	}
}
