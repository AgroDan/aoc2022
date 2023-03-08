package main

import (
	"math"
	"strconv"
	"strings"
)

func ReverseString(s string) string {
	// this will reverse a string
	runes := []rune(s)
	retval := make([]rune, 0)

	for i := len(runes) - 1; i >= 0; i-- {
		retval = append(retval, runes[i])
	}
	return string(retval)
}

func ConvertToBaseFive(num int) int {
	// this will return a base-5 integer of a numeric number.
	conv := make([]int, 0)
	for num != 0 {
		conv = append(conv, num%5)
		num /= 5
	}
	retval := 0
	for i := len(conv) - 1; i >= 0; i-- {
		retval = (retval * 10) + conv[i]
	}
	return retval
}

func ConvertToBaseFiveSlice(num int) []int {
	// this will return a base-5 integer in which the digits are
	// separated as items within a slice. Note, this will be backwards
	// which saves me time since I will essentially be working backwards.
	conv := make([]int, 0)
	for num != 0 {
		conv = append(conv, num%5)
		num /= 5
	}
	// rev := make([]int, 0)
	// for i := len(conv) - 1; i <= 0; i-- {
	// 	rev = append(rev, conv[i])
	// }
	return conv
}

func ConvertFromSnafu(snafu string) int {
	// this function will accept a SNAFU number and
	// return an integer.
	results := make([]int, 0)

	for i, v := range ReverseString(snafu) {
		num := 0
		switch v {
		case '2':
			num = 2
		case '1':
			num = 1
		case '0':
			num = 0
		case '-':
			num = -1
		case '=':
			num = -2
		default:
			num = 999999
		}

		digitPlace := int(math.Pow(5, float64(i)))
		results = append(results, digitPlace*num)
	}

	var retval int = 0
	for _, v := range results {
		retval += v
	}
	return retval
}

func ConvertToSnafu(num int) string {
	// this will use the same idea as converting from a balanced ternary
	// system as described here: https://www.geeksforgeeks.org/balanced-ternary-number-system/
	retval := make([]string, 0)

	for num != 0 {
		r := num % 5
		num /= 5

		if r <= 2 {
			retval = append(retval, strconv.Itoa(r))
		} else {
			switch r {
			case 3:
				retval = append(retval, "=")
				// add 1 to the next number working right to left
				num++
			case 4:
				retval = append(retval, "-")
				// add 1 to the next number working right to left
				num++
			}
		}
	}

	s := strings.Join(retval, "")
	return ReverseString(s)
}

// func ConvertToSnafu(numSlice []int) string {
// 	// this will use the same idea as converting from a balanced ternary
// 	// system as described here: https://www.geeksforgeeks.org/balanced-ternary-number-system/
// 	retval := make([]string, 0)
// 	workingNum := 0
// 	for _, num := range numSlice {
// 		workingNum += num
// 		if workingNum <= 2 {
// 			retval = append(retval, strconv.Itoa(workingNum))
// 			workingNum = 0
// 		} else {
// 			// otherwise, either 3 or 4, or -/=
// 			switch workingNum {
// 			case 3:
// 				retval = append(retval, "=")
// 				workingNum = 1
// 			case 4:
// 				retval = append(retval, "-")
// 				workingNum = 1
// 			case 5:
// 				retval = append(retval, "0")
// 				workingNum = 1
// 			case 6:
// 				retval = append(retval, "1")
// 				workingNum = 1
// 			}
// 		}
// 	}
// 	s := strings.Join(retval, "")
// 	return ReverseString(s)
// }

// func ConvertToSnafu(n int) string {
// 	digits := "=-012"
// 	if n == 0 {
// 		return string(digits[0])
// 	}
// 	var digitsList []int
// 	for n != 0 {
// 		rem := n % 5
// 		switch rem {
// 		case 3:
// 			digitsList = append(digitsList, -1)
// 			n += 2
// 		case 4:
// 			digitsList = append(digitsList, 1)
// 			n += 1
// 		default:
// 			digitsList = append(digitsList, rem)
// 		}
// 		n /= 5
// 	}

// 	// now build the string
// 	var retval string
// 	for i := len(digitsList) - 1; i >= 0; i-- {
// 		d := digitsList[i]
// 		retval += string(digits[d+1])
// 	}
// 	return retval
// }
