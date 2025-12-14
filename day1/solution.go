// Package aoc25
package aoc25

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func SolutionDay1(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		_ = file.Close()
	}()

	position := 50
	res := 0

	scnr := bufio.NewScanner(file)
	for scnr.Scan() {
		line := scnr.Text()
		direction := line[0]
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println(err)
			return
		}

		change := 1
		if direction == byte('L') {
			change = -1
		}

		for range amount {
			position += change

			switch position {
			case -1:
				position = 99
			case 100:
				position = 0
			}

			if position == 0 {
				res += 1
			}
		}
	}

	fmt.Println(res)
}
