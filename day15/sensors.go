package main

import (
	"math"
)

type Point struct {
	X, Y int
}

type Sensor struct {
	self, beacon Point
	distance     int
}

func NewSensor(sX, sY, bX, bY int) Sensor {
	s := Point{
		X: sX,
		Y: sY,
	}
	b := Point{
		X: bX,
		Y: bY,
	}
	d := GetRange(s, b)
	se := Sensor{
		self:     s,
		beacon:   b,
		distance: d,
	}

	return se
}

func GetRange(a, b Point) int {
	// given the values of the sensor and beacon,
	// this will determine the range of the sensor.
	// this is mostly an internal function
	r := math.Abs(float64(a.X)-float64(b.X)) + math.Abs(float64(a.Y)-float64(b.Y))
	return int(r)
}

func ContainsRow(s Sensor, row int) bool {
	// this function checks to see if the provided row
	// even touches the sensor at all.
	r := GetRange(s.self, s.beacon)

	highest := s.self.Y + r
	lowest := s.self.Y - r
	return (row <= highest && row >= lowest)
}

func ContainsCol(s Sensor, col int) bool {
	// this checks if the provided column touches
	// the sensor at all
	r := GetRange(s.self, s.beacon)

	highest := s.self.X + r
	lowest := s.self.X - r
	return (col <= highest && col >= lowest)
}

func ContainsPoint(s Sensor, p Point) bool {
	// just combine both. If both are true
	// then we must be within a sensor's range
	return ContainsRow(s, p.Y) && ContainsCol(s, p.X)
}

func HowManyBlocks(s Sensor, row int) int {
	if !ContainsRow(s, row) {
		return 0
	}

	r := GetRange(s.self, s.beacon)
	mostCovered := (r * 2) + 1
	// find out how offset from the highest
	// or lowest point. First, is it higher
	// or lower? Or better yet, if it's the
	// same row then the number is 2x R +1
	if row == s.self.Y {
		return mostCovered
	}

	// for each one offset, subtract 2 from mostCovered
	offset := int(math.Abs(float64(s.self.Y - row)))
	return mostCovered - (2 * offset)
}

func StartAndEndOfY(s Sensor, row int) (int, int) {
	// this function returns the X coord of the left-most covered
	// area of a sensor and the right-most area.
	span := HowManyBlocks(s, row)
	left := s.self.X - ((span - 1) / 2)
	right := s.self.X + ((span - 1) / 2)
	return left, right
}

func HowManyBeacons(sa []Sensor, row int) []Sensor {
	// this function checks each sensor and returns
	// how many sensor objects have a UNIQUE beacon
	// in the provided row.
	retval := make([]Sensor, 0)
	for _, v := range sa {
		if v.beacon.Y == row {
			if len(retval) == 0 {
				retval = append(retval, v)
			} else {
				found := false
				for _, w := range retval {
					if w.beacon.X == v.beacon.X {
						found = true
					}
				}
				if !found {
					retval = append(retval, v)
				}
			}
		}
	}
	return retval
}

func GetBorder(s Sensor) []Point {
	// this function will return every point that is d+1 from a sensor.
	// It will completely reduce my workspace from 4_000_000*4_000_000 to
	// however many points border a sensor's range. Total credit to 0xdf
	// for explaining this on his youtube channel.
	retPoints := make([]Point, 0)

	for i := 0; i <= s.distance; i++ {
		ascLeftMost := Point{
			X: s.self.X - (s.distance - i) - 1,
			Y: s.self.Y - i,
		}
		ascRightMost := Point{
			X: s.self.X + (s.distance - i) + 1,
			Y: s.self.Y - i,
		}
		descLeftMost := Point{
			X: s.self.X - (s.distance - i) - 1,
			Y: s.self.Y + i,
		}
		descRightMost := Point{
			X: s.self.X + (s.distance - i) + 1,
			Y: s.self.Y + i,
		}
		retPoints = append(retPoints, ascLeftMost, ascRightMost, descLeftMost, descRightMost)
	}

	// get top most point
	top := Point{
		X: s.self.X,
		Y: s.self.Y - s.distance - 1,
	}

	// bottom most point
	bottom := Point{
		X: s.self.X,
		Y: s.self.Y + s.distance + 1,
	}

	// left most point
	left := Point{
		X: s.self.X - s.distance - 1,
		Y: s.self.Y,
	}

	// right most point
	right := Point{
		X: s.self.X + s.distance + 1,
		Y: s.self.Y,
	}

	retPoints = append(retPoints, top, bottom, left, right)

	// for i := 1; i <= s.distance; i++ {
	// 	ascLeftSide := Point{
	// 		X: s.self.X - s.distance - i,
	// 		Y: s.self.Y + i,
	// 	}
	// 	ascRightSide := Point{
	// 		X: s.self.X + s.distance + i,
	// 		Y: s.self.Y + i,
	// 	}
	// 	descLeftSide := Point{
	// 		X: s.self.X - s.distance - i,
	// 		Y: s.self.Y - i,
	// 	}
	// 	descRightSide := Point{
	// 		X: s.self.X + s.distance + i,
	// 		Y: s.self.Y - i,
	// 	}
	// 	retPoints = append(retPoints, ascLeftSide, ascRightSide, descLeftSide, descRightSide)
	// }
	return retPoints
}
