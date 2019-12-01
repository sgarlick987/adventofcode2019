package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func moduleFuelCost(mass int) int {
	total := fuelCost(mass)
	cost := total
	for cost > 0 {
		cost = fuelCost(cost)
		total += cost
	}
	return total
}

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
		sum += moduleFuelCost(mass)
	}
	log.Printf("module fuel cost: %d", sum)
}
