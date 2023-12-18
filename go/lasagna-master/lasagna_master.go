package lasagna

func PreparationTime(layers []string, prepTimePerLayer int) int {
	if prepTimePerLayer == 0 {
		prepTimePerLayer = 2
	}

	return len(layers) * prepTimePerLayer
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		} else if layers[i] == "sauce" {
			sauce += 0.2
		}
	}

	return
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	updatedQuantities := make([]float64, len(quantities))
	copy(updatedQuantities, quantities)

	for i := 0; portions != 2 && i < len(quantities); i++ {
		updatedQuantities[i] *= float64(portions) / 2
	}

	return updatedQuantities
}
