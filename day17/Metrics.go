package main

import (
	"fmt"
	"strconv"
	"strings"
)

type DropStat struct {
	rocksDropped int
	height       int
}

type Metric struct {
	// this object gets metrics on each falling piece
	drops map[string][]DropStat
}

func NewMetric() Metric {
	return Metric{
		drops: make(map[string][]DropStat),
	}
}

func (m *Metric) Contains(RockType, JetIdx int) bool {
	index := fmt.Sprintf("%s~%s", strconv.Itoa(RockType), strconv.Itoa(JetIdx))
	_, ok := m.drops[index]
	return ok
}

func (m *Metric) Add(RockType, JetIdx, Height, amtDropped int) {
	// this adds to a Metric type properly
	index := fmt.Sprintf("%s~%s", strconv.Itoa(RockType), strconv.Itoa(JetIdx))
	d := DropStat{
		rocksDropped: amtDropped,
		height:       Height,
	}
	if m.Contains(RockType, JetIdx) {
		m.drops[index] = append(m.drops[index], d)
	} else {
		m.drops[index] = make([]DropStat, 0)
		m.drops[index] = append(m.drops[index], d)
	}
}

func (m *Metric) GetHeights(RockType, JetIdx int) []int {
	index := fmt.Sprintf("%s~%s", strconv.Itoa(RockType), strconv.Itoa(JetIdx))
	retval := make([]int, 0)
	for _, v := range m.drops[index] {
		retval = append(retval, v.height)
	}
	return retval
}

func GetMapping(index string) (int, int) {
	coords := strings.Split(index, "~")
	RockType, err := strconv.Atoi(coords[0])
	if err != nil {
		panic("Could not pull integer")
	}
	JetIdx, err := strconv.Atoi(coords[1])
	if err != nil {
		panic("Could not pull integer")
	}
	return RockType, JetIdx
}

func (m *Metric) GetDeltas(index string) []int {
	// determines the delta of each rocktype combination
	// index := fmt.Sprintf("%s~%s", strconv.Itoa(RockType), strconv.Itoa(JetIdx))
	retval := make([]int, 0)

	RockType, JetIdx := GetMapping(index)
	if m.Contains(RockType, JetIdx) && len(m.drops[index]) > 1 {
		for i := 0; i < len(m.drops[index]); i++ {
			if i+1 == len(m.drops[index]) {
				break
			}
			retval = append(retval, m.drops[index][i+1].height-m.drops[index][i].height)
		}
	}
	return retval
}

func (m *Metric) GetRockDeltas(index string) []int {
	retval := make([]int, 0)
	_, ok := m.drops[index]
	if ok && len(m.drops[index]) > 1 {
		for i := 0; i < len(m.drops[index]); i++ {
			if i+1 == len(m.drops[index]) {
				break
			}
			retval = append(retval, m.drops[index][i+1].rocksDropped-m.drops[index][i].rocksDropped)
		}
	}
	return retval
}

// ^^^^ NEW HOTNESS  vvvv OLD AND BUSTED
type DeltaMap struct {
	deltaList map[string][]int
}

func NewDeltaMap() DeltaMap {
	dm := DeltaMap{
		deltaList: make(map[string][]int),
	}
	return dm
}

func (d *DeltaMap) Add(RockType, JetIdx, Height int) {
	// adds to the deltamap
	index := fmt.Sprintf("%s~%s", strconv.Itoa(RockType), strconv.Itoa(JetIdx))
	if d.Contains(RockType, JetIdx) {
		d.deltaList[index] = append(d.deltaList[index], Height)
	} else {
		d.deltaList[index] = make([]int, 0)
		d.deltaList[index] = append(d.deltaList[index], Height)
	}
	/*
		So let's define a map of deltas
	*/
}

func (d *DeltaMap) Contains(RockType, JetIdx int) bool {
	index := fmt.Sprintf("%s~%s", strconv.Itoa(RockType), strconv.Itoa(JetIdx))
	_, ok := d.deltaList[index]
	// what the hell was I thinking here, this should just
	// return the 'ok' variable. I'm leaving this here so
	// you can see how much of a failure I am
	if ok {
		return true
	} else {
		return false
	}
}

func (d *DeltaMap) GetMapping(index string) (int, int) {
	// this is dumb, why make this a struct object?
	coords := strings.Split(index, "~")
	RockType, err := strconv.Atoi(coords[0])
	if err != nil {
		panic("Could not pull integer")
	}
	JetIdx, err := strconv.Atoi(coords[1])
	if err != nil {
		panic("Could not pull integer")
	}
	return RockType, JetIdx

}

// this doesn't need to be a struct function!! WTH was I thinking
func (d *DeltaMap) GetDeltas(index string) []int {
	// determines the delta of each rocktype combination
	retval := make([]int, 0)

	RockType, JetIdx := d.GetMapping(index)
	if d.Contains(RockType, JetIdx) && len(d.deltaList[index]) > 1 {

		for i := 0; i < len(d.deltaList[index]); i++ {
			if i+1 == len(d.deltaList[index]) {
				break
			}
			retval = append(retval, d.deltaList[index][i+1]-d.deltaList[index][i])
		}
		return retval
	} else {
		return make([]int, 0)
	}
}

func FindDupes(dupeArray []int) []int {
	// this function takes a slice of ints and returns
	// the numbers that are duplicated
	lenSlice := len(dupeArray)

	retVal := make([]int, 0)
	for i := 0; i < lenSlice; i++ {
		for j := i + 1; j < lenSlice; j++ {
			if dupeArray[i] == dupeArray[j] {
				retVal = append(retVal, dupeArray[i])
			}
		}
	}
	return retVal
}
