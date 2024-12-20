package rsa

import (
	eea "ZK-go/04_eea"
	"ZK-go/utils"
)

type Key struct {
	a int64
	b int64
}

func GenPrivateKey() (Key, Key) {
	// 生成两个素数
	p, q := utils.RandomPrime(15), utils.RandomPrime(16)
	// 计算 n
	n := p * q

	// 因为 p q 互质，所以 n 的欧拉函数结果为 (p-1)(q-1)
	euler := (p - 1) * (q - 1)

	// 选择一个 e  与 euler 互质，这里选择一个质数且不超过 euler
	e := utils.RandomPrime(int64(euler))

	// 计算 (d * e) mod euler = 1
	_, _, d := eea.ExtendedEuclideanAlgorithm(e, int64(euler))

	return Key{int64(n), int64(e)}, Key{int64(n), int64(d)}
}

func Encrypt(m int64, key Key) int64 {
	return utils.Pow(m, key.b) % key.a
}

func Decrypt(c int64, key Key) int64 {
	return utils.Pow(c, key.b) % key.a
}
