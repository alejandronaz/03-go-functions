package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateTax(t *testing.T) {

	t.Run("Salary below 50k", func(t *testing.T) {
		salary := 38000.0
		taxExpected := 0.0
		taxCalculated := CalculateTax(salary)

		require.Equal(t, taxExpected, taxCalculated, "Tax calculated is not the expected")
	})

	t.Run("Salary over 50k", func(t *testing.T) {
		salary := 58000.0
		taxExpected := 9860.0
		taxCalculated := CalculateTax(salary)

		require.Equal(t, taxExpected, taxCalculated, "Tax calculated is not the expected")
	})

	t.Run("Salary over 150k", func(t *testing.T) {
		salary := 158000.0
		taxExpected := 42660.0
		taxCalculated := CalculateTax(salary)

		require.Equal(t, taxExpected, taxCalculated, "Tax calculated is not the expected")
	})

}
