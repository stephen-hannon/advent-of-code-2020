package main

import (
  "fmt"
  "bufio"
  "os"
  "regexp"
  "strings"
  "strconv"
)

type validator func(string) bool

func isValid1(fieldsMap map[string]string) (isValid bool) {
  requiredFields := [...]string{
    "byr",
    "iyr",
    "eyr",
    "hgt",
    "hcl",
    "ecl",
    "pid",
  }

  for _, reqField := range requiredFields {
    if _, ok := fieldsMap[reqField]; !ok {
      return false
    }
  }

  fmt.Println(fieldsMap)
  return true
}

func makeRangeValidator(min, max int) validator {
  return func (str string) bool {
    num, err := strconv.Atoi(str)
    return err == nil && num >= min && num <= max
  }
}

func isValidHgt (str string) bool {
  num, err := strconv.Atoi(str[:len(str)-2])
  unit := str[len(str)-2:]
  if err != nil {
    return false
  }
  if unit == "cm" {
    return num >= 150 && num <= 193
  } else if unit == "in" {
    return num >= 59 && num <= 76
  }
  return false
}

func isValidHcl(str string) bool {
  matched, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, str)
  return matched
}

func isValidEcl(str string) bool {
  switch str {
    case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
      return true
  }
  return false
}

func isValidPid(str string) bool {
  matched, _ := regexp.MatchString(`^\d{9}$`, str)
  return matched
}

func isValid2(fieldsMap map[string]string) (isValid bool) {
  requiredFields := map[string]validator{
    "byr": makeRangeValidator(1920, 2002),
    "iyr": makeRangeValidator(2010, 2020),
    "eyr": makeRangeValidator(2020, 2030),
    "hgt": isValidHgt,
    "hcl": isValidHcl,
    "ecl": isValidEcl,
    "pid": isValidPid,
  }

  for reqField, validatorFn := range requiredFields {
    if value, ok := fieldsMap[reqField]; !ok || !validatorFn(value) {
      fmt.Println("invalid", reqField, value)
      return false
    }
  }

  fmt.Println(fieldsMap)
  return true
}

func main() {
  var numValid1, numValid2 int
  fieldsMap := make(map[string]string)
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    line := scanner.Text()

    // Reset fieldsMap if this is a new passport
    if line == "" {
      if isValid1(fieldsMap) {
        numValid1++
        if isValid2(fieldsMap) {
          numValid2++
        }
      }

      fieldsMap = make(map[string]string)
      continue
    }

    fields := strings.Split(line, " ")
    for _, field := range fields {
      fieldSplit := strings.Split(field, ":")
      fieldName := fieldSplit[0]
      fieldValue := fieldSplit[1]
      fieldsMap[fieldName] = fieldValue
    }
  }

  // The loop ends on EOF, so check the last passport
  if isValid1(fieldsMap) {
    numValid1++
    if isValid2(fieldsMap) {
      numValid2++
    }
  }

  fmt.Println("Part 1 valid:", numValid1)
  fmt.Println("Part 2 valid:", numValid2)
}
