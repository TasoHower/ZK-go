package remainder

import (
	eea "ZK-go/04_eea"
)

// 表示 a mod n
type Congruences struct {
	A int
	N int
}

// 中国剩余定理(一元同余方程组)
func ChineseRemainderTheorem(input []Congruences) int {
	// 首先计算模元素的乘积

	var m int = 1
	for _, item := range input {
		m = m * item.N
	}

	// 存储 x
	var sum int
	for _, item := range input {
		// 计算 mi 与 mi 的逆元
		mi := m / item.N
		_, _, mie := eea.ExtendedEuclideanAlgorithm(item.N, mi)
		sum += item.A * mi * int(mie)
	}

	return sum % m
}
