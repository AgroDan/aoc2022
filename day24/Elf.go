package main

/*

This package will handle the functions specific to the
character that we control. This will attempt to find all
possible actions they can take given the location.

*/

type Elf struct {
	loc     Coord
	minutes int
}

func NewElf(c Coord, m int) Elf {
	return Elf{
		loc:     c,
		minutes: m,
	}
}
