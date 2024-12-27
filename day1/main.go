package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"time"
)

//go:embed day1.txt
var inputFile []byte

func main() {
	start := time.Now()
	count := 0
	for _, char := range bytes.TrimSpace(inputFile) {
		if char == '(' {
			count++
		}
		if char == ')' {
			count--
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Part 1:", count, "in", elapsed)
	start = time.Now()

	count = 0
	basementAt := 0
	for i, char := range bytes.TrimSpace(inputFile) {
		if char == '(' {
			count++
		}
		if char == ')' {
			count--
		}
		if count < 0 {
			basementAt = i + 1
			break
		}
	}

	elapsed = time.Since(start)
	fmt.Println("Part 2:", basementAt, "in", elapsed)
}
