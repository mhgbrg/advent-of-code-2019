package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println(isPassword(109165))
	fmt.Println(isPassword(111111))
	fmt.Println(isPassword(223450))
	fmt.Println(isPassword(123789))
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
	for i := 1; i < len(str); i++ {
		if str[i-1] > str[i] {
			return false
		} else if str[i-1] == str[i] {
			adjacent = true
		}
	}
	return adjacent
}
