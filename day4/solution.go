// Package day4
package day4

import (
	"bufio"
	"fmt"
)

func SolutionDay4(scnr *bufio.Scanner) {
	symbols := "abcdefghijklmnaopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#$%^&*()-_=+,<.>;:'[{]}"
	res := 0

	var rollMap [][]byte
	for scnr.Scan() {
		rollMap = append(rollMap, []byte(scnr.Text()))
	}

	for lvl := 0; lvl < len(symbols); lvl++ {
		gained := 0

		for r := 0; r < len(rollMap); r++ {
			for c := 0; c < len(rollMap[r]); c++ {
				if rollMap[r][c] != '@' {
					continue
				}

				// count paper rolls in all 8 adjacent spots
				count := 0
				for dr := -1; dr < 2; dr++ {
					for dc := -1; dc < 2; dc++ {
						nr, nc := r+dr, c+dc
						if dr == 0 && dc == 0 { // looking at ourselves
							continue
						}

						if nr >= 0 && nr < len(rollMap) && nc >= 0 && nc < len(rollMap[r]) {
							if rollMap[nr][nc] == '@' || rollMap[nr][nc] == symbols[lvl] {
								count += 1
							}
						}
					}
				}

				if count < 4 {
					rollMap[r][c] = symbols[lvl]
					gained += 1
				}
			}
		}

		res += gained
		if gained == 0 {
			break
		}

		if lvl == len(symbols)-1 {
			// TODO: "clean up" the graph by removing all non-@'s instead of giving up
			fmt.Println("Ran out of symbols. Map may be unfinished.")
		}
	}

	fmt.Println(res)
}
