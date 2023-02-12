package main

import (
	"errors"
	"fmt"
)

/*
This will contain the "Monkey" that moves around the map.
I assume it's a monkey. Whatever.
*/

const (
	right = iota
	down
	left
	up
)

func BetterMod(x, y int) int {
	// this function is better than any stupid interpretation that Go
	// uses. IMO -1 % 50 SHOULD equal 49. There is some logical reason
	// as to why Go does this but it sounds dumb regardless.
	rem := x % y
	if rem < 0 {
		rem += y
	}
	return rem
}

// type Breadcrumb struct {
// 	c      Coord
// 	facing int
// }

type Monkey struct {
	X, Y   int        // position
	mmap   *MonkeyMap // the monkeymap associated.
	facing int        // what direction we are facing
}

func NewMonkey(X, Y, facing int, mm *MonkeyMap) *Monkey {
	return &Monkey{
		X:      X,
		Y:      Y,
		facing: facing,
		mmap:   mm,
	}
}

// DAN'S NOTE: go over this function with a fine toothed comb
// because I got turned around with the coordinates, so hopefully
// the logic is ok. I thought it was [X][Y], when in reality it's
// [Y][X]. Whoops.
func (m *Monkey) MoveOne() bool {
	// this function moves us around. It will wrap around
	// to the opposite side of the map if it has to. Throws
	// a false if we can't move there because there's a wall
	// or something.
	switch m.facing {
	case up:
		if m.Y <= 0 || m.mmap.pos[m.Y-1][m.X] == "+" {
			// we're at the top or at least at a blank area
			// so work up from the bottom looking for an open
			// spot. Or a wall, whichever comes first (and throw
			// appropriate return value)
			for i := len(m.mmap.pos) - 1; i > 0; i-- {
				if m.mmap.pos[i][m.X] == "#" {
					return false
				} else if m.mmap.pos[i][m.X] == "+" {
					continue
				} else {
					m.Y = i
					return true
				}
			}
		} else if m.mmap.pos[m.Y-1][m.X] == "#" {
			return false
		} else {
			m.Y--
			return true
		}
	case down:
		if m.Y >= len(m.mmap.pos)-1 || m.mmap.pos[m.Y+1][m.X] == "+" {
			// we're at the bottom
			for i := 0; i < m.Y; i++ {
				if m.mmap.pos[i][m.X] == "#" {
					return false
				} else if m.mmap.pos[i][m.X] == "+" {
					continue
				} else {
					m.Y = i
					return true
				}
			}
		} else if m.mmap.pos[m.Y+1][m.X] == "#" {
			return false
		} else {
			m.Y++
			return true
		}
	case left:
		if m.X <= 0 || m.mmap.pos[m.Y][m.X-1] == "+" {
			// we're all the way left
			for i := len(m.mmap.pos[m.Y]) - 1; i > m.X; i-- {
				if m.mmap.pos[m.Y][i] == "#" {
					return false
				} else if m.mmap.pos[m.Y][i] == "+" {
					continue
				} else {
					m.X = i
					return true
				}
			}
		} else if m.mmap.pos[m.Y][m.X-1] == "#" {
			return false
		} else {
			m.X--
			return true
		}
	case right:
		if m.X >= len(m.mmap.pos[m.Y])-1 || m.mmap.pos[m.Y][m.X+1] == "+" {
			// we're all the way right
			for i := 0; i < m.X; i++ {
				if m.mmap.pos[m.Y][i] == "#" {
					return false
				} else if m.mmap.pos[m.Y][i] == "+" {
					continue
				} else {
					m.X = i
					return true
				}
			}
		} else if m.mmap.pos[m.Y][m.X+1] == "#" {
			return false
		} else {
			m.X++
			return true
		}
	}
	// otherwise what the hell do you want me to do
	return false
}

func (m *Monkey) Rotate(dir string) error {
	// this function, given an R or L direction, will rotate
	// the monkey appropriately.
	switch dir {
	case "L":
		switch m.facing {
		case up:
			m.facing = left
		case down:
			m.facing = right
		case left:
			m.facing = down
		case right:
			m.facing = up
		}
	case "R":
		switch m.facing {
		case up:
			m.facing = right
		case down:
			m.facing = left
		case left:
			m.facing = up
		case right:
			m.facing = down
		}
	default:
		return errors.New("Unknown direction")
	}
	return nil
}

