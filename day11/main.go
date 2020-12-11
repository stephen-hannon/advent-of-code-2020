package main

import (
	"fmt"
	"bufio"
	"os"
)

const Floor = '.'
const Empty = 'L'
const Occupied = '#'

type Board [][]rune
type Point struct {
	x, y int
}

func getSafe(board Board, x, y int) rune {
	height := len(board)
	width := len(board[0])
	if x < 0 || x >= width || y < 0 || y >= height {
		return Empty
	}
	return board[y][x]
}

func areBoardsEqual(board1, board2 Board) bool {
	for y, row := range board1 {
		for x, _ := range row {
			if board1[y][x] != board2[y][x] {
				return false
			}
		}
	}
	return true
}

func getDirectedSeat(board Board, x, y int, dir Point, countVisible bool) (cell rune) {
	for {
		x += dir.x
		y += dir.y
		cell = getSafe(board, x, y)
		if !countVisible || cell != Floor {
			break
		}
	}
	return
}

func countOccupiedAdjacent(board Board, x, y int, countVisible bool) (count int) {
	adjacentVectors := [...]Point {
		Point{-1, -1},
		Point{ 0, -1},
		Point{ 1, -1},
		Point{-1,  0},
		Point{ 1,  0},
		Point{-1,  1},
		Point{ 0,  1},
		Point{ 1,  1},
	}

	for _, vector := range adjacentVectors {
		visibleSeat := getDirectedSeat(board, x, y, vector, countVisible)
		if visibleSeat == Occupied {
			count++
		}
	}

	return
}

func countOccupied(board Board) (count int) {
	for _, row := range board {
		for _, cell := range row {
			if cell == Occupied {
				count++
			}
		}
	}
	return
}

func step(board Board, countVisible bool) (newBoard Board) {
	var maxOccupied int
	if countVisible {
		maxOccupied = 5
	} else {
		maxOccupied = 4
	}

	newBoard = make(Board, len(board))
	for y, row := range board {
		newBoard[y] = make([]rune, len(row))

		for x, cell := range row {
			newCell := Empty
			numOccupiedAdjacent := countOccupiedAdjacent(board, x, y, countVisible)
			// fmt.Println(string(cell), x, y, numOccupiedAdjacent)
			if cell == Floor {
				newCell = Floor
			} else if cell == Empty && numOccupiedAdjacent == 0 {
				newCell = Occupied
			} else if cell == Occupied && numOccupiedAdjacent < maxOccupied {
				newCell = Occupied
			}
			newBoard[y][x] = newCell
		}
	}
	return
}

func printBoard(board Board) {
	for _, row := range board {
		fmt.Println(string(row))
	}
	fmt.Println("")
}

func main() {
	var origBoard Board
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		origBoard = append(origBoard, []rune(scanner.Text()))
	}
	printBoard(origBoard)

	board := origBoard
	newBoard := step(board, false)
	for !areBoardsEqual(board, newBoard) {
		// printBoard(newBoard)
		board = newBoard
		newBoard = step(board, false)
	}
	printBoard(newBoard)
	fmt.Println("Occupied", countOccupied(board))

	// Part 2
	board = origBoard
	newBoard = step(board, true)
	for !areBoardsEqual(board, newBoard) {
		// printBoard(newBoard)
		board = newBoard
		newBoard = step(board, true)
	}
	printBoard(newBoard)
	fmt.Println("Occupied", countOccupied(board))
}