package main

import "testing"

// TestEvaluate checks if the pricing problem evaluate function
// gets the right result.
func TestEvaluate(t *testing.T) {
	// Init PricingProblem object
	var f PricingProblem
	f.priceResponseType = []int{1, 1, 0}
	f.priceResponse = [][]float64{
		{98.7779513552688, 0.6057452361319272},
		{72.55095240160495, 0.7835296077043379},
		{55.411193024402074, 4.8510513976513465},
	}
	f.bnds = [][]float64{
		{0.01, 10.0},
		{0.01, 10.0},
		{0.01, 10.0},
	}
	f.impact = [][]float64{
		{0.0, 0.09634427398335131, 0.017438440697874036},
		{0.08245721442561278, 0.0, 0.02502820719206618},
		{0.05338596274594409, 0.06557473295406872, 0.0},
	}

	expected := 220.0
	actual := f.evaluate([]float64{1.0, 1.0, 1.0})

	if expected != actual {
		t.Errorf("Error: PricingProblem.evaluate() | expected %v got %v", expected, actual)
	}
}
