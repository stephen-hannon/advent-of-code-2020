package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strings"
)

type Rule struct {
	deps [][]string
	str string
}

func parseRule(line string) (key string, rule Rule) {
	lineSplit := strings.Split(line, ": ")
	key = lineSplit[0]
	ruleR := lineSplit[1]
	if ruleR[0] == '"' {
		rule.str = ruleR[1:2]
		return
	}

	subRules := strings.Split(ruleR, " | ")
	for _, subRule := range subRules {
		rule.deps = append(rule.deps, strings.Split(subRule, " "))
	}

	return
}

func areDepsBuilt(rules map[string]Rule, key string) bool {
	for _, dep := range rules[key].deps {
		for _, subDep := range dep {
			if rules[subDep].str == "" {
				return false
			}
		}
	}
	return true
}

func buildRegexStr(rules map[string]Rule, key string) string {
	var str strings.Builder
	if len(rules[key].deps) > 1 {
		str.WriteRune('(')
	}
	for index, dep := range rules[key].deps {
		if index != 0 {
			str.WriteRune('|')
		}
		for _, subDep := range dep {
			str.WriteString(rules[subDep].str)
		}
	}
	if len(rules[key].deps) > 1 {
		str.WriteRune(')')
	}
	return str.String()
}

func buildRegex(rules map[string]Rule, start string) string {
	stack := []string{start}
	
	for len(stack) != 0 {
		last := len(stack) - 1
		ruleKey := stack[last]

		if rules[ruleKey].str != "" {
			stack = stack[:last]
			continue
		}

		if areDepsBuilt(rules, ruleKey) {
			rules[ruleKey] = Rule{rules[ruleKey].deps, buildRegexStr(rules, ruleKey)}
			stack = stack[:last]
		} else {
			for _, dep := range rules[ruleKey].deps {
				stack = append(stack, dep...)
			}
		}
	}

	return fmt.Sprintf("^%s$", rules[start].str)
}

func main() {
	rules1 := make(map[string]Rule)
	rules2 := make(map[string]Rule)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		key1, rule1 := parseRule(line)
		rules1[key1] = rule1

		// Really hacky, but I didn't want to build a full parser
		if line == "8: 42" {
			line = "8: 42 | 42 42 | 42 42 42 | 42 42 42 42 | 42 42 42 42 42"
		} else if line == "11: 42 31" {
			line = "11: 42 31 | 42 42 31 31 | 42 42 42 31 31 31 | 42 42 42 42 31 31 31 31 | 42 42 42 42 42 31 31 31 31 31"
		}
		key2, rule2 := parseRule(line)
		rules2[key2] = rule2
	}

	fmt.Println(rules2)
	regex1 := buildRegex(rules1, "0")
	fmt.Println(regex1)

	regex2 := buildRegex(rules2, "0")
	fmt.Println(regex2)

	regex1Parsed := regexp.MustCompile(regex1)
	count1 := 0
	regex2Parsed := regexp.MustCompile(regex2)
	count2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		if regex1Parsed.MatchString(line) {
			fmt.Println("1 matches:", line)
			count1++
		}
		if regex2Parsed.MatchString(line) {
			fmt.Println("2 matches:", line)
			count2++
		}
	}
	fmt.Println("Part 1 count", count1)
	fmt.Println("Part 2 count", count2)
}
