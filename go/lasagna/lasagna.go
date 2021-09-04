package lasagna

func OvenTime() int {
	return 40
}

func RemainingOvenTime(t int) int {
	return 40 - t
}

func PreparationTime(l int) int {
	return l * 2
}

func ElapsedTime(l, t int) int {
	return PreparationTime(l) + (40 - RemainingOvenTime(t))
}
