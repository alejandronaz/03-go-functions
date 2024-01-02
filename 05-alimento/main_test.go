package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDog(t *testing.T) {

	alimento := 10
	expected := 100.0
	result := calculateDog(alimento)
	require.Equal(t, expected, result)

}

func TestCat(t *testing.T) {

	alimento := 10
	expected := 50.0
	result := calculateCat(alimento)
	require.Equal(t, expected, result)

}

func TestHamster(t *testing.T) {

	alimento := 10
	expected := 2.5
	result := calculateHamster(alimento)
	require.Equal(t, expected, result)

}

func TestTarantula(t *testing.T) {

	alimento := 10
	expected := 1.5
	result := calculateTarantula(alimento)
	require.Equal(t, expected, result)

}
