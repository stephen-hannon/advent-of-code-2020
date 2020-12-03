package main

import "fmt"

const TREE = '#'
const NUM_SLOPES = 5

func main() {
  fmt.Println("Hello World")

  slopes := [NUM_SLOPES][2]int{
    {1, 1},
    {3, 1},
    {5, 1},
    {7, 1},
    {1, 2},
  }
  var xArr, numTreesArr [NUM_SLOPES]int

  y := 0
  var line string
  for {
    numScanned, _ := fmt.Scanln(&line)
    if numScanned == 0 {
      break
    }

    for slopeIndex, slope := range slopes {
      dx := slope[0]
      dy := slope[1]

      if y % dy != 0 {
        continue
      }

      xArr[slopeIndex] %= len(line)
      if line[xArr[slopeIndex]] == TREE {
        numTreesArr[slopeIndex]++
      }
      xArr[slopeIndex] += dx
    }

    y++
  }

  fmt.Println("Trees:", numTreesArr)
  product := 1
  for index, numTrees := range numTreesArr {
    fmt.Printf("Trees in slope %d: %d\n", index + 1, numTrees)
    product *= numTrees
  }
  fmt.Println("Product of trees:", product)
}
