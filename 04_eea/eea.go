package eea

import (
	"ZK-go/utils"
	"errors"
)

// 扩展欧几里得引理: 输出依次为 最大公约数、a 的 逆元（a、b 互质），b 的逆元（a、b 互质）
func ExtendedEuclideanAlgorithm[T int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64](a, b T) (T, T, T) {
	var x1, y1, x2, y2 = T(1), T(0), T(0), T(1)

	for b != 0 {
		q := (a / b)
		a, b = b, utils.Mod(a, b)
		x1, x2 = x2, x1-q*x2
		y1, y2 = y2, y1-q*y2
	}

	return a, x1, y1
}

// 计算逆元
func ModInverse(a, m int64) (int64, error) {
	gcd, x, _ := ExtendedEuclideanAlgorithm(a, m)
	if gcd != 1 {
		return 0, errors.New("modular inverse does not exist")
	}
	// 确保返回值为正
	return (x%m + m) % m, nil
}
