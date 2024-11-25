package mathFunc

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// Data struct to store input data and its parameters
type Data struct {
	Data       []float64
	dataSorted []float64
	size       int
	min        float64
	q1         float64
	median     float64
	q3         float64
	max        float64
	iqr        float64
	avg        float64
	sd         float64
	simLineA   float64
	simLineB   float64
	pearCoef   float64
}

// AddDPoint adds a single data point and recalculates the data parameters
func (d *Data) AddDPoint(dPointStr string) {
	if dPoint, err := strconv.ParseFloat(dPointStr, 64); err == nil {
		oldMin, oldMax := d.min, d.max
		d.Data = append(d.Data, dPoint)
		d.SetData(d.Data)

		if dPoint > oldMax || dPoint < oldMin {
			d.FindOutlier()
			return
		}
	}
}

// calculateQuartiles() finds the min, max, IQR, q1, median and q3
func (d *Data) calculateQuartiles() {
	sort.Float64s(d.dataSorted)
	size := len(d.dataSorted)
	if size == 1 {
		d.q1, d.q3, d.median = d.dataSorted[0], d.dataSorted[0], d.dataSorted[0]
		d.min, d.max = d.dataSorted[0], d.dataSorted[0]
		d.iqr = d.q3 - d.q1
		return
	}
	d.min, d.max = d.dataSorted[0], d.dataSorted[d.size-1]

	d.median = GetMedian(d.dataSorted)
	midPoint := size / 2
	d.q1 = GetMedian(d.dataSorted[:midPoint])

	if size%2 == 0 {
		d.q3 = GetMedian(d.dataSorted[midPoint:])
	} else {
		d.q3 = GetMedian(d.dataSorted[midPoint+1:])
	}
	d.iqr = d.q3 - d.q1
}

// FindOutier() uses IQR to calculate the lower and upper bound.
// It removes data points that falls outside the bounds.
func (d *Data) FindOutlier() {
	lwBound := d.q1 - (d.iqr * 3)
	upBound := d.q3 + (d.iqr * 3)
	trimmed := make([]float64, 0)
	for _, dPoint := range d.Data {
		if dPoint >= lwBound && dPoint <= upBound {
			trimmed = append(trimmed, dPoint)
		}
	}
	if len(trimmed) != d.size {
		d.SetData(trimmed)
	}
}

// PrintRange prints the lower and upper limit guess
// using a z-score of 1 and assuming a normal distribution
func (d Data) PrintRange() {
	const zScore = 1.0
	lwRange, upRange := int(d.Data[0]-2), int(d.Data[0]+3)
	if d.size > 1 {
		if d.pearCoef >= -0.35 && d.pearCoef <= 0.35 {

			lwRange = Round(((-zScore * d.sd) + d.median))
			upRange = Round(((zScore * d.sd) + d.median))
		} else {
			nextPoint := d.simLineA + d.simLineB*float64(d.size)
			rangeRatio := (d.max - d.min) / d.sd
			spread := (1 - math.Abs(d.pearCoef)) * rangeRatio

			lwRange = Round((1 - spread) * nextPoint)
			upRange = Round((1 + spread) * nextPoint)
		}
	}
	fmt.Printf("%v %v\n", lwRange, upRange)
}

// SetData() receive the data calculates its parameters
func (d *Data) SetData(input []float64) {
	if len(input) == 0 {
		return
	}
	d.Data = input
	d.size = len(input)

	d.dataSorted = make([]float64, d.size)
	copy(d.dataSorted, d.Data)

	d.calculateQuartiles()
	d.avg = GetAvg(d.Data)
	d.sd = GetStandardDeviation(d.Data)
	d.simLineA, d.simLineB = GetSimpLinearModel(d.Data)
	d.pearCoef = GetPearsonCoef(d.Data)
}
