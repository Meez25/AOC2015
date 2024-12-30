package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"time"
	"unicode/utf8"
)

//go:embed day8.txt
var inputFile []byte

func main() {
	start := time.Now()
	count := 0
	P2 := 0
	for _, line := range bytes.Split(bytes.TrimSpace(inputFile), []byte("\n")) {
		representationCount := utf8.RuneCountInString(string(line))
		inMemory := countInMemory(line)
		count += representationCount - inMemory
		encode := encode(line)
		P2 += encode - representationCount
	}
	fmt.Println("Part 1:", count, "in", time.Since(start))
	fmt.Println("Part 2:", P2, "in", time.Since(start))
}

func countInMemory(input []byte) int {
	count := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '\\' {
			if input[i+1] != 'x' {
				i++
			} else {
				i += 3
			}
		}
		count++
	}
	return count - 2
}

func encode(input []byte) int {
	count := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '"' {
			count++
		}
		if input[i] == '\\' {
			count++
		}
		count++
	}
	return count + 2
}
