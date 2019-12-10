package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	tree := parseInput()
	//fmt.Println(tree)
	orbits := countOrbits(tree, "COM", 0)
	fmt.Println(orbits)

	fmt.Println(time.Since(start))
}

func parseInput() map[string][]string {
	scanner := bufio.NewScanner(os.Stdin)
	tree := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ")")
		center := parts[0]
		orbiter := parts[1]
		tree[center] = append(tree[center], orbiter)
	}
	return tree
}

func countOrbits(tree map[string][]string, node string, height int) int {
	count := 0
	for _, planet := range tree[node] {
		// direct orbits
		count += height + 1
		// indirect orbits
		count += countOrbits(tree, planet, height+1)
	}
	return count
}
