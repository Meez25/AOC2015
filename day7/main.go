package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

//go:embed "day7.txt"
var inputFile []byte

var signals = map[string]func() uint16{}
var cache = map[string]uint16{}

func main() {
	start := time.Now()
	for _, line := range bytes.Split(bytes.TrimSpace(inputFile), []byte("\n")) {
		parseLine(line)
	}
	a := lookup("a")
	fmt.Println("Part 1 :", a, "in", time.Since(start))

	for k := range cache {
		delete(cache, k)
	}

	cache["b"] = a
	a = lookup("a")

	fmt.Println("Part 2 :", a, "in", time.Since(start))
}

func lookup(s string) uint16 {
	if val, ok := cache[s]; ok {
		return val
	}

	if number, err := strconv.Atoi(s); err == nil {
		cache[s] = uint16(number)
		return cache[s]
	}

	value := signals[s]()
	cache[s] = value
	return value
}

func parseLine(lineAsByte []byte) {
	var regAnd = regexp.MustCompile(`(.+) AND (.+) -> (\w+)`)
	var regOr = regexp.MustCompile(`(.+) OR (.+) -> (\w+)`)
	var regLshift = regexp.MustCompile(`(\w+) LSHIFT (.+) -> (\w+)`)
	var regRshift = regexp.MustCompile(`(\w+) RSHIFT (.+) -> (\w+)`)
	var regNot = regexp.MustCompile(`NOT (.+) -> (\w+)`)
	var regFeed = regexp.MustCompile(`(.+) -> (\w+)`)
	line := string(lineAsByte)

	if a := regAnd.FindStringSubmatch(line); a != nil {
		signals[a[3]] = func() uint16 { return lookup(a[1]) & lookup(a[2]) }
	} else if a := regOr.FindStringSubmatch(line); a != nil {
		signals[a[3]] = func() uint16 { return lookup(a[1]) | lookup(a[2]) }
	} else if a := regLshift.FindStringSubmatch(line); a != nil {
		signals[a[3]] = func() uint16 { return lookup(a[1]) << lookup(a[2]) }
	} else if a := regRshift.FindStringSubmatch(line); a != nil {
		signals[a[3]] = func() uint16 { return lookup(a[1]) >> lookup(a[2]) }
	} else if a := regNot.FindStringSubmatch(line); a != nil {
		signals[a[2]] = func() uint16 { return ^lookup(a[1]) }
	} else if a := regFeed.FindStringSubmatch(line); a != nil {
		signals[a[2]] = func() uint16 { return lookup(a[1]) }
	}
}
