package mat

type zeros struct {
}

var ZEROS = zeros{}

func (z zeros) At(i, j int) float64 {
	return 0.0
}
