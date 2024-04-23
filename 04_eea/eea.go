package eea

func ExtendedEuclideanAlgorithm(a, b int64) (int64, int64, int16) {
	var x1, y1, x2, y2 = 1, 0, 0, 1

	for b != 0 {
		q := int(a / b)
		a, b = b, a%b
		x1, x2 = x2, x1-q*x2
		y1, y2 = y2, y1-q*y2
	}

	return a,int64(x1),int16(y1)
}
