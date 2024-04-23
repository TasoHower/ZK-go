package division

import eea "ZK-go/04_eea"

func ModularInverseElement(q, n int64) int64 {
	r, _, y := eea.ExtendedEuclideanAlgorithm(q, n)
	if r == 1 {
		return int64(y)
	} else {
		return -1
	}
}
