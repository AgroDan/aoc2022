package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func countSmallerTrees(compareNum int, dirArray []int) int {
	// will traverse the dirArray and count how many trees
	// we can see
	counter := 0
	for _, v := range dirArray {
		counter++
		if compareNum <= v {
			break
		}
	}
	return counter
}

func getScenicScore(x int, y int, t [][]int) int {
	// builds the same slice of integers as isVisible in day8 part 1,
	// but instead returns a scenic score of how many trees are visible
	// in all 4 directions. Count how many trees until you reach the edge
	// or another tree at the same height or higher than the current tree.
	var up, down, left, right []int
	var leftScore, rightScore, upScore, downScore int
	num := t[y][x] // for simplicity's sake

	// work with left first
	if x == 0 {
		leftScore = 0
	} else {
		for i := x - 1; i >= 0; i-- {
			left = append(left, t[y][i])
		}
		leftScore = countSmallerTrees(num, left)
		// fmt.Printf("LeftScore: %d\n", leftScore)
	}

	// now the right
	if x == len(t[y])-1 {
		rightScore = 0
	} else {
		for i := x + 1; i < len(t[y]); i++ {
			right = append(right, t[y][i])
		}
		rightScore = countSmallerTrees(num, right)
		// fmt.Printf("RightScore: %d\n", rightScore)
	}

	// now up
	if y == 0 {
		upScore = 0
	} else {
		for i := y - 1; i >= 0; i-- {
			up = append(up, t[i][x])
		}
		upScore = countSmallerTrees(num, up)
		// fmt.Printf("UpScore: %d\n", upScore)
	}

	// finally down
	if y == len(t)-1 {
		downScore = 0
	} else {
		for i := y + 1; i < len(t); i++ {
			down = append(down, t[i][x])
		}
		downScore = countSmallerTrees(num, down)
		// fmt.Printf("DownScore: %d\n", downScore)
	}

	return upScore * downScore * leftScore * rightScore
}

func isVisible(x int, y int, t [][]int) bool {
	// given an x/y coordinate, it will build a slice of integers
	// containing the values going straight up, straight down,
	// left and right. Then it will determine the tree's visibility.
	// NOTE: y is resolved FIRST! so t[y][x]

	var up, down, left, right []int

	// quick wins
	if x == 0 || x == len(t[y])-1 {
		return true
	}
	if y == 0 || y == len(t)-1 {
		return true
	}

	// otherwise, build the slices.
	// all the ones above
	for i := y - 1; i >= 0; i-- {
		up = append(up, t[i][x])
	}
	// all below
	for i := y + 1; i < len(t); i++ {
		down = append(down, t[i][x])
	}
	// all to the left
	for i := x - 1; i >= 0; i-- {
		left = append(left, t[y][i])
	}
	// all to the right
	for i := x + 1; i < len(t[y]); i++ {
		right = append(right, t[y][i])
	}

	num := t[y][x] // for simplicity's sake

	dirArray := [4]*[]int{&up, &down, &left, &right}

	for _, v := range dirArray {
		visibleDirection := true
		for _, w := range *v {
			if num <= w {
				visibleDirection = false
				break
			}
		}
		if visibleDirection {
			return true
		}
	}
	return false
}

func main() {
	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	trees := make([][]int, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		temp := make([]int, 0)
		for _, ch := range line {
			num, err := strconv.Atoi(fmt.Sprintf("%c", ch))
			if err != nil {
				fmt.Println("Could not convert!", err)
			}
			temp = append(temp, num)
		}
		trees = append(trees, temp)
	}
	// fmt.Printf("Contents of trees: +%v\n", trees)

	// Now loop through every single number, providing x and y coords
	// fmt.Println("Is 2/2 visible?", isVisible(2, 2, trees))
	// fmt.Printf("Scenic score of x:2 y:3 (value:%d): %d", trees[3][2], getScenicScore(2, 3, trees))

	var highScore int = 0
	for i, v := range trees { // y
		for j := range v { // x
			tryScore := getScenicScore(j, i, trees)
			if tryScore > highScore {
				highScore = tryScore
			}
		}
	}
	fmt.Println("Highest Scenic Score:", highScore)
}
