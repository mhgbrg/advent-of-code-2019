package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/mhgbrg/advent-of-code-2019/util/mathi"
)

const (
	mapSize  = 40
	stationX = 31
	stationY = 20
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

	station := Pos{stationX, stationY}
	asteroids := getAsteroids(m, station)

	angles := make([]Angle, 0)
	for angle, positions := range asteroids {
		angles = append(angles, angle)
		sort.Slice(positions, func(i, j int) bool {
			d1 := distance(station, positions[i])
			d2 := distance(station, positions[j])
			return d1 < d2
		})
	}

	sort.Slice(angles, func(i, j int) bool {
		t1 := angles[i].atan2()
		t2 := angles[j].atan2()
		return t1 < t2
	})

	count := 0
	angleCounts := make(map[Angle]int)
	for _, angle := range angles {
		inLine := asteroids[angle]
		i := angleCounts[angle]
		if i >= len(inLine) {
			continue
		}
		count++
		angleCounts[angle]++
		fmt.Printf("%d (%d, %d) angle=%v\n", count, inLine[i].x, inLine[i].y, angle)
		if count == 200 {
			fmt.Println(100*inLine[i].x + inLine[i].y)
			break
		}
	}

	fmt.Println(time.Since(start))
}

func getAsteroids(m [][]bool, station Pos) map[Angle][]Pos {
	asteroids := make(map[Angle][]Pos)
	for x := 0; x < mapSize; x++ {
		for y := 0; y < mapSize; y++ {
			if x == station.x && y == station.y || !m[y][x] {
				continue
			}
			pos := Pos{x, y}
			angle := calcAngle(station, pos)
			asteroids[angle] = append(asteroids[angle], pos)
		}
	}
	return asteroids
}

// Pos ...
type Pos struct {
	x int
	y int
}

func distance(p1, p2 Pos) float64 {
	return math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2))
}

// Angle ...
type Angle struct {
	dx int
	dy int
}

func calcAngle(p1, p2 Pos) Angle {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
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

func (a Angle) atan2() float64 {
	t := math.Atan2(float64(a.dx), -float64(a.dy))
	if t < 0 {
		t = 2*math.Pi + t
	}
	return t
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
