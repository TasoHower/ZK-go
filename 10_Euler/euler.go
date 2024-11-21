package euler

import (
	euclidean "ZK-go/03_euclidean"
	unit "ZK-go/09_Unit"
	"ZK-go/utils"
	"errors"
)

func EulerTheorem[T int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64](a, n T) (T, error) {
	// 检查 a 和 n 是否互质
	if m := euclidean.GCD(a, n); m != 1 {
		return 0, errors.New("gcd not 1")
	}

	unit := unit.EulerPhi(n)

	return utils.Pow(a, unit) % n, nil
}
