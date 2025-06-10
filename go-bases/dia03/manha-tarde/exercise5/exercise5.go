package exercise5

const (
	dog = "dog"
	cat = "cat"
)

func animalDog(quantity int) float64 {
	return 10.0 * float64(quantity)
}

func animalCat(quantity int) float64 {
	return 5.0 * float64(quantity)
}

func animalHamster(quantity int) float64 {
	return 0.250 * float64(quantity)
}

func animalTarantula(quantity int) float64 {
	return 0.150 * float64(quantity)
}

func animal(animalType string) (func(quantity int) float64, error) {

	switch animalType {
	case dog:
		return animalDog, nil
	case cat:
		return animalCat, nil
	case "hamster":
		return animalHamster, nil
	case "tarantula":
		return animalTarantula, nil
	default:
		return nil, nil
	}
}

/*
func main() {
	animalDog, msg := animal(dog)
	if msg != nil {
		println("Error:", msg)
		return
	}

	animalCat, msg := animal(cat)
	if msg != nil {
		println("Error:", msg)
		return
	}

	fmt.Println("Food necessary for 10 dogs: ", animalDog(10))
	fmt.Println("Food necessary for 10 cats: ", animalCat(10))

}
*/
