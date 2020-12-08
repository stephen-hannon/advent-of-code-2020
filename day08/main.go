package main

import (
  "fmt"
  "bufio"
  "os"
)

type Instr string

const InstrAcc Instr = "acc"
const InstrJmp Instr = "jmp"
const InstrNop Instr = "nop"

type Loc struct {
  instr Instr
  num int
}

func invertLoc(loc *Loc) (changed bool) {
  if loc.instr == InstrJmp {
    loc.instr = InstrNop
    changed = true
  } else if loc.instr == InstrNop {
    loc.instr = InstrJmp
    changed = true
  }
  return
}

func exec(code []Loc) (acc int, halts bool) {
  var line int
  numLoc := len(code)
  visitedLines := make([]bool, numLoc)
  for {
    if line > numLoc {
      fmt.Println("Line too high", line)
      halts = false
      break
    } else if line == numLoc {
      halts = true
      break
    } else if visitedLines[line] {
      halts = false
      break
    }
    visitedLines[line] = true
    loc := code[line]

    if loc.instr == InstrAcc {
      acc += loc.num
      line++
    } else if loc.instr == InstrJmp {
      line += loc.num
    } else {
      line++
    }
  }

  return
}

func main() {
  var code []Loc
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    var instr string
    var num int
    fmt.Sscanf(scanner.Text(), "%s %d", &instr, &num)
    code = append(code, Loc{Instr(instr), num})
  }

  acc, _ := exec(code)

  fmt.Println("part 1 accumulator:", acc)

  for index, _ := range code {
    changed := invertLoc(&code[index])
    if !changed {
      continue
    }

    acc2, halts := exec(code)
    if halts {
      fmt.Printf("Flipping line %d halts! Accumulator: %d\n", index, acc2)
      invertLoc(&code[index])
      break
    }
    invertLoc(&code[index])
  }
}
