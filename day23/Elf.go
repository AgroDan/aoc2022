package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// type Coord struct {
// 	r, c int
// }

var exists = struct{}{}

type Elves struct {
	set      map[string]struct{}
	proposed map[string]string // map[elf] = "proposed dest"
}

func NewElfStack() Elves {
	return Elves{
		set:      make(map[string]struct{}),
		proposed: make(map[string]string),
	}
}

func (e *Elves) ClearProposed() {
	// this function removes all the proposals.
	for k := range e.proposed {
		delete(e.proposed, k)
	}
}

func (e *Elves) SubmitProposal(fromRow, fromCol, toRow, toCol int) error {
	// this will submit a new proposal.
	setFromMap := fmt.Sprintf("%s~%s", strconv.Itoa(fromRow), strconv.Itoa(fromCol))
	setToMap := fmt.Sprintf("%s~%s", strconv.Itoa(toRow), strconv.Itoa(toCol))
	if _, ok := e.proposed[setFromMap]; ok {
		return errors.New("proposal already submitted")
	}
	e.proposed[setFromMap] = setToMap
	return nil
}

func (e *Elves) Add(R, C int) error {
	setMap := fmt.Sprintf("%s~%s", strconv.Itoa(R), strconv.Itoa(C))
	if _, ok := e.set[setMap]; ok {
		return errors.New("already exists")
	}
	// otherwise...
	e.set[setMap] = exists
	return nil
}

func (e *Elves) Remove(R, C int) error {
	setMap := fmt.Sprintf("%s~%s", strconv.Itoa(R), strconv.Itoa(C))
	if _, ok := e.set[setMap]; ok {
		delete(e.set, setMap)
		return nil
	}
	return errors.New("elf does not exist at that mapping")
}

func (e *Elves) Exists(R, C int) bool {
	// checks to see if an elf exists at the given coords
	setMap := fmt.Sprintf("%s~%s", strconv.Itoa(R), strconv.Itoa(C))
	if _, ok := e.set[setMap]; ok {
		return true
	} else {
		return false
	}
}

func (e *Elves) Move(fromR, fromC, toR, toC int) error {
	// this will move one elf from coordinates to additional
	// coordinates. Will throw an error if an elf exists in that
	// position already.
	if e.Exists(toR, toC) {
		return errors.New("cannot move elf, already exists")
	}
	e.Remove(fromR, fromC)
	e.Add(toR, toC)
	return nil
}

func (e *Elves) IsAlone(R, C int) bool {
	// this will check all 8 directions for the existance of an
	// elf. If one exists, return False.
	iterate := [8]string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}

	for _, v := range iterate {
		diffR, diffC := GetDirection(v)
		if e.Exists(R+diffR, C+diffC) {
			return false
		}
	}
	return true
}

func (e *Elves) GetFurthest(dir string) int {
	// this function will loop through all the known
	// elves and determine what is the highest number
	// in that direction, or lowest if going West or North.

	num := 0
	for k := range e.set {
		row, col, err := ExtractCoords(k)
		if err != nil {
			fmt.Printf("Error extracting coords: %s\n", err)
		}

		switch dir {
		case "N":
			if row < num {
				num = row
			}
		case "S":
			if row > num {
				num = row
			}
		case "E":
			if col > num {
				num = col
			}
		case "W":
			if col < num {
				num = col
			}
		}
	}
	return num
}

func (e *Elves) PrintElves() {
	// this will print the elves in a map, or at least try its best to.
	N := e.GetFurthest("N")
	S := e.GetFurthest("S")
	W := e.GetFurthest("W")
	E := e.GetFurthest("E")
	// fmt.Printf("N: %d, S: %d, W: %d, E: %d\n", N, S, W, E)
	height := (S - N) + 1
	width := (E - W) + 1
	grove := make([][]rune, 0)

	for i := 0; i < height; i++ {
		c := make([]rune, 0)
		for j := 0; j < width; j++ {
			c = append(c, '.')
		}
		grove = append(grove, c)
	}

	// now we have to find the offset of the lowest possible Western and
	// Northern elf. We can use this to find the "zero" position.
	HOffset := W * -1
	VOffset := N * -1

	// fmt.Printf("HOffset: %d\nVOffset: %d\n", HOffset, VOffset)
	// now let's loop over every elf
	for k := range e.set {
		row, col, err := ExtractCoords(k)
		if err != nil {
			fmt.Printf("Error extracting coords: %s\n", err)
		}
		grove[row+VOffset][col+HOffset] = '#'
	}

	for _, v := range grove {
		for _, w := range v {
			fmt.Printf("%c", w)
		}
		fmt.Printf("\n")
	}
}

func (e *Elves) GetEmptySpace() int {
	// given the amount of elves present in this struct, found out the
	// area of the spaces given up, and subtract by the amount of elves
	// present.
	N := e.GetFurthest("N")
	S := e.GetFurthest("S")
	W := e.GetFurthest("W")
	E := e.GetFurthest("E")
	// fmt.Printf("N: %d, S: %d, W: %d, E: %d\n", N, S, W, E)
	height := (S - N) + 1
	width := (E - W) + 1

	return (height * width) - len(e.set)
}

func ExtractCoords(coordString string) (row, col int, e error) {
	items := strings.Split(coordString, "~")
	row, col = 0, 0
	e = nil
	if len(items) > 1 {
		// working
		row, err := strconv.Atoi(items[0])
		if err != nil {
			e = errors.New("Could not convert row coords")
			return 0, 0, e
		}
		col, err := strconv.Atoi(items[1])
		if err != nil {
			e = errors.New("Could not convert col coords")
			return 0, 0, e
		}
		return row, col, e
	} else {
		e = errors.New("Could not split coords")
		return
	}
}

func GetDirection(which string) (row, col int) {
	row, col = 0, 0
	switch which {
	case "N":
		row = -1
		col = 0
	case "S":
		row = 1
		col = 0
	case "W":
		row = 0
		col = -1
	case "E":
		row = 0
		col = 1
	case "NW":
		row = -1
		col = -1
	case "NE":
		row = -1
		col = 1
	case "SW":
		row = 1
		col = -1
	case "SE":
		row = 1
		col = 1
	}
	return
}

func CheckDirection(which string) [3]string {
	// this function is a helper function that just returns
	// the three directiosn to check if we are going in one
	// particular direction.
	switch which {
	case "N":
		return [3]string{"NW", "N", "NE"}
	case "S":
		return [3]string{"SW", "S", "SE"}
	case "E":
		return [3]string{"NE", "E", "SE"}
	case "W":
		return [3]string{"NW", "W", "SW"}
	}
	return [3]string{}
}
