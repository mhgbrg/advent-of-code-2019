package main

import (
	"fmt"
	"time"

	"github.com/mhgbrg/advent-of-code-2019/util/mathi"
)

const (
	numSteps = 1000
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

func (m *Moon) energy() int {
	potential := mathi.Abs(m.x) + mathi.Abs(m.y) + mathi.Abs(m.z)
	kinetic := mathi.Abs(m.vx) + mathi.Abs(m.vy) + mathi.Abs(m.vz)
	return potential * kinetic
}

func main() {
	start := time.Now()

	moons := []*Moon{
		//{-1, 0, 2, 0, 0, 0},
		//{2, -10, -7, 0, 0, 0},
		//{4, -8, 8, 0, 0, 0},
		//{3, 5, -1, 0, 0, 0},
		{-10, -13, 7, 0, 0, 0},
		{1, 2, 1, 0, 0, 0},
		{-15, -3, 13, 0, 0, 0},
		{3, 7, -4, 0, 0, 0},
	}

	for step := 0; step < numSteps; step++ {
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
				if m1.y < m2.y {
					m1.vy++
					m2.vy--
				} else if m1.y > m2.y {
					m1.vy--
					m2.vy++
				}
				if m1.z < m2.z {
					m1.vz++
					m2.vz--
				} else if m1.z > m2.z {
					m1.vz--
					m2.vz++
				}
			}
		}
		for _, moon := range moons {
			moon.x += moon.vx
			moon.y += moon.vy
			moon.z += moon.vz
		}
	}

	energy := 0
	for _, moon := range moons {
		energy += moon.energy()
	}
	fmt.Println(energy)

	fmt.Println(time.Since(start))
}
