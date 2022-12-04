package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// could probably make an initialization function but ¯\_(ツ)_/¯

type ElfAssignment struct {
	leftNum  int
	rightNum int
}

func (e *ElfAssignment) CreateAssignment(sections string) {
	var err error
	s := strings.Split(sections, "-")
	e.leftNum, err = strconv.Atoi(s[0])
	if err != nil {
		fmt.Println("Error on leftNum decl:", err)
	}

	e.rightNum, err = strconv.Atoi(s[1])
	if err != nil {
		fmt.Println("Error on rightNum decl:", err)
	}
}

func IsFullyOverlapped(leftElf ElfAssignment, rightElf ElfAssignment) bool {
	// this will return True if one elf is fully overlapped

	if leftElf.leftNum >= rightElf.leftNum && leftElf.rightNum <= rightElf.rightNum {
		return true
	}

	if rightElf.leftNum >= leftElf.leftNum && rightElf.rightNum <= leftElf.rightNum {
		return true
	}
	return false
}

// this is an atrocity and I am so sorry this exists. I will not remove it
// as punishment for my terrible coding practices.
//
// func splitIntoNumbers(elfAssignments string) (leftSide[] int, rightSide[] int) {
// 	// This takes the string and parses the line into integers
// 	// that can be worked on in a separate function
// 	sections := strings.Split(elfAssignments, ",")
// 	leftSideString := strings.Split(sections[0], "-")
// 	leftSide[0], err := strconv.Atoi(leftSideString[0])
// 	leftSide[1], err := strconv.Atoi(leftSideString[1])
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	rightSideString := strings.Split(sections[1], "-")
// 	rightSide[0], err := strconv.Atoi(rightSideString[0])
// 	rightSide[1], err := strconv.Atoi(rightSideString[1])
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	return
// }

func main() {
	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	readFile.Close()

	// test := strings.Split(elfAssignments[0], ",")
	// fmt.Println("Left side:", test[0], "; right side:", test[1])

	var counter int

	for _, v := range lines {
		a := strings.Split(v, ",")
		firstElf := ElfAssignment{}
		secondElf := ElfAssignment{}

		firstElf.CreateAssignment(a[0])
		secondElf.CreateAssignment(a[1])
		if IsFullyOverlapped(firstElf, secondElf) {
			counter++
		}
	}

	fmt.Printf("Count of overlapped assignments: %d\n", counter)
	// a := strings.Split(lines[2], ",")
	// b1 := ElfAssignment{}
	// b2 := ElfAssignment{}
	// b1.CreateAssignment(a[0])
	// b2.CreateAssignment(a[1])

	// fmt.Println("Do they overlap?")

	// if IsFullyOverlapped(b1, b2) {
	// 	fmt.Println("Yes they are.")
	// } else {
	// 	fmt.Println("Nope")
	// }
}
