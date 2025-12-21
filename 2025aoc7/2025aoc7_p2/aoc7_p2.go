package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("./2025aoc7/aoc7.txt")

	if err != nil {
		panic(err)
	}

	var row []bool
	scanner := bufio.NewScanner(file)
	split := 0
	for scanner.Scan() {
		currentRow := scanner.Text()

		if row == nil {
			row = make([]bool, len(currentRow))
		}
		next := make([]bool, len(currentRow))

		for index, value := range currentRow {
			if value == '.' {
				next[index] = next[index] || row[index]
			} else if value == 'S' {
				next[index] = true
			} else if value == '^' {
				if row[index] {
					split += 1
					setIfInBounds(next, index+1, true)
					setIfInBounds(next, index-1, true)
				}
			} else {
				panic("Invalid character in row")
			}
		}

		row = next
	}

	println(split)
}

func setIfInBounds(row []bool, index int, value bool) {
	if index > len(row) || index < 0 {
		return
	}

	row[index] = value
}
