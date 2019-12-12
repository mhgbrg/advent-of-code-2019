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
	phaseMin      = 5
	phaseMax      = 9
)

func main() {
	start := time.Now()

	program := readProgram()

	seq := []int{phaseMin, phaseMin, phaseMin, phaseMin, phaseMin}
	maxOutput := math.MinInt32
	bestSeq := make([]int, 5)
	for seq[0] = phaseMin; seq[0] <= phaseMax; seq[0]++ {
		for seq[1] = phaseMin; seq[1] <= phaseMax; seq[1]++ {
			for seq[2] = phaseMin; seq[2] <= phaseMax; seq[2]++ {
				for seq[3] = phaseMin; seq[3] <= phaseMax; seq[3]++ {
					for seq[4] = phaseMin; seq[4] <= phaseMax; seq[4]++ {
						if repeating(seq) {
							continue
						}
						fmt.Println(seq)
						output := runWithSeq(program, seq)
						fmt.Println(output)
						if output > maxOutput {
							maxOutput = output
							copy(bestSeq, seq)
						}
					}
				}
			}
		}
	}
	fmt.Println(bestSeq)
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

func runWithSeq(program []int, seq []int) int {
	// init amps
	amps := make([]*machine, 5)
	for i := 0; i < 5; i++ {
		amps[i] = initMachine(program, seq[i])
	}
	// run programs
	signal := 0
	var finalOutput int
	var done bool
	for {
		for i := 0; i < 5; i++ {
			signal, done = amps[i].run(signal)
			if i == 4 {
				finalOutput = signal
			}
			if done {
				return finalOutput
			}
		}
	}
}

type machine struct {
	memory []int
	ptr    int
}

func initMachine(program []int, input int) *machine {
	if program[0] != 3 {
		log.Panic("program does not begin with input operation")
	}
	storeInputAt := program[1]
	memory := make([]int, len(program))
	copy(memory, program)
	memory[storeInputAt] = input
	return &machine{memory, 2}
}

func (m *machine) run(input int) (int, bool) {
	for ; m.ptr < len(m.memory); m.ptr++ {
		opCode := m.memory[m.ptr] % 100
		modes := m.memory[m.ptr] / 100
		if opCode == 1 {
			// x + y
			val1 := m.getVal(1, modes)
			val2 := m.getVal(2, modes)
			storeAt := m.memory[m.ptr+3]
			m.memory[storeAt] = val1 + val2
			m.ptr += 3
		} else if opCode == 2 {
			// x * y
			val1 := m.getVal(1, modes)
			val2 := m.getVal(2, modes)
			storeAt := m.memory[m.ptr+3]
			m.memory[storeAt] = val1 * val2
			m.ptr += 3
		} else if opCode == 3 {
			// input
			storeAt := m.memory[m.ptr+1]
			m.memory[storeAt] = input
			m.ptr++
		} else if opCode == 4 {
			// output
			val := m.getVal(1, modes)
			m.ptr += 2
			return val, false
		} else if opCode == 5 {
			// jumm.ptrf-true
			test := m.getVal(1, modes)
			if test != 0 {
				jumpTo := m.getVal(2, modes)
				m.ptr = jumpTo - 1
			} else {
				m.ptr += 2
			}
		} else if opCode == 6 {
			// jumm.ptrf-false
			test := m.getVal(1, modes)
			if test == 0 {
				jumpTo := m.getVal(2, modes)
				m.ptr = jumpTo - 1
			} else {
				m.ptr += 2
			}
		} else if opCode == 7 {
			// x < y
			val1 := m.getVal(1, modes)
			val2 := m.getVal(2, modes)
			storeAt := m.memory[m.ptr+3]
			if val1 < val2 {
				m.memory[storeAt] = 1
			} else {
				m.memory[storeAt] = 0
			}
			m.ptr += 3
		} else if opCode == 8 {
			// x== y
			val1 := m.getVal(1, modes)
			val2 := m.getVal(2, modes)
			storeAt := m.memory[m.ptr+3]
			if val1 == val2 {
				m.memory[storeAt] = 1
			} else {
				m.memory[storeAt] = 0
			}
			m.ptr += 3
		} else if opCode == 99 {
			return 0, true
		} else {
			log.Panicf("invalid opcode: %d", opCode)
		}
	}
	log.Panic("got to end of machine without exit operation")
	return 0, false
}

func (m *machine) getVal(offset, modes int) int {
	mode := modes / int(math.Pow10(offset-1)) % 10
	param := m.memory[m.ptr+offset]
	if mode == positionMode {
		return m.memory[param]
	} else if mode == immediateMode {
		return param
	}
	log.Panicf("invalid mode: %d", mode)
	return 0
}
