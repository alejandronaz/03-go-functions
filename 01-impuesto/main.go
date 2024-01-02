package main

import "fmt"

func main() {

	fmt.Println(CalculateTax(160000))
	fmt.Println(CalculateTax(110000))
	fmt.Println(CalculateTax(38000))

}

func CalculateTax(salary float64) (tax float64) {
	switch {
	case salary > 150000:
		tax = salary * 0.1
		fallthrough
	case salary > 50000:
		tax += salary * 0.17
	}
	return
}
