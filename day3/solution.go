// Package day3
package day3

import (
	"bufio"
	"fmt"
)

func SolutionDay3p1(scnr *bufio.Scanner) {
	res := 0

	for scnr.Scan() {
		line := scnr.Text()

		var left [10]int
		var right [10]int

		// fill in right count
		for _, c := range line {
			val := int(c - '0')
			right[val] += 1
		}

		maxn := 0

		// move pointer along
		for i, c := range line {
			val := int(c - '0')
			right[val] -= 1

			if i == 0 {
				left[val] += 1
				continue
			}

			// find best starting digit
			for i := 9; i >= 0; i-- {
				if left[i] > 0 {
					num := (i * 10) + val
					if num > maxn {
						maxn = num
					}
					break
				}
			}

			// find best ending digit
			for i := 9; i >= 0; i-- {
				if right[i] > 0 {
					num := (val * 10) + i
					if num > maxn {
						maxn = num
					}
					break
				}
			}

			left[val] += 1
		}

		res += maxn
	}

	fmt.Println(res)
}
