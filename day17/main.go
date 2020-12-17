package main

import (
	"fmt"
	"bufio"
	"os"
)

type Cell bool
type Row []Cell
type Plane []Row
type Space []Plane
type Hyper []Space

const NumCycles = 6
const Active = '#'
const Inactive = '.'

//// Printing ////
func printPlane(plane Plane) {
	for _, row := range plane {
		for _, cell := range row {
			if cell {
				fmt.Printf(string(Active))
			} else {
				fmt.Printf(string(Inactive))
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func printSpace(space Space) {
	for z, plane := range space {
		fmt.Printf("z=%d\n", z)
		printPlane(plane)
	}
}

//// Making ////
func makeHyper(wlen, depth, height, width int) (hyper Hyper) {
	hyper = make(Hyper, wlen)
	for w := range hyper {
		hyper[w] = makeSpace(depth, height, width)
	}
	return
}

func makeSpace(depth, height, width int) (space Space) {
	space = make(Space, depth)
	for z := range space {
		space[z] = makePlane(height, width)
	}
	return
}

func makePlane(height, width int) (plane Plane) {
	plane = make(Plane, height)
	for y := range plane {
		plane[y] = make(Row, width)
	}
	return
}

//// Expanding ////
func expandRow(row Row, num int) (newRow Row) {
	width := len(row)
	newRow = make(Row, width + num * 2)
	for x := range row {
		newRow[x + num] = row[x]
	}
	return
}

func expandPlane(plane Plane, num int) (newPlane Plane) {
	height := len(plane)
	width := len(plane[0])

	newPlane = make(Plane, height + num * 2)
	for y := range plane {
		newPlane[y + num] = expandRow(plane[y], num)
	}

	for i := 0; i < num; i++ {
		newPlane[i] = make(Row, width + num * 2)
		newPlane[num + height + i] = make(Row, width + num * 2)
	}
	return
}

func expandSpace(space Space, num int) (newSpace Space) {
	depth := len(space)
	height := len(space[0])
	width := len(space[0][0])

	newSpace = make(Space, depth + num * 2)
	for z := range space {
		newSpace[z + num] = expandPlane(space[z], num)
	}

	for i := 0; i < num; i++ {
		newSpace[i] = makePlane(height + num * 2, width + num * 2)
		newSpace[num + depth + i] = makePlane(height + num * 2, width + num * 2)
	}

	return
}

func expandHyper(hyper Hyper, num int) (newHyper Hyper) {
	wlen := len(hyper)
	depth := len(hyper[0])
	height := len(hyper[0][0])
	width := len(hyper[0][0][0])

	newHyper = make(Hyper, wlen + num * 2)
	for w := range hyper {
		newHyper[w + num] = expandSpace(hyper[w], num)
	}

	for i := 0; i < num; i++ {
		newHyper[i] = makeSpace(depth + num * 2, height + num * 2, width + num * 2)
		newHyper[num + wlen + i] = makeSpace(depth + num * 2, height + num * 2, width + num * 2)
	}

	return
}

//// Counting ////
func countTotalActive(space Space) (count int) {
	for _, plane := range space {
		for _, row := range plane {
			for _, cell := range row {
				if cell {
					count++
				}
			}
		}
	}
	return
}

func countTotalActiveHyper(hyper Hyper) (count int) {
	for _, space := range hyper {
		count += countTotalActive(space)
	}
	return
}

func countAdjacentActive(space Space, x, y, z int, dwIsZero bool) (count int) {
	for dz := -1; dz <= 1; dz++ {
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dwIsZero && dz == 0 && dy == 0 && dx == 0 {
					// Skip the center cell
					continue
				}
				if space[z + dz][y + dy][x + dx] {
					count++
				}
			}
		}
	}
	return
}

func countAdjacentActiveHyper(hyper Hyper, x, y, z, w int) (count int) {
	for dw := -1; dw <= 1; dw++ {
		count += countAdjacentActive(hyper[w + dw], x, y, z, dw == 0)
	}
	return
}

//// Step ////
func step(space Space) (newSpace Space) {
	depth := len(space)
	height := len(space[0])
	width := len(space[0][0])

	newSpace = makeSpace(depth, height, width)
	for z := 1; z < depth - 1; z++ {
		for y := 1; y < height - 1; y++ {
			for x := 1; x < width - 1; x++ {
				numAdjacentActive := countAdjacentActive(space, x, y, z, true)
				if space[z][y][x] {
					newSpace[z][y][x] = (numAdjacentActive == 2 || numAdjacentActive == 3)
				} else {
					newSpace[z][y][x] = numAdjacentActive == 3
				}
			}
		}
	}
	return
}

func stepHyper(hyper Hyper) (newHyper Hyper) {
	wlen := len(hyper)
	depth := len(hyper[0])
	height := len(hyper[0][0])
	width := len(hyper[0][0][0])

	newHyper = makeHyper(wlen, depth, height, width)
	for w := 1; w < wlen - 1; w++ {
		for z := 1; z < depth - 1; z++ {
			for y := 1; y < height - 1; y++ {
				for x := 1; x < width - 1; x++ {
					numAdjacentActive := countAdjacentActiveHyper(hyper, x, y, z, w)
					if hyper[w][z][y][x] {
						newHyper[w][z][y][x] = (numAdjacentActive == 2 || numAdjacentActive == 3)
					} else {
						newHyper[w][z][y][x] = numAdjacentActive == 3
					}
				}
			}
		}
	}
	return
}

func main() {
	var space Space
	var hyper Hyper

	var startPlane Plane
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var row Row
		for _, char := range line {
			var cell Cell
			if char == Active {
				cell = true
			}
			row = append(row, cell)
		}
		startPlane = append(startPlane, row)
	}
	space = Space{startPlane}
	printSpace(space)

	hyper = Hyper{space}

	space = expandSpace(space, NumCycles + 1)

	for i := 0; i < NumCycles; i++ {
		space = step(space)
	}
	printSpace(space)

	totalActive := countTotalActive(space)
	fmt.Println("part 1 total active:", totalActive)

	// Part 2
	hyper = expandHyper(hyper, NumCycles + 1)

	for i := 0; i < NumCycles; i++ {
		hyper = stepHyper(hyper)
	}

	totalActiveHyper := countTotalActiveHyper(hyper)
	fmt.Println("part 2 total active:", totalActiveHyper)
}

