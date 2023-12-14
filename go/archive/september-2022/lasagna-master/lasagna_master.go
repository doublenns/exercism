package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += .2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList, myList []string) []string {
	newList := make([]string, len(myList))
	copy(newList, myList)
	newList = append(newList, friendsList[len(friendsList)-1])
	return newList
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	// res := make([]float64, len(quantities))
	var res []float64
	for _, v := range quantities {
		res = append(res, v*float64(portions)/2.0)
	}
	return res
}
