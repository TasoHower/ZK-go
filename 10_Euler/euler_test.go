package euler_test

import (
	"testing"

	euler "ZK-go/10_Euler"
)

func TestEulerTheorem(t *testing.T) {
	a, e := euler.EulerTheorem(3, 11)

	t.Log(a)
	t.Log(e)
}
