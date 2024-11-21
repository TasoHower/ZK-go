package division

import (
	eea "ZK-go/04_eea"
	"errors"
)

/*
	逆元：如果存在整数 w 使得 w*y ≡ 1(mod n)，那我们称 w 为 y 在模 n 下的逆元，写为 y^-1。
	当且仅当 y 和 n 互质时（即gcd(y,n)=1），逆元存在
*/

// 求模逆元
func ModularInverseElement[T int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64](q, n T) (T, error) {
	r, _, y := eea.ExtendedEuclideanAlgorithm(q, n)
	if r == 1 {
		return y, nil
	} else {
		return y, errors.New("failed to get ModularInverseElement")
	}
}
