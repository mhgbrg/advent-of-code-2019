package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	graph := parseInput()
	//fmt.Println(graph)
	count := countOrbitalTransfers(graph)
	fmt.Println(count)

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
	graph := make(map[string][]string)
	for center, orbiters := range tree {
		graph[center] = append(graph[center], orbiters...)
		for _, orbiter := range orbiters {
			graph[orbiter] = append(graph[orbiter], center)
		}
	}
	return graph
}

// BFSQueueElement ...
type BFSQueueElement struct {
	planet   string
	distance int
}

func countOrbitalTransfers(graph map[string][]string) int {
	start, end := "YOU", "SAN"
	queue := make([]BFSQueueElement, 0)
	for _, planet := range graph[start] {
		queue = append(queue, BFSQueueElement{planet, 1})
	}
	visited := make(map[string]bool)
	visited[start] = true
	for len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]
		if elem.planet == end {
			return elem.distance - 2
		}
		visited[elem.planet] = true
		for _, planet := range graph[elem.planet] {
			if !visited[planet] {
				queue = append(queue, BFSQueueElement{planet, elem.distance + 1})
			}
		}
	}
	log.Panic("couldn't find SAN :(")
	return 0
}
