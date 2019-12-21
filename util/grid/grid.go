package grid

import (
	"fmt"
	"log"
	"math"
)

const (
	north = 1
	south = 2
	west  = 3
	east  = 4
)

// Position ...
type Position struct {
	X int
	Y int
}

// Move ...
func (pos Position) Move(direction int) Position {
	switch direction {
	case north:
		return Position{pos.X, pos.Y - 1}
	case south:
		return Position{pos.X, pos.Y + 1}
	case west:
		return Position{pos.X - 1, pos.Y}
	case east:
		return Position{pos.X + 1, pos.Y}
	default:
		panic("invalid direction")
	}
}

// String ...
func (pos Position) String() string {
	return fmt.Sprintf("(%d, %d)", pos.X, pos.Y)
}

// Around1 ...
func (pos Position) Around1() []Position {
	return []Position{
		Position{pos.X, pos.Y - 1},
		Position{pos.X - 1, pos.Y},
		Position{pos.X + 1, pos.Y},
		Position{pos.X, pos.Y + 1},
	}
}

// Around2 ...
func (pos Position) Around2() []Position {
	return []Position{
		Position{pos.X - 1, pos.Y - 1},
		Position{pos.X, pos.Y - 1},
		Position{pos.X + 1, pos.Y - 1},
		Position{pos.X - 1, pos.Y},
		Position{pos.X + 1, pos.Y},
		Position{pos.X - 1, pos.Y + 1},
		Position{pos.X, pos.Y + 1},
		Position{pos.X + 1, pos.Y + 1},
	}
}

// Limits ...
type Limits struct {
	XMin int
	XMax int
	YMin int
	YMax int
}

// CalcLimits ...
func CalcLimits(posMap map[Position]bool) Limits {
	xMin, xMax, yMin, yMax := math.MaxInt32, -math.MaxInt32, math.MaxInt32, -math.MaxInt32
	for p := range posMap {
		if p.X < xMin {
			xMin = p.X
		}
		if p.X > xMax {
			xMax = p.X
		}
		if p.Y < yMin {
			yMin = p.Y
		}
		if p.Y > yMax {
			yMax = p.Y
		}
	}
	return Limits{xMin, xMax, yMin, yMax}
}

// CalcLimitsInt ...
func CalcLimitsInt(posMap map[Position]int) Limits {
	xMin, xMax, yMin, yMax := math.MaxInt32, -math.MaxInt32, math.MaxInt32, -math.MaxInt32
	for p := range posMap {
		if p.X < xMin {
			xMin = p.X
		}
		if p.X > xMax {
			xMax = p.X
		}
		if p.Y < yMin {
			yMin = p.Y
		}
		if p.Y > yMax {
			yMax = p.Y
		}
	}
	return Limits{xMin, xMax, yMin, yMax}
}

// Print ...
func Print(posMap map[Position]int, origo rune, runeMap map[int]rune) {
	limits := CalcLimitsInt(posMap)
	for y := limits.YMin; y <= limits.YMax; y++ {
		for x := limits.XMin; x <= limits.XMax; x++ {
			if x == 0 && y == 0 {
				fmt.Printf("%c", origo)
				continue
			}
			val, ok := posMap[Position{x, y}]
			if !ok {
				fmt.Print(" ")
			} else {
				r, ok := runeMap[val]
				if !ok {
					log.Panicf("no rune for %d", val)
				}
				fmt.Printf("%c", r)
			}
		}
		fmt.Println()
	}
}
