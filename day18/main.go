package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func tokenize(str string) []string {
	re := regexp.MustCompile(`\d+|[+*()]`)
	return re.FindAllString(str, -1)
}

func applyInstr(instrStack []string, numStack []int, num2 int) []int {
	if len(instrStack) == 0 {
		return append(numStack, num2)
	}
	instr := instrStack[len(instrStack) - 1]
	lastNum := len(numStack) - 1
	if instr == "(" {
		numStack = append(numStack, num2)
	} else {
		if instr == "+" {
			numStack[lastNum] += num2
		} else {
			numStack[lastNum] *= num2
		}
	}
	return numStack
}

func eval(str string) int {
	str = fmt.Sprintf("(%s)", str)
	tokens := tokenize(str)

	var instrStack []string
	var numStack []int
	for _, token := range tokens {
		switch token {
		case "+", "*", "(":
			instrStack = append(instrStack, token)
		case ")":
			lastNum := len(numStack) - 1
			num2 := numStack[lastNum]
			numStack = numStack[:lastNum]

			numStack = applyInstr(instrStack, numStack, num2)
			if len(instrStack) != 0 {
				instrStack = instrStack[:len(instrStack) - 1]
			}
		default:
			num2, _ := strconv.Atoi(token)
			numStack = applyInstr(instrStack, numStack, num2)
			instrStack = instrStack[:len(instrStack) - 1]
		}
		fmt.Println(token, instrStack, numStack)
	}
	return numStack[0]
}

func main() {
	total := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		result := eval(line)
		total += result
		fmt.Println(result, "=", line)
	}

	fmt.Println("total", total)
}

