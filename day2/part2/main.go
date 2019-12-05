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

func main() {
	start := time.Now()

	originalProgram := readProgram()
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			program := make([]int, len(originalProgram))
			copy(program, originalProgram)
			program[1] = noun
			program[2] = verb
			runProgram(program)
			if program[0] == 19690720 {
				fmt.Printf("noun: %d, verb: %d\n", noun, verb)
				fmt.Println(100*noun + verb)
			}
		}
	}

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
		opCode := program[i]
		if opCode == 1 {
			pos1 := program[i+1]
			pos2 := program[i+2]
			storeAt := program[i+3]
			program[storeAt] = program[pos1] + program[pos2]
			i += 3
		} else if opCode == 2 {
			pos1 := program[i+1]
			pos2 := program[i+2]
			storeAt := program[i+3]
			program[storeAt] = program[pos1] * program[pos2]
			i += 3
		} else if opCode == 99 {
			return
		} else {
			log.Panicf("invalid opcode: %d", opCode)
		}
	}
}
