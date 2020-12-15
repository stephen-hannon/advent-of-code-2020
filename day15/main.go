package main

import (
	"fmt"
)

const NumTurns = 30000000

func main() {
	ages := make(map[int]int)
	var input []int
	input = []int{0, 3, 6}
	input = []int{3, 1, 2}
	input = []int{0,14,6,20,1,4}

	turn := 1
	for ; turn <= len(input); turn++ {
		num := input[turn - 1]
		ages[num] = turn
	}

	prevNum := input[len(input) - 1]
	for ; turn <= NumTurns; turn++ {
		lastSeen, wasSeen := ages[prevNum]
		var newNumber int
		if wasSeen {
			newNumber = turn - 1 - lastSeen
		} else {
			newNumber = 0
		}
		if (turn & 0xfffff) == 0 || turn == 2020 || turn == NumTurns {
			fmt.Printf("%d.\t%d -> %d\n", turn, prevNum, newNumber)
		}
		ages[prevNum] = turn - 1
		prevNum = newNumber
	}
}

