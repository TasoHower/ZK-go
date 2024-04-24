package unit

import prime "ZK-go/02_prime"

func EulerPhi[T int | int16 | int32 | int64](n T) T {

	ret := n
	// 首先获得 n 的质因数
	primes := prime.PrimeFactors(n)

	tmp := make(map[T]bool)

	for _, p := range primes {
		tmp[p] = true
	}

	for k := range tmp {
		ret = ret - ret/k
	}

	return ret
}
