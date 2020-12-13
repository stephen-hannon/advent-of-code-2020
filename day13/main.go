package main

import (
	"fmt"
	"strings"
	"strconv"
)

type Bus struct {
	offset, id int
}

func getWaitingTime(earliestDepart, id int) int {
	return (earliestDepart / id + 1) * id - earliestDepart
}

func getMaxBusById(buses []Bus) (maxBus Bus) {
	maxBus = buses[0]
	for _, bus := range buses {
		if bus.id > maxBus.id {
			maxBus = bus
		}
	}
	return
}

func isValidStartForOne(buses []Bus, start int) (int, bool) {
	for index, bus := range buses {
		if (start + bus.offset) % bus.id == 0 {
			fmt.Printf("%d is valid for %d!\n", start, bus.id)
			return index, true
		}
	}
	return -1, false
}

func product(buses []Bus) (product int) {
	product = 1
	for _, bus := range buses {
		product *= bus.id
	}

	return
}

func remove(slice []Bus, s int) []Bus {
	return append(slice[:s], slice[s+1:]...)
}

func main()  {
	var earliestDepart int
	fmt.Scanf("%d", &earliestDepart)
	var str string
	fmt.Scanf("%s", &str)

	minId := -1
	var buses []Bus
	ids := strings.Split(str, ",")
	for index, idStr := range ids {
		if idStr == "x" {
			continue
		}
		id, _ := strconv.Atoi(idStr)
		buses = append(buses, Bus{index, id})

		if minId == -1 || getWaitingTime(earliestDepart, id) < getWaitingTime(earliestDepart, minId) {
			minId = id
		}
	}

	minWaitingTime := getWaitingTime(earliestDepart, minId)
	fmt.Printf("%d * %d = %d\n", minId, minWaitingTime, minId * minWaitingTime)
	fmt.Println(buses)

	// Part 2
	upperBound := product(buses)
	fmt.Println("Upper bound", upperBound)
	factor := 1
	offset := 0
	for i := 0; i < upperBound / factor; i++ {
		start := offset + factor * i
		// fmt.Printf("Checking %d %e\n", start, float64(start))

		goodIndex, isGood := isValidStartForOne(buses, start)
		if !isGood {
			continue
		}

		goodBus := buses[goodIndex]
		factor *= goodBus.id
		offset = start
		i = 1
		buses = remove(buses, goodIndex)
		// fmt.Println(goodBus, buses)
		if len(buses) == 0 {
			break
		}
	}
}
