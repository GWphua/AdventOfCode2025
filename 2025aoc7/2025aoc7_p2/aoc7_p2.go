package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./2025aoc7/aoc7.txt")

	if err != nil {
		panic(err)
	}

	var row []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentRow := scanner.Text()

		if row == nil {
			row = make([]int, len(currentRow))
		}
		next := make([]int, len(currentRow))

		for index, value := range currentRow {
			if value == '.' {
				next[index] = next[index] + row[index]
			} else if value == 'S' {
				next[index] = 1
			} else if value == '^' {
				if row[index] > 0 {
					addIfInBounds(next, index+1, row[index])
					addIfInBounds(next, index-1, row[index])
				}
			} else {
				panic("Invalid character in row")
			}
		}

		row = next
	}

	timelines := int64(0)

	for _, value := range row {
		timelines += int64(value)
	}

	fmt.Println(timelines)
}

func addIfInBounds(row []int, index int, value int) {
	if index > len(row) || index < 0 {
		return
	}

	row[index] += value
}
