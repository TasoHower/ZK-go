package rsa_test

import (
	"ZK-go/utils"
	"testing"

	rsa "ZK-go/MS01_RSA"
)

func TestRandomPrime(t *testing.T) {
	p := utils.RandomPrime(10000)

	t.Log(p)
}

func TestRsa(t *testing.T) {
	pub, priv := rsa.GenPrivateKey()
	t.Logf("pub:%v------priv:%v", pub, priv)

	c := rsa.Encrypt(4, pub)

	t.Logf("c:%d", c)

	m := rsa.Decrypt(c, priv)
	t.Logf("m:%d", m)
}
