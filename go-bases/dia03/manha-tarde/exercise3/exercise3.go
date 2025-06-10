package exercise3

func calcSalary(minutesWorked int64, category string) float64 {
	hoursWorked := float64(minutesWorked) / 60.0
	var salary float64

	switch category {
	case "C":
		salary = hoursWorked * 1000
	case "B":
		base := hoursWorked * 1500
		salary = base + (base * 0.20)
	case "A":
		base := hoursWorked * 3000
		salary = base + (base * 0.50)
	}
	return salary
}
