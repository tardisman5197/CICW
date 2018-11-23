package main

import (
	"fmt"
	"math/rand"
)

const executeTime = 30

func main() {
	prices, revenue := PSO(5, 10)
	fmt.Printf("Prices: %v\nRevenue %v\n\n", prices, revenue)
	fitnessTester()
}

// fitnessTester runs a random search on the pricing problem
func fitnessTester() {
	numberOfGoods := 20

	var f PricingProblem

	f.PricingProblem(numberOfGoods, 0)

	prices := make([]float64, numberOfGoods)
	newPrices := make([]float64, numberOfGoods)

	for i := 0; i < numberOfGoods; i++ {
		prices[i] = rand.Float64() * 10
	}

	bestRevenue := f.evaluate(prices)
	newRevenue := 0.0

	for i := 0; i < 100; i++ {
		// fmt.Printf("Best revenue so far: %v\n", bestRevenue)

		// Generate more!
		for j := 0; j < numberOfGoods; j++ {
			newPrices[j] = rand.Float64() * 10
		}

		newRevenue = f.evaluate(newPrices)
		if newRevenue > bestRevenue {
			for j := 0; j < len(prices); j++ {
				prices[j] = newPrices[j]
			}
			bestRevenue = newRevenue
		}
	}

	fmt.Printf("Prices: %v\nFinal best revenue: %v", prices, bestRevenue)
}

// randomPrices generates a list of random prices
func randomPrices(noOfGoods int) (prices []float64) {
	for i := 0; i < noOfGoods; i++ {
		prices = append(prices, rand.Float64()*10)
	}
	return
}
