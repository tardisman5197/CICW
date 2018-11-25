package main

import (
	"fmt"
	"time"
)

// PSO (Particle Swarm Optimisation) finds the best design
// by creating a population of valid solutions, then flocking
// towards the best soloution in the population.
func PSO(noOfGoods int, seed int64) (prices []float64, revenue float64) {
	fmt.Printf("Starting PSO\n")

	// Init antennaArray
	var p PricingProblem
	p.PricingProblem(noOfGoods, seed)

	var population []Particle

	// INITIALISE population, with random designs
	for i := 0; i < 100; i++ {
		currentParticle := Particle{}
		// Find a random design
		currentParticle.currentPostion = randomPrices(noOfGoods)
		// Set the best peak values
		currentParticle.pBest = currentParticle.currentPostion
		currentParticle.pBestRevenue = p.evaluate(currentParticle.currentPostion)
		// Set the inital velocity to the differnce between the init position
		// and another random position divided by 2
		currentParticle.currentVelocity = make([]float64, noOfGoods)
		tmp := randomPrices(noOfGoods)
		for i, pos := range currentParticle.currentPostion {
			currentParticle.currentVelocity[i] = (tmp[i] - pos) / 2
		}

		population = append(population, currentParticle)
	}

	start := time.Now()

	var gBest []float64
	gBestRevenue := 0.0

	// Loop until time termination
	for i := 0; i >= 0; i++ {
		// Update global best
		// for j, cParticle := range population {
		for j := 0; j < len(population); j++ {
			if population[j].pBestRevenue > gBestRevenue {
				gBest = make([]float64, noOfGoods)
				copy(gBest, population[j].currentPostion)
				gBestRevenue = population[j].pBestRevenue
			}
		}

		// 1. UPDATE velocity and position
		// 2. EVALUATE new position
		// 3. UPDATE personal best
		for j := 0; j < len(population); j++ {
			currentParticle := population[j]
			currentParticle.update(gBest)

			// evaluate also updates the personal best
			currentParticle.evalulate(p)

			// if j == 0 && i%500 == 0 {
			// 	fmt.Printf("\r%v: %v : %v", i, currentParticle.pBest, currentParticle.pBestRevenue)
			// }

			population[j] = currentParticle
		}

		// Termination condition
		now := time.Now()
		if now.Sub(start).Seconds() >= executeTime {
			fmt.Printf("\nExecute Time Acheived\n")
			break
		}
	}

	return gBest, gBestRevenue
}
