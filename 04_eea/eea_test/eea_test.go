package eea_test

import (
	"testing"

	eea "ZK-go/04_eea"
)

func TestEEA(t *testing.T) {
	var a, b int64 = 69, 7
	gcd, x1, y1 := eea.ExtendedEuclideanAlgorithm(a, b)

	t.Logf("%d 与 %d 的最大公约数为%d", a, b, gcd)
	t.Logf("满足贝祖等式的整数解为：%d*%d + %d*%d = %d", a, x1, b, y1, gcd)
}
