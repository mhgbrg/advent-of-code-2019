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
	input         = 5
)

func main() {
	start := time.Now()

	program := readProgram()
	runProgram(program, input)

	fmt.Println(time.Since(start))
}

func readProgram() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	programStr := scanner.Text()
	return slices.Atoi(strings.Split(programStr, ","))
}

func runProgram(program []int, input int) {
	memory := make([]int, memorySize)
	copy(memory, program)
	relativeBase := 0
	for i := 0; i < len(memory); i++ {
		opCode := memory[i] % 100
		modes := memory[i] / 100
		fmt.Printf("i=%d relativeBase=%d opCode=%d params=%d %d %d\n", i, relativeBase, memory[i], memory[i+1], memory[i+2], memory[i+3])
		if opCode == 1 {
			x := getVal(memory, modes, relativeBase, i, 1)
			y := getVal(memory, modes, relativeBase, i, 2)
			storeAt := getMemLoc(memory, modes, relativeBase, i, 3)
			memory[storeAt] = x + y
			fmt.Printf("memory[%d] = %d + %d = %d\n", storeAt, x, y, memory[storeAt])
			i += 3
		} else if opCode == 2 {
			x := getVal(memory, modes, relativeBase, i, 1)
			y := getVal(memory, modes, relativeBase, i, 2)
			storeAt := getMemLoc(memory, modes, relativeBase, i, 3)
			memory[storeAt] = x * y
			fmt.Printf("%d * %d = %d => memory[%d]\n", x, y, memory[storeAt], storeAt)
			i += 3
		} else if opCode == 3 {
			storeAt := getMemLoc(memory, modes, relativeBase, i, 1)
			memory[storeAt] = input
			fmt.Printf("%d => memory[%d]\n", input, storeAt)
			i++
		} else if opCode == 4 {
			x := getVal(memory, modes, relativeBase, i, 1)
			fmt.Printf("OUTPUT: %d\n", x)
			i++
		} else if opCode == 5 {
			test := getVal(memory, modes, relativeBase, i, 1)
			if test != 0 {
				jumpTo := getVal(memory, modes, relativeBase, i, 2)
				i = jumpTo - 1
				fmt.Printf("test != 0, i => %d\n", jumpTo)
			} else {
				fmt.Println("test == 0")
				i += 2
			}
		} else if opCode == 6 {
			test := getVal(memory, modes, relativeBase, i, 1)
			if test == 0 {
				jumpTo := getVal(memory, modes, relativeBase, i, 2)
				i = jumpTo - 1
				fmt.Printf("test == 0, i => %d\n", jumpTo)
			} else {
				fmt.Println("test != 0")
				i += 2
			}
		} else if opCode == 7 {
			x := getVal(memory, modes, relativeBase, i, 1)
			y := getVal(memory, modes, relativeBase, i, 2)
			storeAt := getMemLoc(memory, modes, relativeBase, i, 3)
			if x < y {
				memory[storeAt] = 1
				fmt.Printf("%d < %d => memory[%d] = 1\n", x, y, storeAt)
			} else {
				memory[storeAt] = 0
				fmt.Printf("%d >= %d => memory[%d] = 0\n", x, y, storeAt)
			}
			i += 3
		} else if opCode == 8 {
			x := getVal(memory, modes, relativeBase, i, 1)
			y := getVal(memory, modes, relativeBase, i, 2)
			storeAt := getMemLoc(memory, modes, relativeBase, i, 3)
			if x == y {
				memory[storeAt] = 1
				fmt.Printf("%d == %d => memory[%d] = 1\n", x, y, storeAt)
			} else {
				memory[storeAt] = 0
				fmt.Printf("%d != %d => memory[%d] = 0\n", x, y, storeAt)
			}
			i += 3
		} else if opCode == 9 {
			offset := getVal(memory, modes, relativeBase, i, 1)
			relativeBase += offset
			fmt.Printf("relativeBase += %d\n", offset)
			i++
		} else if opCode == 99 {
			return
		} else {
			log.Panicf("invalid opcode: %d", opCode)
		}
	}
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
