package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"time"
)

//go:embed day3.txt
var inputFile []byte

type position struct {
	x int
	y int
}

func main() {
	p1()
	p2()
}

func p1() {
	start := time.Now()
	visited := make(map[position]int, 0)
	firstPos := position{x: 0, y: 0}
	visited[firstPos] = 1
	currentX := 0
	currentY := 0
	for _, char := range bytes.TrimSpace(inputFile) {
		switch char {
		case '>':
			currentX += 1
		case '<':
			currentX -= 1
		case '^':
			currentY -= 1
		case 'v':
			currentY += 1
		}
		visited[position{x: currentX, y: currentY}] += 1
	}
	elapsed := time.Since(start)
	fmt.Println("Part 1", len(visited), "in", elapsed)
}

func p2() {
	start := time.Now()
	visited := make(map[position]int, 0)
	SantaX := 0
	SantaY := 0
	RobotX := 0
	RobotY := 0
	firstPos := position{x: 0, y: 0}
	visited[firstPos] = 2
	for i, char := range bytes.TrimSpace(inputFile) {
		if i%2 == 0 {
			switch char {
			case '>':
				SantaX += 1
			case '<':
				SantaX -= 1
			case '^':
				SantaY -= 1
			case 'v':
				SantaY += 1
			}
			visited[position{x: SantaX, y: SantaY}] += 1
		} else {
			switch char {
			case '>':
				RobotX += 1
			case '<':
				RobotX -= 1
			case '^':
				RobotY -= 1
			case 'v':
				RobotY += 1
			}
			visited[position{x: RobotX, y: RobotY}] += 1
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Part 2", len(visited), "in", elapsed)
}
