// Package day6
package day6

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func insert(list [][]int, numStrs []string) bool {
	for i := range numStrs {
		num, err := strconv.Atoi(numStrs[i])
		if err != nil {
			return false
		}

		list[i] = append(list[i], num)
	}

	return true
}

func SolutionDay6(scnr *bufio.Scanner) {
	// use first line to generate 2d array
	scnr.Scan()
	line := scnr.Text()
	numStrs := strings.Fields(line)

	nums := make([][]int, len(numStrs))
	insert(nums, numStrs)

	res := 0
	for scnr.Scan() {
		line := scnr.Text()
		strs := strings.Fields(line)
		if !insert(nums, strs) {
			for i, op := range strs {
				opResult := 0
				if op == "*" {
					opResult = 1
					for _, num := range nums[i] {
						opResult *= num
					}
				} else {
					for _, num := range nums[i] {
						opResult += num
					}
				}

				res += opResult
			}
		}
	}

	fmt.Println(res)
}
