package main

import "fmt"

type Coord struct {
	row, col int
}

type Valley struct {
	TopLeft     Coord // this is totally unnecessary but I'm keeping it for sanity's sake
	BottomRight Coord
	Entrance    Coord
	Exit        Coord
	Blizzards   []Blizzard
	Minutes     int
	// this is where i would put a list of blizzards
}

func (v *Valley) DeepCopy() Valley {
	// this will create a new copy of a valley object
	newBlizz := make([]Blizzard, len(v.Blizzards))
	copy(newBlizz, v.Blizzards)

	return Valley{
		TopLeft:     v.TopLeft,
		BottomRight: v.BottomRight,
		Entrance:    v.Entrance,
		Exit:        v.Exit,
		Blizzards:   newBlizz,
		Minutes:     v.Minutes,
	}
}

func NewValley(m []string) Valley {
	// this will read in the valley data line by line and make the
	// determinations to build a valley object.

	v := Valley{
		Blizzards: make([]Blizzard, 0),
		Minutes:   0,
	}

	// topleft is easy
	v.TopLeft = Coord{
		row: 0,
		col: 0,
	}

	// since it's always a rectangle, we know the col value of the BottomRight
	v.BottomRight = Coord{
		row: len(m) - 1,
		col: len(m[0]) - 1,
	}

	// get the entrance
	for i, r := range m[0] {
		if r == '.' {
			v.Entrance = Coord{
				row: 0,
				col: i,
			}
			break
		}
	}

	// exit
	for i, r := range m[len(m)-1] {
		if r == '.' {
			v.Exit = Coord{
				row: len(m) - 1,
				col: i,
			}
			break
		}
	}

	// now get the blizzards
	for R, line := range m {
		for C, char := range line {
			if char == '>' || char == '<' || char == '^' || char == 'v' {
				// this is where I would put the blizzard objects
				b := NewBlizzard(char, Coord{row: R, col: C})
				v.Blizzards = append(v.Blizzards, b)
			}
		}
	}
	return v
}

func (v *Valley) IsBorder(c Coord) bool {
	// this function will determine if the coordinate provided is on the
	// border of the valley. This can be used to determine if a blizzard
	// is about to hit a border and can thus move appropriately

	// two separate if statements for readability's sake
	if c.row <= 0 || c.row >= v.BottomRight.row {
		return true
	}

	if c.col <= 0 || c.col >= v.BottomRight.col {
		return true
	}

	return false
}

func (v *Valley) Wrap(c Coord) Coord {
	// this function will wrap around to wherever we are and give the
	// opposing side where a blizzard _should_ end up after wrapping
	// around. It is up to the caller to ENSURE we are at a border before
	// wrapping. Also note: ONLY BLIZZARDS WILL WRAP! Elves cannot.

	// check if entrance or exit
	if c == v.Entrance {
		return v.Entrance
	}

	if c == v.Exit {
		return v.Exit
	}

	if c.row <= 0 {
		return Coord{
			col: c.col,
			row: v.BottomRight.row - 1,
		}
	}

	if c.row >= v.BottomRight.row {
		return Coord{
			col: c.col,
			row: 1,
		}
	}

	if c.col <= 0 {
		return Coord{
			col: v.BottomRight.col - 1,
			row: c.row,
		}
	}

	if c.col >= v.BottomRight.col {
		return Coord{
			col: 1,
			row: c.row,
		}
	}

	// otherwise we're good, return the same coord
	return c
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

func (v *Valley) PrintValley(elf Coord) {
	// this will print the entire map including all of the blizzards.
	// to do this, we draw a blank map first, then add all the blizzards.
	asciiZero := 48 // omg why

	mapVis := make([][]rune, v.BottomRight.row+1)

	// build the map
	for i := 0; i <= v.BottomRight.row; i++ {
		rowLine := make([]rune, v.BottomRight.col+1)
		for j := 0; j <= v.BottomRight.col; j++ {
			if i == 0 || i == v.BottomRight.row {
				// top and bottom walls
				rowLine[j] = '#'
			} else if j == 0 || j == v.BottomRight.col {
				// left and right walls
				rowLine[j] = '#'
			} else {
				// otherwise, blank spaces
				rowLine[j] = '.'
			}
		}
		mapVis[i] = rowLine
	}

	// set the entrance and exits
	mapVis[v.Entrance.row][v.Entrance.col] = '.'
	mapVis[v.Exit.row][v.Exit.col] = '.'

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
		mapVis[blizz.pos.row][blizz.pos.col] = char
	}

	mapVis[elf.row][elf.col] = 'E'

	// now print
	for _, row := range mapVis {
		for _, col := range row {
			fmt.Printf("%c", col)
		}
		fmt.Printf("\n")
	}
}

