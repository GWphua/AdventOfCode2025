package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./2025aoc1/aoc1.txt")
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
	dial := 50
	counter := 0
	clicks := 0
	for scanner.Scan() {
		turns := getKnobValue(scanner.Text())

		wasNotZero := dial != 0
		dial += turns

		if dial <= 0 {
			if wasNotZero {
				clicks += 1 - dial/100
			} else {
				clicks -= dial / 100
			}
		} else if dial >= 100 {
			clicks += dial / 100
		}
		dial = ((dial % 100) + 100) % 100

		if dial == 0 {
			counter++
		}
	}

	println(clicks)
}

func getKnobValue(text string) int {
	direction := 1
	sum := 0
	for index, char := range text {
		if index == 0 {
			if char == 'L' {
				direction = -1
			}

			continue
		}

		sum = 10*sum + int(char-'0')
	}

	return direction * sum
}
