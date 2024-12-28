package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"time"
)

//go:embed day5.txt
var inputFile []byte

func main() {
	start := time.Now()
	count := 0
	lines := bytes.Split(bytes.TrimSpace(inputFile), []byte("\n"))
	for _, line := range lines {
		if doesNotContains(line) {
			if check3Vowels(line) && doublechar(line) {
				count++
			}
		}
	}
	fmt.Println("Part 1:", count, "in", time.Since(start))
}

func check3Vowels(input []byte) bool {
	count := 0
	vowels := []byte{'a', 'e', 'o', 'i', 'u'}
	for _, char := range input {
		for _, vowel := range vowels {
			if char == vowel {
				count++
			}
		}
	}
	return count >= 3
}

func doublechar(input []byte) bool {
	var lastChar byte
	for _, char := range input {
		if lastChar != 0 {
			if char == lastChar {
				return true
			}
			lastChar = char
		} else {
			lastChar = char
		}
	}
	return false
}

func doesNotContains(input []byte) bool {
	forbiddens := [][]byte{[]byte("ab"), []byte("cd"), []byte("pq"), []byte("xy")}
	for _, forbidden := range forbiddens {
		if bytes.Contains(input, forbidden) {
			return false
		}
	}
	return true
}
