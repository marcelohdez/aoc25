// Package day4
package day4

import (
	"bufio"
	"fmt"
)

func SolutionDay4(scnr *bufio.Scanner) {
	res := 0

	var rollMap []string
	for scnr.Scan() {
		rollMap = append(rollMap, scnr.Text())
	}

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
						if rollMap[nr][nc] == '@' {
							count += 1
						}
					}
				}
			}

			if count < 4 {
				fmt.Println(r, c)
				res += 1
			}
		}
	}

	fmt.Println(res)
}
