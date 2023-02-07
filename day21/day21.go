package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

func ParseLines(l []string) map[string]*Monkey {
	// this function parses all the lines in the
	// file and creates the MonkeyMap object

	MonkeyMap := make(map[string]*Monkey)
	for _, v := range l {
		m := ReadMonkey(v)
		MonkeyMap[m.Name] = &m
	}
	return MonkeyMap
}

func GetIntended(knownResult, primary, secondary int, whichVar bool, operation rune) int {
	// this is a weird (but necessary) function. Given the operation, it will
	// determine at the primary or secondary number should be to equal the
	// knownResult. the whichVar parameter points to which is the _empty_ variable
	// that we need to find out. 0 is primary, 1 is secondary.
	switch operation {
	case '+':
		if !whichVar { // == 0
			// x + secondary = knownResult
			return knownResult - secondary
		} else {
			// primary + x == knownResult
			return knownResult - primary
		}
	case '-':
		if !whichVar { // == 0
			// x - secondary = knownResult
			return knownResult + secondary
		} else {
			// primary - x = knownResult
			return (knownResult - primary) / -1
		}
	case '/':
		if !whichVar { // == 0
			// x / secondary = knownResult
			return secondary * knownResult
		} else {
			// primary / x = knownResult
			return primary / knownResult
		}
	case '*':
		if !whichVar { // == 0
			// x * secondary = knownResult
			return knownResult / secondary
		} else {
			// primary * x = knownResult
			return knownResult / primary
		}
	default:
		return 0
	}
}

func Dive(m map[string]*Monkey, whichMonkey string, knownValue int) int {
	// this function will traverse the entire string to
	// determine the root value of the humn item.
	if whichMonkey == "humn" {
		return knownValue
	}

	// otherwise, let's dive.
	if m[whichMonkey].JobStatus {
		panic("brick wall")
	}

	var primOrSecondary bool
	if HumanElement(m[m[whichMonkey].PrimaryMonkey].Name, m) {
		// solve for x on the primary monkey side
		primOrSecondary = false
	} else if HumanElement(m[m[whichMonkey].SecondaryMonkey].Name, m) {
		// solve for x on the secondary monkey side
		primOrSecondary = true
	} else {
		panic("No human found")
	}

	var newKnownValue int
	if !primOrSecondary {
		// this is the numeric value of this monkey specifically
		secondaryValue := findResult(m, m[m[whichMonkey].SecondaryMonkey].Name)
		newKnownValue = GetIntended(knownValue, 0, secondaryValue, primOrSecondary, m[whichMonkey].Operation)
		return Dive(m, m[m[whichMonkey].PrimaryMonkey].Name, newKnownValue)
	} else {
		primaryValue := findResult(m, m[m[whichMonkey].PrimaryMonkey].Name)
		newKnownValue = GetIntended(knownValue, primaryValue, 0, primOrSecondary, m[whichMonkey].Operation)
		return Dive(m, m[m[whichMonkey].SecondaryMonkey].Name, newKnownValue)
	}
}

func monkeyOperate(primary, secondary int, op rune) int {
	// this function will determine what to do when given a rune
	// to read as the operation in general
	switch op {
	case '+':
		return primary + secondary
	case '-':
		return primary - secondary
	case '/':
		return primary / secondary
	case '*':
		return primary * secondary
	}
	return 0
}

func findResult(m map[string]*Monkey, start string) int {
	// this function will be a recursive function that will
	// go down the list to find the end result of the provided
	// monkey name.
	if _, ok := m[start]; !ok {
		panic("That monkey does not exist!")
	}

	if m[start].JobStatus {
		// job's done!
		return m[start].Result
	} else {
		// otherwise, find out if we're waiting on something
		if m[m[start].PrimaryMonkey].JobStatus && m[m[start].SecondaryMonkey].JobStatus {
			return monkeyOperate(m[m[start].PrimaryMonkey].Result, m[m[start].SecondaryMonkey].Result, m[start].Operation)
		} else if !m[m[start].PrimaryMonkey].JobStatus && m[m[start].SecondaryMonkey].JobStatus {
			return monkeyOperate(findResult(m, m[m[start].PrimaryMonkey].Name), m[m[start].SecondaryMonkey].Result, m[start].Operation)
		} else if m[m[start].PrimaryMonkey].JobStatus && !m[m[start].SecondaryMonkey].JobStatus {
			return monkeyOperate(m[m[start].PrimaryMonkey].Result, findResult(m, m[m[start].SecondaryMonkey].Name), m[start].Operation)
		} else {
			return monkeyOperate(findResult(m, m[m[start].PrimaryMonkey].Name), findResult(m, m[m[start].SecondaryMonkey].Name), m[start].Operation)
		}
	}
}

func HumanElement(name string, m map[string]*Monkey) bool {
	// this function acts like a DFS algorithm to search through the
	// entire list of monkeys, looking for the string "humn".

	if name == "humn" {
		return true
	}
	var exists = struct{}{}
	visited := make(map[string]struct{})

	visited[name] = exists
	s := Stack{}

	s.Push(name)

	for len(s.val) > 0 {
		thisMonkey, err := s.Pop()
		if err != nil {
			break
		}
		if thisMonkey == "humn" {
			return true
		}
		if m[thisMonkey].PrimaryMonkey != "" {
			s.Push(m[thisMonkey].PrimaryMonkey)
		}
		if m[thisMonkey].SecondaryMonkey != "" {
			s.Push(m[thisMonkey].SecondaryMonkey)
		}
	}
	return false
}

func FindHumanInFork(m map[string]*Monkey, left, right string) string {
	// finds which fork has a human
	if HumanElement(left, m) {
		return left
	} else if HumanElement(right, m) {
		return right
	} else {
		panic("could not find human")
	}
}

func PartTwoSplit(m map[string]*Monkey, whichMonkey string) (string, string) {
	fmt.Printf("")
	// finds the two monkeys we need to compare.
	if m[whichMonkey].JobStatus {
		// no! bad monke!
		panic("No fork from starting monke")
	}

	return m[m[whichMonkey].PrimaryMonkey].Name, m[m[whichMonkey].SecondaryMonkey].Name
}

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")

	flag.Parse()
	readFile, err := os.Open(*filePtr)

	if err != nil {
		fmt.Println("Fatal:", err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	// Insert code here
	m := ParseLines(lines)
	part1Result := findResult(m, "root")

	fmt.Printf("Result of \"root\" Monke: %d\n", part1Result)

	// now we have to find out what 'humn' should say.
	// find out what branch requires the 'humn' element.

	left, right := PartTwoSplit(m, "root")

	var compareScore int
	var which bool
	if HumanElement(left, m) {
		fmt.Printf("Detected human element from this monke: %s\n", left)
		compareScore = findResult(m, right)
		fmt.Printf("Total score from %s is %d\n", right, compareScore)
		which = false
	} else if HumanElement(right, m) {
		fmt.Printf("Detected human element from this monke: %s\n", right)
		compareScore = findResult(m, left)
		fmt.Printf("Total score from %s is %d\n", left, compareScore)
		which = true
	} else {
		panic("Too many routes back to human")
	}

	fmt.Printf("Comparing score to %d...\n", compareScore)
	var part2Result int
	if !which {
		part2Result = Dive(m, left, compareScore)
	} else {
		part2Result = Dive(m, right, compareScore)
	}

	fmt.Printf("'humn' value to state to get the end result of %d: %d\n", compareScore, part2Result)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
