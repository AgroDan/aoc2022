package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

const KEY int = 811589153

func ParseLines(l []string) *LinkedList {
	// parses the lines and creates a massive
	// linked list. Since it doesn't really matter
	// what we do, we'll just return the latest
	// linkedlist member
	num, err := strconv.Atoi(l[0])
	if err != nil {
		panic("could not convert")
	}
	head := NewLinkedList(num)
	current := &head

	for _, v := range l[1:] {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic("could not convert")
		}

		n := NewLinkedList(num)
		InsertAfter(current, &n)
		current = current.next
	}
	return current.next
}

func ParseLinesPart2(l []string) *LinkedList {
	// this basically does the same thing as the original
	// function, but this time it multiplies each number
	// by the decryption key (listed as a constant).
	num, err := strconv.Atoi(l[0])
	if err != nil {
		panic("could not convert")
	}
	head := NewLinkedList(num * KEY)
	current := &head

	for _, v := range l[1:] {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic("could not convert")
		}

		n := NewLinkedList(num * KEY)
		InsertAfter(current, &n)
		current = current.next
	}
	return current.next
}

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")
	debugPtr := flag.Bool("debug", false, "Debug flag, will be extra verbose")

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

	// "l" is the list that will be manipulated.
	l := ParseLines(lines)
	// PrintLL(l, *debugPtr)

	// idx is a slice showing the order of the linked list.
	idx := GetAddresses(l)

	for _, v := range idx {
		newLocation := Traverse(v.coord, v)
		if v.coord > 0 {
			if *debugPtr {
				fmt.Printf("Val: %d, moving right, placing after %d\n", v.coord, newLocation.coord)
			}
			UpdateAfter(newLocation, v)
		} else if v.coord < 0 {
			if *debugPtr {
				fmt.Printf("Val: %d, moving left, placing before %d\n", v.coord, newLocation.coord)
			}
			UpdateBefore(newLocation, v)
		} else {
			if *debugPtr {
				fmt.Printf("Val: %d, no movement.\n", v.coord)
			}
		}
	}

	// fmt.Println("New LL:")
	// PrintLL(l, *debugPtr)

	// now find the 1000th, 2000th, and 3000th items and add them together
	zeroFinder := FindCoord(0, l)

	// using StepTraverse because it also counts the original pointer
	// when traversing the LL
	gpsA := StepTraverse(1000, zeroFinder)
	gpsB := StepTraverse(2000, zeroFinder)
	gpsC := StepTraverse(3000, zeroFinder)

	fmt.Printf("Part 1 - Coord a: %d, b: %d, c: %d, sum: %d\n", gpsA.coord, gpsB.coord, gpsC.coord, gpsA.coord+gpsB.coord+gpsC.coord)

	fmt.Printf("Now for part 2...\n")

	p2 := ParseLinesPart2(lines)

	// fmt.Printf("Current list:\n")
	// PrintLL(p2, *debugPtr)

	idx = nil
	idx = GetAddresses(p2)
	for i := 0; i < 10; i++ {
		for _, v := range idx {
			steps := v.coord % (len(idx) - 1)
			newLocation := Traverse(steps, v)

			if steps > 0 {
				if *debugPtr {
					fmt.Printf("Steps Val: %d, placing after %d\n", steps, newLocation.coord)
				}
				UpdateAfter(newLocation, v)
			} else if steps < 0 {
				if *debugPtr {
					fmt.Printf("Steps Val: %d, placing before %d\n", steps, newLocation.coord)
				}
				UpdateBefore(newLocation, v)
			} else {
				if *debugPtr {
					fmt.Printf("Steps Val: %d, no movement.\n", steps)
				}
			}
		}
		// fmt.Printf("After %d cycles:\n", i+1)
		// PrintLL(p2.next, *debugPtr)
	}
	zeroFinder = FindCoord(0, p2)
	gpsA = StepTraverse(1000, zeroFinder)
	gpsB = StepTraverse(2000, zeroFinder)
	gpsC = StepTraverse(3000, zeroFinder)
	fmt.Printf("Part 2 - Coord a: %d, b: %d, c: %d, sum: %d\n", gpsA.coord, gpsB.coord, gpsC.coord, gpsA.coord+gpsB.coord+gpsC.coord)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
