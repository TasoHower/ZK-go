package unit_test

import (
	"testing"

	unit "ZK-go/09_Unit"
)

func TestEulerPhi(t *testing.T) {
	r := unit.EulerPhi(15)
	t.Log(r)
}
