package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var valuesInString [][]string
	var values [][]int64
	var op []string
	var accumulatedValues []int64

	file, err := os.Open("./2025aoc6/aoc6.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		homeworkValues := strings.Fields(line)

		if checkIsOperations(homeworkValues) {
			op = homeworkValues
		} else {
			valuesInString = append(valuesInString, homeworkValues)
			values = append(values, convertAllToNumbers(homeworkValues))
		}
	}

	for rowNumber := 0; rowNumber < len(values); rowNumber++ {
		row := values[rowNumber]
		for i := 0; i < len(row); i++ {
			if len(accumulatedValues) != len(row) {
				accumulatedValues = append(accumulatedValues, row[i])
			} else {
				accumulatedValues[i] = conductOperation(row[i], accumulatedValues[i], op[i])
			}
		}
	}

	var sum int64 = 0
	for i := 0; i < len(accumulatedValues); i++ {
		sum += accumulatedValues[i]
	}

	println(sum)
}

func checkIsOperations(homeworkValues []string) bool {
	str := homeworkValues[0]
	return strings.Contains(str, "+") || strings.Contains(str, "*")
}

func convertAllToNumbers(homeworkValues []string) []int64 {
	result := make([]int64, len(homeworkValues))
	for i := 0; i < len(homeworkValues); i++ {
		convertedValue, err := strconv.Atoi(homeworkValues[i])

		if err != nil {
			panic(err)
		}

		result[i] = int64(convertedValue)
	}

	return result
}

func conductOperation(a, b int64, op string) int64 {
	if op == "*" {
		return a * b
	} else if op == "+" {
		return a + b
	} else {
		return -1
	}
}
