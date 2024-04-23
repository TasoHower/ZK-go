package utils

func Min(a,b int64)int64{
	if a>b {
		return b
	}

	return a
}

func Max(a,b int64) int64 {
	if a>b {
		return a
	}
	return b
}

func EuclideanDivision(a ,b int64)(int64,int64){
	var quotient,remainder int64

	quotient = a/b
	remainder = a%b

	return quotient,remainder
}