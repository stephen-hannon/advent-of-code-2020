package main

import (
  "fmt"
  "bufio"
  "os"
)

type questionMap map[rune]int

func countTrues(questions questionMap) (count int) {
  for _, value := range questions {
    if value != 0 {
      count++
    }
  }
  return
}

func countTruesN(questions questionMap, members int) (count int) {
  for _, value := range questions {
    if value == members {
      count++
    }
  }
  return
}

func main() {
  totalCount1 := 0
  totalCount2 := 0
  questions := make(questionMap)
  members := 0
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    line := scanner.Text()

    if line == "" {
      totalCount1 += countTrues(questions)
      totalCount2 += countTruesN(questions, members)
      questions = make(questionMap)
      members = 0
      continue
    }

    members++
    for _, letter := range line {
      questions[letter]++
    }
  }

  // The loop ends on EOF, so check the last one
  totalCount1 += countTrues(questions)
  totalCount2 += countTruesN(questions, members)

  fmt.Println("total count part 1:", totalCount1)
  fmt.Println("total count part 2:", totalCount2)
}
