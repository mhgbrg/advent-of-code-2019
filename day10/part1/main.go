package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mhgbrg/advent-of-code-2019/util/mathi"
)

const (
	mapSize = 40
)

func main() {
	start := time.Now()

	scanner := bufio.NewScanner(os.Stdin)
	m := make([][]bool, mapSize)
	for i := 0; scanner.Scan(); i++ {
		m[i] = make([]bool, mapSize)
		row := scanner.Text()
		for j, pos := range strings.Split(row, "") {
			if pos == "#" {
				m[i][j] = true
			}
		}
	}

	maxCount := 0
	for y := 0; y < mapSize; y++ {
		for x := 0; x < mapSize; x++ {
			if !m[y][x] {
				continue
			}
			count := countDetected(m, x, y)
			if count > maxCount {
				fmt.Printf("(%d, %d) count=%d\n", x, y, count)
				maxCount = count
			}
		}
	}

	fmt.Println(time.Since(start))
}

// Angle ...
type Angle struct {
	dx int
	dy int
}

func calcAngle(x1, y1, x2, y2 int) Angle {
	dx := x2 - x1
	dy := y2 - y1
	if dx == 0 {
		return Angle{
			dx: 0,
			dy: dy / mathi.Abs(dy),
		}
	} else if dy == 0 {
		return Angle{
			dx: dx / mathi.Abs(dx),
			dy: 0,
		}
	}
	gcd := gcd(mathi.Abs(dx), mathi.Abs(dy))
	return Angle{
		dx: dx / gcd,
		dy: dy / gcd,
	}
}

func gcd(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

func countDetected(m [][]bool, x, y int) int {
	count := 0
	obstructed := make(map[Angle]bool)
	for y2 := 0; y2 < mapSize; y2++ {
		for x2 := 0; x2 < mapSize; x2++ {
			if x2 == x && y2 == y || !m[y2][x2] {
				continue
			}
			angle := calcAngle(x, y, x2, y2)
			if !obstructed[angle] {
				count++
				obstructed[angle] = true
			}
		}
	}
	return count
}
