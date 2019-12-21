package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/mhgbrg/advent-of-code-2019/util/grid"
	"github.com/mhgbrg/advent-of-code-2019/util/slices"
)

const (
	positionMode  = 0
	immediateMode = 1
	relativeMode  = 2
	memorySize    = 10000
	debug         = false
	north         = 1
	south         = 2
	west          = 3
	east          = 4
	wall          = 0
	nothing       = 1
	oxygenSystem  = 2
)

func main() {
	start := time.Now()

	program := readProgram()
	machine := initMachine(program)
	discovered := map[grid.Position]int{
		grid.Position{X: 0, Y: 0}: 1,
	}
	distance, foundOxygen := discoverArea(machine, grid.Position{X: 0, Y: 0}, discovered)
	grid.Print(discovered, 'o', map[int]rune{
		wall:         '#',
		nothing:      '.',
		oxygenSystem: 'x',
	})
	fmt.Printf("%t %d\n", foundOxygen, distance)

	fmt.Println(time.Since(start))
}

func discoverArea(machine *Machine, pos grid.Position, discovered map[grid.Position]int) (int, bool) {
	for direction := 1; direction <= 4; direction++ {
		destination := pos.Move(direction)
		if _, ok := discovered[destination]; ok {
			continue
		}
		result, done := machine.run(direction)
		if done {
			panic("done on discover")
		}
		discovered[destination] = result
		if result == oxygenSystem {
			fmt.Printf("found oxygen at %s\n", destination)
			return 1, true
		} else if result == wall {
			continue
		}
		distance, foundOxygen := discoverArea(machine, destination, discovered)
		if foundOxygen {
			return distance + 1, true
		}
		var oppositeDirection int
		if direction == north || direction == west {
			oppositeDirection = direction + 1
		} else {
			oppositeDirection = direction - 1
		}
		result, done = machine.run(oppositeDirection)
		if done {
			panic("done on return")
		} else if result != discovered[pos] {
			log.Panicf("result not the same on return. got=%d, expected=%d\n", result, discovered[pos])
		}
	}
	return 0, false
}

func readProgram() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	programStr := scanner.Text()
	return slices.Atoi(strings.Split(programStr, ","))
}

// Machine ...
type Machine struct {
	memory       []int
	ptr          int
	relativeBase int
}

func initMachine(program []int) *Machine {
	memory := make([]int, memorySize)
	copy(memory, program)
	return &Machine{
		memory:       memory,
		ptr:          0,
		relativeBase: 0,
	}
}

