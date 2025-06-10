package exercise4

func minValue(values ...float64) float64 {
	if len(values) == 0 {
		return 0
	}
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}

func maxValue(values ...float64) float64 {
	if len(values) == 0 {
		return 0
	}
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

func averageValue(values ...float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sum := 0.0
	for _, value := range values {
		sum += value
	}
	return float64(sum) / float64(len(values))
}

func operation(operator string) (func(numbers ...float64) float64, error) {
	switch operator {
	case "minimum":
		return minValue, nil
	case "maximum":
		return maxValue, nil
	case "average":
		return averageValue, nil
	default:
		return nil, nil
	}

}

/*
func main() {
	minFunc, err := operation("minimum")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	averageFunc, err := operation("average")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	maxFunc, err := operation("maximum")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Minimum:", minFunc(1.0, 2.0, 3.0, 4.0, 5.0))
	fmt.Println("Average:", averageFunc(1.0, 2.0, 3.0, 4.0, 5.0))
	fmt.Println("Maximum:", maxFunc(1.0, 2.0, 3.0, 4.0, 5.0))
}
*/
