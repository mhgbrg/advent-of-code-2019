package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mhgbrg/advent-of-code-2019/util/slices"
)

const (
	positionMode  = 0
	immediateMode = 1
)

func main() {
	start := time.Now()

	program := readProgram()
	runProgram(program)

	fmt.Println(time.Since(start))
}

func readProgram() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	programStr := scanner.Text()
	return slices.Atoi(strings.Split(programStr, ","))
}

func runProgram(program []int) {
	for i := 0; i < len(program); i++ {
		opCode := program[i] % 100
		modes := program[i] / 100
		fmt.Printf("%d: opCode=%d modes=%d program=%v\n", i, opCode, modes, program)
		if opCode == 1 {
			val1 := getVal(program, program[i+1], modes%10)
			val2 := getVal(program, program[i+2], modes/10%10)
			storeAt := program[i+3]
			program[storeAt] = val1 + val2
			i += 3
		} else if opCode == 2 {
			val1 := getVal(program, program[i+1], modes%10)
			val2 := getVal(program, program[i+2], modes/10%10)
			storeAt := program[i+3]
			program[storeAt] = val1 * val2
			i += 3
		} else if opCode == 3 {
			input := 5
			storeAt := program[i+1]
			program[storeAt] = input
			i++
		} else if opCode == 4 {
			val := getVal(program, program[i+1], modes%10)
			fmt.Printf("OUTPUT: %d\n", val)
			i++
		} else if opCode == 5 {
			test := getVal(program, program[i+1], modes%10)
			if test != 0 {
				jumpTo := getVal(program, program[i+2], modes/10%10)
				i = jumpTo - 1
			} else {
				i += 2
			}
		} else if opCode == 6 {
			test := getVal(program, program[i+1], modes%10)
			if test == 0 {
				jumpTo := getVal(program, program[i+2], modes/10%10)
				i = jumpTo - 1
			} else {
				i += 2
			}
		} else if opCode == 7 {
			val1 := getVal(program, program[i+1], modes%10)
			val2 := getVal(program, program[i+2], modes/10%10)
			storeAt := program[i+3]
			if val1 < val2 {
				program[storeAt] = 1
			} else {
				program[storeAt] = 0
			}
			i += 3
		} else if opCode == 8 {
			val1 := getVal(program, program[i+1], modes%10)
			val2 := getVal(program, program[i+2], modes/10%10)
			storeAt := program[i+3]
			if val1 == val2 {
				program[storeAt] = 1
			} else {
				program[storeAt] = 0
			}
			i += 3
		} else if opCode == 99 {
			return
		} else {
			log.Panicf("invalid opcode: %d", opCode)
		}
	}
}

func getVal(program []int, param, mode int) int {
	if mode == positionMode {
		return program[param]
	} else if mode == immediateMode {
		return param
	}
	log.Panicf("invalid mode: %d", mode)
	return 0
}
