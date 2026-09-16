package main

// Scoring weights for the four answers. Tuned across the season beta.

const (
	weightSolo   = 0.9
	weightSpeed  = 1.1
	weightRisk   = 1.0
	weightActive = 0.4
)

// score combines the class base with the answer match matrix.
func score(c Class, a Answers) float64 {
	return c.Base +
		weightSolo*float64(c.Solo[a.Solo]) +
		weightSpeed*float64(c.Speed[a.Speed]) +
		weightRisk*float64(c.Risk[a.Risk]) +
		weightActive*float64(c.Active[a.Active])
}
