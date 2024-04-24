package eea

import (
	"ZK-go/utils"
)

// 扩展欧几里得引理: 输出依次为 最大公约数、a 的 逆元（a、b 互质），b 的逆元（a、b 互质）
func ExtendedEuclideanAlgorithm[T int | int16 | int32 | int64](a, b T) (T, T, T) {
	var x1, y1, x2, y2 = T(1), T(0), T(0), T(1)

	for b != 0 {
		q := T(a / b)
		a, b = b, utils.Mod(a,b)
		x1, x2 = x2, x1-q*x2
		y1, y2 = y2, y1-q*y2
	}

	return a,x1,y1
}

