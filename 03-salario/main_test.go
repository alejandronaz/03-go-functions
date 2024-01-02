package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateTax(t *testing.T) {

	t.Run("Category A", func(t *testing.T) {
		minutes := 60.0
		salaryExpected := 1440.0
		salaryCalculated := calculateSalary(minutes, catA)

		require.Equal(t, salaryExpected, salaryCalculated, "Tax calculated is not the expected")
	})

	t.Run("Category B", func(t *testing.T) {
		minutes := 60.0
		salaryExpected := 1800.0
		salaryCalculated := calculateSalary(minutes, catB)

		require.Equal(t, salaryExpected, salaryCalculated, "Tax calculated is not the expected")
	})

	t.Run("Category C", func(t *testing.T) {
		minutes := 60.0
		salaryExpected := 3000.0
		salaryCalculated := calculateSalary(minutes, catC)

		require.Equal(t, salaryExpected, salaryCalculated, "Tax calculated is not the expected")
	})
}
