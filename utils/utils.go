package utils

func Min(a, b int64) int64 {
	if a > b {
		return b
	}

	return a
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func EuclideanDivision(a, b int64) (int64, int64) {
	var quotient, remainder int64

	quotient = a / b
	remainder = a % b

	return quotient, remainder
}

func Mod[T int | int16 | int32 | int64](a, b T) T {
	return (a%b + b) % b
}


func Pow[T int | int16 | int32 | int64](a, b T) T {
	var ret T =1
	for b > 0 {
		ret = ret *a
		b--
	}

	return ret
}
