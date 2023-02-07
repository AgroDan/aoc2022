package main

import (
	"strconv"
	"strings"
)

// will use type map[string]*Monkey to point to a monkey name

type Monkey struct {
	Name            string
	JobStatus       bool
	PrimaryMonkey   string
	SecondaryMonkey string
	Operation       rune
	Result          int
}

func ReadMonkey(l string) Monkey {
	// this function will read a line
	// and build a Monkey object out of the data

	// root: pppw + sjmn
	// dbpl: 5
	m := Monkey{}
	words := strings.Split(l, " ")
	m.Name = strings.TrimSuffix(words[0], ":")

	num, err := strconv.Atoi(words[1])
	if err != nil {
		// this is not just a number, it is an operations string.
		// continue as needed.
		pMonke := words[1]
		sMonke := words[3]

		// this is so stupid. Go can't cast a single string character
		// to a rune without dumping into a rune slice? Bad Go, bad!
		op := []rune(words[2])

		m.JobStatus = false
		m.PrimaryMonkey = pMonke
		m.SecondaryMonkey = sMonke
		m.Operation = op[0]
	} else {
		// if we got here, then the Monke's job is just to shout the
		// number.

		m.JobStatus = true
		m.Result = num
	}
	return m
}
