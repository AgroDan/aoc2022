package main

import (
	"fmt"
	"io"
	"os"
)

// Now you're thinking with structs!

type SigBuffer struct {
	first  string
	second string
	third  string
	fourth string
}

func (s *SigBuffer) push(char string) {
	// Will push an item into the struct
	s.fourth = s.third
	s.third = s.second
	s.second = s.first
	s.first = char
}

func (s *SigBuffer) allValuesFilled() bool {
	// returns true if all items in buffer contain data
	if s.first != "" && s.second != "" && s.third != "" && s.fourth != "" {
		return true
	} else {
		return false
	}
}

func (s *SigBuffer) newSignal() bool {
	// Will check to see if a new pattern exists
	if !s.allValuesFilled() {
		return false
	}
	switch s.first {
	case s.second, s.third, s.fourth:
		return false
	}
	switch s.second {
	case s.third, s.fourth:
		return false
	}
	switch s.third {
	case s.fourth:
		return false
	}
	return true
}

func main() {
	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	content, err := io.ReadAll(readFile)
	if err != nil {
		fmt.Println("Fatal:", err)
	}
	s := SigBuffer{}
	counter := 1

	for _, v := range content {
		s.push(fmt.Sprintf("%c", v))
		if s.newSignal() {
			break
		}
		counter++
	}

	fmt.Printf("Signal digit count: %d\n", counter)

}
