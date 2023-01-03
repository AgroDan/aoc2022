package main

/*
	This is a merge sort algorithm written in go
	because I always wanted to write one.
*/

func MergeSort(u [][]interface{}) [][]interface{} {
	if len(u) == 1 {
		return u
	} else {
		half := len(u) / 2
		left := u[:half]
		right := u[half:]
		lsort := MergeSort(left)
		rsort := MergeSort(right)
		return merge(lsort, rsort)
	}
}

func merge(left, right [][]interface{}) [][]interface{} {
	res := [][]interface{}{}
	for {
		if len(left) == 0 || len(right) == 0 {
			break
		}
		ordered, _ := IsInOrder(left[0], right[0])
		if ordered {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}

	if len(left) > 0 {
		for {
			if len(left) == 0 {
				break
			}
			res = append(res, left[0])
			left = left[1:]
		}
	}
	if len(right) > 0 {
		for {
			if len(right) == 0 {
				break
			}
			res = append(res, right[0])
			right = right[1:]
		}
	}
	return res
}
