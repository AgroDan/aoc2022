package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type MonkeyMap struct {
	pos        [][]string
	directions []string
}

func (m *MonkeyMap) ParseDirections(d string) {
	// this function will parse a direction line
	// also i'd like to find out how to do this with
	// regex, but my google-fu is weak apparently.
	var buf string
	for _, c := range d {
		// for each character...
		if c == 'L' || c == 'R' {
			if len(buf) > 0 {
				m.directions = append(m.directions, buf)
				buf = ""
			}
			m.directions = append(m.directions, string(c))
		} else {
			buf += string(c)
		}
	}
	if len(buf) > 0 {
		m.directions = append(m.directions, buf)
	}
}

func ReadMap(input []string) MonkeyMap {
	// REMEMBER -- to reference the map, it is [Y][X]
	m := MonkeyMap{}
	widest := 0

	for _, line := range input {
		if len(line) == 0 {
			continue
		}

		match, _ := regexp.MatchString(`[0-9LR]`, line)
		if match {
			m.ParseDirections(line)
		} else {
			r := make([]string, 0)
			for i, row := range line {
				// for every letter in every row
				if string(row) == " " {
					r = append(r, "+")
				} else {
					r = append(r, string(row))
				}
				if i > widest {
					widest = i
				}
			}
			m.pos = append(m.pos, r)
		}
		// fmt.Printf("length of col: %d\n", len(r))
	}

	// now to make things nice and pretty, let's
	// go back and loop through to make sure the
	// map is a nice square.
	for i, v := range m.pos {
		if len(v)-1 < widest {
			diff := widest - (len(v) - 1)
			for i := 0; i < diff; i++ {
				v = append(v, "+")
			}
			m.pos[i] = v
		}
	}
	return m
}

func PrintMonkeyMap(m *MonkeyMap, monke *Monkey, breadcrumbs map[Coord]int) {
	// this function simply prints out the monkeymap.
	// Using this for debugging purposes only. This also
	// prints out the Monkey on the map at it's position
	// including the direction.
	for Y, row := range m.pos {
		for X, col := range row {
			c := Coord{
				X: X,
				Y: Y,
			}
			if _, ok := breadcrumbs[c]; ok {
				switch breadcrumbs[c] {
				case left:
					fmt.Printf("<")
				case right:
					fmt.Printf(">")
				case up:
					fmt.Printf("^")
				case down:
					fmt.Printf("v")
				}
			} else {
				fmt.Printf("%s", col)
			}
		}
		fmt.Printf("\n")
	}
	for _, dir := range m.directions {
		if dir == "R" || dir == "L" {
			fmt.Printf("Turn %s, ", dir)
		} else {
			fmt.Printf("Move %s paces forward, ", dir)
		}
	}
	fmt.Printf("\n")
}

func FindStart(m *MonkeyMap) (X, Y int) {
	// this function simply returns the X and Y
	// coordinate of where the starting point of
	// the map is. REMEMBER ROWS AND COLS START AT
	// 1 BECAUSE THE GUYS THAT RUN AOC ARE SADISTS
	for i, v := range m.pos {
		for j, w := range v {
			if w == "." {
				return j, i
			}
		}
	}
	return -1, -1
}

func Travel(m *MonkeyMap, monke *Monkey) {
	// this function iterates through all the instructions
	// given, and updates the monke object
	for _, v := range m.directions {
		num, err := strconv.Atoi(v)
		if err != nil {
			// most likely it is R or L
			e := monke.Rotate(v)
			if e != nil {
				panic("Unknown direction")
			}
		} else {
			for i := 0; i < num; i++ {
				monke.MoveOne()
			}
		}
	}
}

// this is mostly for part 2 now. Yikes
type Coord struct {
	X, Y int
}

// type Side struct {
// 	// this will specify the top left, top right, bottom
// 	// left and bottom right corners. Technically then this
// 	// means that if the proposed coordinate is within this
// 	// object then it will continue as normal. However if we
// 	// are outside of this side, it should point to the side
// 	// that we are currently in. I'll have to figure out how
// 	// to automatically build this though.
// 	TL, TR, BL, BR        Coord
// 	Up, Down, Left, Right *Side
// }

// hardcoded psuedocode
// [0]Coord{X: 50, Y: 0}
// [1]Coord{X: 100, Y: 0}
// [2]Coord{X: 50, Y: 50}
// [3]Coord{X: 0, Y: 100}
// [4]Coord{X: 50, Y: 100}
// [5]Coord{X: 0, Y: 150}

// this is a sextant object which defines which
// sides are adjacent to it in which direction.
type Sextant struct {
	loc                   Coord
	up, down, left, right int
}

type Cube struct {
	Sides [6]Sextant // the coordinates of the top left corner of every side
	Size  int        // how many units each side is
}

func BuildDebugCube() Cube {
	c := [6]Sextant{
		// side 1
		{
			loc:   Coord{X: 8, Y: 0},
			up:    1,
			down:  3,
			left:  2,
			right: 5,
		},
		// side 2
		{loc: Coord{X: 0, Y: 4},
			up:    0,
			down:  4,
			left:  5,
			right: 2,
		},
		// side 3
		{loc: Coord{X: 4, Y: 4},
			up:    0,
			down:  4,
			left:  1,
			right: 3,
		},
		// side 4
		{loc: Coord{X: 8, Y: 4},
			up:    0,
			down:  4,
			left:  2,
			right: 5,
		},
		// side 5
		{loc: Coord{X: 8, Y: 8},
			up:    3,
			down:  1,
			left:  2,
			right: 5,
		},
		// side 6
		{loc: Coord{X: 12, Y: 8},
			up:    3,
			down:  1,
			left:  4,
			right: 0,
		},
	}
	return Cube{
		Sides: c,
		Size:  4,
	}
}

