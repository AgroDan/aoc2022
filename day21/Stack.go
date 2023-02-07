package main

import "errors"

// this file contains the items necessary to operate
// a stack of monkeys

type Stack struct {
	val []string
}

func (s *Stack) Push(st string) {
	// pushes item onto the stack
	s.val = append(s.val, st)
}

func (s *Stack) Pop() (string, error) {
	// pops item off the stack
	if len(s.val) == 0 {
		return "", errors.New("empty stack")
	}
	r := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]
	return r, nil
}

func (s *Stack) Peek() string {
	return s.val[len(s.val)-1]
}
