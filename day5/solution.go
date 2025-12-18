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

func binarySearch(in []Interval, query int) bool {
	l, r := 0, len(in)-1
	for l <= r {
		mid := (l + r) / 2
		if in[mid].start > query {
			r = mid - 1
		} else if in[mid].end < query {
			l = mid + 1
		} else {
			return true
		}
	}

	return false
}

func SolutionDay5(scnr *bufio.Scanner) {
	res := 0
	var intervals []Interval

	inRanges := true
	for scnr.Scan() {
		line := scnr.Text()
		if line == "" {
			inRanges = false
			// sort ranges we have
			sort.Slice(intervals, func(i, j int) bool {
				return intervals[i].start < intervals[j].start
			})
			intervals = merge(intervals)
			continue
		}

		if inRanges {
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
		} else {
			query, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Failed to convert query", line)
				break
			}

			if binarySearch(intervals, query) {
				res += 1
			}
		}
	}

	fmt.Println(res)
}
