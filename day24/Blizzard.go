package main

import "fmt"

const (
	north = iota
	south
	east
	west
)

type Blizzard struct {
	direction int
	pos       Coord
}

func (b *Blizzard) String() string {
	// returns a printable string, though this is used
	// mainly for the poor man's hash I'm about to write
	return fmt.Sprintf("%d~%d", b.pos.col, b.pos.row)
}

func NewBlizzard(char rune, loc Coord) Blizzard {
	// given the character and the coordinate, will return
	// a Blizzard object
	b := Blizzard{
		pos: loc,
	}
	switch char {
	case '>':
		b.direction = east
	case '<':
		b.direction = west
	case '^':
		b.direction = north
	case 'v':
		b.direction = south
	default:
		b.direction = 99 // this is an error
	}
	return b
}

func (b Blizzard) Peek() Coord {
	// this will not MOVE the blizzard, but rather will return
	// the proposed coordinate of where the blizzard will move.
	// note that this will NOT wrap around the map, that is the
	// job of the map to place the blizzard accordingly.
	c := Coord{
		row: b.pos.row,
		col: b.pos.col,
	}
	switch b.direction {
	case north:
		c.row--
	case south:
		c.row++
	case east:
		c.col++
	case west:
		c.col--
	}
	return c
}

func (b *Blizzard) Move(c Coord) {
	// this function simply sets the blizzard to the new location.
	// the map will ultimately decide if this is a valid move or not
	b.pos = c

	// this is probably not necessary.
	// One line of code for a function? come on dan
}