func (m *Machine) run(input int) (int, bool) {
	for ; m.ptr < len(m.memory); m.ptr++ {
		opCode := m.memory[m.ptr] % 100
		modes := m.memory[m.ptr] / 100
		if debug {
			fmt.Printf(
				"m.ptr=%d m.relativeBase=%d opCode=%d params=%d %d %d\n",
				m.ptr,
				m.relativeBase,
				m.memory[m.ptr],
				m.memory[m.ptr+1],
				m.memory[m.ptr+2],
				m.memory[m.ptr+3],
			)
		}
		if opCode == 1 {
			x := getVal(m.memory, modes, m.relativeBase, m.ptr, 1)
			y := getVal(m.memory, modes, m.relativeBase, m.ptr, 2)
			storeAt := getMemLoc(m.memory, modes, m.relativeBase, m.ptr, 3)
			m.memory[storeAt] = x + y
			if debug {
				fmt.Printf("m.memory[%d] = %d + %d = %d\n", storeAt, x, y, m.memory[storeAt])
			}
			m.ptr += 3
		} else if opCode == 2 {
			x := getVal(m.memory, modes, m.relativeBase, m.ptr, 1)
			y := getVal(m.memory, modes, m.relativeBase, m.ptr, 2)
			storeAt := getMemLoc(m.memory, modes, m.relativeBase, m.ptr, 3)
			m.memory[storeAt] = x * y
			if debug {
				fmt.Printf("%d * %d = %d => m.memory[%d]\n", x, y, m.memory[storeAt], storeAt)
			}
			m.ptr += 3
		} else if opCode == 3 {
			storeAt := getMemLoc(m.memory, modes, m.relativeBase, m.ptr, 1)
			m.memory[storeAt] = input
			if debug {
				fmt.Printf("%d => m.memory[%d]\n", input, storeAt)
			}
			m.ptr++
		} else if opCode == 4 {
			x := getVal(m.memory, modes, m.relativeBase, m.ptr, 1)
			if debug {
				fmt.Printf("OUTPUT: %d\n", x)
			}
			m.ptr += 2
			return x, false
		} else if opCode == 5 {
			test := getVal(m.memory, modes, m.relativeBase, m.ptr, 1)
			if test != 0 {
				jumpTo := getVal(m.memory, modes, m.relativeBase, m.ptr, 2)
				m.ptr = jumpTo - 1
				if debug {
					fmt.Printf("test != 0, m.ptr => %d\n", jumpTo)
				}
			} else {
				if debug {
					fmt.Println("test == 0")
				}
				m.ptr += 2
			}
		} else if opCode == 6 {
			test := getVal(m.memory, modes, m.relativeBase, m.ptr, 1)
			if test == 0 {
				jumpTo := getVal(m.memory, modes, m.relativeBase, m.ptr, 2)
				m.ptr = jumpTo - 1
				if debug {
					fmt.Printf("test == 0, m.ptr => %d\n", jumpTo)
				}
			} else {
				if debug {
					fmt.Println("test != 0")
				}
				m.ptr += 2
			}
		} else if opCode == 7 {
			x := getVal(m.memory, modes, m.relativeBase, m.ptr, 1)
			y := getVal(m.memory, modes, m.relativeBase, m.ptr, 2)
			storeAt := getMemLoc(m.memory, modes, m.relativeBase, m.ptr, 3)
			if x < y {
				m.memory[storeAt] = 1
				if debug {
					fmt.Printf("%d < %d => m.memory[%d] = 1\n", x, y, storeAt)
				}
			} else {
				m.memory[storeAt] = 0
				if debug {
					fmt.Printf("%d >= %d => m.memory[%d] = 0\n", x, y, storeAt)
				}
			}
			m.ptr += 3
		} else if opCode == 8 {
			x := getVal(m.memory, modes, m.relativeBase, m.ptr, 1)
			y := getVal(m.memory, modes, m.relativeBase, m.ptr, 2)
			storeAt := getMemLoc(m.memory, modes, m.relativeBase, m.ptr, 3)
			if x == y {
				m.memory[storeAt] = 1
				if debug {
					fmt.Printf("%d == %d => m.memory[%d] = 1\n", x, y, storeAt)
				}
			} else {
				m.memory[storeAt] = 0
				if debug {
					fmt.Printf("%d != %d => m.memory[%d] = 0\n", x, y, storeAt)
				}
			}
			m.ptr += 3
		} else if opCode == 9 {
			offset := getVal(m.memory, modes, m.relativeBase, m.ptr, 1)
			m.relativeBase += offset
			if debug {
				fmt.Printf("m.relativeBase += %d\n", offset)
			}
			m.ptr++
		} else if opCode == 99 {
			return 0, true
		} else {
			log.Panicf("invalid opcode: %d", opCode)
		}
	}
	panic("ran out of program")
}

func getVal(memory []int, modes, relativeBase, i, offset int) int {
	mode := modes / int(math.Pow10(offset-1)) % 10
	param := memory[i+offset]
	if mode == positionMode {
		return memory[param]
	} else if mode == immediateMode {
		return param
	} else if mode == relativeMode {
		return memory[relativeBase+param]
	}
	log.Panicf("invalid mode: %d", mode)
	return 0
}

func getMemLoc(memory []int, modes, relativeBase, i, offset int) int {
	mode := modes / int(math.Pow10(offset-1)) % 10
	param := memory[i+offset]
	if mode == positionMode {
		return param
	} else if mode == relativeMode {
		return relativeBase + param
	}
	log.Panicf("invalid mode: %d", mode)
	return 0
}

func mod(a, b int) int {
	return (a%b + b) % b
}
