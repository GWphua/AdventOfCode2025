package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./2025aoc1/aoc2.txt")
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

	for scanner.Scan() {
		text := scanner.Text()
		
	}
	
	println(clicks)
}
