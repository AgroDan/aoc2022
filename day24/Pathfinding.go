package main

import (
	"errors"
	"fmt"
)

/*
Since a lot of my previous files were getting fairly large, I
decided to split off my pathfinding algorithms into a separate
file. Let's go!
*/

// Breadth-First-Search

//
// define State type
//

type State struct {
	loc Coord  // location of the elf
	v   Valley // valley object, the whole map at this slice
}

func NewState(c Coord, val Valley) State {
	return State{
		loc: c,
		v:   val,
	}
}

//
// define queue
//

type Queue struct {
	q []State
}

func (q *Queue) Length() int {
	return len(q.q)
}

func (q *Queue) Push(s State) {
	// adds explorer to queue
	q.q = append(q.q, s)
}

func (q *Queue) Pop() (State, error) {
	if len(q.q) == 0 {
		return State{}, errors.New("empty queue")
	}

	ne := q.q[0]
	q.q = q.q[1:]
	return ne, nil
}

//
// Stack
//

type Stack struct {
	s []State
}

func (s *Stack) Length() int {
	return len(s.s)
}

func (s *Stack) Push(st State) {
	// adds state to stack
	s.s = append(s.s, st)
}

func (s *Stack) Pop() (State, error) {
	if len(s.s) == 0 {
		return State{}, errors.New("empty stack")
	}

	ne := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return ne, nil
}

//
// define breadcrumb
//

type exists struct{}

type Breadcrumb struct {
	crumb map[string]exists
}

func (b *Breadcrumb) Add(hash string, c Coord) bool {
	// this will place a breadcrumb. Will return false ONLY if
	// the breadcrumb already exists!

	// create the hash
	newHash := fmt.Sprintf("%s|%d~%d", hash, c.row, c.col)
	if _, ok := b.crumb[newHash]; ok {
		return false
	}

	b.crumb[newHash] = exists{}
	return true
}

//
// BFS
//

func Traverse(start Coord, val Valley) int {
	// second time's a charm!
	b := Breadcrumb{
		crumb: make(map[string]exists),
	}

	// hello little explorer!
	agr0 := NewState(start, val)

	q := Queue{}
	q.Push(agr0)

	for q.Length() > 0 {
		this, err := q.Pop()
		if err != nil {
			fmt.Printf("Empty queue!\n")
			break
		}

		// let's do a quick check for exit status
		if this.loc == this.v.Exit {
			// we're at the exit, return total.
			// fmt.Printf("We made it!\n")
			return this.v.Minutes
		}

		// add and check the breadcrumb
		if !b.Add(this.v.GetHash(), this.loc) {
			// we are in the same state as previously defined, this would
			// yield no additional positive outcomes, so bomb this one out.
			// fmt.Printf("Hit a breadcrumb, blow this away\n")
			continue
		}

		// otherwise, let's eval our options
		options := this.v.GetValidDirections(this.loc)
		// fmt.Printf("Found options from this loc: c:%d r:%d : %#v\n", this.loc.col, this.loc.row, options)
		for _, value := range options {
			// split off the realities!
			altUniverse := NewState(value, this.v.DeepCopy())
			altUniverse.v.MoveOne()
			q.Push(altUniverse)
		}
	}
	return 0
}

func TraverseAndReturn(currentState State, start, end Coord) (int, State) {
	// this function operates similarly to the Traverse function, but it
	// is now made to start at arbitrary positions.
	b := Breadcrumb{
		crumb: make(map[string]exists),
	}

	// // hello little explorer!
	// agr0 := NewState(start, val)

	q := Queue{}
	q.Push(currentState)
	// iter := 0

	for q.Length() > 0 {
		// iter++
		this, err := q.Pop()
		if err != nil {
			fmt.Printf("Empty queue!\n")
			break
		}

		// let's do a quick check for exit status
		if this.loc == end {
			// we're at the exit, return total.
			// fmt.Printf("We made it!\n")
			return this.v.Minutes, this
		}

		// add and check the breadcrumb
		if !b.Add(this.v.GetHash(), this.loc) {
			// we are in the same state as previously defined, this would
			// yield no additional positive outcomes, so bomb this one out.
			// fmt.Printf("Hit a breadcrumb, blow this away\n")
			continue
		}

		// otherwise, let's eval our options
		options := this.v.GetValidDirectionsPartTwo(this.loc, start, end)
		// fmt.Printf("Found options from this loc: c:%d r:%d : %#v\n", this.loc.col, this.loc.row, options)
		for _, value := range options {
			// split off the realities!
			altUniverse := NewState(value, this.v.DeepCopy())
			altUniverse.v.MoveOne()
			q.Push(altUniverse)
		}
	}
	// fmt.Printf("Iterated: %d\n", iter)
	return 0, State{} // we got nothin
}
