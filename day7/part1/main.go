package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/mhgbrg/advent-of-code-2019/util/mathi"
	"github.com/mhgbrg/advent-of-code-2019/util/slices"
)

const (
	positionMode  = 0
	immediateMode = 1
)

func main() {
	start := time.Now()

	originalProgram := readProgram()
	seq := []int{0, 0, 0, 0, 0}
	maxOutput := math.MinInt32
	for seq[0] = 0; seq[0] < 5; seq[0]++ {
		for seq[1] = 0; seq[1] < 5; seq[1]++ {
			for seq[2] = 0; seq[2] < 5; seq[2]++ {
				for seq[3] = 0; seq[3] < 5; seq[3]++ {
					for seq[4] = 0; seq[4] < 5; seq[4]++ {
						if repeating(seq) {
							continue
						}
						fmt.Println(seq)
						output := 0
						program := make([]int, len(originalProgram))
						copy(program, originalProgram)
						for i := 0; i < 5; i++ {
							output = runProgram(program, []int{seq[i], output})
						}
						fmt.Println(output)
						maxOutput = mathi.Max(maxOutput, output)
					}
				}
			}
		}
	}
	fmt.Println(maxOutput)

	fmt.Println(time.Since(start))
}

func readProgram() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	programStr := scanner.Text()
	return slices.Atoi(strings.Split(programStr, ","))
}

func repeating(seq []int) bool {
	found := make(map[int]bool)
	for _, i := range seq {
		if found[i] {
			return true
		}
		found[i] = true
	}
	return false
}

func runProgram(program []int, inputs []int) int {
	var output int
	j := 0
	for i := 0; i < len(program); i++ {
		opCode := program[i] % 100
		modes := program[i] / 100
		//fmt.Printf("%d: opCode=%d modes=%d program=%v\n", i, opCode, modes, program)
		if opCode == 1 {
			// x + y
			val1 := getVal(program, program[i+1], modes%10)
			val2 := getVal(program, program[i+2], modes/10%10)
			storeAt := program[i+3]
			program[storeAt] = val1 + val2
			i += 3
		} else if opCode == 2 {
			// x * y
			val1 := getVal(program, program[i+1], modes%10)
			val2 := getVal(program, program[i+2], modes/10%10)
			storeAt := program[i+3]
			program[storeAt] = val1 * val2
			i += 3
		} else if opCode == 3 {
			// input
			storeAt := program[i+1]
			program[storeAt] = inputs[j]
			j++
			i++
		} else if opCode == 4 {
			// output
			val := getVal(program, program[i+1], modes%10)
			output = val
			i++
		} else if opCode == 5 {
			// jump-if-true
			test := getVal(program, program[i+1], modes%10)
			if test != 0 {
				jumpTo := getVal(program, program[i+2], modes/10%10)
				i = jumpTo - 1
			} else {
				i += 2
			}
		} else if opCode == 6 {
			// jump-if-false
			test := getVal(program, program[i+1], modes%10)
			if test == 0 {
				jumpTo := getVal(program, program[i+2], modes/10%10)
				i = jumpTo - 1
			} else {
				i += 2
			}
		} else if opCode == 7 {
			// x < y
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
			// x== y
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
			return output
		} else {
			log.Panicf("invalid opcode: %d", opCode)
		}
	}
	log.Panic("got to end of program without exit operation")
	return 0
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
