package mathFunc

import "math"

func GetPearsonCoef(data []float64) (coef float64) {
	size := float64(len(data))
	sumX, sumY, sumXY, sumXsq, sumYsq := calcSums(data)
	coef = (size*sumXY - sumX*sumY) /
		math.Sqrt((size*sumXsq-sumX*sumX)*(size*sumYsq-sumY*sumY))

	return
}
