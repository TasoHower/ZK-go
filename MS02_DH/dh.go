package ms02dh

import (
	"ZK-go/utils"
	"math"
	"math/rand"
)

/*
密钥交换算法
Diffie-Hellman 算法的核心思想是利用离散对数问题，即给定大素数 p、底数 g 和 A ≡g^a mod p，要计算出 a 的值是一个困难问题。算法步骤如下：
密钥生成
选择参数： 选择两个大素数 p 和它的原根（乘法生成元） g。p 和 g 都是公开的。
私钥生成： 每个通信方选择一个私钥。假设 Alice 选择私钥 a，Bob 选择私钥 b, a 和 b 是隐私的。
计算公钥： Alice 计算 A =g^a mod p，Bob 计算 B = g^b mod p。
交换公钥： Alice 将 A 发送给 Bob，Bob 将 B 发送给 Alice，这一步 A 和 B 都是公开的。这一步，公钥 p,g,A,B 是公开的，而私钥 a,b 是隐私的。

2.2 密钥协商
计算会话密钥： Alice 收到 Bob 的公钥 B 后，计算 K=B^a mod p；同样，Bob 收到 Alice 的公钥 A 后，计算 K=A^b mod p。
生成共享密钥： Alice 和 Bob 现在都有相同的密钥 K，该值可以作为会话密钥用于加密通信。由于窃听者没有 a 和 b 的信息，无法计算出 K，因此密钥 K 是隐私的。
*/
func GenerateParams() (int64, int64) {
	p := utils.RandomPrime(3000)
	g := rand.Int63n(p-2) + 2
	return p, g
}

func GeneratePrivateKey() int64 {
	return rand.Int63n(int64(math.Pow(2, 16))) + 2
}

func GeneratePublicKey(prv, p, g int64) int64 {
	return utils.ModPow(g, prv, p)
}

func GenerateSharedSecret(prv, pub, p int64) int64 {
	return utils.ModPow(pub, prv, p)
}
