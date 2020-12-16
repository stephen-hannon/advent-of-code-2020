package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
	"strings"
	"strconv"
)

const MaxNum = 1000

type Range struct {
	min, max int
}

type IndexedSlice struct {
	index int
	slice []int
}

func markValid(valid *[MaxNum]bool, r Range) {
	for i := r.min; i <= r.max; i++ {
		valid[i] = true
	}
}

func isValidField(field [2]Range, num int) bool {
	return (num >= field[0].min && num <= field[0].max) || (num >= field[1].min && num <= field[1].max)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func pprint(slice []IndexedSlice) {
	for _, subslice := range slice {
		fmt.Println(subslice)
	}
	fmt.Println("")
}

func main() {
	var fields [][2]Range
	var valid [MaxNum]bool

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		splitLine := strings.Split(line, ": ")

		rangesStr := splitLine[1]
		var range1, range2 Range
		fmt.Sscanf(rangesStr, "%d-%d or %d-%d", &range1.min, &range1.max, &range2.min, &range2.max)
		fields = append(fields, [2]Range{range1, range2})
		markValid(&valid, range1)
		markValid(&valid, range2)
	}

	scanner.Scan() // your ticket
	scanner.Scan() // your ticket numbers

	line := scanner.Text()
	splitLine := strings.Split(line, ",")
	validFields := make([]IndexedSlice, len(splitLine))
	yourTicket := make([]int, len(splitLine))
	for index, str := range splitLine {
		num, _ := strconv.Atoi(str)
		yourTicket[index] = num
		validFields[index].index = index
		for fieldIndex, field := range fields {
			if isValidField(field, num) {
				validFields[index].slice = append(validFields[index].slice, fieldIndex)
			}
		}
	}

	scanner.Scan() // blank line
	scanner.Scan() // nearby tickets

	var validTickets [][]int
	errorRate := 0
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ",")

		var ticket []int
		isValidTicket := true

		for _, str := range splitLine {
			num, _ := strconv.Atoi(str)
			ticket = append(ticket, num)
			if !valid[num] {
				errorRate += num
				isValidTicket = false
			}
		}
		if isValidTicket {
			validTickets = append(validTickets, ticket)
		}
	}
	fmt.Println("error rate", errorRate)

	// Part 2
	for _, ticket := range validTickets {
		for index, num := range ticket {
			var newValidFields []int
			for _, field := range validFields[index].slice {
				if isValidField(fields[field], num) {
					newValidFields = append(newValidFields, field)
				} else {
					// fmt.Printf("validTickets[%d][%d] = %d:\telim field %d\n", ticketIndex, index, num, field)
				}
			}
			validFields[index].slice = newValidFields
		}
	}

	sort.Slice(validFields, func (i, j int) bool {
		return len(validFields[i].slice) < len(validFields[j].slice)
	})

	fieldsTaken := make([]bool, len(fields))
	departureProduct := 1
	for _, validField := range validFields {
		for _, candidateField := range validField.slice {
			if fieldsTaken[candidateField] {
				continue
			}

			if candidateField < 6 {
				departureProduct *= yourTicket[validField.index]
			}
			// fmt.Printf("Field #%d -> rule #%d\n", validField.index, candidateField)
			fieldsTaken[candidateField] = true
		}
	}
	fmt.Println("Departure product", departureProduct)
}

