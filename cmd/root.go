package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	p202401 "github.com/yvesbonami/advent-of-code-go/2024/day01"
	p202402 "github.com/yvesbonami/advent-of-code-go/2024/day02"
	p202403 "github.com/yvesbonami/advent-of-code-go/2024/day03"
	p202404 "github.com/yvesbonami/advent-of-code-go/2024/day04"
)

var (
	day, year int
	rootCmd   = &cobra.Command{
		Use:   "advent-of-code-go",
		Short: "A collection of Advent of Code puzzles solved in go",
		Long: `A collection of Avent of Code puzzles solved in go.
A solution can be run by specifying year and day in this cli`,
		Run: func(cmd *cobra.Command, args []string) {
			err := runSolution(year, day)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&year, "year", "y", time.Now().Year(), "The year of the puzzle to solve")
	rootCmd.PersistentFlags().IntVarP(&day, "day", "d", 0, "The day of the puzzle to solve")
	rootCmd.MarkPersistentFlagRequired("day")
}

func runSolution(year int, day int) error {
	switch year {
	case 2024:
		switch day {
		case 1:
			p202401.Solve()
		case 2:
			p202402.Solve()
		case 3:
			p202403.Solve()
		case 4:
			p202404.Solve()
		default:
			return errors.New(fmt.Sprintf("No solutions for year %v - day %v", year, day))
		}
	default:
		return errors.New(fmt.Sprintf("No solutions for year %v", year))
	}
	return nil
}
