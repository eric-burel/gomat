
package vec

type Vector interface {
	At(i int) float64
}

type Length interface {
	Length() int
}

type LengthVector interface {
	Length
	Vector
}


