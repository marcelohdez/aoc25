// Package day3
package day3

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func SolutionDay3p2(scnr *bufio.Scanner) {
	res := 0

	for scnr.Scan() {
		line := scnr.Text()

		var numStr strings.Builder
		starti := 0

		numDigits := 0
		for numDigits < 12 {
			best := 0
			for i := starti; i < len(line)-(11-numDigits); i++ {
				val := int(line[i] - '0')

				if val > best {
					best = val
					starti = i + 1
				}
			}

			numStr.WriteString(strconv.Itoa(best))
			numDigits += 1
		}

		num, err := strconv.Atoi(numStr.String())
		if err != nil {
			fmt.Println(err)
			return
		}

		res += num
	}

	fmt.Println(res)
}
