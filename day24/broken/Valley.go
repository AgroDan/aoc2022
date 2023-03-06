package main

import (
	"errors"
	"fmt"
)

type Coord struct {
	row, col int
}

type Valley struct {
	m         [][]rune
	Entrance  Coord
	Exit      Coord
	Blizzards []Blizzard
	// this is where i would put a list of blizzards
}

func (v *Valley) DeepCopy() Valley {
	// this will create a new copy of a valley object
	newBlizz := make([]Blizzard, len(v.Blizzards))
	copy(newBlizz, v.Blizzards)

	// now the map object
	newMap := make([][]rune, len(v.m))
	for i := range v.m {
		newMap[i] = make([]rune, len(v.m[i]))
		copy(newMap[i], v.m[i])
	}

	return Valley{
		m:         newMap,
		Entrance:  v.Entrance,
		Exit:      v.Exit,
		Blizzards: newBlizz,
	}
}

func NewValley(m []string) Valley {
	// this function creates a new valley object
	// based on the []string data that is read
	// from a file.
	v := Valley{
		m:         make([][]rune, 0),
		Blizzards: make([]Blizzard, 0),
	}

	for R, line := range m {
		row := make([]rune, 0)
		for C, char := range line {
			if char == '#' || char == '.' {
				row = append(row, char)
			} else if char == '>' || char == '<' || char == '^' || char == 'v' {
				// this is where I would put the blizzard objects
				b := NewBlizzard(char, Coord{row: R, col: C})
				v.Blizzards = append(v.Blizzards, b)
				row = append(row, '.')
			} else {
				// this is probably a blank space. Fill it with ? just so i can come back
				// and fix this at a later time
				row = append(row, '?')
			}
		}
		v.m = append(v.m, row)
	}

	// now let's find the entrance
	for inc, val := range v.m[0] {
		if val == '.' {
			v.Entrance = Coord{
				row: 0,
				col: inc,
			}
			break
		}
	}

	// exit now
	for inc, val := range v.m[len(v.m)-1] {
		if val == '.' {
			v.Exit = Coord{
				row: len(v.m) - 1,
				col: inc,
			}
			break
		}
	}
	return v
}

func (v *Valley) IsBorder(c Coord) bool {
	// this function will determine if the coordinate provided is on the
	// border of the valley. This can be used to determine if a blizzard
	// is about to hit a border and can thus move appropriately

	// two separate if statements for readability's sake
	if c.row <= 0 || c.row >= len(v.m)-1 {
		return true
	}

	if c.col <= 0 || c.col >= len(v.m[c.row])-1 {
		return true
	}

	return false
}

func (v *Valley) Wrap(c Coord) (Coord, error) {
	// this function will wrap around to wherever we are and give the
	// opposing side where a blizzard _should_ end up after wrapping
	// around. Will throw an error if the coordinate provided is not
	// a border.
	if !v.IsBorder(c) {
		return Coord{}, errors.New("not a border spot")
	}

	// check if entrance or exit
	if c == v.Entrance {
		return Coord{}, errors.New("entrance")
	}

	if c == v.Exit {
		return Coord{}, errors.New("exit")
	}

	if c.row <= 0 {
		return Coord{
			col: c.col,
			row: len(v.m) - 2,
		}, nil
	}

	if c.row >= len(v.m)-1 {
		return Coord{
			col: c.col,
			row: 1,
		}, nil
	}

	if c.col <= 0 {
		return Coord{
			col: len(v.m[c.row]) - 2, // don't forget the border on the opp side
			row: c.row,
		}, nil
	}

	if c.col >= len(v.m[c.row])-1 {
		return Coord{
			col: 1,
			row: c.row,
		}, nil
	}

	newErr := fmt.Sprintf("unknown spot, given coord: %#v", c)
	return Coord{}, errors.New(newErr)
}

func (v *Valley) CountSameSpace(c Coord) int {
	// this function will count how many blizzards are on the same space.
	count := 0
	for _, blizz := range v.Blizzards {
		if c == blizz.pos {
			count++
		}
	}
	return count
}

