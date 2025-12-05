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

	sum := int64(0)
	for scanner.Scan() {
		text := scanner.Text()
		ranges := strings.Split(text, ",")

		for _, rangeString := range ranges {
			sum += findInvalidCount(rangeString)
		}
	}

	println(sum)
}

func findInvalidCount(rangeString string) int64 {
	numberRange := strings.Split(rangeString, "-")
	low, _ := strconv.Atoi(numberRange[0])
	high, _ := strconv.Atoi(numberRange[1])

	invalidCount := int64(0)
	for i := low; i <= high; i++ {
		if isInvalidId(i) {
			invalidCount += int64(i)
		}
	}

	return invalidCount
}

func isInvalidId(id int) bool {
	numberOfDigits := int(math.Log10(float64(id))) + 1

	for i := 1; i <= numberOfDigits/2; i++ {
		if numberOfDigits%i == 0 {
			reductionFactor := int(math.Pow10(i)) // 100
			copyOfId := id
			target := copyOfId % reductionFactor // 42

			for copyOfId > 0 {
				if target != copyOfId%reductionFactor {
					break
				}

				copyOfId /= reductionFactor
			}

			if copyOfId == 0 {
				return true
			}
		}
	}

	return false
}
