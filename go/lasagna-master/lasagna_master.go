package lasagna

func PreparationTime(layers []string, prepTimePerLayer int) int {
	if prepTimePerLayer == 0 {
		prepTimePerLayer = 2
	}
	return len(layers) * prepTimePerLayer
}

func Quantities(layers []string) (noodles int, sauce float64) {
	noodlesPerLayer, saucePerLayer := 50, 0.2
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += noodlesPerLayer
		} else if layers[i] == "sauce" {
			sauce += saucePerLayer
		}
	}

	return noodles, sauce
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := make([]float64, 0, len(quantities))
	for i := 0; i <= len(quantities)-1; i++ {
		scaledQuantities = append(scaledQuantities, quantities[i]*float64(portions)/2)
	}

	return scaledQuantities
}
