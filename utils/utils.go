package utils

import (
	prime "ZK-go/02_prime"
	"math/rand"
)

func Min(a, b int64) int64 {
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
