package euclidean

// 欧几里得算法求最大公约数
func EuclideanAlgorithm[T int | int16 | int32 | int64](a, b T) T {
	// 保证 a 大于 b
	if a < b {
		a, b = b, a
	}

	var ret T = b
	for {
		r := a % b
		if r == 0 {
			break
		}

		a = b
		b = r
		ret = r
	}

	return ret
}
