package p202404

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Solve() {
	fmt.Println("Solution for 2024 - day 04")
	fmt.Printf("Part one: %v\n", partOne(input))
	fmt.Printf("Part two: %v\n", partTwo(input))
}

func partOne(input string) string {
	var result int

	// check horizontal
	var horizontal [][]rune
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		horizontal = append(horizontal, []rune(line))
	}

	for _, line := range horizontal {
		result += countOverlapping(string(line), "XMAS", "SAMX")
	}

	// check vertical
	vertical := make([][]rune, len(horizontal[0]))
	for i := range vertical {
		vertical[i] = make([]rune, len(horizontal))
	}
	for i, line := range horizontal {
		for j, char := range line {
			vertical[j][i] = char
		}
	}

	for _, line := range vertical {
		result += countOverlapping(string(line), "XMAS", "SAMX")
	}

	// check diagonal 1
	var diagonal [][]rune

	// determine max diagonal length
	max := len(horizontal)
	if max > len(horizontal[0]) {
		max = len(horizontal[0])
	}

	for i, line := range horizontal {
		for j := range line {
			var dLine []rune
			if i != 0 && j > 0 {
				continue
			}

			k := 0
			for range max {
				if i+k < len(horizontal[0]) && j+k < len(horizontal) {
					dLine = append(dLine, horizontal[i+k][j+k])
				}
				k++

			}
			diagonal = append(diagonal, dLine)
		}
	}

	for _, line := range diagonal {
		result += countOverlapping(string(line), "XMAS", "SAMX")
	}

	// check diagonal 2
	var diagonalRev [][]rune
	for i, line := range horizontal {
		for j := range line {
			j = len(line) - j - 1
			var dLine []rune
			if i != 0 && j < len(line)-1 {
				continue
			}

			k := 0
			for range max {
				if i+k < len(horizontal[0]) && j-k >= 0 {
					dLine = append(dLine, horizontal[i+k][j-k])
				}
				k++
			}
			diagonalRev = append(diagonalRev, dLine)
		}
	}

	for _, line := range diagonalRev {
		result += countOverlapping(string(line), "XMAS", "SAMX")
	}

	return strconv.Itoa(result)
}

func partTwo(input string) string {
	var result int

	var horizontal [][]rune
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		horizontal = append(horizontal, []rune(line))
	}

	for i, line := range horizontal {
		for j, char := range line {
			if char != 'A' ||
				i == 0 || i == len(horizontal)-1 ||
				j == 0 || j == len(horizontal[0])-1 {
				continue
			}

			// check previous
			prevA := horizontal[i-1][j-1]
			if prevA != 'M' && prevA != 'S' {
				continue
			}
			prevB := horizontal[i-1][j+1]
			if prevB != 'M' && prevB != 'S' {
				continue
			}

			// check next
			nextA := horizontal[i+1][j-1]
			if nextA == prevB || (nextA != 'M' && nextA != 'S') {
				continue
			}
			nextB := horizontal[i+1][j+1]
			if nextB == prevA || (nextB != 'M' && nextB != 'S') {
				continue
			}

			result++
		}
	}

	return strconv.Itoa(result)
}

func countOverlapping(input string, substring string, subReversed string) int {
	var c int
	if len(input) < len(substring) && len(input) < len(subReversed) {
		return 0
	}
	for d := range input {
		if strings.HasPrefix(input[d:], substring) || strings.HasPrefix(input[d:], subReversed) {
			c++
		}
	}
	return c
}