func (m *Monkey) MoveOnePartTwo(c *Cube, debug bool) bool {
	// this function will move the monkey specific to part 2's
	// requirements, which is to say that for each move forward
	// this will check the next square to determine the usual differences,
	// only this time it's not a simple wrap-around, but rather
	// a cube traversal! I have to confirm mappings for each step.

	// I will kinda jerry-rig this new coord object in here
	// since I made this after I finished part 1 but it does
	// basically the same thing
	currentCoord := Coord{
		X: m.X,
		Y: m.Y,
	}
	checkBox := currentCoord
	relCoord := c.GetRelative(currentCoord)

	// Need this to see if we need to change directions
	// DirectionMatrix[from][to] == new direction. -1 if no change
	// DirectionMatrix := [6][6]int{
	// 	{-1, -1, -1, right, -1, right},
	// 	{-1, -1, left, -1, left, -1},
	// 	{-1, up, -1, down, -1, -1},
	// 	{right, -1, right, -1, -1, -1},
	// 	{-1, left, -1, -1, -1, left},
	// 	{down, -1, -1, -1, up, -1},
	// }
	DirectionMatrix := [6][6]int{
		{-1, right, down, right, -1, right},
		{left, -1, left, -1, left, up},
		{up, up, -1, down, down, -1},
		{right, -1, right, -1, right, down},
		{-1, left, up, left, -1, left},
		{down, down, -1, up, up, -1},
	}

	DebugDirectionMatrix := [6][6]int{
		{-1, down, down, down, -1, left},
		{down, -1, right, -1, up, up},
		{right, left, -1, right, right, -1},
		{up, -1, left, -1, down, down},
		{-1, up, up, up, -1, right},
		{left, right, -1, left, left, -1},
	}

	currentSextant := c.whatSextant(currentCoord)

	switch m.facing {
	case up:
		checkBox.Y--
		relCoord.Y = BetterMod(relCoord.Y-1, c.Size)
	case down:
		checkBox.Y++
		relCoord.Y = BetterMod(relCoord.Y+1, c.Size)
	case left:
		checkBox.X--
		relCoord.X = BetterMod(relCoord.X-1, c.Size)
	case right:
		checkBox.X++
		relCoord.X = BetterMod(relCoord.X+1, c.Size)
	}

	// // check if the next box is valid to move onto, quick win
	// if m.mmap.pos[checkBox.Y][checkBox.X] == "." {
	// 	m.Y = checkBox.Y
	// 	m.X = checkBox.X
	// 	return true
	// } else if m.mmap.pos[checkBox.Y][checkBox.X] == "#" {
	// 	return false // we hit a wall
	// }

	// now it gets complicated. We didn't hit a wall, and we didn't hit
	// an open space. Which means the next spot is the only other possibility,
	// which means we moved to another side. Now we have to determine what
	// side that is. We have the Sextant object which should map what side
	// we are traveling to. We can use the relative coordinate functions
	// to determineIMPORTANT: If we DO NOT change directions, then the
	// intended X/Y coordinates must be flipped!

	var newSextant int
	// first, let's see if we've hit the actual boundary of the map
	if checkBox.X < 0 {
		// we are moving left from the box
		// which sextent?
		newSextant = c.Sides[currentSextant].left

	} else if checkBox.X > len(m.mmap.pos[m.Y])-1 {
		// we are moving right now
		newSextant = c.Sides[currentSextant].right
	} else if checkBox.Y < 0 {
		// we are moving up
		newSextant = c.Sides[currentSextant].up
	} else if checkBox.Y > len(m.mmap.pos)-1 {
		// moving down
		newSextant = c.Sides[currentSextant].down
	} else {
		// otherwise this is within bounds. Last remaining possibility
		// is this is a "+" or a "." character, meaning we just round
		// out where we are relatively speaking

		// first, is this legit:
		if m.mmap.pos[checkBox.Y][checkBox.X] == "#" {
			// we hit a wall
			return false
		} else if m.mmap.pos[checkBox.Y][checkBox.X] == "." {
			// this is legit
			m.X = checkBox.X
			m.Y = checkBox.Y
			return true
		}

		// otherwise, the last possible thing this can be is a "+" char
		// meaning we need to determine if we're flipping around or anything
		switch m.facing {
		case up:
			newSextant = c.Sides[currentSextant].up
		case down:
			newSextant = c.Sides[currentSextant].down
		case left:
			newSextant = c.Sides[currentSextant].left
		case right:
			newSextant = c.Sides[currentSextant].right
		}

	}

	// I THINK THIS NEEDS TO BE MOVED DOWN
	newCoord := c.GetAbsolute(relCoord, newSextant)

	// now the tricky part! Do we change directions? If so flip the X/Y!
	var dir int
	if debug {
		dir = DebugDirectionMatrix[currentSextant][newSextant]
	} else {
		dir = DirectionMatrix[currentSextant][newSextant]
	}
	var invert bool
	if dir >= 0 {
		if m.facing == left {
			switch dir {
			case up, right:
				invert = true
			default:
				invert = false
			}
		} else if m.facing == right {
			switch dir {
			case down, left:
				invert = true
			default:
				invert = false
			}
		} else if m.facing == up {
			switch dir {
			case down, left:
				invert = true
			default:
				invert = false
			}
		} else if m.facing == down {
			switch dir {
			case up, right:
				invert = true
			default:
				invert = false
			}
		}

		// find the adjacent side of the next sextant -- it will be
		// the opposite of whatever the sideMatrix says it is
		var oppSide int
		switch dir {
		case up:
			oppSide = down
		case down:
			oppSide = up
		case left:
			oppSide = right
		case right:
			oppSide = left
		}

		// NOW...get the edges of the two sides.
		currEdge := c.Sides[currentSextant].GetSide(m.facing, c.Size)
		targetEdge := c.Sides[newSextant].GetSide(oppSide, c.Size)

		// get where we are on the current edge
		var idx int
		for i, v := range currEdge {
			if v == currentCoord {
				idx = i
				break
			}
		}

		// now based on the direction, map accordingly
		if !invert {
			// just a simple mapping
			newCoord = targetEdge[idx]
		} else {
			// we are inverting
			offset := (len(targetEdge) - 1) - idx
			newCoord = targetEdge[offset]
		}

	}

	// now can we go there?
	if m.mmap.pos[newCoord.Y][newCoord.X] == "#" {
		return false
	} else if m.mmap.pos[newCoord.Y][newCoord.X] == "." {
		m.Y = newCoord.Y
		m.X = newCoord.X
		if dir >= 0 {
			m.facing = dir
		}
		return true
	} else {
		fmt.Printf("The item in question is %s\n", m.mmap.pos[newCoord.Y][newCoord.X])
		fmt.Printf("The coords: %#v\n", newCoord)
		fmt.Printf("Direction: %d\n", m.facing)
		panic("unknown character in map")
	}
}
