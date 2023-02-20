package main

import (
	"fmt"
	"strconv"
)

// var dirQueue = []string{"N", "S", "W", "E"}

func RotateQueue(q []string) []string {
	temp := q[0]
	q = q[1:]
	q = append(q, temp)
	return q
}

type Proposals struct {
	set map[string]int
}

func NewProposals() Proposals {
	return Proposals{
		set: make(map[string]int),
	}
}

func (p *Proposals) ClearProposals() {
	// clears out all proposals
	for k := range p.set {
		delete(p.set, k)
	}
}

func (p *Proposals) DeclareProposal(Row, Col int) {
	// this function will create a new proposal if possible,
	// but if it already exists, then increment the proposal by 1
	setMap := fmt.Sprintf("%s~%s", strconv.Itoa(Row), strconv.Itoa(Col))
	if _, ok := p.set[setMap]; ok {
		// this exists already, increment
		p.set[setMap]++
	} else {
		p.set[setMap] = 1
	}
}

func (p *Proposals) ValidProposal(Row, Col int) bool {
	// this function will return true if the proposal is a valid one,
	// meaning that this has not been proposed previously.
	setMap := fmt.Sprintf("%s~%s", strconv.Itoa(Row), strconv.Itoa(Col))
	if _, ok := p.set[setMap]; ok {
		// fmt.Printf("Valid proposal, it exists!\n")
		if p.set[setMap] > 1 {
			return false
		} else {
			return true
		}
	} else {
		fmt.Printf("Something went wrong!\n")
		return false // this has not been declared!
	}
}
