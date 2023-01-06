package main

import "fmt"

type Cave struct {
	coords     [][]string
	mapX, mapY int
}

// func NewCave(Xnum, Ynum int) Cave {
// 	// This function builds a cave of X*Y dimensions.
// 	// Fills it up with "." to signify air.
// 	c := Cave{}
// 	for i := 0; i < Ynum; i++ {
// 		X := make([]string, Xnum)
// 		for j := range X {
// 			X[j] = "."
// 		}
// 		c.coords = append(c.coords, X)
// 	}
// 	return c
// }

func (c Cave) MapCoords(fromX, fromY int) (toX, toY int) {
	// this function takes provided coordinates and maps
	// it to "real" coordinates.
	toX = fromX - c.mapX
	// toY = fromY - c.mapY
	toY = fromY
	return
}

func (c *Cave) DrawRocks(rocks [][2]int) {
	// this is an internal function which takes
	// the result of a rock pattern's DrawRock()
	// function and simply applies it.
	for _, v := range rocks {
		x, y := c.MapCoords(v[0], v[1])
		c.coords[y][x] = "#"
	}
}

func NewCave(r Rocks) Cave {
	// this function properly builds a cave based on
	// the rocks that are given to it. It will set up
	// the cave object properly after performing some
	// calculations on the rocks to determine the mapping.
	c := Cave{}
	var err error
	err = r.FindLowest()
	if err != nil {
		panic(fmt.Sprintf("Received error: %s", err))
	}
	c.mapX, c.mapY = r.lX, r.lY

	highX, highY, err := r.FindHighest()
	if err != nil {
		panic(fmt.Sprintf("Received error: %s", err))
	}
	xTop, _ := c.MapCoords(highX, highY)
	for y := 0; y < highY; y++ {
		air := make([]string, xTop)
		for z := range air {
			air[z] = "."
		}
		c.coords = append(c.coords, air)
	}
	// now draw the rocks
	// fmt.Printf("Length of Y: %d\n", len(c.coords))
	for _, v := range r.rset {
		drawThese := v.DrawRock()
		c.DrawRocks(drawThese)
	}

	return c
}

func NewCavePart2(r Rocks) Cave {
	// this function properly builds a cave based on
	// the rocks that are given to it. It will set up
	// the cave object properly after performing some
	// calculations on the rocks to determine the mapping.
	// NOTE: This is specific to Part 2, which also gives it
	// a ground which is +2 from the highest Y level.
	c := Cave{}
	var err error
	err = r.FindLowest()
	if err != nil {
		panic(fmt.Sprintf("Received error: %s", err))
	}
	c.mapX, c.mapY = r.lX, r.lY

	highX, highY, err := r.FindHighest()
	if err != nil {
		panic(fmt.Sprintf("Received error: %s", err))
	}
	// Add 1 more layer of air
	highY++

	xTop, _ := c.MapCoords(highX, highY)
	for y := 0; y < highY; y++ {
		air := make([]string, xTop)
		for z := range air {
			air[z] = "."
		}
		c.coords = append(c.coords, air)
	}

	// draw the ground
	ground := make([]string, xTop)
	for z := range ground {
		ground[z] = "#"
	}
	c.coords = append(c.coords, ground)

	// now draw the rocks
	// fmt.Printf("Length of Y: %d\n", len(c.coords))
	for _, v := range r.rset {
		drawThese := v.DrawRock()
		c.DrawRocks(drawThese)
	}

	return c
}

func DrawLeft(c *Cave) {
	// this function adds 1 column to the left of
	// the cave, including the ground. It also
	// updates the offset properly. This is necessary
	// since we are re-building the map and the indices
	// will be different.
	for i, v := range c.coords {
		// looping through the Y axis
		var a []string
		if i == len(c.coords)-1 {
			// ground
			a = append(a, "#")
		} else {
			a = append(a, ".")
		}
		a = append(a, v...)
		c.coords[i] = a
	}
	c.mapX-- // update the offset
}

func DrawRight(c *Cave) {
	// this function adds 1 column to the right of the
	// cave, including the ground. As this doesn't matter
	// much in terms of offsetting, we just add to the column
	// to the right but won't need to update the offset.
	for i, v := range c.coords {
		var a string
		if i == len(c.coords)-1 {
			// ground
			a = "#"
		} else {
			a = "."
		}
		v = append(v, a)
		c.coords[i] = v
	}
}

func PrintCave(c *Cave) {
	// prints the cave
	var caveRow string
	for _, v := range c.coords {
		for _, w := range v {
			caveRow += w
		}
		fmt.Printf("%s\n", caveRow)
		caveRow = ""
	}
}

func CountSand(c *Cave) int {
	// counts the grains of sand in the
	// current drawn map
	counter := 0
	for _, v := range c.coords {
		for _, w := range v {
			if w == "o" {
				counter++
			}
		}
	}
	return counter
}