func BuildCube() Cube {
	// this just builds a hard-coded cube. I will figure this out
	// later, maybe.
	c := [6]Sextant{
		// side 1, remember sides start at 0 here to add to the confusion of course
		{loc: Coord{X: 50, Y: 0},
			up:    5,
			down:  2,
			left:  3,
			right: 1,
		},
		// side 2
		{loc: Coord{X: 100, Y: 0},
			up:    5,
			down:  2,
			left:  0,
			right: 4,
		},
		// side 3
		{
			loc:   Coord{X: 50, Y: 50},
			up:    0,
			down:  4,
			left:  3,
			right: 1,
		},
		// side 4
		{
			loc:   Coord{X: 0, Y: 100},
			up:    2,
			down:  5,
			left:  0,
			right: 4,
		},
		// side 5
		{
			loc:   Coord{X: 50, Y: 100},
			up:    2,
			down:  5,
			left:  3,
			right: 1,
		},
		// side 6
		{
			loc:   Coord{X: 0, Y: 150},
			up:    3,
			down:  1,
			left:  0,
			right: 4,
		},
	}
	return Cube{
		Sides: c,
		Size:  50,
	}
}

func (s Sextant) InSextant(comp Coord, size int) bool {
	// this function, given a coordinate, will check to see if the provided
	// coordinate is inside of this particular sextant.
	if (comp.X >= s.loc.X && comp.X < s.loc.X+size) &&
		(comp.Y >= s.loc.Y && comp.Y < s.loc.Y+size) {
		return true
	}
	return false
}

func (s Sextant) GetSide(side, size int) []Coord {
	// this function, given a direction, will
	// return the absolute coordinates of every
	// item on the given side. You must provide
	// the size of the square for this to be useful
	// the order will always be left to right, or
	// top to bottom.
	var retval []Coord
	switch side {
	case up:
		for i := s.loc.X; i < s.loc.X+size; i++ {
			retval = append(retval, Coord{
				X: i,
				Y: s.loc.Y,
			})
		}
	case down:
		// X: 12, Y: 8 edge of 6
		for i := s.loc.X; i < s.loc.X+size; i++ {
			retval = append(retval, Coord{
				X: i,
				Y: s.loc.Y + (size - 1),
			})
		}
	case left:
		for i := s.loc.Y; i < s.loc.Y+size; i++ {
			retval = append(retval, Coord{
				X: s.loc.X,
				Y: i,
			})
		}
	case right:
		for i := s.loc.Y; i < s.loc.Y+size; i++ {
			retval = append(retval, Coord{
				X: s.loc.X + (size - 1),
				Y: i,
			})
		}
	}
	return retval
}

func (c Cube) whatSextant(comp Coord) int {
	// this will determine what sextant the provided coordinate is inside of.
	// the number provided will refer to the index of the sextants of the cube.
	for i, v := range c.Sides {
		if v.InSextant(comp, c.Size) {
			return i
		}
	}
	return -1
}

// func (c Cube) TransferSide(from, to Coord) int {
// 	// this function has a little bit of hard coding involved. I have created
// 	// the DirectionMatrix object, which states the

// 	DirectionMatrix := [6][6]int{
// 		{0, 0, 0, down, -1, right},
// 		{0, 0, left, -1, left, 0},
// 		{0, up, 0, down, 0, -1},
// 		{right, -1, right, 0, 0},
// 		{-1, left, 0, 0, 0, left},
// 		{down, 0, -1, 0, up, 0},
// 	}
// }

func (c Cube) GetAbsolute(comp Coord, sideNum int) Coord {
	// Given RELATIVE coordinates, specific to the side in question,
	// this will return the ACTUAL coordinates that can be used in the
	// MonkeyMap object
	return Coord{
		X: c.Sides[sideNum].loc.X + comp.X,
		Y: c.Sides[sideNum].loc.Y + comp.Y,
	}
}

func (c Cube) GetRelative(comp Coord) Coord {
	// given the absolute coordinates, this will find out the relative
	// coordinates for the side this is in.
	num := c.whatSextant(comp)
	return Coord{
		X: c.Sides[num].loc.X % c.Size,
		Y: c.Sides[num].loc.Y % c.Size,
	}
}

// func (c Cube) GetMirror(comp Coord, dest Sextant) Coord {
// 	// this returns the ABSOLUTE COORDINATE of the "mirror" of the given
// 	// coordinate on this side. NOTE: This does NOT do any checking for
// 	// whether direction needs to be changed, you must figure this out
// 	// elsewhere.
// 	relCoord := c.GetRelative(comp)

// }

func TravelPart2(m *MonkeyMap, c *Cube, monke *Monkey, debug bool) map[Coord]int {
	// this function iterates through all the instructions
	// given, and updates the monke object
	breadcrumbs := make(map[Coord]int)
	for _, v := range m.directions {
		num, err := strconv.Atoi(v)
		if err != nil {
			// most likely it is R or L
			e := monke.Rotate(v)
			breadcrumbs[Coord{X: monke.X, Y: monke.Y}] = monke.facing
			if e != nil {
				panic("Unknown direction")
			}
		} else {
			for i := 0; i < num; i++ {
				monke.MoveOnePartTwo(c, debug)
				breadcrumbs[Coord{X: monke.X, Y: monke.Y}] = monke.facing
			}
		}
	}
	return breadcrumbs
}
