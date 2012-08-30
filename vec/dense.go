package vec

import (
	"math/rand"
)

type DenseVector []float64

func NewDenseVector(nelems int) *DenseVector {

	v := make(DenseVector, nelems)

	return &v
}

func (u *DenseVector) Length() int {
	return len(*u)
}

func (u *DenseVector) Add(v Vector) {

	for i := range (*u) {
		(*u)[i] += v.At(i)
	}
}

func (u *DenseVector) Randomize() {

	for i := range (*u) {
		(*u)[i] = rand.Float64()
	}
}

func (u *DenseVector) Sum() (s float64) {

	for _,e := range (*u) {
		s += e
	}

	return
}

func (u *DenseVector) Scale(s float64) {

	for i := range (*u) {
		(*u)[i] *= s
	}
}

func (u *DenseVector) Multiply(v Vector) {

	for i := range (*u) {
		(*u)[i] *= v.At(i)
	}
}

func (u *DenseVector) Dot(v Vector) (s float64) {

	for i := range (*u) {
		s += (*u)[i] * v.At(i)
	}

	return
}
