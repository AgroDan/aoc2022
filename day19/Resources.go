package main

import "errors"

type Resources struct {
	ore, clay, obsidian, geode int
}

func NewResources() Resources {
	return Resources{
		ore:      0,
		clay:     0,
		obsidian: 0,
		geode:    0,
	}
}

func (r *Resources) Copy() Resources {
	// this function copies the contents of the current
	// resources object into a new object for use when
	// splitting decisions.
	nr := NewResources()
	nr.ore = r.ore
	nr.clay = r.clay
	nr.obsidian = r.obsidian
	nr.geode = r.geode
	return nr
}

// turns out I never used this function anyway, oh well
func (r *Resources) AddOre(ore, clay, obsidian, geode int) {
	// adds resources to the pool.
	r.ore += ore
	r.clay += clay
	r.obsidian += obsidian
	r.geode += geode
}

func (r *Resources) Spend(ore, clay, obsidian int) error {
	// subtracts resources from the collection. Since we won't
	// be spending Geodes they are not included here. returns
	// an error if spending of any item is less than zero.
	if r.ore < ore {
		return errors.New("ore cost is too expensive")
	}
	if r.clay < clay {
		return errors.New("clay cost is too expensive")
	}
	if r.obsidian < obsidian {
		return errors.New("obsidian cost is too expensive")
	}
	r.ore -= ore
	r.clay -= clay
	r.obsidian -= obsidian
	return nil
}