func (v *Valley) PrintValley(e Elf) {
	// this will print the entire map including all of the blizzards.
	// to do this, we draw a blank map first, then add all the blizzards.
	asciiZero := 48 // omg why

	copyM := make([][]rune, len(v.m))
	for i := range v.m {
		copyM[i] = make([]rune, len(v.m[i]))
		copy(copyM[i], v.m[i])
	}

	// now insert the blizzards

	for _, blizz := range v.Blizzards {
		var char rune
		switch blizz.direction {
		case north:
			char = '^'
		case south:
			char = 'v'
		case east:
			char = '>'
		case west:
			char = '<'
		default:
			char = '?'
		}

		blizzAmt := v.CountSameSpace(Coord{
			row: blizz.pos.row,
			col: blizz.pos.col,
		})
		// fmt.Printf("BlizzAmt: %d\n", blizzAmt)
		if blizzAmt > 9 {
			char = 'M'
		} else if blizzAmt > 1 {
			char = rune(blizzAmt + asciiZero)
		}
		// fmt.Printf("Char: %c\n", char)
		copyM[blizz.pos.row][blizz.pos.col] = char
	}

	copyM[e.loc.row][e.loc.col] = 'E'

	// now print
	for _, row := range copyM {
		for _, col := range row {
			fmt.Printf("%c", col)
		}
		fmt.Printf("\n")
	}
}

func (v *Valley) MoveOne() {
	// fmt.Printf("Total rows: %d, total cols: %d\n", len(v.m), len(v.m[0]))
	// this function moves the map by one minute.
	for i, blizz := range v.Blizzards {
		proposed := blizz.Peek()
		var err error
		// fmt.Printf("Proposed: r: %d, c: %d\n", proposed.row, proposed.col)
		if v.IsBorder(proposed) {
			// fmt.Printf("Border at r: %d, c: %d\n", proposed.row, proposed.col)
			proposed, err = v.Wrap(proposed)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		}
		v.Blizzards[i].pos = proposed
		// blizz.pos = proposed
	}
}

func (v *Valley) GetBlizzard(c Coord) ([]*Blizzard, error) {
	// this function returns the blizzards that are at a specific coordinate. Throws
	// an error if no blizzard exists at that coord
	retval := make([]*Blizzard, 0)
	for _, blizz := range v.Blizzards {
		if c == blizz.pos {
			fmt.Printf("Found a blizzard at %#v\n", blizz.pos)
			retval = append(retval, &blizz)
		}
	}

	if len(retval) > 0 {
		return retval, nil
	} else {
		return retval, errors.New("no blizzards at this loc")
	}
}

func (v *Valley) GetValidDirections(c Coord) []Coord {
	// this will return a list of valid directions, including the wait
	// direction (which is the same coord)
	dir := make([]Coord, 0)

	if c == v.Entrance {
		// we can always wait here, since it's the entrance.
		dir = append(dir, c)

		// we can only go south, so let's see if we even can.
		checkSouth := c
		checkSouth.row++
		if v.CheckLoc(checkSouth) {
			dir = append(dir, checkSouth)
		}

		return dir
	}

	if c == v.Exit {
		return dir // empty, we made it
	}

	// let's see if we can wait in our current location
	if v.CheckLoc(c) {
		dir = append(dir, c)
	}

	// going to simplify this...

	surrounding := v.GetSurrounding(c)

	for _, value := range surrounding {
		if value == v.Exit {
			return []Coord{value}
		}
		if v.CheckLoc(value) {
			dir = append(dir, value)
		}
	}

	return dir
}

func (v *Valley) CheckLoc(c Coord) bool {
	// checks if the given coordinate is valid. What determines validity is
	// checking all 4 directions _around_ the coordinate, looking for a blizzard
	// whose _next_ step is THAT coordinate. Confusing but this will handle it all.

	// first, let's see if the given coordinate is the entrance or exit
	if c == v.Entrance {
		return false // we don't want to go back to the entrance. Maybe?
	}

	if c == v.Exit {
		return true // if we can, we always want to go to the exit
	}

	// check to see if any blizzards' next step will be this coordinate
	for _, blizz := range v.Blizzards {
		if blizz.Peek() == c {
			return false
		}
	}

	return true
}

