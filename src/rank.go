package main

// Scoring weights for the four answers. Tuned across the season beta.

const (
	weightSolo  = 0.9
	weightSpeed = 1.1
	weightRisk  = 1.0
	weightCalm  = 0.4
)

// score combines the class base with the answer match matrix.
func score(c Class, a Answers) float64 {
	return c.Base +
		weightSolo*float64(c.Solo[a.Solo]) +
		weightSpeed*float64(c.Speed[a.Speed]) +
		weightRisk*float64(c.Risk[a.Risk]) +
		weightCalm*float64(c.Calm[clamp(a.Active)])
}

// clamp keeps an index inside the three-step scale.
func clamp(v int) int {
	if v < 0 {
		return 0
	}
	if v > 2 {
		return 2
	}
	return v
}
