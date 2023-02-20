package main

import (
	"errors"
	"fmt"
)

/*
This package will define the Elf object. When parsing the map,
it will come across a '#' character. This will symbolize the elf.
Considering that elves will be ephemeral (when it comes to location)
since the map will grow potentially with all movements, the location
will be useful only for the parsing and moving. After that it should
be forgotten and re-parsed.
*/

type Coord struct {
	X, Y int
}

func GetDirection(c Coord, dir string) (Coord, error) {
	newc := Coord{}

	switch dir {
	case "N":
		newc.X = c.X
		newc.Y = c.Y - 1
	case "NW":
		newc.X = c.X - 1
		newc.Y = c.Y - 1
	case "NE":
		newc.X = c.X + 1
		newc.Y = c.Y - 1
	case "E":
		newc.X = c.X + 1
		newc.Y = c.Y
	case "W":
		newc.X = c.X - 1
		newc.Y = c.Y
	case "S":
		newc.X = c.X
		newc.Y = c.Y + 1
	case "SW":
		newc.X = c.X - 1
		newc.Y = c.Y + 1
	case "SE":
		newc.X = c.X + 1
		newc.Y = c.Y + 1
	default:
		return newc, errors.New("Invalid direction")
	}
	return newc, nil
}

type Elf struct {
	Loc         Coord  // current location
	Propose     Coord  // proposed new movement.
	SetProposal bool   // have we made a proposal yet? (otherwise coord is 0,0)
	Move        bool   // do we move on this round?
	MoveQueue   []rune // this is a queue
}

func NewElf(X, Y int) Elf {
	return Elf{
		Loc: Coord{
			X: X,
			Y: Y,
		},
		Propose:     Coord{},
		SetProposal: false,
		Move:        false,
		MoveQueue:   []rune{'N', 'S', 'W', 'E'},
	}
}

func (e *Elf) PrintElf() {
	fmt.Printf("Elf Location X: %d, Y: %d || Set Proposal? %t || Will Move? %t\n", e.Loc.X, e.Loc.Y, e.SetProposal, e.Move)
	fmt.Printf("\tProposed X: %d, Y: %d || Current Next Check: %c\n", e.Propose.X, e.Propose.Y, e.MoveQueue[0])
}

func (e *Elf) Reset() {
	// this function resets the elf's proposals back
	// to what it was. This does NOT change the location.
	e.Propose.X, e.Propose.Y = 0, 0
	e.SetProposal, e.Move = false, false
}

func (e *Elf) Dequeue() rune {
	// this function pops the item out of the queue,
	// puts it back at the end of the queue, and returns it.
	retval := e.MoveQueue[0]
	e.MoveQueue = e.MoveQueue[1:]
	e.MoveQueue = append(e.MoveQueue, retval)
	return retval
}

func (e *Elf) Alone(g *GroveMap) bool {
	// given the elf, it will check for the surrounding
	// area around the elf and determine if we are alone.
	// if true, then we shouldn't move.
	directions := []string{"N", "NW", "NE", "E", "W", "S", "SW", "SE"}

	for _, v := range directions {
		checkDir, err := GetDirection(e.Loc, v)
		if err != nil {
			fmt.Printf("Bad direction!\n")
		}
		if g.CheckCoord(checkDir) == '#' {
			return false
		}
	}
	return true
}

func (e *Elf) ProposeDirection(g *GroveMap) {
	// this will determine what to do based on the particular
	// elf's next queued action.
	if e.Alone(g) {
		e.Move = false
		e.SetProposal = true
	}

	for i := 0; i < len(g.MoveQueue); i++ {
		// while we don't have a proposal set yet...
		if e.SetProposal {
			break
		}
		var checkArea [3]string
		switch g.MoveQueue[i] {
		case 'N':
			checkArea = [3]string{"N", "NE", "NW"} // check above
		case 'S':
			checkArea = [3]string{"S", "SE", "SW"} // check below
		case 'E':
			checkArea = [3]string{"E", "NE", "SE"} // check right
		case 'W':
			checkArea = [3]string{"W", "NW", "SW"} // check left
		}
		hasElf := false
		for _, v := range checkArea {
			checkDir, err := GetDirection(e.Loc, v)
			if err != nil {
				fmt.Printf("Bad direction!\n")
			}
			if g.CheckCoord(checkDir) == '#' {
				hasElf = true
				break
			}
		}
		if !hasElf {
			// no elves in the direction, go in proposed dir
			e.Move = true
			var err error
			e.Propose, err = GetDirection(e.Loc, string(g.MoveQueue[i]))
			if err != nil {
				fmt.Printf("Bad direction!")
			}
			// fmt.Printf("Proposing this elf goes to X: %d, Y: %d\n", e.Propose.X, e.Propose.Y)
			e.SetProposal = true
		}
	}
}

// func CheckForProposals(elves []*Elf) {
// 	// this loops through the entire elf list and determines
// 	// if they are moving or not.
// 	for _, e := range elves {
// 		if e.Move { // this elf is moving, check for doubles
// 			for _, check := range elves {
// 				if check.Move && e != check {
// 					if e.Propose == check.Propose {
// 						e.Move = false
// 						check.Move = false
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// The elfstack is the list of elves we know about.
type ElfStack []*Elf

func (e *ElfStack) GrowMap(size int) {
	// this function will iterate through every
	// elf on the elfstack and increase their X/Y
	// coordinates by the proposed size
	for _, v := range *e {
		v.Loc.X += size
		v.Loc.Y += size
	}
}

func (e *ElfStack) ReviewProposals() {
	// this function will iterate through the entire Elfstack
	// and determine if each elf can move based on other proposals
	for _, outerElf := range *e {
		// outerElf.PrintElf()
		if !outerElf.Move {
			// current elf is not expected to move
			continue
		}
		checkP := Coord{
			X: outerElf.Propose.X,
			Y: outerElf.Propose.Y,
		}
		for _, innerElf := range *e {
			if outerElf == innerElf {
				// working with the same elf
				continue
			}
			if !innerElf.Move {
				// compared elf is not expected to move
				continue
			}

			if innerElf.Propose.X == checkP.X && innerElf.Propose.Y == checkP.Y {
				// same proposal, reset the elves!
				outerElf.Reset()
				innerElf.Reset()
			}
		}
	}
}

func (e *ElfStack) ResetElves() {
	// goes through the list and resets all elves
	for _, v := range *e {
		v.Reset()
	}
}
