package weasel

func Search(variants []string, sample string) string {
	var result string
	minFitness := len(sample)
	for _, variant := range variants {
		fitness := Fitness(variant, sample)
		if minFitness > fitness {
			minFitness = fitness
			result = variant
		}
	}

	return result
}
