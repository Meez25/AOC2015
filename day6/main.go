package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const (
	TOGGLE = iota
	OFF
	ON
)

//go:embed day6.txt
var inputFile []byte

type Square struct {
	startX, startY, endX, endY int
}

func main() {
	start := time.Now()

	grid := [1000][1000]bool{}
	for _, line := range bytes.Split(bytes.TrimSpace(inputFile), []byte("\n")) {
		instruction, square := decodeInstruction(line)
		switch instruction {
		case ON:
			for y := square.startY; y <= square.endY; y++ {
				for x := square.startX; x <= square.endX; x++ {
					grid[y][x] = true
				}
			}
		case OFF:
			for y := square.startY; y <= square.endY; y++ {
				for x := square.startX; x <= square.endX; x++ {
					grid[y][x] = false
				}
			}
		case TOGGLE:
			for y := square.startY; y <= square.endY; y++ {
				for x := square.startX; x <= square.endX; x++ {
					grid[y][x] = !grid[y][x]
				}
			}
		}
	}
	count := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] {
				count++
			}
		}
	}
	fmt.Println("Part 1", count, "in", time.Since(start))
	grid2 := [1000][1000]int{}
	for _, line := range bytes.Split(bytes.TrimSpace(inputFile), []byte("\n")) {
		instruction, square := decodeInstruction(line)
		switch instruction {
		case ON:
			for y := square.startY; y <= square.endY; y++ {
				for x := square.startX; x <= square.endX; x++ {
					grid2[y][x]++
				}
			}
		case OFF:
			for y := square.startY; y <= square.endY; y++ {
				for x := square.startX; x <= square.endX; x++ {
					if grid2[y][x] > 0 {
						grid2[y][x]--
					}
				}
			}
		case TOGGLE:
			for y := square.startY; y <= square.endY; y++ {
				for x := square.startX; x <= square.endX; x++ {
					grid2[y][x] += 2
				}
			}
		}
	}
	count2 := 0
	for y := range grid2 {
		for x := range grid2[y] {
			count2 += grid2[y][x]
		}
	}
	fmt.Println("Part 2", count2, "in", time.Since(start))
}

func decodeInstruction(input []byte) (int, Square) {
	var instruction int
	re := regexp.MustCompile(`\d+`)
	if input[6] == 'n' {
		instruction = ON
	}
	if input[6] == 'f' {
		instruction = OFF
	}
	if input[6] == ' ' {
		instruction = TOGGLE
	}

	results := re.FindAllSubmatch(input, -1)
	startX, _ := strconv.Atoi(string(results[0][0]))
	startY, _ := strconv.Atoi(string(results[1][0]))
	endX, _ := strconv.Atoi(string(results[2][0]))
	endY, _ := strconv.Atoi(string(results[3][0]))
	return instruction, Square{startX: startX, startY: startY, endX: endX, endY: endY}
}
