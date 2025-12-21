package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var homeworkSheet []string

	file, err := os.Open("./2025aoc6/aoc6.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		homeworkSheet = append(homeworkSheet, scanner.Text())
	}

	if file.Close() != nil {
		return
	}

	lastLine := len(homeworkSheet) - 1
	var operation func(int64, int64) int64
	result := int64(0)
	curr := int64(0)

	for i := 0; i < len(homeworkSheet[lastLine]); i++ {
		theresNothingInThisLine := true
		if homeworkSheet[lastLine][i] == '*' {
			result += curr
			curr = 1
			operation = func(a, b int64) int64 {
				return a * b
			}
		} else if homeworkSheet[lastLine][i] == '+' {
			result += curr
			curr = 0
			operation = func(a, b int64) int64 {
				return a + b
			}
		}

		stringBuilder := strings.Builder{}
		for j := 0; j < len(homeworkSheet)-1; j++ {
			if homeworkSheet[j][i] == ' ' {
				continue
			}

			stringBuilder.WriteByte(homeworkSheet[j][i])
			theresNothingInThisLine = false
		}

		if theresNothingInThisLine {
			continue
		}

		cephalopodsValue, _ := strconv.ParseInt(stringBuilder.String(), 10, 64)
		curr = operation(curr, cephalopodsValue)
		println(curr)
		println(result)
	}
	result += curr

	println(result)
}
