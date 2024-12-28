package main

import (
	"crypto/md5"
	_ "embed"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

//go:embed day4.txt
var inputFile string

func main() {
	start := time.Now()
	inputFile = strings.TrimSpace(inputFile)
	i := 0
	p1Found := false
	for {
		h := md5.New()
		io.WriteString(h, inputFile)
		asStr := strconv.Itoa(i)
		io.WriteString(h, asStr)
		hexa := fmt.Sprintf("%x", h.Sum(nil))
		if hexa[:5] == "00000" && !p1Found {
			p1Found = true
			elapsed := time.Since(start)
			fmt.Println("Part 1:", i, "in", elapsed)
		}
		if hexa[:6] == "000000" {
			elapsed := time.Since(start)
			fmt.Println("Part 2:", i, "in", elapsed)
			break
		}
		i++
	}
}
