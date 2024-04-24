package eea_test

import (
	"testing"

	eea "ZK-go/04_eea"
)

func TestEEA(t *testing.T) {
	a,x1,y1 := eea.ExtendedEuclideanAlgorithm(69,7)
	
	t.Log(a)
	t.Log(x1)
	t.Log(y1)
}
