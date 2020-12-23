package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func parseStack(scanner *bufio.Scanner) (stack []int) {
	scanner.Scan() // skip player name
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		num, _ := strconv.Atoi(line)
		stack = append(stack, num)
	}
	return
}

func playTurn(stack1, stack2 []int) ([]int, []int) {
	card1 := stack1[0]
	stack1 = stack1[1:]

	card2 := stack2[0]
	stack2 = stack2[1:]

	if card1 > card2 {
		stack1 = append(stack1, card1, card2)
	} else {
		stack2 = append(stack2, card2, card1)
	}

	// fmt.Println("1", stack1)
	// fmt.Println("2", stack2)
	return stack1, stack2
}

func score(stack []int) (result int) {
	for index, card := range stack {
		result += card * (len(stack) - index)
	}
	return
}

func stringifyDecks(stack1, stack2 []int) string {
	return fmt.Sprint(stack1, stack2)
}

func playGameRecursive(stack1, stack2 []int) (winner int, winStack []int) {
	arrangementSeen := make(map[string]bool)

	for len(stack1) != 0 && len(stack2) != 0 {
		if arrangementSeen[stringifyDecks(stack1, stack2)] {
			winner = 1
			winStack = stack2
			return
		}
		arrangementSeen[stringifyDecks(stack1, stack2)] = true

		card1 := stack1[0]
		stack1 = stack1[1:]

		card2 := stack2[0]
		stack2 = stack2[1:]

		var turnWinner int
		if card1 <= len(stack1) && card2 <= len(stack2) {
			stack1Copy := make([]int, card1)
			stack2Copy := make([]int, card2)
			copy(stack1Copy, stack1)
			copy(stack2Copy, stack2)
			turnWinner, _ = playGameRecursive(stack1Copy, stack2Copy)
		} else {
			if card1 > card2 {
				turnWinner = 1
			} else {
				turnWinner = 2
			}
		}

		if turnWinner == 1 {
			stack1 = append(stack1, card1, card2)
		} else {
			stack2 = append(stack2, card2, card1)
		}
	}

	if len(stack2) == 0 {
		winner = 1
		winStack = stack1
	} else {
		winner = 2
		winStack = stack2
	}

	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	stack1Orig := parseStack(scanner)
	stack2Orig := parseStack(scanner)
	stack1 := stack1Orig
	stack2 := stack2Orig

	for len(stack1) != 0 && len(stack2) != 0 {
		stack1, stack2 = playTurn(stack1, stack2)
	}
	winStack := stack1
	if len(stack1) == 0 {
		winStack = stack2
	}

	fmt.Println("winner", winStack)
	winScore := score(winStack)
	fmt.Println("score", winScore)

	// Part 2
	winner, winStack2 := playGameRecursive(stack1Orig, stack2Orig)
	fmt.Println("2 winner:", winner, winStack2)
	winScore2 := score(winStack2)
	fmt.Println("score 2:", winScore2)
}

