package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/mhgbrg/advent-of-code-2019/util/conv"
)

func main() {
	start := time.Now()

	scanner := bufio.NewScanner(os.Stdin)
	totalFuel := 0
	for scanner.Scan() {
		str := scanner.Text()
		mass := conv.Atoi(str)
		totalFuel += mass/3 - 2
	}
	fmt.Println(totalFuel)

	fmt.Println(time.Since(start))
}
