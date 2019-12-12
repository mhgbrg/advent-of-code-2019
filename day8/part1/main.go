package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/mhgbrg/advent-of-code-2019/util/slices"
)

const (
	width  = 25
	height = 6
)

func main() {
	start := time.Now()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	pixels := slices.Atoi(strings.Split(scanner.Text(), ""))
	numLayers := len(pixels) / (width * height)

	minZeroCount := math.MaxInt32

	for layer := 0; layer < numLayers; layer++ {
		offset := layer * width * height
		zeroCount, oneCount, twoCount := 0, 0, 0
		for i := 0; i < width*height; i++ {
			pixel := pixels[offset+i]
			if pixel == 0 {
				zeroCount++
			} else if pixel == 1 {
				oneCount++
			} else if pixel == 2 {
				twoCount++
			}
		}
		if zeroCount < minZeroCount {
			minZeroCount = zeroCount
			fmt.Printf("layer=%d zeros=%d ones=%d twos=%d answer=%d\n", layer, zeroCount, oneCount, twoCount, oneCount*twoCount)
		}
	}

	fmt.Println(time.Since(start))
}
