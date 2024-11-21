package ms03ElGamal

import (
	euclidean "ZK-go/03_euclidean"
	eea "ZK-go/04_eea"
	"ZK-go/utils"
	"math/rand"
)

/*
ElGamal 加密算法
1.1 密钥生成
Bob 使用 ElGamal 算法的密钥生成包括以下步骤：
选择大素数 p：选择一个足够大的素数 p 作为 Z^p 的模数。根据原根的存在性， Z^p 为循环群，存在原根。
选择生成元 g： 选择一个模 p 的原根 g，这时 g 的阶是 p − 1，此时离散对数问题是困难的。
选择私钥 x： 随机选择一个私钥x，1 < x < p。
计算公钥 y： 计算公钥 y = g^x mod p。
最终，公钥为(p,g,y)，是公开的；私钥为 x，不公开。

1.2 加密
Alice 在获取公钥(p,g,y) 后使用 ElGamal 加密，过程如下：

选择随机数k： 随机选择一个k，1 < k < p。
计算临时公钥 a 和临时密文 b： 计算 a ≡ g^k mod p 和 b ≡y^k⋅M ≡ g^(x*k)⋅M mod p，其中 M 是要加密的消息明文。
密文： 密文为(a,b)，是公开的。
随机数k 在每次加密都会变换，保证了 ElGamal 算法即使加密相同的明文也会输出不同的临时密文。在加密操作中，私钥 x 和随机数 k 是隐私的，而公钥(p,g,y) 和密文(a,b)，是公开的。

1.3 解密
Bob 在收到密文(a,b) 后，使用 ElGamal 解密过程如下：

计算共享密钥s： 计算s≡a^x mod p。
计算模逆 s^(−1)。
还原明文 M ： 还原明文 M ≡ b⋅s^(−1) mod p。因为b⋅s^(-1)=b⋅a^(−x)=g^(x*k)⋅M⋅g^(−x*k)=M。
ElGamal 算法就巧妙在 b⋅s−1 能够还原出原文，而计算 s 需要私钥 x 的信息。没有私钥，窃听者想得到明文，就要解离散对数问题（解不出来）。
*/

/*
ElGamal 签名算法
2.1 密钥生成
Alice 生成密钥，用于签名：

选择参数： 选择一个大素数 p 和一个原根 g。
生成私钥： 随机选择一个私钥 x，要求 1 < x < p−1。
计算公钥： 计算公钥 y ≡ g^x (mod p)。
密钥生成步骤和 ElGamal 加密算法相同，最终的公钥为 (p, g, y)，公开的；私钥为 x ，隐私的。

2.2 生成签名
Alice 利用密钥和消息哈希生成签名：

选择随机数： 随机选择一个整数 k，确保 1 <k<p−1 且 gcd(k,p−1)=1。这是因为我们之后会计算k^(−1) (mod p−1)，需要k 在模p−1 下存在逆元素。
计算中间值r： 计算 r ≡ g^k(mod p)。
计算签名： 计算s ≡ k^(−1) * (H(M)−xr) (mod −1)，其中H(M) 是消息M 的哈希值。注意，这里用的模数是 p−1。如果 s = 0，那么需要重新生成一个随机数 k 再次计算。最初论文里用的不是哈希H(M) 而是消息M 本身，但这会带来安全问题。最终的签名为(r,s)，公开的。
2.3 验证签名
Bob 可以利用公开的信息(g,p,r,s,M) 来验证签名是否真实。
验证参数： 如果满足 0<r<p 且 0<s<p−1，可以进行下一步。
验证签名： 如果g^H(M) ≡ y^r * r^s (mod p) 成立，则签名有效。

因为y^r * r^s = g^(x*r) * r^s = g^(x*r) * g^(k*s)=g^(xr+ks)，而 xr+ks = xr+k(k^−1 * (H(M)−xr))= H(M) (mod p−1)。所以，如果g^H(M)≡y^r * r^s(mod p)，也就是 xr+ks = H(M) (mod p−1)
 成立，那么签名就是有效的。
*/

type ElGamalPubKey struct {
	p int64
	g int64
	y int64
}

type ElGamalKey struct {
	pubKey ElGamalPubKey
	prvKey int64
}

func GenerateKeys() ElGamalKey {
	var p int64
	for {
		p = rand.Int63n(1000) + 1000
		if utils.IsPrime(p) {
			break
		}
	}

	g := rand.Int63n(p-3) + 2

	x := rand.Int63n(p-3) + 1

	y := utils.ModPow(g, x, p)

	return ElGamalKey{
		pubKey: ElGamalPubKey{
			p: p,
			g: g,
			y: y,
		},

		prvKey: x,
	}
}

func (c *ElGamalKey) Encrypt(msg int64) (int64, int64) {
	k := rand.Int63n(c.pubKey.p-3) + 1

	c1 := utils.ModPow(c.pubKey.g, k, c.pubKey.p)
	c2 := (msg * utils.ModPow(c.pubKey.y, k, c.pubKey.p)) % c.pubKey.p

	return c1, c2
}

func (c *ElGamalKey) Decrypt(c1, c2 int64) (int64, error) {
	var ret int64
	s := utils.ModPow(c1, c.prvKey, c.pubKey.p)
	mi, err := eea.ModInverse(s, c.pubKey.p)
	if err != nil {
		return -1, err
	}
	ret = (c2 * mi) % c.pubKey.p

	return ret, nil
}

func (c *ElGamalKey) Sign(msg int64) (int64, int64, error) {
	var k int64
	for {
		k = rand.Int63n(c.pubKey.p-2) + 1
		if euclidean.GCD(k, c.pubKey.p-1) == 1 {
			break
		}
	}

	r := utils.ModPow(c.pubKey.g, k, c.pubKey.p)
	x := c.prvKey

	mi, err := eea.ModInverse(k, c.pubKey.p-1)
	if err != nil {
		return -1, -1, err
	}
	s := ((msg - x*r) * mi) % (c.pubKey.p - 1)
	if s < 0 {
		s += (c.pubKey.p - 1)
	}
	return r, s, nil
}

func (c *ElGamalKey) Verify(msg, r, s int64) bool {
	if r < 0 || r > c.pubKey.p {
		return false
	}

	if s < 0 || s > c.pubKey.p-1 {
		return false
	}

	return utils.ModPow(c.pubKey.g, msg, c.pubKey.p) == utils.ModPow(c.pubKey.y, r, c.pubKey.p)*utils.ModPow(r, s, c.pubKey.p)%c.pubKey.p
}
