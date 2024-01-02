package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateTax(t *testing.T) {
	numbers := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedAvg := 5.5
	calculatedAvg := calculateAverage(numbers...)
	require.Equal(t, expectedAvg, calculatedAvg, "Average calculated is not the expected")
}
