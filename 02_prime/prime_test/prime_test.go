package prime_test

import (
	"testing"

	prime "ZK-go/02_prime"
)

func TestSieveOfEratosthenes(t *testing.T) {
	prime.SieveOfEratosthenes(200)
}
