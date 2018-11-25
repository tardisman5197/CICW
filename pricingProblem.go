package main

import (
	"fmt"
	"math"
	"math/rand"
)

// PricingProblem is a model that can evaluate an instance of the
// pricing problem
type PricingProblem struct {
	priceResponseType []int
	priceResponse     [][]float64
	impact            [][]float64
	bnds              [][]float64
}

// PricingProblem creates an evaluation model
func (p *PricingProblem) PricingProblem(n int, seed int64) {
	rand.Seed(seed)

	p.impact = [][]float64{}

	for i := 0; i < n; i++ {
		// fmt.Printf("Setting up good %v, with type: ", i)

		currentType := rand.Float64()

		if currentType <= 0.4 {
			// Linear
			p.priceResponseType = append(p.priceResponseType, 0)

			currentPriceResponse := make([]float64, 2)
			currentPriceResponse[0] = p.getRandomTotalDemand()
			currentPriceResponse[1] = p.getRandomSatiatingPrice()
			p.priceResponse = append(p.priceResponse, currentPriceResponse)

			// fmt.Printf(" L (%v / %v)\n", currentPriceResponse[0], currentPriceResponse[1])

		} else if currentType > 0.4 && currentType < 0.9 {
			// Constant elasticity
			p.priceResponseType = append(p.priceResponseType, 1)

			currentPriceResponse := make([]float64, 2)
			currentPriceResponse[0] = p.getRandomTotalDemand()
			currentPriceResponse[1] = p.getRandomElasticity()
			p.priceResponse = append(p.priceResponse, currentPriceResponse)

			// fmt.Printf(" CE (%v / %v)\n", currentPriceResponse[0], currentPriceResponse[1])

		} else {
			// Fixed demand
			p.priceResponseType = append(p.priceResponseType, 2)

			currentPriceResponse := make([]float64, 2)
			currentPriceResponse[0] = p.getRandomTotalDemand()
			p.priceResponse = append(p.priceResponse, currentPriceResponse)

			// fmt.Printf(" FD (%v / %v)\n", currentPriceResponse[0], currentPriceResponse[1])
		}

		impactRow := []float64{}
		for j := 0; j < n; j++ {
			impactRow = append(impactRow, rand.Float64()*0.1)
		}
		impactRow[i] = 0.0

		p.impact = append(p.impact, impactRow)
	}

	p.bnds = make([][]float64, len(p.priceResponse))
	dimBnd := []float64{0.01, 10.0}
	for i := 0; i < len(p.priceResponse); i++ {
		p.bnds[i] = dimBnd
	}
}

// isValid checks to see if the prices given lie within the bounds
func (p *PricingProblem) isValid(prices []float64) bool {
	if len(prices) != len(p.bnds) {
		return false
	}

	for i := 0; i < len(prices); i++ {
		if prices[i] < p.bnds[i][0] || prices[i] > p.bnds[i][1] {
			return false
		}
	}
	return true
}

// evaluate gets the total revenue from pricing the goods as given in the parameter.
func (p *PricingProblem) evaluate(prices []float64) float64 {
	if len(prices) != len(p.bnds) {
		fmt.Printf("PricingProblem::evaluate called on price array of the wrong size. Expected: %v Actual: %v\n", len(p.bnds), len(prices))
		panic("Error")
	}

	if !p.isValid(prices) {
		return 0
	}

	revenue := 0.0
	for i := 0; i < len(prices); i++ {
		revenue += float64(p.getDemand(i, prices)) * prices[i]
	}
	return math.Round(revenue*100.0) / 100.0
}

// getDemand gets the demand for good i at the price p
func (p *PricingProblem) getDemand(i int, prices []float64) (demand int) {
	demand = p.getGoodDemand(i, prices[i]) + p.getResidualDemand(i, prices)

	if float64(demand) > p.priceResponse[i][0] {
		demand = int(math.Round(p.priceResponse[i][0]))
	}
	return
}

func (p *PricingProblem) getGoodDemand(i int, price float64) int {
	demand := 0.0
	switch p.priceResponseType[i] {
	case 0: // Linear
		demand = p.priceResponse[i][0] - ((p.priceResponse[i][0] / p.priceResponse[i][1]) * price)
	case 1: // Constant elasticity
		demand = p.priceResponse[i][0] / math.Pow(price, p.priceResponse[i][1])
	case 2: // Fixed demand
		demand = p.priceResponse[i][0]
	default:
		fmt.Printf("Error: Incorrect price response type\n")
	}

	// Sanity check - cannot have more demand than the market holds
	if demand > p.priceResponse[i][0] {
		demand = math.Round(p.priceResponse[i][0])
	}

	// Or less than 0 demand
	if demand < 0 {
		demand = 0
	}

	return int(math.Round(demand))
}

func (p *PricingProblem) getResidualDemand(i int, prices []float64) int {
	demand := 0.0
	for j := 0; j < len(p.priceResponse); j++ {
		if i != j {
			demand += float64(p.getGoodDemand(j, prices[j])) * p.impact[j][i]
		}
	}
	return int(math.Round(demand))
}

func (p *PricingProblem) getRandomTotalDemand() float64 {
	return rand.Float64() * 100
}

func (p *PricingProblem) getRandomSatiatingPrice() float64 {
	return rand.Float64() * 10
}

func (p *PricingProblem) getRandomElasticity() float64 {
	return rand.Float64()
}
