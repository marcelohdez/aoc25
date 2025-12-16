// Package day2
package day2

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isInvalid(n int) bool {
	numDigits := int(math.Ceil(math.Log10(float64(n))))

	for div := 10; div < n; div *= 10 {
		p := int(math.Log10(float64(div)))

		if p > numDigits/2 {
			break
		}

		if numDigits%p != 0 {
			continue
		}

		clone := n
		cmp := clone % div

		res := true
		for clone >= div {
			clone /= div
			if clone%div != cmp {
				res = false
				break
			}
		}

		if res {
			return true
		}

	}

	return false
}

func splitOnComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if i := bytes.IndexByte(data, ','); i >= 0 { // we found another comma
		return i + 1, data[:i], nil
	}

	if atEOF && len(data) > 0 {
		return len(data), data, bufio.ErrFinalToken
	}

	return 0, nil, nil
}

func SolutionDay2(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		_ = file.Close()
	}()

	res := 0

	scnr := bufio.NewScanner(file)
	scnr.Split(splitOnComma)
	for scnr.Scan() {
		rangeStr := strings.TrimSpace(scnr.Text())
		if rangeStr == "" {
			break
		}

		before, after, _ := strings.Cut(rangeStr, "-")

		start, err := strconv.Atoi(before)
		if err != nil {
			fmt.Println(err)
			return
		}

		end, err := strconv.Atoi(after)
		if err != nil {
			fmt.Println(err)
			return
		}

		for n := start; n <= end; n++ {
			if isInvalid(n) {
				res += n
			}
		}
	}

	fmt.Println(res)
}
