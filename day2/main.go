package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"time"
)

//go:embed day2.txt
var inputFile []byte

func main() {
	start := time.Now()
	count := 0
	totalRibbon := 0
	for _, line := range bytes.Split(bytes.TrimSpace(inputFile), []byte("\n")) {
		values := bytes.Split(line, []byte("x"))
		l := values[0]
		w := values[1]
		h := values[2]
		lNum, _ := strconv.Atoi(string(l))
		wNum, _ := strconv.Atoi(string(w))
		hNum, _ := strconv.Atoi(string(h))
		surface := (2 * wNum * lNum) + (2 * wNum * hNum) + (2 * hNum * lNum)
		ribbon := lNum*2 + wNum*2 + hNum*2
		biggestToRemove := findBiggest([]int{lNum, wNum, hNum})
		ribbon -= 2 * biggestToRemove
		ribbon += lNum * wNum * hNum
		count += findSmallest([]int{lNum * wNum, wNum * hNum, hNum * lNum})
		count += surface
		totalRibbon += ribbon

	}
	elapsed := time.Since(start)
	fmt.Println("Part 1:", count, "Part 2:", totalRibbon, "in", elapsed)
}

func findSmallest(input []int) int {
	return slices.Min(input)
}

func findBiggest(input []int) int {
	return slices.Max(input)
}
