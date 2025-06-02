package exercise2

func calcAverage(numbers ...int64) float64 {
	sum := int64(0)
	for _, number := range numbers {
		if number < 0 {
			continue
		}
		sum += number
	}
	if len(numbers) == 0 {
		return 0
	}
	return float64(sum) / float64(len(numbers))
}
