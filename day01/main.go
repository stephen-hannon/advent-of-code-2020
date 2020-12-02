package main

import (
	"fmt"
)

const TARGET_NUM = 2020

func findSumTo(target int, isIncludedArr []bool) (number int, success bool) {
	for number, isIncluded := range isIncludedArr {
		if !isIncluded {
			continue
		}

		partnerNumber := target - number
		if partnerNumber < 1 || !isIncludedArr[partnerNumber] {
			continue
		}

		return number, true
	}

	return -1, false
}

func main() {
	isIncludedArr := make([]bool, TARGET_NUM+1)

	var num int
	for {
		num_scanned, _ := fmt.Scanf("%d", &num)
		if num_scanned == 0 {
			break
		}
		isIncludedArr[num] = true
	}

	// Part 1
	number, _ := findSumTo(TARGET_NUM, isIncludedArr)
	partnerNumber := TARGET_NUM - number
	fmt.Printf("%d * %d = %d\n", number, partnerNumber, number*partnerNumber)

	// Part 2
	for number, isIncluded := range isIncludedArr {
		if !isIncluded {
			continue
		}

		partnerNumber := TARGET_NUM - number
		secondNumber, success := findSumTo(partnerNumber, isIncludedArr)
		if !success {
			continue
		}

		thirdNumber := partnerNumber - secondNumber
		fmt.Printf("%d * %d * %d = %d\n", number, secondNumber, thirdNumber, number*secondNumber*thirdNumber)
		break
	}
}
