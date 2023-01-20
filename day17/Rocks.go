package main

/*
	This file will hold all the rocks. A rock is
	built with a known "focal point" combined with
	a set of offsets to show how to print each rock.
	So a rock will contain an offset for each block
	that a rock takes up which is to be calculated
	from the focal point.
*/

type Point struct {
	X, Y int
}

type Rock struct {
	offSets []Point
}

// remember that the "UP" direction is +Y, and down is -Y.

func NewPoint(setX, setY int) Point {
	return Point{
		X: setX,
		Y: setY,
	}
}

func NewRock(offsets ...Point) Rock {
	// This builds a rock object.
	p := make([]Point, 0)
	for _, v := range offsets {
		p = append(p, v)
	}
	return Rock{
		offSets: p,
	}
}

func (r *Rock) GreatestLeftOffset() int {
	// this function returns the greatest left offset of the
	// rock. With some calculation you can determine if we are
	// up against a wall or not.
	lowest := 0
	for _, v := range r.offSets {
		if v.X < lowest {
			lowest = v.X
		}
	}
	return lowest
}

func (r *Rock) GreatestRightOffset() int {
	// opposite of GreatestLeftOffset(). Doing this in 2 separate
	// functions for simplicity
	highest := 0
	for _, v := range r.offSets {
		if v.X > highest {
			highest = v.X
		}
	}
	return highest
}

func (r *Rock) LowestOffset() int {
	// returns the lowest point of the rock, useful for determining
	// if we hit the ground (or another rock) or not.
	lowest := 0
	for _, v := range r.offSets {
		if v.Y < lowest {
			lowest = v.Y
		}
	}
	return lowest
}

func (r *Rock) HighestOffset() int {
	// returns the highest point of the rock, useful for determining
	// if we need to bump up the height of the cavern or not
	highest := 0
	for _, v := range r.offSets {
		if v.Y > highest {
			highest = v.Y
		}
	}
	return highest
}

func HorizontalLine() Rock {
	// this function sets up a horizontal line
	o := []Point{
		NewPoint(0, 0),
		NewPoint(1, 0),
		NewPoint(2, 0),
		NewPoint(3, 0),
	}
	return NewRock(o...)
}

func PlusSign() Rock {
	// this function sets up the plus sign + rock
	o := []Point{
		NewPoint(0, 0),
		NewPoint(1, 0),
		NewPoint(-1, 0),
		NewPoint(0, 1),
		NewPoint(0, -1),
	}
	return NewRock(o...)
}

func Corner() Rock {
	// this function sets up the "Corner" backwards L rock
	o := []Point{
		NewPoint(0, 0),
		NewPoint(-1, 0),
		NewPoint(-2, 0),
		NewPoint(0, 1),
		NewPoint(0, 2),
	}
	return NewRock(o...)
}
func VerticalLine() Rock {
	// this function sets up the Vertical line
	o := []Point{
		NewPoint(0, 0),
		NewPoint(0, 1),
		NewPoint(0, 2),
		NewPoint(0, 3),
	}
	return NewRock(o...)
}

func Square() Rock {
	// this function sets up a square block
	o := []Point{
		NewPoint(0, 0),
		NewPoint(-1, 0),
		NewPoint(0, 1),
		NewPoint(-1, 1),
	}
	return NewRock(o...)
}
