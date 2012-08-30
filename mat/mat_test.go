package mat

import (
	"testing"
)

func TestProduct(t *testing.T) {

	a := New(5, 5)
	b := New(5, 6)
	c := New(6, 5)

	a.Set(1, 1, 4.0)
	b.Set(1, 1, 6.0)

	a.Product(b, c)
}

func TestAdd(t *testing.T) {

	a := New(5, 6)
	b := New(5, 6)

	a.Set(1, 1, 4.0)
	b.Set(1, 1, 6.0)

	a.Add(b)

	if a.At(1, 1) != 10.0 {
		t.Fatal(a.At(1, 1))
	}

}

func TestOnes(t *testing.T) {

	a := New(1, 1)

	a.Set(0, 0, 10.0)

	a.Add(ONES)

	if a.At(0, 0) != 11.0 {
		t.Fatal(a.At(1, 1))
	}
}

func TestIdentity(t *testing.T) {

	a := New(10, 10)

	a.Set(0, 0, 10.0)
	a.Set(1, 0, 10.0)

	a.Add(IDENTITY)

	if a.At(0, 0) != 11.0 || a.At(1, 0) != 10.0 {
		t.Fatal()
	}
}

func TestZeros(t *testing.T) {

	a := New(1, 1)

	a.Set(0, 0, 10.0)

	a.Add(ZEROS)

	if a.At(0, 0) != 10.0 {
		t.Fatal(a.At(1, 1))
	}
}
