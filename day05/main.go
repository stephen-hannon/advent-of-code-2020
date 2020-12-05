package main

import (
  "fmt"
  "strconv"
  "strings"
)

const MAX_ID = 848

func getSeatId(pass string) int64 {
  rowPass := pass[:7]
  colPass := pass[7:]

  rowBinary := strings.ReplaceAll(rowPass, "F", "0")
  rowBinary = strings.ReplaceAll(rowBinary, "B", "1")
  rowNum, _ := strconv.ParseInt(rowBinary, 2, 8)

  colBinary := strings.ReplaceAll(colPass, "L", "0")
  colBinary = strings.ReplaceAll(colBinary, "R", "1")
  colNum, _ := strconv.ParseInt(colBinary, 2, 8)

  return rowNum * 8 + colNum
}

func main() {
  var line string
  var maxId int64
  var ids [MAX_ID + 1]bool
  for {
    numScanned, _ := fmt.Scan(&line)
    if numScanned == 0 {
      break
    }
    seatId := getSeatId(line)
    ids[seatId] = true
    if seatId > maxId {
      maxId = seatId
    }
    fmt.Println(line, seatId)
  }
  fmt.Println("Max ID:", maxId)

  for id, value := range ids {
    if value {
      continue
    }
    fmt.Printf("ids[%d] = false\n", id)
  }
}
