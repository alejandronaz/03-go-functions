package main

import "fmt"

func main() {

	fmt.Println(calculateSalary(60, catA))
	fmt.Println(calculateSalary(60, catB))
	fmt.Println(calculateSalary(60, catC))

}

const catA = "A"
const catB = "B"
const catC = "C"

func calculateSalary(minutes float64, category string) (salary float64) {

	switch category {
	case catA:
		salary = minutes * (1000 / 60)
		salary *= 1.5
	case catB:
		salary = minutes * (1500 / 60)
		salary *= 1.2
	case catC:
		salary = minutes * (3000 / 60)
	default:
		salary = 0
	}

	return
}
