package euler

import (
	euclidean "ZK-go/03_euclidean"
	unit "ZK-go/09_Unit"
	"ZK-go/utils"
)

func EulerTheorem[T int | int16 | int32 | int64](a, n T) T {
	// 检查 a 和 n 是否互质
	if m := euclidean.EuclideanAlgorithm(a, n); m != 1 {
		return -1 
	}

	unit := unit.EulerPhi(n)


	return utils.Pow(a,unit)%n
}
