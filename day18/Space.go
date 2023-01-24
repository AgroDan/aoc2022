package main

import "errors"

// set up a queue object
// shamelessly stolen from...me. Day 12.
type Queue struct {
	q []Point
}

func (q *Queue) Push(p Point) {
	q.q = append(q.q, p)
}

func (q *Queue) Pop() (Point, error) {
	if len(q.q) == 0 {
		return Point{}, errors.New("empty queue")
	}
	np := q.q[0]
	q.q = q.q[1:]
	return np, nil
}

type Space struct {
	area [][][]rune
}

func NewSpace(size int) Space {
	// this function creates a new "Space"
	// object in the dimesional size you
	// want.
	s := Space{}
	s.area = make([][][]rune, 0)
	for i := 0; i < size; i++ {
		a := make([][]rune, 0)
		for j := 0; j < size; j++ {
			b := make([]rune, 0)
			for k := 0; k < size; k++ {
				b = append(b, '.')
			}
			a = append(a, b)
		}
		s.area = append(s.area, a)
	}
	return s
}

func (s *Space) GrowSpace(amount int) {
	// this used to be a lot more complicated. But
	// instead of growing onto another array, I will
	// simply create a new one, populate it with every
	// object in the current array, and throw away the
	// old array object.
	n := NewSpace(len(s.area) + amount)

	for i := 0; i < len(s.area); i++ {
		for j := 0; j < len(s.area); j++ {
			for k := 0; k < len(s.area); k++ {
				n.area[i][j][k] = s.area[i][j][k]
			}
		}
	}
	s.area = nil
	s.area = n.area
	n.area = nil // deallocate
}

func (s *Space) CheckPoint(p Point, isClosed bool) int {
	// given a point, this checks +1 or -1 of each
	// side to see if it is open or closed. This assumes
	// that the point given is already marked in the spacemap.
	// isClosed = true, look for anything but '.'
	// isClosed = false, look for anything but '#'

	var check rune
	if isClosed {
		check = '.'
	} else {
		check = '#'
	}
	totalSides := 0
	// first we'll work with the X axis.
	if p.X+1 >= len(s.area) {
		totalSides++
	} else if s.area[p.X+1][p.Y][p.Z] != check {
		totalSides++
	}

	if p.X-1 < 0 {
		totalSides++
	} else if s.area[p.X-1][p.Y][p.Z] != check {
		totalSides++
	}

	// now the Y axis
	if p.Y+1 >= len(s.area[p.X]) {
		totalSides++
	} else if s.area[p.X][p.Y+1][p.Z] != check {
		totalSides++
	}

	if p.Y-1 < 0 {
		totalSides++
	} else if s.area[p.X][p.Y-1][p.Z] != check {
		totalSides++
	}

	// finally the Z axis
	if p.Z+1 >= len(s.area[p.X][p.Y]) {
		totalSides++
	} else if s.area[p.X][p.Y][p.Z+1] != check {
		totalSides++
	}

	if p.Z-1 < 0 {
		totalSides++
	} else if s.area[p.X][p.Y][p.Z-1] != check {
		totalSides++
	}
	return totalSides
}

func (s *Space) IsEdge(p Point) bool {
	// this will check to see if the given point
	// is at an edge or not.

	if p.X >= len(s.area)-1 || p.X <= 0 {
		return true
	}
	if p.Y >= len(s.area[p.X])-1 || p.Y <= 0 {
		return true
	}
	if p.Z >= len(s.area[p.X][p.Y])-1 || p.Z <= 0 {
		return true
	}
	return false
}

func (s *Space) GetEmptyEdges(p Point) []Point {
	// given a particular point, this returns all
	// "available" points around it. It returns any
	// point around the given point that contains an
	// empty "air" unit that ISN'T a border. This
	// function can be used in conjunction with a
	// Breadth-First-Search function to search specifically
	// if that air object is enclosed or open. Essentially
	// looking for "bubbles" within an enclosed set
	// of "lava" units.
	air := make([]Point, 0)

	// one day I will find out how to make this better.
	// seems like there are WAY too many if statements
	// and this really annoys my Don't-Repeat-Yourself annoyance.

	// check around the X axis

	if !(p.X+1 >= len(s.area)) && s.area[p.X+1][p.Y][p.Z] == '.' {
		air = append(air, Point{
			X: p.X + 1,
			Y: p.Y,
			Z: p.Z,
		})
	}
	if !(p.X-1 < 0) && s.area[p.X-1][p.Y][p.Z] == '.' {
		air = append(air, Point{
			X: p.X - 1,
			Y: p.Y,
			Z: p.Z,
		})
	}

	// check around the Y axis

	if !(p.Y+1 >= len(s.area[p.X])) && s.area[p.X][p.Y+1][p.Z] == '.' {
		air = append(air, Point{
			X: p.X,
			Y: p.Y + 1,
			Z: p.Z,
		})
	}
	if !(p.Y-1 < 0) && s.area[p.X][p.Y-1][p.Z] == '.' {
		air = append(air, Point{
			X: p.X,
			Y: p.Y - 1,
			Z: p.Z,
		})
	}

	// finally, check around the Z axis

	if !(p.Z+1 >= len(s.area[p.X][p.Y])) && s.area[p.X][p.Y][p.Z+1] == '.' {
		air = append(air, Point{
			X: p.X,
			Y: p.Y,
			Z: p.Z + 1,
		})
	}
	if !(p.Z-1 < 0) && s.area[p.X][p.Y][p.Z-1] == '.' {
		air = append(air, Point{
			X: p.X,
			Y: p.Y,
			Z: p.Z - 1,
		})
	}

	return air
}

type Point struct {
	X, Y, Z int
}

func PointsEqual(a, b Point) bool {
	// returns if two points are
	// exactly equal. This is for
	// convenience.
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z
}

func RemoveDuplicatePoints(pointSlice []Point) []Point {
	// this function removes duplicate points from a
	// Point{} slice.

	var exists = struct{}{}
	keys := make(map[Point]struct{})
	retval := make([]Point, 0)

	for _, v := range pointSlice {
		if _, e := keys[v]; !e {
			keys[v] = exists
			retval = append(retval, v)
		}
	}
	return retval
}

func (s *Space) DetectBubble(p Point) []Point {
	// this performs a depth-first-search function to
	// check if the point is a bubble. If it manages to
	// reach the edge of the measured 3-D space, then it
	// isn't a bubble and will return an empty []Point
	// object. If it IS a bubble, it will return a slice
	// of points that encompass the entire bubble.

	// first, is this point a lava piece? If so bomb out
	// immediately ya turkey
	if s.area[p.X][p.Y][p.Z] == '#' {
		return make([]Point, 0)
	}

	retval := make([]Point, 0)

	// create our queue, set our starting point as the first
	q := Queue{}
	q.Push(p)

	// create a finite set
	var exists = struct{}{}
	visited := make(map[Point]struct{})

	// mark the starting point as visited
	visited[p] = exists

	for {
		workingPoint, err := q.Pop()
		if err != nil {
			// empty queue, drop out
			break
		}

		// check if workingpoint is an edge of the map
		if s.IsEdge(workingPoint) {
			// if we hit an edge, this isn't a bubble.
			return make([]Point, 0)
		}

		// we made it here, we're air and not an edge.
		retval = append(retval, workingPoint)

		// get additional edges
		for _, v := range s.GetEmptyEdges(workingPoint) {
			_, didVisit := visited[v]
			if !didVisit {
				q.Push(v)
				visited[v] = exists
			}
		}
	}
	return retval
}
