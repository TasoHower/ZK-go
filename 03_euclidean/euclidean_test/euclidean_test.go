package euclidean_test

import (
	"testing"

	euclidean "ZK-go/03_euclidean"
)

func TestEuclidean(t *testing.T) {
	b := euclidean.EuclideanAlgorithm(1234172,12374)
	t.Log(b)
}
