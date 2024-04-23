package modular_test

import (
	"testing"

	modular "ZK-go/05_modular"
)

func TestQuadraticResidue(t *testing.T) {
	list := modular.QuadraticResidue(5,11)
	t.Log(list)
}
