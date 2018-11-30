package main

import (
	"fmt"
	"math/rand"
)

const executeTime = 3
const numberOfGoods = 20
const seed = 0

// Artificial Immnune System
const popSize = 5
const cloneSizeFactor = 2
const bestFitness = 8000
const replacementSize = 2

func main() {

	rand.Seed(seed)
	var p PricingProblem

	p.PricingProblem(numberOfGoods, 0)

	// Run given random search
	fitnessTester(p, seed)
	fmt.Printf("\n\n")

	// Run Artificial Immune
	prices, revenue := artificialImmuneSystem(numberOfGoods, p, seed)
	fmt.Printf("Prices: %v\nRevenue %v\n\n", prices, revenue)

	// Run PSO algorithm
	prices, revenue = PSO(numberOfGoods, p, seed)
	fmt.Printf("Prices: %v\nRevenue %v\n\n", prices, revenue)

}

// randomPrices generates a list of random prices
func randomPrices(noOfGoods int) (prices []float64) {
	for i := 0; i < noOfGoods; i++ {
		prices = append(prices, rand.Float64()*10)
	}
	return
}

// generateRandomPopulation creates a population of prices
func generateRandomPopulation(noOfGoods int, popSize int, p PricingProblem) (population []Prices) {
	for i := 0; i < popSize; i++ {
		currentPrices := Prices{}
		currentPrices.prices = randomPrices(noOfGoods)
		currentPrices.revenue = p.evaluate(currentPrices.prices)
		population = append(population, currentPrices)
	}
	return
}
