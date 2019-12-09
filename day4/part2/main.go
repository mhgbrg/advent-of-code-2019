package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	lower, upper := 109165, 576723
	count := 0
	for i := lower; i <= upper; i++ {
		if isPassword(i) {
			count++
		}
	}
	fmt.Println(count)

	fmt.Println(time.Since(start))
}

func isPassword(num int) bool {
	str := strconv.Itoa(num)
	adjacent := false
	for i := 0; i < len(str); i++ {
		if i > 0 && str[i-1] > str[i] {
			return false
		}
		count := 1
		for ; i < len(str)-1 && str[i] == str[i+1]; i++ {
			count++
		}
		if count == 2 {
			adjacent = true
		}
	}
	return adjacent
}
