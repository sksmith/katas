package popgrowth

import "math"

// NbYear count the number of years until the target p (population)
func NbYear(p0 int, percent float64, aug int, p int) int {
	years := 0
	if p0 >= p {
		return years
	}

	percent = percent / 100
	for currentPop := p0; currentPop <= p; currentPop = updatePop(currentPop, percent, aug) {
		years++
	}
	return years
}

func updatePop(startPop int, percent float64, aug int) (endPop int) {
	return startPop + int(math.Round(float64(startPop)*percent)) + aug
}
