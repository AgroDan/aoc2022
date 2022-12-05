package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Container struct {
	contents []string
}

func reverse(s []string) []string {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (c *Container) push(d string) {
	// pushes an item onto the contents stack
	c.contents = append(c.contents, d)
}

func (c *Container) pop() string {
	// pops an item off the contents stack, returns the item
	switch len(c.contents) {
	case 0:
		return ""
	case 1:
		retval := c.contents[0]
		c.contents = nil
		return retval
	default:
		retval := c.contents[len(c.contents)-1]
		c.contents = c.contents[:len(c.contents)-1]
		return retval
	}
}

func (c *Container) remove(howMany int) []string {
	// will return a slice of items that are removed
	// from the stack in the order they reside in the stack,
	// with the top-most item being at the beginning of the slice.

	r := []string{}
	for i := 0; i < howMany; i++ {
		r = append(r, c.pop())
	}
	return reverse(r)
}

func (c *Container) add(r []string) {
	// adds [the output of the remove() func] to the stack
	// in the order it was in previously.
	for _, v := range r {
		c.push(v)
	}
}

func PrintTop(c *Container) string {
	if len(c.contents) == 0 {
		return "Empty container!"
	} else {
		// return fmt.Sprintf("%s", c.contents[len(c.contents)-1])
		return c.contents[len(c.contents)-1]
	}
}

func PrintAllTop(m map[string]*Container) (retString string) {

	for i := 1; i <= len(m); i++ {
		stringInt := strconv.Itoa(i)
		retString += fmt.Sprintf("Top of Container %d: %s\n", i, PrintTop(m[stringInt]))
	}
	return
}

func PrintContainer(c Container) {
	if len(c.contents) == 0 {
		fmt.Printf("Empty container!")
	} else {
		for i, v := range c.contents {
			fmt.Printf("Stack %d: Contents: %s\n", i, v)
		}
	}
}

// [N]         [C]     [Z]
// [Q] [G]     [V]     [S]         [V]
// [L] [C]     [M]     [T]     [W] [L]
// [S] [H]     [L]     [C] [D] [H] [S]
// [C] [V] [F] [D]     [D] [B] [Q] [F]
// [Z] [T] [Z] [T] [C] [J] [G] [S] [Q]
// [P] [P] [C] [W] [W] [F] [W] [J] [C]
// [T] [L] [D] [G] [P] [P] [V] [N] [R]
//  1   2   3   4   5   6   7   8   9

// Now create the contrived example.
func BuildState() map[string]*Container {
	// this builds the container
	m := make(map[string]*Container)
	m["1"] = &Container{}
	m["1"].push("T")
	m["1"].push("P")
	m["1"].push("Z")
	m["1"].push("C")
	m["1"].push("S")
	m["1"].push("L")
	m["1"].push("Q")
	m["1"].push("N")

	m["2"] = &Container{}
	m["2"].push("L")
	m["2"].push("P")
	m["2"].push("T")
	m["2"].push("V")
	m["2"].push("H")
	m["2"].push("C")
	m["2"].push("G")

	m["3"] = &Container{}
	m["3"].push("D")
	m["3"].push("C")
	m["3"].push("Z")
	m["3"].push("F")

	m["4"] = &Container{}
	m["4"].push("G")
	m["4"].push("W")
	m["4"].push("T")
	m["4"].push("D")
	m["4"].push("L")
	m["4"].push("M")
	m["4"].push("V")
	m["4"].push("C")

	m["5"] = &Container{}
	m["5"].push("P")
	m["5"].push("W")
	m["5"].push("C")

	m["6"] = &Container{}
	m["6"].push("P")
	m["6"].push("F")
	m["6"].push("J")
	m["6"].push("D")
	m["6"].push("C")
	m["6"].push("T")
	m["6"].push("S")
	m["6"].push("Z")

	m["7"] = &Container{}
	m["7"].push("V")
	m["7"].push("W")
	m["7"].push("G")
	m["7"].push("B")
	m["7"].push("D")

	m["8"] = &Container{}
	m["8"].push("N")
	m["8"].push("J")
	m["8"].push("S")
	m["8"].push("Q")
	m["8"].push("H")
	m["8"].push("W")

	m["9"] = &Container{}
	m["9"].push("R")
	m["9"].push("C")
	m["9"].push("Q")
	m["9"].push("F")
	m["9"].push("S")
	m["9"].push("L")
	m["9"].push("V")

	return m
}

func ParseCommand(command string, s map[string]*Container) {
	// parses a command from a line, then executes the command
	// on the provided stack

	// first, split the string into a slice

	com := strings.Split(command, " ")
	// move # from A to B
	howMany, err := strconv.Atoi(com[1])
	from := com[3]
	to := com[5]
	if err != nil {
		fmt.Println("Error!:", err)
	}

	temp := s[from].remove(howMany)
	s[to].add(temp)
	// for i := 0; i < howMany; i++ {
	// 	temp := s[from].pop()
	// 	s[to].push(temp)
	// }

}

func main() {
	// this one is going to be weird, because I'm going to ignore the first input
	// lines in the interest of time and hard-code the state. Only care about lines
	// that start with the word "move"
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

	// now set up the valid commands
	var commands []string

	for _, v := range lines {
		if strings.HasPrefix(v, "move") {
			commands = append(commands, v)
		}
	}

	stack := BuildState()

	for _, v := range commands {
		ParseCommand(v, stack)
	}
	// ParseCommand("move 6 from 2 to 1", stack)

	fmt.Printf("%s", PrintAllTop(stack))
	// fmt.Println("Container 1:")
	// PrintContainer(*stack["1"])

	// fmt.Println("Container 2:")
	// PrintContainer(*stack["2"])
}
