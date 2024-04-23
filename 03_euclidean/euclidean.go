package euclidean

func EuclideanAlgorithm(a, b int64) int64 {
	// 保证 a 大于 b
	if a < b {
		a, b = b, a
	}

	var ret int64 = b
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
