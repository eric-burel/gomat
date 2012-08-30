package mat

type identity struct {
}

var IDENTITY = identity{}

func (a identity) At(i, j int) float64 {

	if i == j {
		return 1.0
	}

	return 0.0
}
