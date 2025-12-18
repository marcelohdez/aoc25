package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/marcelohdez/aoc25/day1"
	"github.com/marcelohdez/aoc25/day2"
	"github.com/marcelohdez/aoc25/day3"
	"github.com/marcelohdez/aoc25/day4"
	"github.com/marcelohdez/aoc25/day5"
	"github.com/marcelohdez/aoc25/day6"
)

func help(msg ...any) {
	if len(msg) > 0 {
		fmt.Println(msg...)
		fmt.Println()
	}

	fmt.Println("Usage:")
	fmt.Println("\tgo run . <DAY> [FILES...]")
	fmt.Println()
	fmt.Println("Where FILES is optional, denoting specific input files to run")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("\tgo run . 2 ./day2/self-test.txt")

	os.Exit(1)
}

func getPaths(day int) ([]string, error) {
	// build filepaths we will read
	var paths []string

	if len(os.Args) > 2 { // use given files only
		paths = os.Args[2:]
	} else { // use all .txt files in day directory
		dir := fmt.Sprintf("day%d", day)
		entries, err := os.ReadDir(dir)
		if err != nil {
			return nil, err
		}

		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".txt") {
				paths = append(paths, path.Join(dir, entry.Name()))
			}
		}
	}

	return paths, nil
}

func run(paths []string, solution func(*bufio.Scanner)) {
	// run solution with each path's scanner
	for _, filepath := range paths {
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer func() {
			err = file.Close()
			if err != nil {
				fmt.Println(err)
			}
		}()

		fmt.Println("----", path.Base(filepath))
		scnr := bufio.NewScanner(file)
		solution(scnr)
	}

	fmt.Println("----")
}

func main() {
	if len(os.Args) == 1 {
		help("Not enough arguments.")
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		help("DAY was not an integer.")
	}

	solutions := []func(*bufio.Scanner){
		day1.SolutionDay1,
		day2.SolutionDay2,
		day3.SolutionDay3,
		day4.SolutionDay4,
		day5.SolutionDay5,
		day6.SolutionDay6,
	}

	if day <= 0 || day > len(solutions) {
		help("Day", day, "does not exist")
	}

	paths, err := getPaths(day)
	if err != nil {
		fmt.Println(err)
		return
	}

	run(paths, solutions[day-1])
}
