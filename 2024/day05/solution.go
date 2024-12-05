package p202405

import (
	_ "embed"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Solve() {
	fmt.Println("Solution for 2024 - day 05")
	fmt.Printf("Part one: %v\n", partOne(input))
	fmt.Printf("Part two: %v\n", partTwo(input))
}

func partOne(input string) string {
	var result int
	rules := make(map[string][]string)
	lines := strings.Split(input, "\n")

	// process rules
	for _, line := range lines {
		if len(line) == 0 {
			break
		}

		rule := strings.Split(line, "|")

		rules[rule[0]] = append(rules[rule[0]], rule[1])
	}

	startUpdate := false
	for _, line := range lines {
		if startUpdate && len(line) > 0 {
			update := strings.Split(line, ",")
			value, _ := checkUpdate(update, rules)
			result += value
		}

		if len(line) == 0 {
			startUpdate = true
		}
	}

	return strconv.Itoa(result)
}

func partTwo(input string) string {
	var result int
	rules := make(map[string][]string)
	lines := strings.Split(input, "\n")

	// process rules
	for _, line := range lines {
		if len(line) == 0 {
			break
		}

		rule := strings.Split(line, "|")

		rules[rule[0]] = append(rules[rule[0]], rule[1])
	}

	startUpdate := false
	for _, line := range lines {
		if startUpdate && len(line) > 0 {
			update := strings.Split(line, ",")
			_, err := checkUpdate(update, rules)

			if err != nil {
				// order correctly
				ordered := order(update, rules)
				value, _ := checkUpdate(ordered, rules)
				result += value
			}
		}

		if len(line) == 0 {
			startUpdate = true
		}
	}
	return strconv.Itoa(result)
}

func checkUpdate(update []string, rules map[string][]string) (int, error) {
	for i, u := range update {
		if rules[u] != nil {
			for _, val := range rules[u] {
				if slices.Contains(update[:i], val) {
					return 0, errors.New("not in order")
				}
			}
		}
	}

	var length int
	if len(update)%2 == 1 {
		length = (len(update) - 1) / 2
	} else {
		length = len(update) / 2
	}

	middleElement, _ := strconv.Atoi(update[length])
	return middleElement, nil
}

func order(update []string, rules map[string][]string) []string {
	newUpdate := make([]string, 0)
	newUpdate = append(newUpdate, update...)
	for i, u := range update {
		if rules[u] != nil {
			index := i
			for _, val := range rules[u] {

				for slices.Contains(newUpdate[:index], val) {
					// remove
					newUpdate = append(newUpdate[:index], newUpdate[index+1:]...)
					// add
					newUpdate = append(newUpdate[:index-1], append([]string{u}, newUpdate[index-1:]...)...)
					index--
				}
			}
		}
	}

	return newUpdate
}
