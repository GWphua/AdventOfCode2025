package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./2025aoc5/aoc5.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var freshIngredients [][2]int64

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		ranges := strings.Split(line, "-")
		lowerRange, _ := strconv.ParseInt(ranges[0], 10, 64)
		upperRange, _ := strconv.ParseInt(ranges[1], 10, 64)
		freshIngredients = append(freshIngredients, [2]int64{lowerRange, upperRange})
	}

	fmt.Println(freshIngredients)
	freshIngredients = mergeFreshIngredientIntervals(freshIngredients)
	fmt.Println(freshIngredients)

	rangeCount := int64(0)
	for i := 0; i < len(freshIngredients); i++ {
		rangeCount += freshIngredients[i][1] - freshIngredients[i][0] + 1
	}

	println(rangeCount)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()

		ingredient, _ := strconv.ParseInt(line, 10, 64)

		for _, ingredientRange := range freshIngredients {
			lowerRange := ingredientRange[0]
			upperRange := ingredientRange[1]

			if ingredient >= lowerRange && ingredient <= upperRange {
				count++
			}
		}
	}

	println(count)

	if file.Close() != nil {
		return
	}
}

func mergeFreshIngredientIntervals(freshIngredients [][2]int64) [][2]int64 {
	sort.Slice(freshIngredients, func(i, j int) bool {
		return freshIngredients[i][0] < freshIngredients[j][0]
	})

	var result [][2]int64
	prev := freshIngredients[0]

	for i := 1; i < len(freshIngredients); i++ {
		if overlappingRange(freshIngredients[i], prev) {
			lower := min(prev[0], freshIngredients[i][0])
			upper := max(prev[1], freshIngredients[i][1])
			prev = [2]int64{lower, upper}
		} else {
			result = append(result, prev)
			prev = freshIngredients[i]
		}
	}

	result = append(result, prev)
	return result
}

func overlappingRange(range1, range2 [2]int64) bool {
	if range1[0] < range2[0] && range1[1] < range2[0] {
		return false
	} else if range1[0] > range2[1] && range1[1] > range2[1] {
		return false
	}

	return true
}
