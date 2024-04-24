package unit

import prime "ZK-go/02_prime"

func EulerPhi[T int | int16 | int32 | int64](n T) T {

	ret := n
	// 首先获得 n 的质因数
	primes := prime.PrimeFactors(n)

	for _,p := range primes {
		ret = ret - ret / p
	}	

	return ret
}
