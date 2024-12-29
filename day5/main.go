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
	countP2 := 0
	lines := bytes.Split(bytes.TrimSpace(inputFile), []byte("\n"))
	for _, line := range lines {
		if doesNotContains(line) && check3Vowels(line) && doublechar(line) {
			count++
		}

		if part2SecondRule(line) && part2FirstRule(line) {
			countP2++
		}
	}
	fmt.Println("Part 1:", count, "in", time.Since(start))
	fmt.Println("Part 2:", countP2, "in", time.Since(start))
}

func check3Vowels(input []byte) bool {
	count := 0
	vowels := []byte{'a', 'e', 'o', 'i', 'u'}
	for _, char := range input {
		for _, vowel := range vowels {
			if char == vowel {
				count++
				if count >= 3 {
					return true
				}
			}
		}
	}
	return false
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

type Position struct {
	start int
	value []byte
}

func part2FirstRule(input []byte) bool {
	pairMap := make(map[string][]int, len(input)-1)

	for i := 0; i < len(input)-1; i++ {
		pair := string(input[i : i+2])
		pairMap[pair] = append(pairMap[pair], i)
	}

	for _, positions := range pairMap {
		if len(positions) < 2 {
			continue
		}
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				if positions[j]-positions[i] > 1 {
					return true
				}
			}
		}
	}
	return false
}

func part2SecondRule(input []byte) bool {
	for i := 0; i < len(input)-2; i++ {
		if input[i] == input[i+2] {
			return true
		}
	}
	return false
}
