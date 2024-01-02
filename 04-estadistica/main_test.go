package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOperations(t *testing.T) {

	values := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var expected float64
	var result float64

	t.Run("calculate minimun", func(t *testing.T) {
		expected = 1
		result = calculateMinimun(values...)
		require.Equal(t, expected, result)
	})

	t.Run("calculate average", func(t *testing.T) {
		expected = 5.5
		result = calculateAverage(values...)
		require.Equal(t, expected, result)
	})

	t.Run("calculate maximum", func(t *testing.T) {
		expected = 10
		result = calculateMaximum(values...)
		require.Equal(t, expected, result)
	})
}
