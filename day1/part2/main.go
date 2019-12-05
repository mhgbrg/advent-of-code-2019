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
		fuelForMass := 0
		str := scanner.Text()
		mass := conv.Atoi(str)
		fuel := mass/3 - 2
		fuelForMass += fuel
		for {
			fuel = fuel/3 - 2
			if fuel <= 0 {
				break
			}
			fuelForMass += fuel
		}
		totalFuel += fuelForMass
	}
	fmt.Println(totalFuel)

	fmt.Println(time.Since(start))
}
