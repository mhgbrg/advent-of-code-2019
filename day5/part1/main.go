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
	runProgram(program, 1)

	fmt.Println(time.Since(start))
}

func readProgram() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	programStr := scanner.Text()
	return slices.Atoi(strings.Split(programStr, ","))
}

func runProgram(program []int, input int) {
	for i := 0; i < len(program); i++ {
		opCode := program[i] % 100
		modes := program[i] / 100
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
			storeAt := program[i+1]
			program[storeAt] = input
			i++
		} else if opCode == 4 {
			val := getVal(program, program[i+1], modes%10)
			fmt.Printf("OUTPUT: %d\n", val)
			i++
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
