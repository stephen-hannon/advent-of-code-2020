package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func parseLine(line string) (ingredients, allergens []string) {
	lineSplit := strings.Split(line, " (contains ")
	ingredients = strings.Split(lineSplit[0], " ")
	allergensStr := lineSplit[1][:len(lineSplit[1]) - 1]
	allergens = strings.Split(allergensStr, ", ")
	return
}

func intersect(indices []int, strings [][]string) (intersection map[string]int) {
	intersection = make(map[string]int)
	for _, str := range strings[indices[0]] {
		intersection[str] = 1
	}

	for _, index := range indices[1:] {
		for _, str := range strings[index] {
			intersection[str]++
		}
	}
	return
}

func main() {
	var foods [][]string
	allergens := make(map[string][]int)
	ingredients := make(map[string][]string)
	ingredientCounts := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		foodIngredients, foodAllergens := parseLine(line)

		foods = append(foods, foodIngredients)
		for _, allergen := range foodAllergens {
			allergens[allergen] = append(allergens[allergen], lineNum)
		}
		for _, ingredient := range foodIngredients {
			ingredients[ingredient] = []string{}
			ingredientCounts[ingredient]++
		}
		lineNum++
	}

	for allergen, allergenFoods := range allergens {
		intersection := intersect(allergenFoods, foods)
		for ingredient, count := range intersection {
			if count != len(allergenFoods) {
				continue
			}
			ingredients[ingredient] = append(ingredients[ingredient], allergen)
		}
	}

	noAllergenCount := 0
	for ingredient, possibleAllergens := range ingredients {
		if len(possibleAllergens) == 0 {
			noAllergenCount += ingredientCounts[ingredient]
		} else {
			fmt.Println(ingredient, "\t", possibleAllergens)
		}
	}
	fmt.Println("no allergens:", noAllergenCount)
}

