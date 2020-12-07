package main

import (
  "fmt"
  "bufio"
  "os"
  "regexp"
  "strconv"
  "strings"
)

type adjList map[string][]string
type nodeWeighted struct {
  weight int
  label string
}
type adjListWeighted map[string][]nodeWeighted

func countDescendents(dag adjList, source string) (count int) {
  visited := make(map[string]bool)
  queue := []string{source}

  for len(queue) > 0 {
    node := queue[0]
    queue = queue[1:]
    if visited[node] {
      continue
    }

    count++
    queue = append(queue, dag[node]...)
    visited[node] = true
  }

  return count - 1 // subtract 1 to ignore the source
}

func countChildren(dag adjListWeighted, source string) (count int) {
  children, ok := dag[source]
  count = 1 // self
  if !ok {
    return
  }
  for _, child := range children {
    count += child.weight * countChildren(dag, child.label)
  }
  return
}

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  re := regexp.MustCompile(`(\d+) (\w+ \w+) bags?`)
  graphReversed := make(adjList)
  graph := make(adjListWeighted)

  for scanner.Scan() {
    line := scanner.Text()
    lineSplit := strings.Split(line, " bags contain ")
    parentBag := lineSplit[0]
    childBags := lineSplit[1]
    childBags = strings.TrimSuffix(childBags, ".")

    if childBags == "no other bags" {
      continue
    }

    childBagsSplit := strings.Split(childBags, ", ")
    for _, child := range childBagsSplit {
      matches := re.FindStringSubmatch(child)
      numChildBag, _ := strconv.Atoi(matches[1])
      childBag := matches[2]

      // fmt.Println(parentBag, "->", childBag)
      graphReversed[childBag] = append(graphReversed[childBag], parentBag)
      graph[parentBag] = append(graph[parentBag], nodeWeighted{numChildBag, childBag})
    }
  }

  fmt.Println(graphReversed)
  fmt.Println(graph)

  numDescendents := countDescendents(graphReversed, "shiny gold")
  fmt.Println("Containers for shiny gold:", numDescendents)

  // subtract 1 to ignore the source
  numChildren := countChildren(graph, "shiny gold") - 1
  fmt.Println("Bags in shiny gold:", numChildren)
}