func (v *Valley) MoveOne() {
	// this function moves the map by one minute.
	for i, blizz := range v.Blizzards {
		proposed := blizz.Peek()
		if v.IsBorder(proposed) {
			proposed = v.Wrap(proposed)
		}
		v.Blizzards[i].pos = proposed
	}
	v.Minutes++
}

func (v *Valley) CheckLoc(c Coord) bool {
	// checks if the given coordinate is valid. What determines validity
	// is checking the NEXT step of each blizzard and determining if they
	// would be moving to this coordinate.

	// first, let's see if the given coordinate is the entrance or exit
	if c == v.Entrance {
		return false // we don't want to go back to the entrance. Maybe?
	}

	if c == v.Exit {
		return true // if we can, we always want to go to the exit
	}

	// check to see if any blizzards' next step will be this coordinate
	for _, blizz := range v.Blizzards {
		nextStep := blizz.Peek()

		if v.IsBorder(nextStep) {
			nextStep = v.Wrap(nextStep)
		}

		if nextStep == c {
			return false
		}
	}

	return true
}

func (v *Valley) CheckLocPartTwo(c, start, end Coord) bool {
	// checks if the given coordinate is valid. What determines validity
	// is checking the NEXT step of each blizzard and determining if they
	// would be moving to this coordinate.

	// first, let's see if the given coordinate is the entrance or exit
	if c == start {
		return false // we don't want to go back to the entrance. Maybe?
	}

	if c == end {
		return true // if we can, we always want to go to the exit
	}

	// check to see if any blizzards' next step will be this coordinate
	for _, blizz := range v.Blizzards {
		nextStep := blizz.Peek()

		if v.IsBorder(nextStep) {
			nextStep = v.Wrap(nextStep)
		}

		if nextStep == c {
			return false
		}
	}

	return true
}

func (v *Valley) GetSurroundingForBlizzard(c Coord) []Coord {
	// this will return all 4 directions of the given coordinate.
	// This should ONLY be used for blizzard movement, because this
	// will wrap around if necessary.

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
			check = v.Wrap(check)
		}
		retval = append(retval, check)
	}
	return retval
}

func (v *Valley) GetSurroundingForElf(c Coord) []Coord {
	// this will return all four coordinates THAT AREN'T A WALL.
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

		if v.IsBorder(check) && check != v.Exit && check != v.Entrance {
			// this is a wall, forget it
			continue
		}
		retval = append(retval, check)
	}
	return retval
}

func (v *Valley) GetSurroundingForElfPartTwo(c, start, end Coord) []Coord {
	// this will return all four coordinates THAT AREN'T A WALL.
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

		if v.IsBorder(check) && check != end && check != start {
			// this is a wall, forget it
			continue
		}
		retval = append(retval, check)
	}
	return retval
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

	surrounding := v.GetSurroundingForElf(c)

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

func (v *Valley) GetValidDirectionsPartTwo(c, start, end Coord) []Coord {
	// this will return a list of valid directions, including the wait
	// direction (which is the same coord)
	dir := make([]Coord, 0)

	if c == start {
		if start == v.Entrance {
			// we can always wait here, since it's the entrance.
			dir = append(dir, c)

			// we can only go south, so let's see if we even can.
			checkSouth := c
			checkSouth.row++
			if v.CheckLocPartTwo(checkSouth, start, end) {
				dir = append(dir, checkSouth)
			}

			return dir
		} else if start == v.Exit {
			// we can only go north
			dir = append(dir, c)

			checkNorth := c
			checkNorth.row--
			if v.CheckLocPartTwo(checkNorth, start, end) {
				dir = append(dir, checkNorth)
			}

			return dir
		}
	}

	if c == end {
		return dir // empty, we made it
	}

	// let's see if we can wait in our current location
	if v.CheckLocPartTwo(c, start, end) {
		dir = append(dir, c)
	}

	// going to simplify this...

	surrounding := v.GetSurroundingForElfPartTwo(c, start, end)

	for _, value := range surrounding {
		if value == end {
			return []Coord{value}
		}
		if v.CheckLocPartTwo(value, start, end) {
			dir = append(dir, value)
		}
	}

	return dir
}

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
