package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./2025aoc2/aoc2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		ranges := strings.Split(text, ",")

		for _, rangeString := range ranges {
			sum += findInvalidCount(rangeString)
		}
	}
	
	println(sum)
}

func findInvalidCount(rangeString string) int {
	numberRange := strings.Split(rangeString, "-")
	low, _ := strconv.Atoi(numberRange[0])
	high,_ := strconv.Atoi(numberRange[1])

	invalidCount := 0
	for i := low; i <= high; i++ {
		if isInvalidId(i) {
			invalidCount++
		}
	}

	return invalidCount
}

func isInvalidId(i int) bool {
	numberOfDigits := int(math.Log10(float64(i))) + 1

	if numberOfDigits % 2 == 1 {
		return false
	}

	power := int(math.Pow10(numberOfDigits / 2))

	return i % power == i / power
}
