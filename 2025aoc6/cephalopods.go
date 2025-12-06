package main

import (
	"bufio"
	"os"
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

}
