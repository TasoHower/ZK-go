package division_test

import (
	"testing"

	division "ZK-go/06_Division"
)

func TestDivision(t *testing.T) {
	a, _ := division.ModularInverseElement(69, 70)
	t.Log(a)
}
