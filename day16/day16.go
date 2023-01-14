package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseValves(lines []string) map[string]Valve {
	// this function parses each line
	// Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
	m := make(map[string]Valve)
	for _, line := range lines {
		words := strings.Split(line, " ")
		valveName := words[1]
		r := strings.Split(words[4], "=")[1]
		rate, err := strconv.Atoi(strings.TrimRight(r, ";"))
		if err != nil {
			panic("Could not convert")
		}
		va := strings.Join(words[9:], " ")
		valves := strings.FieldsFunc(va, func(r rune) bool {
			if r == ',' || r == ' ' {
				return true
			}
			return false
		})

		v := NewValve(rate, valves...)
		m[valveName] = v
	}
	return m
}

func BuildMatrix(m *map[string]Valve) map[string]map[string]int {
	// this function returns a matrix of all the "quickest routes"
	// to each destination using the Floyd Warshall Algorithm.
	matrix := make(map[string]map[string]int, 0)
	inf := 99

	for k := range *m {
		matrix[k] = make(map[string]int)
		for j := range *m {
			if k == j {
				matrix[k][j] = 0
			} else {
				matrix[k][j] = inf
			}
		}
	}
	// at this point, the matrix is completed and can be reference with
	// matrix["from"]["to"]. Now work on the edges that we know of.
	for k, v := range *m {
		for _, tunnel := range v.tunnels {
			matrix[k][tunnel] = 1
		}
	}

	// edges are built at this point. Now IN THEORY we can build a new
	// matrix out of this
	for k := range matrix {
		for i := range matrix {
			for j := range matrix {
				if matrix[i][j] > matrix[i][k]+matrix[k][j] {
					matrix[i][j] = matrix[i][k] + matrix[k][j]
				}
			}
		}
	}
	return matrix
}

func PrintMatrix(matrix *map[string]map[string]int) {
	// this function simply prints the contents of the matrix.
	// NOTE: This function is wonky. It works kinda, but it
	// never matches up due to the un-ordered nature of maps in Go.
	// So it looks pretty but it's wrong, though it can be fixed if
	// I cared enough.
	fmt.Printf("     ")
	for k := range *matrix {
		fmt.Printf("%s  ", k)
	}
	fmt.Printf("\n")

	for k, v := range *matrix {
		fmt.Printf("%s: ", k)
		for _, w := range v {
			fmt.Printf(" %02d ", w)
		}
		fmt.Printf("\n")
	}
}

func bestChoice(current string, minutes int, matrix map[string]map[string]int, m *map[string]Valve) string {
	// this function will take the current position, the current amount of minutes
	// left, the built matrix, and the full map to determine the next place we
	// should be moving. This will NOT do the proper path-finding, if that is even
	// necessary. It will simply return the next place we should be moving to.
	//
	// NOTE: this was the wrong way of going about doing this. This function was
	// unused but I kept it anyway

	var choice string
	var highestFlow int = 0
	for k, v := range *m {
		if v.isOpen {
			// valve is already open, ignore
			continue
		}

		// get the distance
		distance := matrix[current][k]
		totalFlow := ((minutes - distance) - 1) * v.flowRate // additional -1 to open the valve once we're there
		if totalFlow > highestFlow {
			highestFlow = totalFlow
			choice = k
			fmt.Printf("Highest so far: %s with a total flow of %d and a distance of %d\n", choice, highestFlow, distance)
		}
	}
	return choice
}

func decideValves(paths []string, m map[string]Valve) (string, int) {
	// given the superset of valves as m, and which valve
	// we are currently have access to get to, this looks
	// at each flowrate of the given valves against the m
	// superset and decides which valve offers the best
	// path _as well as_ the flowrate.
	bestRate := 0
	bestValve := ""

	for _, v := range paths {
		if m[v].flowRate > bestRate && !m[v].isOpen {
			bestRate = m[v].flowRate
			bestValve = v
		}
	}
	return bestValve, bestRate
}

// depth first search
func DFS(current string, minutes int, matrix map[string]map[string]int, thisPath Path, m map[string]Valve) []Path {
	var p []Path
	for k, v := range m {
		distance := minutes - matrix[current][k] - 1
		if distance <= 0 || thisPath.BeenThere(k) || v.flowRate == 0 {
			continue
			// we've been there or it's way too far or it's just not worth it
		}
		thisFlow := v.flowRate * distance
		// copy the path, this screwed me up before I realized this was necessary!
		newPath := thisPath.CopyPath()
		// otherwise, we haven't been there, so let's see what happens
		newPath.AddToPath(k, thisFlow)
		p = append(p, DFS(k, distance, matrix, newPath, m)...)
	}
	p = append(p, thisPath)
	return p
}

func FindTwoHighest(pathList []Path) (Path, Path) {
	// first, find the unique pairs. 2-dimensional array the holds the
	// index of the unique pairs
	uniquePairs := make([][2]int, 0)

	for o, outer := range pathList {
		for i, inner := range pathList {
			unique := true
			for k := range outer.breadCrumbs.location {
				if inner.breadCrumbs.Contains(k) {
					// this is not unique
					unique = false
				}
			}
			if unique {
				var pair [2]int
				pair[0] = o
				pair[1] = i
				uniquePairs = append(uniquePairs, pair)
			}
		}
	}

	var highest int = -1
	var topScore int = 0
	var combinedScore int = 0
	for i, v := range uniquePairs {
		combinedScore = pathList[v[0]].totalPressure + pathList[v[1]].totalPressure
		if combinedScore > topScore {
			topScore = combinedScore
			highest = i
		}
	}
	return pathList[uniquePairs[highest][0]], pathList[uniquePairs[highest][1]]

}

func main() {
	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	readFile.Close()

	// time this
	t1 := time.Now()

	// parse the valves from the input
	m := parseValves(lines)

	// part 1, first build the matrix
	matrix := BuildMatrix(&m)
	paths := DFS("AA", 30, matrix, NewPath(), m)
	highest := 0
	for _, v := range paths {
		if v.totalPressure > highest {
			highest = v.totalPressure
		}
	}
	fmt.Printf("Highest from Part 1: %d\n", highest)
	fmt.Printf("Time elapsed: %s\n", time.Since(t1))

	// Now for part 2.
	t2 := time.Now()
	pathsPartTwo := DFS("AA", 26, matrix, NewPath(), m)
	t, s := FindTwoHighest(pathsPartTwo)

	fmt.Printf("First: %d, Second: %d, Total: %d\n", t.totalPressure, s.totalPressure, s.totalPressure+t.totalPressure)
	fmt.Printf("Time elapsed: %s\n", time.Since(t2))
}
