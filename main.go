package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	aoc25 "github.com/marcelohdez/aoc25/day1"
)

func help(exitCode ...int) {
	fmt.Println("Usage:")
	fmt.Println("\tgo run . <DAY>")

	if len(exitCode) > 0 {
		os.Exit(exitCode[0])
	}
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		help()
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		help()
	}

	solutions := []func(filename string){aoc25.SolutionDay1}

	if day < 0 || day > len(solutions) {
		fmt.Println("Day ", day, " does not exist.")
		fmt.Println()
		help()
	}

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
