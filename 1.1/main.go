package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(cwd + "/1.1/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		total += fuelRequired(val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Total: %v", total)
}

func fuelRequired(mass int) int {
	return int(math.Trunc(float64(mass/3))) - 2
}
