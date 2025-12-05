package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("./2025aoc4/aoc4.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var papers [][]bool

	for scanner.Scan() {
		line := scanner.Text()
		papers = append(papers, createPapersForRow(line))
	}

	removed := 0
	currRemoved := 0

	for {
		currRemoved = 0
		for i := 0; i < len(papers); i++ {
			for j := 0; j < len(papers[i]); j++ {
				if accessibleByForklift(i, j, papers) {
					currRemoved++
					papers[i][j] = false
				}
			}
		}

		if currRemoved == 0 {
			break
		}
		removed += currRemoved
	}

	println(removed)
}

func createPapersForRow(papers string) []bool {
	var rowOfPapers []bool
	for i := 0; i < len(papers); i++ {
		if papers[i] == '@' {
			rowOfPapers = append(rowOfPapers, true)
		} else {
			rowOfPapers = append(rowOfPapers, false)
		}
	}

	return rowOfPapers
}

func accessibleByForklift(i int, j int, papers [][]bool) bool {
	if !papers[i][j] {
		return false
	}

	surroundingPaperCount := 0
	for row := i - 1; row <= i+1; row++ {
		for col := j - 1; col <= j+1; col++ {
			if isPaper(row, col, papers) {
				surroundingPaperCount++
			}
		}
	}

	// Remember to count yourself.
	return surroundingPaperCount-1 < 4
}

func isPaper(i int, j int, papers [][]bool) bool {
	if i < 0 || j < 0 || i >= len(papers) || j >= len(papers[0]) {
		return false
	}

	return papers[i][j]
}
