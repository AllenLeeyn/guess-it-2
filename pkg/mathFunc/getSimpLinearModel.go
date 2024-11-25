package mathFunc

// getSimpleLinearModel() finds the slope (b) and y-intercept (a)
// for the simple linear regression model (y = a + bx).
func GetSimpLinearModel(data []float64) (a, b float64) {
	size := float64(len(data))

	sumX, sumY, sumXY, sumXsq, _ := calcSums(data)
	b = ((size * sumXY) - (sumX * sumY)) / ((size * sumXsq) - (sumX * sumX))
	a = (sumY - (b * sumX)) / size
	return
}

// calcSums() calculate the different sums required for
// the simple linear model calculation.
func calcSums(data []float64) (sumX, sumY, sumXY, sumXsq, sumYsq float64) {
	sumX, sumY, sumXY, sumXsq = 0.0, 0.0, 0.0, 0.0
	for x, y := range data {
		sumX += float64(x)
		sumY += y
		sumXY += float64(x) * y
		sumXsq += float64(x) * float64(x)
		sumYsq += y * y
	}
	return
}
