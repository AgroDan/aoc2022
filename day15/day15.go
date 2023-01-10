package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	// YLINE int = 10
	YLINE int = 2000000
	LOW   int = 0
	// HIGH int = 20
	HIGH int = 4000000
)

func main() {
	// readFile, err := os.Open("testinput")
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

	// parse the scans and beacons
	knownSensors := make([]Sensor, 0)
	for _, v := range lines {
		s := strings.FieldsFunc(v, func(r rune) bool {
			if r == '=' || r == ',' || r == ':' {
				return true
			} else {
				return false
			}
		})
		/*
		   Fields are: sensor X: s[1]
		               sensor Y: s[3]
		               beacon X: s[5]
		               beacon Y: s[7]
		*/
		sensorX, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Println("Could not read sensor X:", err)
		}

		sensorY, err := strconv.Atoi(s[3])
		if err != nil {
			fmt.Println("Could not read sensor Y:", err)
		}

		beaconX, err := strconv.Atoi(s[5])
		if err != nil {
			fmt.Println("Could not read beacon X:", err)
		}
		beaconY, err := strconv.Atoi(s[7])
		if err != nil {
			fmt.Println("Could not read beacon Y:", err)
		}

		knownSensors = append(knownSensors, NewSensor(sensorX, sensorY, beaconX, beaconY))
	}

	// probably don't need to loop again but whatever

	s := NewSet()

	for _, v := range knownSensors {
		// fmt.Printf("Does the sensor touch row 10? %t\n", ContainsRow(v, 10))
		if !ContainsRow(v, YLINE) {
			continue
		}
		left, right := StartAndEndOfY(v, YLINE)
		for i := left; i <= right; i++ {
			num := strconv.Itoa(i)
			s.Add(num)
		}
		// fmt.Printf("%d-%d\n", left, right)
	}

	counter := 0
	for range s.m {
		counter++
	}

	// Now check to see if there is a beacon in the row!
	// If so, subtract one from the counter
	// for _, v := range knownSensors {
	// 	if v.beacon.Y == 10 {
	// 		num := strconv.Itoa(v.beacon.X)
	// 		if s.Contains(num) {
	// 			counter--
	// 			fmt.Printf("Got one: %s\n", num)
	// 		}
	// 	}
	// }

	// let's remove any beacons on that line
	beacons := HowManyBeacons(knownSensors, YLINE)

	fmt.Printf("Count: %d\n", counter-len(beacons))

	// get sensor 12

	// fmt.Printf("Sensor 12: %v\n", knownSensors[12])
	// border := GetBorder(knownSensors[12])
	// fmt.Printf("Border of 12: %v\n", border)
	testGroup := make([]Point, 0)
	for _, s := range knownSensors {
		testGroup = append(testGroup, GetBorder(s)...)
	}

	for _, v := range testGroup {
		// skip if this is out of bounds
		if v.X > HIGH || v.X < LOW || v.Y > HIGH || v.Y < LOW {
			continue
		}
		found := false
		for _, w := range knownSensors {
			distance := GetRange(w.self, v)
			if distance <= w.distance {
				found = true
			}
		}
		if !found {
			fmt.Printf("\nFound potential: X: %d, Y: %d\n", v.X, v.Y)
			fmt.Printf("Tuning Frequency: %d\n", ((v.X * 4000000) + v.Y))
		}
	}

	// c := 0
	// for i := LOW; i <= HIGH; i++ {
	// 	for j := LOW; j <= HIGH; j++ {
	// 		p := Point{
	// 			X: j,
	// 			Y: i,
	// 		}
	// 		found := false
	// 		for _, v := range knownSensors {
	// 			distance := GetRange(v.self, p)
	// 			if distance <= v.distance {
	// 				found = true
	// 			}
	// 		}
	// 		if !found {
	// 			fmt.Printf("\nFound potential: X: %d, Y: %d\n", p.X, p.Y)
	// 			fmt.Printf("Tuning Frequency: %d\n", ((p.X * 4000000) + p.Y))
	// 		}
	// 		fmt.Printf("Tick: %d\r", c)
	// 		c++
	// 	}
	// }

	// fmt.Printf("Set: %v\n", s.m)
	// fmt.Printf("Sensors: %v\n", knownSensors)
}
