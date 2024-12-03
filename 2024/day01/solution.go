package p202401

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Solve() {
	fmt.Println("Solution for 2024 - day 01")
	fmt.Printf("Part one: %v\n", partOne(input))
	fmt.Printf("Part two: %v\n", partTwo(input))
}

func partOne(input string) string {
	var result int
	var listA, listB []int

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		nums := strings.Split(line, "   ")
		if len(nums) != 2 {
			continue
		}
		valA, _ := strconv.Atoi(nums[0])
		valB, _ := strconv.Atoi(nums[1])

		listA = append(listA, valA)
		listB = append(listB, valB)
	}

	sort.Ints(listA)
	sort.Ints(listB)

	for i := range len(listA) {
		if listA[i] < listB[i] {
			result += listB[i] - listA[i]
		} else {
			result += listA[i] - listB[i]
		}
	}

	return strconv.Itoa(result)
}

func partTwo(input string) string {
	var result int
	var listA, listB []int

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		nums := strings.Split(line, "   ")
		if len(nums) != 2 {
			continue
		}
		valA, _ := strconv.Atoi(nums[0])
		valB, _ := strconv.Atoi(nums[1])

		listA = append(listA, valA)
		listB = append(listB, valB)
	}

	for _, valA := range listA {
		for _, valB := range listB {
			if valB == valA {
				result += valA
			}
		}
	}

	return strconv.Itoa(result)
}
