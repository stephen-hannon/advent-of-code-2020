package main

import (
	"fmt"
	"sort"
)

func findPathsTo(nums []int, max int) (count int) {
	numPaths := make([]int, max + 1)
	numPaths[max] = 1

	for i := len(nums) - 2; i >= 0; i-- {
		num := nums[i]
		var numPaths1, numPaths2, numPaths3 int
		if num + 1 <= max {
			numPaths1 = numPaths[num + 1]
		}
		if num + 2 <= max {
			numPaths2 = numPaths[num + 2]
		}
		if num + 3 <= max {
			numPaths3 = numPaths[num + 3]
		}
		numPaths[num] = numPaths1 + numPaths2 + numPaths3
	}
	fmt.Println(numPaths)
	return numPaths[0]
}

func main() {
	input := []int{0}
	for {
		var num int
		numScanned, _ := fmt.Scanf("%d", &num)
		if numScanned == 0 {
			break
		}

		input = append(input, num)
	}

	sort.Ints(input)
	fmt.Println(input)

	var diffs1, diffs3 int
	for index, num := range input {
		if index == 0 {
			continue
		}

		diff := num - input[index - 1]
		// fmt.Printf("%d - %d = %d\n", num, input[index-1], diff)
		if diff == 1 {
			diffs1++
		} else if diff == 3 {
			diffs3++
		}
	}
	diffs3++ // account for built-in adapter

	fmt.Printf("%d * %d = %d\n", diffs1, diffs3, diffs1 * diffs3)

	maxNum := input[len(input) - 1]
	pathsFrom0 := findPathsTo(input, maxNum)
	fmt.Println("Paths from 0:", pathsFrom0)
}

