package main

import (
	"fmt"
	"time"
)

// Moon ...
type Moon struct {
	x  int
	y  int
	z  int
	vx int
	vy int
	vz int
}

// SimpleMoon ...
type SimpleMoon struct {
	x  int
	vx int
}

func main() {
	start := time.Now()

	moons := []*Moon{
		//{-1, 0, 2, 0, 0, 0},
		//{2, -10, -7, 0, 0, 0},
		//{4, -8, 8, 0, 0, 0},
		//{3, 5, -1, 0, 0, 0},

		//{-8, -10, 0, 0, 0, 0},
		//{5, 5, 10, 0, 0, 0},
		//{2, -7, 3, 0, 0, 0},
		//{9, -8, -3, 0, 0, 0},

		{-10, -13, 7, 0, 0, 0},
		{1, 2, 1, 0, 0, 0},
		{-15, -3, 13, 0, 0, 0},
		{3, 7, -4, 0, 0, 0},
	}

	x := findLoop([]*SimpleMoon{
		{moons[0].x, 0},
		{moons[1].x, 0},
		{moons[2].x, 0},
		{moons[3].x, 0},
	})
	y := findLoop([]*SimpleMoon{
		{moons[0].y, 0},
		{moons[1].y, 0},
		{moons[2].y, 0},
		{moons[3].y, 0},
	})
	z := findLoop([]*SimpleMoon{
		{moons[0].z, 0},
		{moons[1].z, 0},
		{moons[2].z, 0},
		{moons[3].z, 0},
	})
	fmt.Printf("%d %d %d\n", x, y, z)

	fmt.Println(time.Since(start))
}

// State ...
type State struct {
	x1  int
	x2  int
	x3  int
	x4  int
	vx1 int
	vx2 int
	vx3 int
	vx4 int
}

func findLoop(moons []*SimpleMoon) int {
	states := make(map[State]int)
	for step := 0; ; step++ {
		state := State{
			moons[0].x,
			moons[1].x,
			moons[2].x,
			moons[3].x,
			moons[0].vx,
			moons[1].vx,
			moons[2].vx,
			moons[3].vx,
		}
		if last, ok := states[state]; ok {
			fmt.Printf("loop from %d to %d => %d\n", last, step, step-last)
			return step
		}
		states[state] = step

		for i := 0; i < len(moons); i++ {
			m1 := moons[i]
			for j := i + 1; j < len(moons); j++ {
				m2 := moons[j]
				if m1.x < m2.x {
					m1.vx++
					m2.vx--
				} else if m1.x > m2.x {
					m1.vx--
					m2.vx++
				}
			}
		}
		for _, moon := range moons {
			moon.x += moon.vx
		}
	}
}
