package conv

import (
	"fmt"
	"log"
	"strconv"
)

// Atoi ...
func Atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("invalid int %s", str))
	}
	return i
}

// ParseFloat ...
func ParseFloat(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Panicf("invalid float %s", str)
	}
	return f
}
