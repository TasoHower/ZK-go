package exp_test

import (
	exp "ZK-go/07_exp"
	"math"
	"testing"
)

func TestModularExponentiation(t *testing.T) {
	mol := exp.ModularExponentiation(7,5,13)
	t.Log(mol)

	mol2 := int64(math.Pow(7,5)) % 13
	t.Log(mol-mol2)
}

func TestFermatSLittleTheorem(t *testing.T) {
	mol := exp.FermatSLittleTheorem(7)
	t.Log(mol)
}

func TestFermatSLittleTheorem2(t *testing.T) {
	mol := exp.FermatSLittleTheorem2(5,7)
	t.Log(mol)
}

func TestFermatSLittleTheorem3(t *testing.T) {
	mol := exp.FermatSLittleTheorem3(3,5)
	t.Log(mol)
}