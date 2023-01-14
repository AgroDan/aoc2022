package main

/*
	This file will contain unordered finite sets used by
	the day16 challenge.
*/

var exists = struct{}{}

// trying something different
type Path struct {
	totalPressure int
	breadCrumbs   Visited
}

func NewPath() Path {
	return Path{
		totalPressure: 0,
		breadCrumbs:   NewVisited(),
	}
}

func (p *Path) AddToPath(loc string, pressure int) {
	p.breadCrumbs.Add(loc)
	p.totalPressure += pressure
}

func (p *Path) BeenThere(loc string) bool {
	// checks if we've been there before.
	// this may be redundant but whatever
	return p.breadCrumbs.Contains(loc)
}

func (p *Path) CopyPath() Path {
	np := Path{
		totalPressure: 0,
		breadCrumbs:   NewVisited(),
	}
	np.totalPressure = p.totalPressure
	for k := range p.breadCrumbs.location {
		np.breadCrumbs.Add(k)
	}
	return np
}

type Visited struct {
	location map[string]struct{}
}

func NewVisited() Visited {
	v := Visited{}
	v.location = make(map[string]struct{})
	return v
}

func (v *Visited) Add(value string) {
	v.location[value] = exists
}

func (v *Visited) Contains(value string) bool {
	_, ok := v.location[value]
	return ok
}
