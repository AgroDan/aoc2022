package main

import (
	"fmt"
	"io"
	"os"
)

// Now you're thinking with structs!
// going to hard-code an array just so it works faster, I guess?

type SigBuffer struct {
	buffer [14]byte
}

func (s *SigBuffer) push(char byte) {
	for i := 13; i > 0; i-- {
		s.buffer[i] = s.buffer[i-1]
	}
	s.buffer[0] = char
}

func (s *SigBuffer) allValuesFilled() bool {
	// returns true if all items in buffer contain data
	for _, v := range s.buffer {
		if v == 0 {
			return false
		}
	}
	return true
}

func (s *SigBuffer) newSignal() bool {
	// Will check to see if a new pattern exists
	if !s.allValuesFilled() {
		return false
	}

	for i := 0; i < len(s.buffer); i++ {
		for j := i + 1; j < len(s.buffer); j++ {
			if s.buffer[i] == s.buffer[j] {
				return false
			}
		}
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
		s.push(v)
		if s.newSignal() {
			break
		}
		counter++
	}

	fmt.Printf("Signal digit count: %d\n", counter)

}
