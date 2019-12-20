package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/mhgbrg/advent-of-code-2019/util/conv"
)

const (
	availableOre = 1000000000000
)

func main() {
	start := time.Now()

	scanner := bufio.NewScanner(os.Stdin)
	recipies := make(map[string]Recipe)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " => ")
		inputsUnparsed := strings.Split(parts[0], ", ")
		inputs := make([]Ingredient, len(inputsUnparsed))
		for i, str := range inputsUnparsed {
			inputs[i] = parseIngredient(str)
		}
		output := parseIngredient(parts[1])
		recipies[output.chemical] = Recipe{
			inputs: inputs,
			output: output,
		}
	}

	ingredients := makeChemical(recipies, Ingredient{"FUEL", 1})
	fmt.Println("per fuel")
	fmt.Println(ingredients)

	fmt.Println("desired")
	amounts := make(map[string]int)
	multiplier := availableOre / ingredients["ORE"]
	for chemical, quantity := range ingredients {
		desired := quantity * multiplier
		amounts[chemical] = int(desired)
		fmt.Printf("%s %f\n", chemical, desired)
	}

	fmt.Println("achievable")
	min := math.MaxInt32
	for _, ingredient := range recipies["FUEL"].inputs {
		achievable := amounts[ingredient.chemical] - amounts[ingredient.chemical]%int(recipies[ingredient.chemical].output.quantity)
		fuel := achievable / int(ingredient.quantity)
		if fuel < min {
			min = fuel
		}
		fmt.Printf("%s %d\n", ingredient.chemical, fuel)
	}
	fmt.Println(min)

	fmt.Println(time.Since(start))
}

// Ingredient ...
type Ingredient struct {
	chemical string
	quantity float64
}

func parseIngredient(str string) Ingredient {
	parts := strings.Split(str, " ")
	return Ingredient{
		chemical: parts[1],
		quantity: conv.ParseFloat(parts[0]),
	}
}

// Recipe ...
type Recipe struct {
	inputs []Ingredient
	output Ingredient
}

func makeChemical(recipies map[string]Recipe, ingredient Ingredient) map[string]float64 {
	ingredients := make(map[string]float64)
	if ingredient.chemical == "ORE" {
		return ingredients
	}
	recipe := recipies[ingredient.chemical]
	multiplier := float64(ingredient.quantity) / float64(recipe.output.quantity)
	for _, input := range recipe.inputs {
		amountNeeded := float64(input.quantity) * multiplier
		ingredients[input.chemical] += amountNeeded
		subIngredients := makeChemical(recipies, Ingredient{input.chemical, amountNeeded})
		for chemical, quantity := range subIngredients {
			ingredients[chemical] += quantity
		}
	}
	return ingredients
}
