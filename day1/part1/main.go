package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func fuelCost(mass int) int {
	cost := mass/3 - 2
	if cost < 0 {
		return 0
	}
	return cost
}

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fuel := fuelCost(mass)
		sum += fuel
	}
	log.Printf("module fuel cost: %d", sum)
}
