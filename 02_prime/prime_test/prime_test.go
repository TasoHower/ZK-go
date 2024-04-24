package prime_test

import (
	"testing"

	prime "ZK-go/02_prime"
)

func TestSieveOfEratosthenes(t *testing.T) {
	prime.SieveOfEratosthenes(200)
}


func TestPrimeFactors(t *testing.T) {
	list := prime.PrimeFactors(168234551512341)
	t.Log(list)
}

func TestIsPrime(t *testing.T) {
	is := prime.IsPrime(20)
	t.Log(is)
}
