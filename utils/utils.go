package utils

import (
	prime "ZK-go/02_prime"
	"math/big"
	"math/rand"
)

func Min[T int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64](a, b T) T {
	if a > b {
		return b
	}

	return a
}

func Max[T int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func EuclideanDivision[T int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64](a, b T) (T, T) {
	var quotient, remainder T

	quotient = a / b
	remainder = a % b

	return quotient, remainder
}

func Mod[T int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64](a, b T) T {
	return (a%b + b) % b
}

func Pow[T int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64](a, b T) T {
	var ret T = 1
	for b > 0 {
		ret = ret * a
		b--
	}

	return ret
}

// 返回一个随机质数
func RandomPrime(a int64) int64 {
	for {
		a := rand.Int63n(a)
		if prime.IsPrime(a) {
			return a
		}
	}
}

func ModPow(base, exponent, modulus int64) int64 {
	// Convert inputs to *big.Int
	b := big.NewInt(base)
	e := big.NewInt(exponent)
	m := big.NewInt(modulus)

	// Calculate (base^exponent) % modulus
	result := new(big.Int).Exp(b, e, m)
	return result.Int64()
}

func IsPrime[T int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64](a T) bool {
	if a < 2 {
		return false
	}
	for i := T(2); i*i < a+1; i++ {
		if a%i == 0 {
			return false
		}
	}

	return true
}