// func (v Valley) CheckLoc(c Coord) bool {
// 	// checks if the given coordinate is valid. What determines validity is
// 	// checking all 4 directions _around_ the coordinate, looking for a blizzard
// 	// whose _next_ step is THAT coordinate. Confusing but this will handle it all.

// 	// first, let's see if the given coordinate is the entrance or exit
// 	if c == v.Entrance {
// 		return false // we don't want to go back to the entrance. Maybe?
// 	}

// 	if c == v.Exit {
// 		return true // if we can, we always want to go to the exit
// 	}

// 	// this is weird but bear with me
// 	dir := map[rune]*Coord{}

// 	// just creating something to loop through all directions with...
// 	dir['N'], dir['S'], dir['E'], dir['W'] = &Coord{}, &Coord{}, &Coord{}, &Coord{}
// 	for k := range dir {
// 		dir[k].col, dir[k].row = c.col, c.row
// 		switch k {
// 		case 'N':
// 			dir[k].row--
// 		case 'S':
// 			dir[k].row++
// 		case 'E':
// 			dir[k].col++
// 		case 'W':
// 			dir[k].col--
// 		}

// 		// now check if we have to wrap...
// 		var err error
// 		target := *dir[k]

// 		if target == v.Entrance {
// 			continue
// 			// if we are looking at the entrance, don't
// 			// even consider it as an option
// 		}
// 		if v.IsBorder(*dir[k]) {
// 			target, err = v.Wrap(*dir[k])
// 			if err != nil {
// 				fmt.Printf("Error: %s\n", err)
// 			}
// 		}

// 		// if we have to wrap then we did already. otherwise
// 		// the standard direction is in the target variable.

// 		// THIS IS WRONG! WE NEED TO CHECK EVERY BLIZZARD TO SEE IF ANY NEXT POSITION
// 		// IS THE CURRENT COORDINATE!

// 		// now get the blizzards currently in this coordinate
// 		blizz, err := v.GetBlizzard(target)
// 		fmt.Printf("Blizzards at r:%d, c:%d -- %#v\n", target.row, target.col, blizz)
// 		if err == nil {
// 			// loop over the blizzards we found
// 			for _, b := range blizz {
// 				// check the next step of each blizzard
// 				check := b.Peek()
// 				// aaaaand check if we need to wrap!
// 				if v.IsBorder(check) {
// 					check, err = v.Wrap(check)
// 					if err != nil {
// 						fmt.Printf("Error: %s\n", err)
// 					}
// 				}

// 				// is the blizzard's next step the coordinate in question?
// 				fmt.Printf("does %d == %d and %d == %d?\n", c.col, check.col, c.row, check.row)
// 				if check == c {
// 					fmt.Printf("Invalid coord: row %d, col %d\n", check.row, check.col)
// 					return false
// 				}
// 			}
// 		}
// 	}
// 	// if we got here without returning, then this is valid
// 	return true
// }

func (v *Valley) GetHash() string {
	// there is probably a better way to do this but whatever
	var s string
	for _, v := range v.Blizzards {
		s += v.String()
	}
	return s
}

func (v *Valley) StepsToExit(c Coord) int {
	// this function will accept a coordinate and return how many steps
	// possible to get to the exit. The exit should always be the highest
	// number col and row since it is in the bottom right corner.
	return (v.Exit.col - c.col) + (v.Exit.row - c.row)
}

func (v *Valley) GetSurrounding(c Coord) []Coord {
	// this will return all 4 directions of the
	// given coordinate. This will wrap around if
	// need be

	retval := make([]Coord, 0)

	iter := []rune{'N', 'S', 'E', 'W'}
	for _, val := range iter {
		check := c
		switch val {
		case 'N':
			check.row--
		case 'S':
			check.row++
		case 'E':
			check.col++
		case 'W':
			check.col--
		}

		if v.IsBorder(check) {
			// if we're at a border...
			var err error
			check, err = v.Wrap(check)
			if err != nil {
				// most likely hit the entrance or something
				// fmt.Printf("Error: %s\n", err)
				continue
			}
		}
		retval = append(retval, check)
	}
	return retval
}
