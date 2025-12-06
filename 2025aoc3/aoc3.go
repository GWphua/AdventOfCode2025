package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./2025aoc3/aoc3.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var sum int64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		maximumJoltage := getMaximumJoltage(line)

		println(maximumJoltage)

		sum += maximumJoltage
	}

	fmt.Println(sum)

	if file.Close() != nil {
		return
	}
}

func getMaximumJoltage(line string) int64 {
	const length = 12
	var joltages [length]int8
	for i := range joltages {
		joltages[i] = -1
	}

	for _, char := range line {
		changed := false

		for i := 0; i < length-1; i++ {
			i1 := joltages[i]
			i2 := joltages[i+1]

			if i1 < i2 || changed {
				joltages[i] = i2
				changed = true
			}
		}

		jolt := int8(char - '0')
		if changed {
			joltages[length-1] = jolt
		} else {
			if joltages[length-1] < jolt {
				joltages[length-1] = jolt
			}
		}
	}

	sum := int64(0)
	for i := 0; i < length; i++ {
		sum *= 10
		sum += int64(joltages[i])
	}

	return sum
}
