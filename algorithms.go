package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
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

// artificialImmuneSystem finds a solution for the pricing problem, by using methods similar to
// an immune system. The steps that this algorithm takes are as follows:
// 	1. Initiation, create random soloutions
//	2. Cloning, make beta amount of copies
//	3. Mutation, inverse proportional hyper-mutation
//	4. Selection, choose the best mu for the next population
//	5. Metadynamics, repace the worst d with random solutions
//	6. Repeat until termination condition
func artificialImmuneSystem(noOfGoods int, seed int64) (bestPrices []float64, bestRevenue float64) {
	fmt.Printf("Starting Artificial Immune System\n")

	var p PricingProblem
	p.PricingProblem(noOfGoods, seed)

	start := time.Now()

	var population []Prices

	// Init population with random routes and Eval the routes.
	population = generateRandomPopulation(noOfGoods, popSize, p)

	// Repeat until terminating condition (executeTime)
	for {
		// Cloning
		// Create cloneSizeFactor amount of copies of each route in the population
		var clonePool []Prices
		for j := 0; j < len(population); j++ {
			for k := 0; k < cloneSizeFactor; k++ {
				currentClone := Prices{}
				currentClone.revenue = population[j].revenue
				currentClone.prices = make([]float64, len(population[j].prices))
				copy(currentClone.prices, population[j].prices)
				clonePool = append(clonePool, currentClone)
			}
		}

		// Mutation
		// For each clone in the pool:
		// 	1. Choose a random hotspot
		// 	2. Calc the length of section based on its fitness comapred to the best
		//  3. Reverse the section selected and place back into route
		for j := 0; j < len(clonePool); j++ {
			// Random hotspot
			start := rand.Intn(len(clonePool[j].prices))

			size := len(clonePool[j].prices)

			// length = routeLength * exp(-p*f/fBest)
			inv := math.Exp(-0.5 * (clonePool[j].revenue / bestFitness))
			lengthFloat := inv * float64(size)
			length := int(lengthFloat)

			// Reverse the section
			var tmp []float64
			for k := 0; k < length; k++ {
				tmp = append(tmp, clonePool[j].prices[(k+start)%size])
			}

			for k := 0; k < length; k++ {
				clonePool[j].prices[(k+start)%size] = tmp[len(tmp)-(k+1)]
			}

			// Get the new cost of the clone
			clonePool[j].revenue = p.evaluate(clonePool[j].prices)
		}

		// Selection
		// 1. Combine the population and clone pool
		// 2. Sort the population by fitness
		// 3. Remove worst routes to get back to population size

		// Combine clones and original population
		population = append(population, clonePool...)

		//Sort by fitness
		sort.SliceStable(population, func(i, j int) bool { return population[i].revenue > population[j].revenue })

		// Remove the worst routes
		population = population[:popSize]

		// Metadynamics
		// Swap the worst kth routes with new random routes
		for j := 1; j <= replacementSize; j++ {
			var currentPrices Prices
			currentPrices.prices = randomPrices(noOfGoods)
			currentPrices.revenue = p.evaluate(currentPrices.prices)
			population[len(population)-j] = currentPrices
		}

		bestPrices = population[0].prices
		bestRevenue = population[0].revenue

		// Check if ran out of time
		now := time.Now()
		if now.Sub(start).Seconds() >= executeTime {
			fmt.Printf("Execute Time Acheived\n")
			break
		}
	}
	return
}
