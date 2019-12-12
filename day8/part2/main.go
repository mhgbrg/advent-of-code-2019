package main

import (
	"bufio"
	"fmt"
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
	picture := make([]int, width*height)
	for i := range picture {
		picture[i] = 2
	}
	for layer := 0; layer < numLayers; layer++ {
		offset := layer * width * height
		for i := 0; i < width*height; i++ {
			if picture[i] == 2 {
				picture[i] = pixels[offset+i]
			}
		}
	}
	printPicture(picture)

	fmt.Println(time.Since(start))
}

func printPicture(picture []int) {
	i := 0
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			pixel := picture[i]
			if pixel == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
			i++
		}
		fmt.Println()
	}
}
