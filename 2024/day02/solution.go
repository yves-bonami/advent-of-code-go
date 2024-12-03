package p202402

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Solve() {
	fmt.Println("Solution for 2024 - day 02")
	fmt.Printf("Part one: %v\n", partOne(input))
	fmt.Printf("Part two: %v\n", partTwo(input))
}

func partOne(input string) string {
	var result int

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		nums := strings.Split(line, " ")
		if len(nums) < 2 {
			continue
		}

		safe := isSafe(nums)
		if safe {
			result += 1
		}

	}
	return strconv.Itoa(result)
}

func partTwo(input string) string {
	var result int

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		nums := strings.Split(line, " ")
		if len(nums) < 2 {
			continue
		}

		safe := isSafe(nums)
		if safe {
			result += 1
		} else {
			for i := range len(nums) {
				newNums := removeIndex(nums, i)
				safe := isSafe(newNums)
				if safe {
					result += 1
					break
				}
			}
		}

	}

	return strconv.Itoa(result)
}

func isSafe(nums []string) bool {
	prev, _ := strconv.Atoi(nums[0])
	nxt, _ := strconv.Atoi(nums[1])
	ascending := prev < nxt

	result := true

	for i, num := range nums {
		if i == 0 {
			continue
		}

		val, _ := strconv.Atoi(num)
		if (ascending && (val < prev || val-prev > 3 || val-prev < 1)) ||
			(!ascending && (val > prev || prev-val > 3 || prev-val < 1)) {
			result = false
			break
		}
		prev = val
	}

	return result
}

func removeIndex(array []string, index int) []string {
	var newArray []string
	for i, val := range array {
		if i == index {
			continue
		}
		newArray = append(newArray, val)
	}
	return newArray
}
