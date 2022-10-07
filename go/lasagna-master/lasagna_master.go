package lasagna

func PreparationTime(l []string, t int) int {
	if t == 0 {
		t = 2
	}
	return len(l) * t
}

func Quantities(l []string) (int, float64) {
	var n int
	var s float64
	for _, v := range l {
		if v == "noodles" {
			n += 50
		} else if v == "sauce" {
			s += .2
		}
	}
	return n, s
}

func AddSecretIngredient(f, m []string) {
	m[len(m)-1] = f[len(f)-1]
}

func ScaleRecipe(q []float64, n int) []float64 {
	results := []float64{}
	for _, v := range q {
		results = append(results, v*(float64(n)/2))
	}
	return results
}
