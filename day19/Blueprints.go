package main

import (
	"errors"
	"fmt"
)

type Cost struct {
	ore, clay, obsidian int
}

func NewCost(ore, clay, obs int) Cost {
	// this function generates a new Cost object
	return Cost{
		ore:      ore,
		clay:     clay,
		obsidian: obs,
	}
}

type Robot struct {
	roboType string
	worth    Cost
}

func NewRobot(ore, clay, obs int, rType string) (Robot, error) {
	// Generates a new robot
	switch rType {
	case "ore", "clay", "obsidian", "geode":
		return Robot{
			roboType: rType,
			worth:    NewCost(ore, clay, obs),
		}, nil
	default:
		return Robot{}, errors.New("Unknown robot type")
	}
}

type Blueprint struct {
	action                     int
	Index                      int
	Ore, Clay, Obsidian, Geode Robot
}

func (b Blueprint) String() string {
	// this function allows me to print the contents of the blueprint.
	t := fmt.Sprintf("Blueprint %d:\n", b.Index)
	t += fmt.Sprintf("\t Ore Robot costs %d ore\n", b.Ore.worth.ore)
	t += fmt.Sprintf("\t Clay Robot costs %d ore\n", b.Clay.worth.ore)
	t += fmt.Sprintf("\t Obsidian Robot costs %d ore, %d clay\n", b.Obsidian.worth.ore, b.Obsidian.worth.clay)
	t += fmt.Sprintf("\t Geode Robot costs %d ore, %d obsidian\n", b.Geode.worth.ore, b.Geode.worth.obsidian)
	return t
}
