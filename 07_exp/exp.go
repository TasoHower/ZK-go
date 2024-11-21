package exp

import (
	"math"
	"math/rand"
)

/*
模幂运算
计算 ret = a^c mod n
*/
func ModularExponentiation[T int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64](a, c, n T) T {
	var result T = 1
	for c > 0 {
		result = (result * a) % n
		c--
	}

	return result
}

// 当 p 为素数，a^p - a 可以被 p 整除，a 为任意整数
func FermatSLittleTheorem(p int) float64 {
	// 随机整数 a<10 防止精度丢失
	a := rand.Intn(10)

	return (math.Pow(float64(a), float64(p)) - float64(a)) / float64(p)
}

// 当 a,p 互质，a^(p-1) - 1 /p 为整数
func FermatSLittleTheorem2(a, p int) float64 {
	return (math.Pow(float64(a), float64(p-1)) - 1) / float64(p)
}

// 如果 p 是素数，且 a 不可被 p 整除，那么 a^(p-2)就是 a 在模 p 意义下的逆元
func FermatSLittleTheorem3(a, p int) int {
	return (int(math.Pow(float64(a), float64(p-2))) * a) % p
}
