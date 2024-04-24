package unit_test

import (
	unit "ZK-go/09_Unit"
	"testing"
)

func TestEulerPhi(t *testing.T) {
	r := unit.EulerPhi(9)
	t.Log(r)
}