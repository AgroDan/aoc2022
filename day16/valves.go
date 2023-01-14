package main

/*
	This module will be all code regarding the valves
	and the functions that surround them.
*/

type Valve struct {
	flowRate int
	tunnels  []string
	isOpen   bool
}

func NewValve(f int, tunnels ...string) Valve {
	// this creates a new valve
	t := make([]string, 0)
	for _, v := range tunnels {
		t = append(t, v)
	}
	return Valve{
		flowRate: f,
		tunnels:  t,
		isOpen:   false,
	}
}
