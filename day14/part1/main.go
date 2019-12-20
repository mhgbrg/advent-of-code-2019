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

	store := make(map[string]int)
	cost := makeChemical(recipies, store, Ingredient{"FUEL", 1})
	fmt.Println(cost)

	fmt.Println(time.Since(start))
}

// Ingredient ...
type Ingredient struct {
	chemical string
	quantity int
}

func parseIngredient(str string) Ingredient {
	parts := strings.Split(str, " ")
	return Ingredient{
		chemical: parts[1],
		quantity: conv.Atoi(parts[0]),
	}
}

// Recipe ...
type Recipe struct {
	inputs []Ingredient
	output Ingredient
}

func makeChemical(recipies map[string]Recipe, store map[string]int, ingredient Ingredient) int {
	// If ORE, we're done.
	if ingredient.chemical == "ORE" {
		return ingredient.quantity
	}

	// Use existing supply if possible.
	existingSupply := store[ingredient.chemical]
	if existingSupply >= ingredient.quantity {
		store[ingredient.chemical] -= ingredient.quantity
		return 0
	}
	needed := ingredient.quantity - existingSupply
	store[ingredient.chemical] = 0

	// Make needed quantity.
	recipe := recipies[ingredient.chemical]
	repetitions := int(math.Ceil(float64(needed) / float64(recipe.output.quantity)))
	cost := 0
	for i := 0; i < repetitions; i++ {
		for _, input := range recipe.inputs {
			cost += makeChemical(recipies, store, input)
		}
	}

	// Save leftovers for later.
	leftOvers := repetitions*recipe.output.quantity - needed
	store[ingredient.chemical] = leftOvers
	return cost
}
