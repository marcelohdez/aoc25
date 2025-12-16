package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	day1 "github.com/marcelohdez/aoc25/day1"
	day2 "github.com/marcelohdez/aoc25/day2"
	"github.com/marcelohdez/aoc25/day3"
)

func help(exitCode ...int) {
	fmt.Println("Usage:")
	fmt.Println("\tgo run . <DAY> [FILES...]")
	fmt.Println()
	fmt.Println("Where FILES is optional, denoting specific input files to run")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("\tgo run . 2 ./day2/self-test.txt")

	if len(exitCode) > 0 {
		os.Exit(exitCode[0])
	}
	os.Exit(1)
}

func main() {
	if len(os.Args) == 1 {
		help()
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		help()
	}

	solutions := []func(filename string){day1.SolutionDay1, day2.SolutionDay2, day3.SolutionDay3}

	if day < 0 || day > len(solutions) {
		fmt.Println("Day ", day, " does not exist.")
		fmt.Println()
		help()
	}

	// run specific files
	if len(os.Args) > 2 {
		for i := 2; i < len(os.Args); i++ {
			fmt.Println("----", os.Args[i])
			solutions[day-1](os.Args[i])
		}
		return
	}

	// run files in day's folder
	dir := fmt.Sprintf("day%d", day)
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".txt") {
			filepath := path.Join(dir, entry.Name())

			fmt.Println("----", entry.Name())
			solutions[day-1](filepath)
		}
	}
	fmt.Println("----")
}
