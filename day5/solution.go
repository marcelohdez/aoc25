// Package day5
package day5

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	start int
	end   int
}

func merge(slice []Interval) []Interval {
	res := slice[:1]

	for i := 1; i < len(slice); i++ {
		last := &res[len(res)-1]

		if last.end >= slice[i].start {
			last.end = max(last.end, slice[i].end)
		} else {
			res = append(res, slice[i])
		}
	}

	return res
}

func count(slice []Interval) int {
	res := 0
	for _, interval := range slice {
		res += (interval.end + 1) - interval.start
	}
	return res
}

func SolutionDay5(scnr *bufio.Scanner) {
	var intervals []Interval

	for scnr.Scan() {
		line := scnr.Text()
		if line == "" {
			// sort ranges we have
			sort.Slice(intervals, func(i, j int) bool {
				return intervals[i].start < intervals[j].start
			})
			intervals = merge(intervals)
			break
		}

		startStr, endStr, found := strings.Cut(line, "-")
		if !found {
			fmt.Println("Split not found in ranges?")
			break
		}

		start, err := strconv.Atoi(startStr)
		if err != nil {
			fmt.Println("Failed to convert start to int", startStr)
		}

		end, err := strconv.Atoi(endStr)
		if err != nil {
			fmt.Println("Failed to convert start to int", endStr)
		}

		intervals = append(intervals, Interval{start, end})
	}

	fmt.Println(count(intervals))
}
