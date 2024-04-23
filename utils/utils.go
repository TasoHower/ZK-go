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