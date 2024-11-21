package modular

/*
二次剩余：
若整数 q 与某个数的平方模 n 同余，即存在 x 满足 pow(x) ≡ q(mod n)，那么称 q 是 n 的二次剩余（Quadratic Residues），否则是非二次剩余。
通俗来讲，就是看 q 在模 n 的情况下会不会是某个数的平方.
*/
func QuadraticResidue[T int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64](q, n T) []T {
	var ret []T

	for i := T(1); i < q; i++ {
		if (i*i)%n == q {
			ret = append(ret, i)
		}
	}

	return ret
}
