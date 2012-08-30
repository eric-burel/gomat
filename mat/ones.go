package mat

type ones struct {
}

var ONES = ones{}

func (a ones) At(i, j int) float64 {

	return 1.0
}
