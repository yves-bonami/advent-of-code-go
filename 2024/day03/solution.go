package p202403

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

func Solve() {
	fmt.Println("Solution for 2024 - day 03")
	fmt.Printf("Part one: %v\n", partOne(input))
	fmt.Printf("Part two: %v\n", partTwo(input))
}

func partOne(input string) string {
	var result int

	r := regexp.MustCompile(`mul\((?P<valA>\d+),(?P<valB>\d+)\)`)
	matches := r.FindAllStringSubmatch(input, -1)

	for _, m := range matches {
		valA, _ := strconv.Atoi(m[1])
		valB, _ := strconv.Atoi(m[2])
		result += valA * valB
	}

	return strconv.Itoa(result)
}

func partTwo(input string) string {
	var result int

	r := regexp.MustCompile(`mul\((?P<valA>\d+),(?P<valB>\d+)\)|(do\(\))|(don't\(\))`)
	matches := r.FindAllStringSubmatch(input, -1)
	enabled := true
	for _, m := range matches {
		if m[0] == "do()" {
			enabled = true
		} else if m[0] == "don't()" {
			enabled = false
		} else if enabled {
			valA, _ := strconv.Atoi(m[1])
			valB, _ := strconv.Atoi(m[2])
			result += valA * valB
		}
	}

	return strconv.Itoa(result)
}
