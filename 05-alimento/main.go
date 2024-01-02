package main

import (
	"errors"
	"fmt"
)

func main() {
	animalFunc, errAnimal := Animal(dog)
	if errAnimal == nil {
		fmt.Println("Cantidad de alimento:", animalFunc(10), "kg.")
	}

	animalFunc, errAnimal = Animal(cat)
	if errAnimal == nil {
		fmt.Println("Cantidad de alimento:", animalFunc(10), "kg.")
	}

	animalFunc, errAnimal = Animal(hamster)
	if errAnimal == nil {
		fmt.Println("Cantidad de alimento:", animalFunc(10), "kg.")
	}

	animalFunc, errAnimal = Animal(tarantula)
	if errAnimal == nil {
		fmt.Println("Cantidad de alimento:", animalFunc(10), "kg.")
	}

	_, errNot := Animal("mouse")
	if errNot != nil {
		fmt.Println(errNot)
	}

}

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func Animal(tipo string) (func(int) float64, error) {
	switch tipo {
	case dog:
		return calculateDog, nil
	case cat:
		return calculateCat, nil
	case hamster:
		return calculateHamster, nil
	case tarantula:
		return calculateTarantula, nil
	default:
		return nil, errors.New("Animal desconocido")
	}
}

func calculateDog(totalDog int) (food float64) {
	food = float64(totalDog) * 10
	return
}

func calculateCat(totalCat int) (food float64) {
	food = float64(totalCat) * 5
	return
}

func calculateHamster(totalHamster int) (food float64) {
	food = float64(totalHamster) * 0.25
	return
}

func calculateTarantula(totalTarantula int) (food float64) {
	food = float64(totalTarantula) * 0.15
	return
}
