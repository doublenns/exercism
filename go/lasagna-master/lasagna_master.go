package lasagna

const (
	// defaultTimePerLayer is how long takes to prep a layer whenever the time is unknown.
	defaultTimePerLayer = 2
	// noodlesQuantityPerLayerGrams is how many grams of noodles is needed for each noodles layer.
	noodlesQuantityPerLayerGrams = 50
	// sauceQuantityPerLayerLiters is how many liters of sauce is needed for each sauce layer.
	sauceQuantityPerLayerLiters = .2
)

func PreparationTime(l []string, t int) int {
	if t == 0 {
		t = defaultTimePerLayer
	}
	return len(l) * t
}

func Quantities(l []string) (int, float64) {
	var n int
	var s float64
	for _, v := range l {
		if v == "noodles" {
			n += noodlesQuantityPerLayerGrams
		} else if v == "sauce" {
			s += sauceQuantityPerLayerLiters
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
