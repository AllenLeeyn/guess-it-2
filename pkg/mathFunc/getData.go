package mathFunc

import (
	"fmt"
	"strconv"
	"strings"
)

// Get data from file content.
// Ideally, one numeric will be on a new line.
// Ignore empty newlines. Skip data points that has problem converting to float64.
func GetData(fileContent string) []float64 {
	var data []float64
	fileContSplited := strings.Split(fileContent, "\n")
	fileContLen := len(fileContSplited)

	for i := 0; i < fileContLen; i++ {
		if fileContSplited[i] == "" {
			continue
		}

		dataPoint, err := strconv.ParseFloat(fileContSplited[i], 64)
		if err != nil {
			fmt.Printf("Error reading data point (line %v). Skipped.\n", i)
			continue
		}
		data = append(data, dataPoint)
	}

	return data
}
