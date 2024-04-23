package integer_test

import (
	"testing"

	integer "ZK-go/01_integer"
)

func TestAdditiveCommutativeLaw(t *testing.T) {
	integer.AdditiveCommutativeLaw()
}

func TestAdditiveLaw(t *testing.T) {
	integer.AdditiveLaw()
}

func TestMultiplicationCommutativeLaw(t *testing.T) {
	integer.MultiplicationCommutativeLaw()
}

func TestMultiplicationLaw(t *testing.T) {
	integer.MultiplicationLaw()
}

func TestEuclideanDivision(t *testing.T) {
	var a,b integer.Integer

	a.Value = 7
	b.Value = 2

	q,r := a.EuclideanDivision(b)
	t.Log(q.Value)
	t.Log(r.Value)
}
