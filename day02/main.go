package main

import (
  "fmt"
  "strings"
)

func isValid1(minCount, maxCount int, letter, password string) (isValid bool) {
  count := strings.Count(password, letter)
  return count >= minCount && count <= maxCount
}

func isValid2(index1, index2 int, letter byte, password string) (isValid bool) {
  char1Matches := password[index1 - 1] == letter
  char2Matches := password[index2 - 1] == letter
  return char1Matches != char2Matches
}

func main() {
  fmt.Println("Hello World")

  numValid1, numValid2 := 0, 0
  var num1, num2 int
  var letter byte
  var password string
  for {
    numScanned, _ := fmt.Scanf("%d-%d %c: %s", &num1, &num2, &letter, &password)
    if numScanned == 0 {
      break
    }

    letterStr := string([]byte{letter})

    if isValid1(num1, num2, letterStr, password) {
      // fmt.Println("1 valid:", password)
      numValid1++
    }

    if isValid2(num1, num2, letter, password) {
      // fmt.Println("2 valid:", password)
      numValid2++
    }
  }

  fmt.Println("Part 1 valid passwords:", numValid1)
  fmt.Println("Part 1 valid passwords:", numValid2)
}
