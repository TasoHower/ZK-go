package modular

func QuadraticResidue(q,n int)[]int{
	var ret []int
	for i := 1;i <q;i++ {
		if (i*i) % n == q {
			ret = append(ret, i)
		}
	}

	return ret
}