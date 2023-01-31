package main

import "fmt"

/*
	This file will maintain the State object, which
	will attach a Blueprint to a Resources object
	and attempt to interact with everything.
*/

const (
	BAKE = iota
	BUILD
)

type MiningState struct {
	bp                      *Blueprint // link to specific blueprint
	res                     *Resources // specific resources attached to this state
	minutes                 int        // minutes left at this state
	OreRobots               int        // how many ore robots
	ClayRobots              int        // how many clay robots
	ObsidianRobots          int        // how many obsidian robots
	GeodeRobots             int        // how many geode robots
	maxOre, maxClay, maxObs int        // the max amount of ore/clay/obs robots we should build
	build                   string     // if we are building a new robot, specify this one here
	action                  int
}

func (m *MiningState) String() string {
	r := fmt.Sprintf("========================\n")
	r += fmt.Sprintf("Mining State:\n")
	r += fmt.Sprintf("\tAmt of Ore robots: %d\n", m.OreRobots)
	r += fmt.Sprintf("\tAmt of Clay robots: %d\n", m.ClayRobots)
	r += fmt.Sprintf("\tAmt of Obs robots: %d\n", m.ObsidianRobots)
	r += fmt.Sprintf("\tAmt of Geode robots: %d\n", m.GeodeRobots)
	r += fmt.Sprintf("Totals:\n")
	r += fmt.Sprintf("\tOres: %d\n", m.res.ore)
	r += fmt.Sprintf("\tClay: %d\n", m.res.clay)
	r += fmt.Sprintf("\tObs: %d\n", m.res.obsidian)
	r += fmt.Sprintf("\tGeodes: %d\n", m.res.geode)
	r += fmt.Sprintf("Max:\n")
	r += fmt.Sprintf("\tMax Ores: %d\n", m.maxOre)
	r += fmt.Sprintf("\tMax Clay: %d\n", m.maxClay)
	r += fmt.Sprintf("\tMax Obs: %d\n", m.maxObs)
	r += fmt.Sprintf("========================\n")
	return r
}

func NewMininigState(bp *Blueprint, min, ore, clay, obs, geode int, build string, action int) MiningState {
	r := NewResources()
	mo, mc, mob := 0, 0, 0

	var robotArray [4]*Robot
	robotArray[0] = &bp.Ore
	robotArray[1] = &bp.Clay
	robotArray[2] = &bp.Obsidian
	robotArray[3] = &bp.Geode

	for _, v := range robotArray {
		if v.worth.ore > mo {
			mo = v.worth.ore
			fmt.Printf("Max ore is %d\n", mo)
		}
		if v.worth.clay > mc {
			mc = v.worth.clay
			fmt.Printf("Max clay is %d\n", mc)
		}
		if v.worth.obsidian > mob {
			mob = v.worth.obsidian
			fmt.Printf("Max obs is %d\n", mob)
		}
	}

	return MiningState{
		bp:             bp,
		res:            &r,
		minutes:        min,
		OreRobots:      ore,
		ClayRobots:     clay,
		ObsidianRobots: obs,
		GeodeRobots:    geode,
		maxOre:         mo,
		maxClay:        mc,
		maxObs:         mob,
		build:          build,
		action:         action,
	}
}

func (m *MiningState) Copy() MiningState {
	// this function copies one mining state to another new
	// object, ALSO copying the resources object so this can
	// be used to split to another choice
	nr := m.res.Copy()
	nm := MiningState{
		bp:             m.bp,
		res:            &nr,
		minutes:        m.minutes,
		OreRobots:      m.OreRobots,
		ClayRobots:     m.ClayRobots,
		ObsidianRobots: m.ObsidianRobots,
		GeodeRobots:    m.GeodeRobots,
		maxOre:         m.maxOre,
		maxClay:        m.maxClay,
		maxObs:         m.maxObs,
		build:          m.build,
		action:         m.action,
	}
	return nm
}

func (m *MiningState) Gather() {
	// this function will gather as many ores as there
	// are robots available. This is to simulate one minute
	// of the challenge.
	m.res.ore += m.OreRobots
	m.res.clay += m.ClayRobots
	m.res.obsidian += m.ObsidianRobots
	m.res.geode += m.GeodeRobots
}

func (m *MiningState) CanWePurchase() []string {
	// this function will return a slice of robot types
	// that we can purchase _if any_. If we can't purchase
	// any robots, the string slice will have a len() of 0

	retval := make([]string, 0)

	// check for an Ore Robot
	// and yes i know this is ridiculous, I was literally
	// just laughing about how stupid this is
	if m.res.ore >= m.bp.Ore.worth.ore && m.maxOre > m.OreRobots {
		retval = append(retval, "ore")
	}

	// check for Clay robot
	if m.res.ore >= m.bp.Clay.worth.ore && m.maxClay > m.ClayRobots {
		retval = append(retval, "clay")
	}

	// check for obsidian robot
	if (m.res.ore >= m.bp.Obsidian.worth.ore && m.res.clay >= m.bp.Obsidian.worth.clay) && m.maxObs > m.ObsidianRobots {
		retval = append(retval, "obsidian")
	}

	// finally check for geode robots
	if m.res.ore >= m.bp.Geode.worth.ore && m.res.obsidian >= m.bp.Geode.worth.obsidian {
		retval = append(retval, "geode")
	}
	return retval
}

