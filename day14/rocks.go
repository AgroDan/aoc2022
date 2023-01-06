package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// This library will house the "rock"
// object, which is parsed though the
// input feed. Each "rock" line will
// have two values: start and end.

type Rock struct {
	pos []map[string]int
}

// find the lowest X value on the map,
// this becomes 0. so if lowest is 500,
// we will map each supplied number to
// num - 500

func NewRock(rockstring string) Rock {
	// this function takes the string that looks like
	// 100,50 -> 100,60 -> 90,60 (etc) and generates a
	// rock object based on this.
	r := Rock{}
	rocks := strings.Split(rockstring, " -> ")

	for _, rock := range rocks {
		// for each rock coord
		c := strings.Split(rock, ",")
		x, err := strconv.Atoi(c[0])
		if err != nil {
			panic(fmt.Sprintf("Could not convert %s", c[0]))
		}
		y, err := strconv.Atoi(c[1])
		if err != nil {
			panic(fmt.Sprintf("Could not convert %s", c[1]))
		}

		xy := make(map[string]int)
		xy["X"] = x
		xy["Y"] = y
		r.pos = append(r.pos, xy)
	}
	return r
}

func (r Rock) DrawRock() [][2]int {
	// this function returns a slice of all possible indices
	// that a rock takes up on the given map. This is designed
	// to loop against and populate the Cave with rock icons.
	// X is [][0], Y is [][1]
	res := make([][2]int, 0)
	for i := range r.pos {
		if i >= len(r.pos)-1 {
			break
		}
		// if you are here and didn't break, then each item in r.pos
		// has a successor. set them to this/that for ease of reference
		this := r.pos[i]
		that := r.pos[i+1]

		if this["X"] != that["X"] {
			// the difference is the X index
			if this["X"] > that["X"] {
				// x is descending
				for j := this["X"]; j >= that["X"]; j-- {
					x := j
					y := this["Y"]
					temp := [2]int{x, y}
					res = append(res, temp)
				}
			} else {
				// x is ascending
				for j := this["X"]; j <= that["X"]; j++ {
					x := j
					y := this["Y"]
					temp := [2]int{x, y}
					res = append(res, temp)
				}
			}
		}
		if this["Y"] != that["Y"] {
			// the difference is the Y index
			if this["Y"] > that["Y"] {
				// y is descending
				for j := this["Y"]; j >= that["Y"]; j-- {
					x := this["X"]
					y := j
					temp := [2]int{x, y}
					res = append(res, temp)
				}
			} else {
				// y is ascending
				for j := this["Y"]; j <= that["Y"]; j++ {
					x := this["X"]
					y := j
					temp := [2]int{x, y}
					res = append(res, temp)
				}
			}
		}
	}
	return res
}

// Rocks are a collection of the Rock type.
// This allows me to make some calculations on the whole.
type Rocks struct {
	rset   []*Rock
	lX, lY int // lowest X and Y coordinates in this set
}

func (r *Rocks) FindLowest() error {
	// This function finds the smallest
	// X or Y coord of the set of rocks assigned to it.
	// This can be used to map a number to X=0 or Y=0
	if len(r.rset) == 0 {
		return errors.New("blank rock set")
	}
	if len(r.rset[0].pos) == 0 {
		return errors.New("invalid rock defined")
	}

	xfind := r.rset[0].pos[0]["X"]
	yfind := r.rset[0].pos[0]["Y"]
	for _, v := range r.rset {
		for _, w := range v.pos {
			if w["X"] < xfind {
				xfind = w["X"]
			}
			if w["Y"] < yfind {
				yfind = w["Y"]
			}
		}
	}
	r.lX, r.lY = xfind, yfind
	return nil
}

func (r Rocks) FindHighest() (X, Y int, e error) {
	// this will find the highest possible X/Y
	// coordinates
	if len(r.rset) == 0 {
		X, Y = 0, 0
		e = errors.New("blank rock set")
		return
	}
	if len(r.rset[0].pos) == 0 {
		X, Y = 0, 0
		e = errors.New("invalid rock defined")
		return
	}

	xfind := 0
	yfind := 0
	for _, v := range r.rset {
		for _, w := range v.pos {
			if w["X"] > xfind {
				xfind = w["X"]
			}
			if w["Y"] > yfind {
				yfind = w["Y"]
			}
		}
	}
	X, Y = (xfind + 1), (yfind + 1) // because index starts at 0
	e = nil
	return
}

func NewRockCollection(input []byte) Rocks {
	// this takes the full byte string
	r := Rocks{}
	str := string(input)
	// trim the excess
	str = strings.TrimSpace(str)
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		rock := NewRock(line)
		r.rset = append(r.rset, &rock)
	}
	err := r.FindLowest()
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err))
	}
	return r
}
