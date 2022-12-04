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

func AnyOverlapAtAll(leftElf ElfAssignment, rightElf ElfAssignment) bool {
	// this will return True if there is ANY overlap at all.
	if leftElf.leftNum >= rightElf.leftNum && leftElf.leftNum <= rightElf.rightNum {
		return true
	}
	if leftElf.rightNum >= rightElf.leftNum && leftElf.rightNum <= rightElf.rightNum {
		return true
	}

	if rightElf.leftNum >= leftElf.leftNum && rightElf.leftNum <= leftElf.rightNum {
		return true
	}
	if rightElf.rightNum >= leftElf.leftNum && rightElf.rightNum <= leftElf.rightNum {
		return true
	}
	return false
}

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

	var counter int

	for _, v := range lines {
		a := strings.Split(v, ",")
		firstElf := ElfAssignment{}
		secondElf := ElfAssignment{}

		firstElf.CreateAssignment(a[0])
		secondElf.CreateAssignment(a[1])
		if AnyOverlapAtAll(firstElf, secondElf) {
			counter++
		}
	}

	fmt.Printf("Count of ANY overlapped assignments: %d\n", counter)

}
