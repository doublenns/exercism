package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var r float32 = 0.0
	switch {
	case balance < 0:
		r = 3.213
	case balance < 1000:
		r = 0.5
	case balance < 5000:
		r = 1.621
	case balance >= 5000:
		r = 2.475
	}
	return r
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	r := float64(InterestRate(balance))
	return balance * (r / 100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	y := 0
	for ; balance < targetBalance; y++ {
		balance = AnnualBalanceUpdate(balance)
	}
	return y
}
