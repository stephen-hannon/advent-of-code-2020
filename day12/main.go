package main

import (
	"fmt"
)

type Pair struct {
	x, y int
}

func applyMove(point *Pair, direction Pair, distance int) {
	point.x += direction.x * distance
	point.y += direction.y * distance
	fmt.Println("moving to", point)
}

func applyTurn2(direction Pair, deltaDirection int) (newDir Pair) {
	fmt.Println("turning to", newDir)
	deltaDirection %= 4
	if deltaDirection == 0 {
		return direction
	}
	if abs(deltaDirection) == 2 {
		newDir.x = -direction.x
		newDir.y = -direction.y
		fmt.Println("turning to", newDir)
		return
	}
	if abs(deltaDirection) == 3 {
		deltaDirection = -sign(deltaDirection)
	}
	newDir.x = direction.y * deltaDirection
	newDir.y = direction.x * deltaDirection * -1
	fmt.Println("turning to", newDir)
	return
}

func sign(n int) int {
	if n < 0 {
		return -1
	}
	return 1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	North := Pair{ 0,  1}
	East  := Pair{ 1,  0}
	South := Pair{ 0, -1}
	West  := Pair{-1,  0}

	direction := East
	point := Pair{0, 0}

	ship := Pair{0, 0}
	waypoint := Pair{10, 1}
	for {
		var action byte
		var num int
		numScanned, _ := fmt.Scanf("%c%d", &action, &num)
		if numScanned == 0 {
			break
		}

		switch action {
		case 'N':
			applyMove(&point, North, num)
			applyMove(&waypoint, North, num)
		case 'E':
			applyMove(&point, East, num)
			applyMove(&waypoint, East, num)
		case 'S':
			applyMove(&point, South, num)
			applyMove(&waypoint, South, num)
		case 'W':
			applyMove(&point, West, num)
			applyMove(&waypoint, West, num)
		case 'L':
			direction = applyTurn2(direction, num / -90)
			waypoint = applyTurn2(waypoint, num / -90)
		case 'R':
			direction = applyTurn2(direction, num / 90)
			waypoint = applyTurn2(waypoint, num / 90)
		case 'F':
			applyMove(&point, direction, num)
			applyMove(&ship, waypoint, num)
		}
	}

	fmt.Println("Part 1")
	fmt.Println("Ending point", point)
	fmt.Println("manhattan distance", abs(point.x) + abs(point.y))

	fmt.Println("Part 2")
	fmt.Println("Ending ship", ship)
	fmt.Println("Ending waypoint", waypoint)
	fmt.Println("manhattan distance", abs(ship.x) + abs(ship.y))
}
