package main

import (
  "fmt"
)

const WindowSize = 25
const InputSize = 1000

func findMinAndMax(slice []int) (min, max int) {
  min = slice[0]
  max = slice[0]
  for _, num := range slice {
    if num > max {
      max = num
    }
    if num < min {
      min = num
    }
  }
  return
}

func has2Sum(slice []int, target int) bool {
  for index, num1 := range slice {
    for _, num2 := range slice[index:] {
      if num1 + num2 == target {
        return true
      }
    }
  }

  return false
}

func findSliceSum(slice []int, target int) (tail, head int) {
  tailIndex := 0
  sum := 0
  for headIndex, head := range slice {
    sum += head
    if sum == target {
      return tailIndex, headIndex
    }

    for sum > target {
      sum -= slice[tailIndex]
      tailIndex++
    }
    if sum == target {
      return tailIndex, headIndex
    }
  }
  return -1, -1
}

func main() {
  window := make([]int, WindowSize)
  input := make([]int, InputSize)
  // Build preamble
  for i := 0; i < WindowSize; i++ {
    fmt.Scanf("%d", &window[i])
    input[i] = window[i]
  }
  fmt.Println(window)

  offset := 0
  readIndex := WindowSize
  part1Answer := 0
  for {
    var newNum int
    numScanned, _ := fmt.Scanf("%d", &newNum)
    if numScanned == 0 {
      break
    }

    input[readIndex] = newNum
    readIndex++

    if part1Answer != 0 {
      continue
    }
    if !has2Sum(window, newNum) {
      fmt.Println("2-sum not found for", newNum)
      part1Answer = newNum
      continue
    }

    window[offset] = newNum
    offset++
    offset %= WindowSize
  }

  tail, head := findSliceSum(input, part1Answer)
  fmt.Println("part 2", input[tail:head + 1])
  min, max := findMinAndMax(input[tail:head + 1])
  fmt.Printf("%d + %d = %d\n", min, max, min + max)
}
