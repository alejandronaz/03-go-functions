package main

import "fmt"

func main() {
	fmt.Println(calculateAverage(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(calculateAverage(1, 2, -3, 4, -5, 6, -7, 8, 9, 10))
}

func calculateAverage(numbers ...float64) float64 {
	totalPositives := 0.0
	sum := 0.0
	for _, num := range numbers {
		if num > 0 {
			totalPositives++
			sum += num
		}
	}
	return sum / totalPositives
}
