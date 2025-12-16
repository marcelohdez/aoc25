// Package day1
package day1

import (
	"bufio"
	"fmt"
	"strconv"
)

func SolutionDay1p2(scnr *bufio.Scanner) {
	position := 50
	res := 0

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
