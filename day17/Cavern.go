package main

import (
	"errors"
	"fmt"
	"math"
	"math/bits"
)

/*
	This object will hold all the functions surrounding
	the "cavern" object, which will serve as an expandable
	2-dimensional array/slice.

	A cavern will always be a width of 7 and an infinite length
*/

const width int = 7

type Cavern struct {
	tunnel [][width]string
}

func NewCavern() Cavern {
	// constructs the Cavern object
	c := Cavern{
		tunnel: make([][width]string, 0),
	}
	// we will start with a layer of 5, because why not
	c.GrowCavern(5)
	return c
}

func (c *Cavern) IsFree(p Point) bool {
	// this function is a boolean which, given a point,
	// will determine if the point is free or not. Additionally,
	// if the point is higher than the allocated Cavern, it will
	// automatically grow the cavern to accommodate (because if
	// someone calls this function we are checking to see if we
	// can fill that space with a rock)
	if (len(c.tunnel) - 1) < p.Y {
		// if we are above the height of the cavern,
		// then by default the cavern is empty so grow
		// the cavern by as many times as necessary to
		// grow it and return true
		c.GrowCavern(p.Y - (len(c.tunnel) - 1))
		return true
	}

	if p.X < 0 || p.Y < 0 || p.X > (width-1) {
		// is this point at or beyond the left wall, ground or right wall
		return false
	}

	// now that we're in bounds, let's see if a rock exists here
	if c.tunnel[p.Y][p.X] == "#" {
		return false
	}

	return true
}

// objects will fall down the Y cavern to 0

func (c *Cavern) GrowCavern(addHeight int) {
	// grows the cavern by addHeight
	for i := 0; i < addHeight; i++ {
		row := [7]string{".", ".", ".", ".", ".", ".", "."}
		c.tunnel = append(c.tunnel, row)
	}
}

func (c *Cavern) DrawRock(t TransitionalRock) {
	for _, v := range t.rockPieces {
		c.tunnel[v.Y][v.X] = "#"
	}
}

func (c *Cavern) PrintCavern() {
	// this function simply prints out the cavern as
	// visualized. Nothing more.

	// print the ceiling
	fmt.Printf("       |")
	for i := 0; i < width; i++ {
		fmt.Printf("~")
	}
	fmt.Printf("|\n")

	// now each item in the tunnel
	for i := len(c.tunnel) - 1; i >= 0; i-- {
		fmt.Printf("%6d ", i+1)
		fmt.Printf("|") // wall
		for j := 0; j < width; j++ {
			fmt.Printf("%s", c.tunnel[i][j])
		}
		fmt.Printf("|\n")
	}

	// just print the floor, i'm tired don't make me loop anymore
	fmt.Printf("     0 +-------+\n")
}

/*
	Allow me to introduce the concept of the "Transitional Rock".
	This object interacts directly with the Cavern object because
	it maintains the state and current position of EVERY part of
	a rock and where it exists in the Cavern as it falls.
*/

type TransitionalRock struct {
	rockPieces []Point // X,Y coordinates of every single part of the rock
	c          *Cavern // the cavern this rock is a part of
}

func NewTransitionalRock(highestY int, thisRock Rock, ca *Cavern) TransitionalRock {
	// this will calculate everything necessary to build the
	// proper transitional Rock. We know that it will start building
	// the rock 2 points away from the left wall and 3 points higher
	// than the highest point in the map including fallen rocks. It
	// is not the responsibility of this function to determine the
	// highest rock in the map.

	// X starting point is 2
	// Y starting point is highestY
	// Find the focal point
	XFocal := int(math.Abs(float64(thisRock.GreatestLeftOffset()))) + 2
	YFocal := (int(math.Abs(float64(thisRock.LowestOffset()))) + 3) + highestY

	rockHeight := YFocal + thisRock.HighestOffset()

	if rockHeight > (len(ca.tunnel) - 1) {
		disp := rockHeight - (len(ca.tunnel) - 1)
		ca.GrowCavern(disp)
	}

	p := make([]Point, 0)
	// with this focal point, we can build everything
	for _, v := range thisRock.offSets {
		p = append(p, Point{
			X: v.X + XFocal,
			Y: v.Y + YFocal,
		})
	}
	return TransitionalRock{
		rockPieces: p,
		c:          ca,
	}
}

func (t *TransitionalRock) Move(dir int) bool {
	// this will try to move the Transitional Rock in
	// the given direction. It will return true if the
	// rock was able to move in the provided direction,
	// and false if there is something obstructing it.
	p := make([]Point, 0)
	switch dir {
	case Right:
		// to move right, we add 1 to all the X values
		for _, v := range t.rockPieces {
			tp := Point{
				X: v.X + 1,
				Y: v.Y,
			}
			if !t.c.IsFree(tp) {
				return false // stop processing, we hit a wall
			}
			p = append(p, tp)
		}
	case Left:
		// to move left, we subtract 1 from all the X values
		for _, v := range t.rockPieces {
			tp := Point{
				X: v.X - 1,
				Y: v.Y,
			}
			if !t.c.IsFree(tp) {
				return false // stop processing, hit the left wall
			}
			p = append(p, tp)
		}
	case Down:
		// to move down, subtract 1 from all the Y values
		for _, v := range t.rockPieces {
			tp := Point{
				X: v.X,
				Y: v.Y - 1,
			}
			if !t.c.IsFree(tp) {
				return false // stop processing, we hit a rock or the ground
			}
			p = append(p, tp)
		}
	default:
		// if you got here you errored out, don't do anything
		return false
	}

	// if you got here, update t with the new coords
	t.rockPieces = nil
	for _, v := range p {
		t.rockPieces = append(t.rockPieces, v)
	}
	return true
}

func (c *Cavern) FindHighest() (int, error) {
	// this will check to see what is the
	// highest peak.
	for i := 0; i < len(c.tunnel); i++ {
		found := false
		for j := 0; j < width; j++ {
			if c.tunnel[i][j] == "#" {
				found = true
				break
			}
		}
		if !found {
			return i, nil
		}
	}
	return 0, errors.New("could not compute highest")
}

/*
	Signatures section

	This is where I will develop the concept of rock
	"signatures". Since the only thing I care about are
	spaces where rocks actually exist, I can use binary
	strings to obtain a signature of the rock row using
	some bitwise math.
*/

func RowSig(row [7]string) uint {
	// this function takes one row and returns
	// the binary value of it as an integer. It
	// is essentially a "poor man's hash" of sorts
	// that becomes more accurate the more it is
	// concatenated
	var retval uint = 0b0
	for _, v := range row {
		retval <<= 1
		if v == "#" {
			// rock
			retval |= 1
		} else {
			retval |= 0
		}
	}
	return retval
}

func ConcatSig(a, b uint) uint {
	// simply concatenates two signatures
	// concat a to b, shift Len(a)
	return (a << bits.Len(b)) | b
}

// and now, get a signature of the caverns

func (c *Cavern) GetSignature(rows, start int) uint {
	// given the start row, get the signature of
	// as many rows as provided in the "rows" variable
	var retsig uint = 0b0
	for i := start; i < start+rows; i++ {
		r := RowSig(c.tunnel[i])
		retsig = ConcatSig(retsig, r)
	}
	return retsig
}
