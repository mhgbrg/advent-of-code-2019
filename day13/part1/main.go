package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/mhgbrg/advent-of-code-2019/util/slices"
)

const (
	positionMode  = 0
	immediateMode = 1
	relativeMode  = 2
	memorySize    = 10000
	debug         = false
)

func main() {
	start := time.Now()

	program := readProgram()
	machine := initMachine(program)

	count := 0
	for {
		x, done := machine.run(0)
		if done {
			break
		}
		y, done := machine.run(0)
		if done {
			panic("done on y")
		}
		tile, done := machine.run(0)
		if done {
			panic("done on tile")
		}
		if tile == 2 {
			count++
		}
		fmt.Printf("x=%d y=%d tile=%d\n", x, y, tile)
	}
	fmt.Println(count)

	fmt.Println(time.Since(start))
}

func readProgram() []int {
	scanner := bufio.NewScanner(os.Stdin)
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