/*
	Build the stack, nice and simple
*/

type Stack struct {
	q []*MiningState
}

func (q *Stack) Push(m *MiningState) {
	q.q = append(q.q, m)
}

func (q *Stack) Pop() *MiningState {
	nq := q.q[len(q.q)-1]
	q.q = q.q[:len(q.q)-1]
	return nq
}
func (q *Stack) Peek() *MiningState {
	// simply reveals the top of the stack without
	// manipulating the stack
	return q.q[len(q.q)-1]
}

func (q *Stack) IsEmpty() bool {
	return len(q.q) == 0
}

func TestBlueprint(b Blueprint, totalMinutes int) int {
	// this function will test the provided blueprint and
	// return an array of possible Geodes produced. It is
	// up to the caller to find the highest amount of Geodes
	// possible

	retval := 0

	// First, create a State with an action. We want to assume
	// that the first action really is "Bake". If not, we can
	// just move on from there.

	// remember we always start with one ore robot
	state := NewMininigState(&b, totalMinutes, 1, 0, 0, 0, "", BAKE)

	// create a new stack and push this state into it
	stack := Stack{}
	stack.Push(&state)

	// now let's start the festivities.

	for !stack.IsEmpty() {
		curr := stack.Pop()

		// first, perform the action queued
		if curr.action == BAKE {
			curr.Gather()
			curr.minutes--
		} else {
			// otherwise we're building. But first, gather
			curr.Gather()
			switch curr.build {
			case "ore":
				// spend the ore necessary
				err := curr.res.Spend(curr.bp.Ore.worth.ore, curr.bp.Ore.worth.clay, curr.bp.Ore.worth.obsidian)
				if err != nil {
					fmt.Printf("Could not afford a new Ore robot!\n")
				} else {
					curr.OreRobots++
				}
			case "clay":
				err := curr.res.Spend(curr.bp.Clay.worth.ore, curr.bp.Clay.worth.clay, curr.bp.Clay.worth.obsidian)
				if err != nil {
					fmt.Printf("Could not afford a new Clay robot!\n")
				} else {
					curr.ClayRobots++
				}
			case "obsidian":
				err := curr.res.Spend(curr.bp.Obsidian.worth.ore, curr.bp.Obsidian.worth.clay, curr.bp.Obsidian.worth.obsidian)
				if err != nil {
					fmt.Printf("Could not afford a new Obsidian robot!\n")
				} else {
					curr.ObsidianRobots++
				}
			case "geode":
				err := curr.res.Spend(curr.bp.Geode.worth.ore, curr.bp.Geode.worth.clay, curr.bp.Geode.worth.obsidian)
				if err != nil {
					fmt.Printf("Could not afford a new Geode robot!\n")
				} else {
					curr.GeodeRobots++
				}
			}
			curr.build = ""
			curr.minutes--
		}

		// get potential future geodes
		futureGeodes := (curr.minutes * curr.GeodeRobots) + curr.res.geode

		// theoreticalGeodes will be how many geodes we can _theoretically_ make
		// if we do nothing but build geode robots and collect geodes for every
		// single turn from here on out.
		theoreticalGeodes := func(minutes, addlGeodes int) (myNewGeodes int) {
			myNewGeodes = 0
			if minutes > 0 {
				for i := 0; i < minutes; i++ {
					myNewGeodes += i
				}
			}
			return myNewGeodes + addlGeodes
		}

		if curr.minutes <= 0 {
			if curr.res.geode > retval {
				retval = curr.res.geode
			}
			// fmt.Printf("%s", curr)
			// } else if futureGeodes >= retval || curr.minutes > 2 {
		} else if theoreticalGeodes(curr.minutes, futureGeodes) > retval {
			// set curr.minutes > 2 because any building during
			// the final 2 minutes is pointless
			//
			// now determine the next steps.
			// We can ALWAYS just bake.
			baker := curr.Copy()
			baker.action = BAKE
			stack.Push(&baker)

			// now let's see what we can build (if anything)
			buildList := curr.CanWePurchase()
			onlyGeode := false
			for _, v := range buildList {
				if v == "geode" {
					onlyGeode = true
					builder := curr.Copy()
					builder.action = BUILD
					builder.build = v
					stack.Push(&builder)
				}
			}
			if !onlyGeode {
				for _, v := range buildList {
					builder := curr.Copy()
					builder.action = BUILD
					builder.build = v
					stack.Push(&builder)
				}
			}
		}
	}
	return retval
}
