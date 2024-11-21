package ms03ElGamal

import "testing"

func TestElGamal(t *testing.T) {
	key := GenerateKeys()
	msg := 123
	c1, c2 := key.Encrypt(int64(msg))
	m, err := key.Decrypt(c1, c2)

	t.Logf("公钥 (p, g, y): %+v", key.pubKey)
	t.Logf("私钥 x:%d", key.prvKey)
	t.Logf("消息明文:%d", msg)
	t.Logf("密文:%d,%d", c1, c2)
	t.Logf("解密还原明文:%d", m)
	t.Logf("解密报错：%v", err)
}

func TestElGamalSign(t *testing.T) {
	key := GenerateKeys()
	t.Logf("公钥 (p, g, y): %+v", key.pubKey)
	t.Logf("私钥 x:%d", key.prvKey)

	msg := 123
	t.Logf("消息明文:%d", msg)

	r, s, err := key.Sign(int64(msg))
	t.Logf("签名  r:%d s:%d", r, s)
	t.Logf("签名报错：%v", err)

	pass := key.Verify(int64(msg), r, s)

	t.Logf("PASS:%t", pass)
}
