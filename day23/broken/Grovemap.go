package main

import "fmt"

type GroveMap struct {
	p         [][]rune
	MoveQueue []rune
}

func NewGroveMap(size int) GroveMap {
	// this function builds a BLANK Grovemap.
	// and by blank, I mean filled with '.'
	// characters. This should not be used to
	// copy a grovemap to another, bigger one.
	rows := make([][]rune, 0)
	for i := 0; i < size; i++ {
		cols := make([]rune, 0)
		for ii := 0; ii < size; ii++ {
			cols = append(cols, '.')
		}
		rows = append(rows, cols)
	}
	return GroveMap{
		p:         rows,
		MoveQueue: []rune{'N', 'S', 'W', 'E'},
	}
}

func (g *GroveMap) Rotate() {
	// this function will rotate the move queue.
	g.MoveQueue = append(g.MoveQueue, g.MoveQueue[0])
	g.MoveQueue = g.MoveQueue[1:]
}

// This function shouldn't really grow too much. because if we do, we
// need to calculate for every known elf. Of course that basically means
// for every grow order we add, we need to add 1 to the X and Y axis of
// all the elves. Maybe that's ok?
func (g *GroveMap) GrowMap(borderSize int) {
	// this function will grow the maps border by borderSize.
	preRows := make([][]rune, 0)
	midRows := make([][]rune, 0)
	postRows := make([][]rune, 0)

	for i := 0; i < borderSize; i++ {
		preCols := make([]rune, 0)
		postCols := make([]rune, 0)
		for ii := 0; ii < (borderSize*2)+len(g.p); ii++ {
			// interesting bug I noted for posterity. I used to just use
			// one variable, "cols" to append a column of dots, then added
			// this variable to both the preRows and postRows. This wound
			// up causing REALLY weird behavior, in which if an elf enters
			// a row within one of these rows it would reflect TWICE. This
			// is because appending these arrays are actually copied by
			// reference, so any changes to one row will be reflected in
			// another row because those rows are pointing to the exact same
			// place in memory! Whoops! That was super difficult to debug for
			// a while.
			preCols = append(preCols, '.')
			postCols = append(postCols, '.')
		}
		preRows = append(preRows, preCols)
		postRows = append(postRows, postCols)
	}

	for _, v := range g.p {
		row := make([]rune, 0)
		for i := 0; i < borderSize; i++ {
			row = append(row, '.')
		}
		for _, w := range v {
			row = append(row, w)
		}
		for i := 0; i < borderSize; i++ {
			row = append(row, '.')
		}
		midRows = append(midRows, row)
	}

	e := make([][]rune, 0)
	e = append(e, preRows...)
	e = append(e, midRows...)
	e = append(e, postRows...)
	g.p = e
}

func (g *GroveMap) ShrinkMap() {
	// this function will shrink the map by 1. It is up to the
	// caller to determine if any elves are at the border! This
	// will blow them away if so!

	// top border
	g.p = g.p[1:]

	// bottom border
	g.p = g.p[:len(g.p)-1]

	for _, v := range g.p {
		// left border
		v = v[1:]

		// right border
		v = v[:len(v)-1]
	}
}

// TODO -- figure out if this is even necessary
func (g GroveMap) ElvesAtBorder() bool {
	// this function checks to see if any elves
	// are on a border square. If so, return true.
	for _, v := range g.p[0] {
		if v == '#' {
			return true
		}
	}
	for _, v := range g.p[1 : len(g.p)-1] {
		if v[0] == '#' || v[len(v)-1] == '#' {
			return true
		}
	}
	for _, v := range g.p[len(g.p)-1] {
		if v == '#' {
			return true
		}
	}
	return false
}

// this will help with the whole "rectangle shape" bit
func (g GroveMap) ElfAtEdge(which rune) bool {
	switch which {
	case 'N':
		for _, v := range g.p[0] {
			if v == '#' {
				return true
			}
		}
	case 'S':
		for _, v := range g.p[len(g.p)-1] {
			if v == '#' {
				return true
			}
		}
	case 'E':
		for _, v := range g.p {
			if v[len(v)-1] == '#' {
				return true
			}
		}
	case 'W':
		for _, v := range g.p {
			if v[0] == '#' {
				return true
			}
		}
	}
	return false
}

func (g *GroveMap) ShaveDirection(which rune, e *ElfStack) {
	// this function will remove a side of the elf map and
	// update all the elves X/Y coords
	// good luck dan!
	switch which {
	case 'N':
		g.p = g.p[1:]
		for _, v := range *e {
			v.Loc.Y--
		}
	case 'S':
		g.p = g.p[:len(g.p)-1]
	case 'E':
		for i := range g.p {
			g.p[i] = g.p[i][:len(g.p[i])-1]
		}
	case 'W':
		for i := range g.p {
			g.p[i] = g.p[i][1:]
		}
		for _, v := range *e {
			v.Loc.X--
		}
	}
}

// func (g *GroveMap) CheckCoord(c Coord) rune {
// 	// this function will NOT ONLY return the value of
// 	// the coordinate, but also grow the map if we are
// 	// trying to access a part of the map that is larger
// 	// than the known map area.
// 	if c.X < 0 {
// 		g.GrowMap(c.X * -1)
// 		return '.'
// 	} else if c.X >= len(g.p[c.Y]) {
// 		g.GrowMap(c.X - (len(g.p[c.Y]) - 1))
// 		return '.'
// 	} else if c.Y < 0 {
// 		g.GrowMap(c.Y * -1)
// 		return '.'
// 	} else if c.Y >= len(g.p) {
// 		g.GrowMap(c.Y - (len(g.p) - 1))
// 		return '.'
// 	}

// 	// by now, we can just check the value of the given
// 	// coordinates.
// 	return g.p[c.Y][c.X]
// }

func (g *GroveMap) CheckCoord(c Coord) rune {
	// this one will NOT grow the map. I figured that shouldn't
	// be the responsibility of this function.
	if c.X < 0 || c.X >= len(g.p[c.Y]) ||
		c.Y < 0 || c.Y >= len(g.p) {
		return '.'
	}
	return g.p[c.Y][c.X]
}

func PrintMap(g *GroveMap) {
	for _, v := range g.p {
		for _, w := range v {
			fmt.Printf("%c", w)
		}
		fmt.Printf("\n")
	}
}

func (g *GroveMap) MoveElf(e *Elf) {
	// this function moves the elf on the grovemap. Note that this
	// does not do any validation checking to determine if an elf
	// _should_ move, just moves it based on the Elf's proposal.
	// fmt.Printf("Moving Elf:\n")
	// e.PrintElf()
	// fmt.Printf("Current Proposal: X: %d, Y: %d\n", e.Propose.X, e.Propose.Y)
	// fmt.Printf("%#v\n", g)
	// fmt.Printf("Length of Y: %d\n", len(g.p))
	g.p[e.Loc.Y][e.Loc.X] = '.'
	g.p[e.Propose.Y][e.Propose.X] = '#'
	// fmt.Printf("%#v\n", g)
	// now reset the Elf
	e.Loc.X = e.Propose.X
	e.Loc.Y = e.Propose.Y

	e.Reset()
}

func (g *GroveMap) CountEmpty() int {
	// counts the empty dots in the map. This is only accurate for
	// the challenge if the map is properly shrunk down.
	count := 0

	for _, v := range g.p {
		for _, w := range v {
			if w == '.' {
				count++
			}
		}
	}
	return count
}
