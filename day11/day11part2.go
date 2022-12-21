package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// first, let's define a monkey object
type Monkey struct {
	items    []int  // the worry levels
	opAction string // the operation action, such as +, *, -, /, etc
	opNum    int    // the integer applied to the opAction
	testNum  int    // the number to test if divisible by
	//	action   map[bool]*Monkey // THIS is weird. If true, throw to monkey address 1, if false, throw to monkey address 2
	action MonkeyAction
}

// this is the struct object for who a monkey throws to for every true or false situation
type MonkeyAction struct {
	trueAction  *Monkey
	falseAction *Monkey
}

// this is just a helper struct I can use to store data about which monkey to throw to when
// parsing the input file.
type MonkeyState struct {
	trueAction  int
	falseAction int
}

func NewMonkey(i []int, opAction string, opNum, testNum int) Monkey {
	m := Monkey{
		items:    i,
		opAction: opAction,
		opNum:    opNum,
		testNum:  testNum,
	}
	//m.action = make(map[bool]*Monkey)
	m.action = MonkeyAction{}
	return m
}

func (m *Monkey) AddTossTarget(which bool, nm *Monkey) {
	// assigns monkey recipient
	//m.action[which] = nm
	if which {
		m.action.trueAction = nm
	} else {
		m.action.falseAction = nm
	}
}

func (m *Monkey) Receive(item int) {
	// This receives a new item
	m.items = append(m.items, item)
}

func (m *Monkey) ModifyWorryLevel(wl int) int {
	// This takes the item being inspected and changes
	// the worry level based on the monkey's defined opAction
	switch m.opAction {
	case "+":
		wl += m.opNum
	case "-":
		wl -= m.opNum
	case "*":
		wl *= m.opNum
	case "/":
		wl /= m.opNum
	case "^":
		wl = wl * wl
	}
	return wl
}

func (m *Monkey) ActOnResult(wl, mod int) error {
	// This function actually moves the item from one monkey to
	// another based on the test.
	// First, are all monkeys ready to receive?
	if !(m.IsReady() && m.action.trueAction.IsReady() && m.action.falseAction.IsReady()) {
		return errors.New("Not all monkeys are prepared!")
	}

	// otherwise, determine where this goes.
	// Remember, AFTER inspection but before the check, the item is
	// divided by 3 and rounded down to the nearest integer. Integer
	// division does this automatically so here we go:
	wl %= mod
	// above commented out because of the restriction. I need to find
	// another way to manage my worry levels...?

	if wl%m.testNum == 0 {
		// toss to the true monkey
		m.action.trueAction.Receive(wl)
	} else {
		// toss to the false monkey
		m.action.falseAction.Receive(wl)
	}
	return nil
}

func (m *Monkey) IsReady() bool {
	// this determines if we are ready to pass to another monkey
	if m.action.trueAction != nil && m.action.falseAction != nil {
		return true
	} else {
		return false
	}
}

func (m *Monkey) Inspect(mod int) {
	// this function loops through every item in its list of items
	// and performs the proper actions on them.
	var currItem int
	for len(m.items) > 0 {
		currItem, m.items = m.items[0], m.items[1:]

		// monkey is inspecting the item.
		newItem := m.ModifyWorryLevel(currItem)

		// new worry level is made from item and actions imposed
		// now act on result. Monkey gets bored with item in this step
		err := m.ActOnResult(newItem, mod)
		if err != nil {
			panic(fmt.Sprintf("Not set up properly! %s\n", err))
		}
	}
}

func ParseMonkeys(lines []string) []*Monkey {
	var throwList []MonkeyState
	var monkeyList []*Monkey
	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], "Monkey") {
			// at this point, we know that i+1 contains the starting items
			// i+2 contains operation
			// i+3 contains Test
			// i+4 contains true action
			// i+5 contains false action

			s := strings.Split(lines[i+1], ": ")
			n := strings.Split(s[1], ",")
			monkeysItems := make([]int, 0)
			for _, v := range n {
				trimmed := strings.TrimSpace(v)
				j, _ := strconv.Atoi(trimmed)
				monkeysItems = append(monkeysItems, j)
			}

			s = strings.Split(lines[i+2], " ")
			operationAction := s[len(s)-2]
			var err error
			var operationNum int
			operationNum, err = strconv.Atoi(s[len(s)-1])
			if err != nil {
				// presumably, this is "old"
				if s[len(s)-1] == "old" {
					operationAction = "^"
					operationNum = 0
				} else {
					fmt.Printf("Could not convert number in action! %s\n", lines[i+2])
				}
			}

			s = strings.Split(lines[i+3], " ")
			testNum, _ := strconv.Atoi(s[len(s)-1])

			s = strings.Split(lines[i+4], " ")
			trueNum, _ := strconv.Atoi(s[len(s)-1])

			s = strings.Split(lines[i+5], " ")
			falseNum, _ := strconv.Atoi(s[len(s)-1])

			thisMonkey := NewMonkey(monkeysItems, operationAction, operationNum, testNum)
			m := MonkeyState{
				trueAction:  trueNum,
				falseAction: falseNum,
			}
			throwList = append(throwList, m)
			monkeyList = append(monkeyList, &thisMonkey)
		}
	}

	// now true up the monkeys
	for i := 0; i < len(monkeyList); i++ {
		monkeyList[i].AddTossTarget(true, monkeyList[throwList[i].trueAction])
		monkeyList[i].AddTossTarget(false, monkeyList[throwList[i].falseAction])
	}
	return monkeyList
}

/*
	Now, to parse -- we have to do two iterations. First, set a list
	of monkeys, and we will use the index to keep track of the monkey
	number:

	var MonkeyList []*Monkey
	monkeyAction := make([]map[bool]int)

	Parse each monkey's true and false values.

	then...
	a := make(map[bool]int)
	a[true] = <trueactionnum>
	a[false] = <falseactionnum>
	monkeyAction = append(monkeyAction, a)

	After all monkeys are made, refer to each monkey action as
	MonkeyList[i].action[true] = MonkeyList[monkeyAction[i][true]]
	MonkeyList[i].action[false] = MonkeyList[monkeyAction[i][false]]
*/

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

	// var MonkeyList []*Monkey
	MonkeyList := ParseMonkeys(lines)

	// fmt.Printf("Full values: +%v\n", MonkeyList)
	// fmt.Printf("List values of monkey 0: +%v\n", *MonkeyList[0])

	scoreBoard := make([]int, 0)

	for range MonkeyList {
		scoreBoard = append(scoreBoard, 0)
	}

	// This is a trick I learned on Reddit.
	var divisors []int
	for _, v := range MonkeyList {
		divisors = append(divisors, v.testNum)
	}

	// now get the number from this...
	var prod int = 1
	for _, v := range divisors {
		prod *= v
	}

	// Now...let's iterate.

	for i := 0; i < 10000; i++ {
		for i, v := range MonkeyList {
			scoreBoard[i] += len(v.items)
			v.Inspect(prod)
		}
	}

	for i, v := range scoreBoard {
		fmt.Printf("Monkey %d inspected %d times.\n", i, v)
	}

	// sortedBoard := sort.IntSlice(scoreBoard)
	// sortedBoard.Sort()
	// a := sort.Reverse(sortedBoard)
	// fmt.Println("Sorted board:")
	sort.Slice(scoreBoard, func(i, j int) bool {
		return scoreBoard[i] > scoreBoard[j]
	})
	// now get what is needed
	fmt.Println("Result:", scoreBoard[0]*scoreBoard[1])
}
