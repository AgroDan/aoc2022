package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

/*
   My logic for this one:

   Create an interface which implements IsInteger()

   Create a struct which holds both an integer and a pointer to another struct of same type

   IsInteger will be a struct function, if pointer is nil, return True. If not nil, return False.

   Now create a []interface{} list of these objects. This will hold either an integer or a pointer to another
   []interface{} slice, which can hold 0 or more objects
*/

// no, let's try something different. No structs.
//
// type Signal interface {
// 	IsInteger() bool
// 	IsEmpty() bool
// }

// type Packet struct {
// 	sig    int
// 	nested *[]Signal
// 	empty  bool
// }

// func (p Packet) IsInteger() bool {
// 	if p.nested == nil {
// 		return true
// 	}
// 	return false
// }

// func (p Packet) IsEmpty() bool {
// 	return p.empty
// }

// now let's see what happens!

func ReadSignal(s string) []interface{} {
	var signal []interface{}
	json.Unmarshal([]byte(s), &signal)
	return signal
}

func IsInOrder(left, right []interface{}) (bool, error) {
	// This function will take two interfaces and determine
	// if they are in the right order in accordance with the
	// advent of code instructions. If an error is presented,
	// then both sides are EQUAL and cannot be technically ordered.

	// first, we will only loop as many times as the
	// smallest of the two indices
	counter := 0
	if len(left) >= len(right) {
		counter = len(right)
	} else {
		counter = len(left)
	}

	// right has more items, making this in order
	// (even if both are the same)
	isRightGreater := len(right) > len(left)

	// we know that if we loop over each and end it with all equal,
	// return isRightGreater

	for i := 0; i < counter; i++ {
		// first, let's assume float64 (what json package interprets numbers as)
		// https://go.dev/tour/methods/15
		leftNum, isLeftNum := left[i].(float64)
		rightNum, isRightNum := right[i].(float64)

		// also assume interfaces
		leftInter, isLeftInter := left[i].([]interface{})
		rightInter, isRightInter := right[i].([]interface{})

		if isLeftNum && isRightNum {
			// if both values are numbers, check if left > right.
			// if so, end it here. It's out of order.
			// fmt.Printf("Compare left: %f, right: %f\n", leftNum, rightNum)
			if leftNum > rightNum {
				// fmt.Printf("Number comparison, left > right - false\n")
				return false, nil
			} else if leftNum < rightNum {
				// however, if left is less than right and we haven't
				// hit a false outcome yet, we're in the right order!
				// fmt.Printf("Left is less than right, return true\n")
				return true, nil
			}

			// otherwise, the numbers are probably equal
		}

		if isLeftInter && isRightInter {
			// both values are lists. Recursion!
			// fmt.Printf("Recursion: left: %v -- right: %v\n", leftInter, rightInter)
			status, err := IsInOrder(leftInter, rightInter)
			if err == nil {
				return status, nil
			}
		}

		if isLeftNum && isRightInter {
			in := []interface{}{leftNum}
			// fmt.Printf("Convert left, recursion: left: %v -- right %v\n", in, rightInter)
			status, err := IsInOrder(in, rightInter)
			if err == nil {
				// definitive end result
				return status, nil
			}
		}

		if isLeftInter && isRightNum {
			in := []interface{}{rightNum}
			// fmt.Printf("Convert right, recursion: left: %v -- right: %v\n", leftInter, in)
			status, err := IsInOrder(leftInter, in)
			if err == nil {
				// definitive end result
				return status, nil
			}
		}
	}
	// fmt.Printf("Is right greater or equal to left: %t\n", isRightGreater)
	if len(left) == len(right) {
		// both are equal, and cannot make comparison.
		return true, errors.New("both sides are equal")
	}
	return isRightGreater, nil
}

func main() {
	// reading the whole thing in memory
	dat, err := os.ReadFile("input")
	if err != nil {
		panic("Could not read file!")
	}

	couples := strings.Split(string(dat), "\n\n")

	coupleCounter := 1
	total := 0
	for _, v := range couples {
		pairs := strings.Split(v, "\n")
		left := ReadSignal(pairs[0])
		right := ReadSignal(pairs[1])

		fmt.Printf("Left side: %v\n", left)
		fmt.Printf("Right side: %v\n", right)
		res, _ := IsInOrder(left, right)
		fmt.Printf("Pair in order? %t\n\n", res)

		if res {
			total += coupleCounter
		}
		coupleCounter++
	}

	fmt.Printf("Total indices: %d\n", total)

	// now for part 2
	// First, create a list of strings
	var signals [][]interface{}
	for _, v := range couples {
		pairs := strings.Split(v, "\n")
		left := ReadSignal(pairs[0])
		right := ReadSignal(pairs[1])
		signals = append(signals, left, right)
	}

	// append the remaining
	signals = append(signals, ReadSignal("[[2]]"), ReadSignal("[[6]]"))

	mySorted := MergeSort(signals)
	dividerPacketIdx := 1
	for i, v := range mySorted {
		if fmt.Sprint(v) == "[[2]]" || fmt.Sprint(v) == "[[6]]" {
			dividerPacketIdx *= (i + 1)
		}
		fmt.Printf("[%02d] %v\n", i+1, v)
	}
	fmt.Printf("Index of distress signals multiplied: %d\n", dividerPacketIdx)
}
