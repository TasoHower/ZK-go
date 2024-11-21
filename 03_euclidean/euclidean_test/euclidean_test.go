package euclidean_test

import (
	"testing"

	euclidean "ZK-go/03_euclidean"
)

func TestEuclidean(t *testing.T) {
	b := euclidean.GCD(1263455123465, 1241326541233)
	t.Log(b)
}
