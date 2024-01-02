package main

import (
	"errors"
	"fmt"
)

func main() {

	minFunc, errMin := operation(minimum)
	if errMin == nil {
		fmt.Println(minFunc(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	}

	averageFunc, errAvg := operation(average)
	if errAvg == nil {
		fmt.Println(averageFunc(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	}

	maxFunc, errMax := operation(maximum)
	if errMax == nil {
		fmt.Println(maxFunc(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	}

	_, errNot := operation("mutiply")
	if errNot != nil {
		fmt.Println(errNot)
	}
}

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func calculateMinimun(values ...float64) (min float64) {
	min = values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return
}

func calculateAverage(values ...float64) (average float64) {
	totalPositives := 0.0
	sum := 0.0
	for _, num := range values {
		if num > 0 {
			totalPositives++
			sum += num
		}
	}
	return sum / totalPositives
}

func calculateMaximum(values ...float64) (max float64) {
	max = values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return
}

func operation(opType string) (func(...float64) float64, error) {
	switch opType {
	case minimum:
		return calculateMinimun, nil
	case average:
		return calculateAverage, nil
	case maximum:
		return calculateMaximum, nil
	default:
		return nil, errors.New("invalid operation type")
	}
}
