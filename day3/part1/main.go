package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/mhgbrg/advent-of-code-2019/util/conv"
	"github.com/mhgbrg/advent-of-code-2019/util/mathi"
)

func main() {
	start := time.Now()

	fmt.Println("parse input")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	wire1Str := scanner.Text()
	scanner.Scan()
	wire2Str := scanner.Text()

	wire1 := strings.Split(wire1Str, ",")
	wire2 := strings.Split(wire2Str, ",")

	fmt.Println("init grid")
	gridSize := 40000
	grid := make([][]bool, gridSize)
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]bool, gridSize)
	}
	origo := gridSize / 2

	fmt.Println("wire 1")
	x, y := origo, origo
	for _, path := range wire1 {
		dir := path[0]
		steps := conv.Atoi(path[1:])
		dx, dy := 0, 0
		switch dir {
		case 'U':
			dy = -1
		case 'D':
			dy = 1
		case 'L':
			dx = -1
		case 'R':
			dx = 1
		}
		for i := 0; i < steps; i++ {
			x += dx
			y += dy
			grid[x][y] = true
		}
	}

	fmt.Println("wire 2")
	x, y = origo, origo
	minDist := math.MaxInt32
	for _, path := range wire2 {
		dir := path[0]
		steps := conv.Atoi(path[1:])
		dx, dy := 0, 0
		switch dir {
		case 'U':
			dy = -1
		case 'D':
			dy = 1
		case 'L':
			dx = -1
		case 'R':
			dx = 1
		}
		for i := 0; i < steps; i++ {
			x += dx
			y += dy
			if grid[x][y] {
				dist := mathi.Abs(origo-x) + mathi.Abs(origo-y)
				fmt.Printf("(%d, %d) dist=%d\n", x, y, dist)
				minDist = mathi.Min(minDist, dist)
			}
		}
	}

	fmt.Println(minDist)

	fmt.Println(time.Since(start))
}
