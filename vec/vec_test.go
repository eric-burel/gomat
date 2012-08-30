package vec

import (
	"testing"
)

func TestAdd(t *testing.T) {

	a := New(5)
	b := New(5)

	a.Randomize()
	b.Randomize()

	old := a.At(2)

	a.Add(b)

	if a.At(2) != old+b.At(2) {
		t.Fatal()
	}
}
