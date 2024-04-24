package rsa_test

import (
	"testing"

	rsa "ZK-go/MS01_RSA"
)

func TestRandomPrime(t *testing.T) {
	p := rsa.RandomPrime()

	t.Log(p)
}

func TestRsa(t *testing.T) {
	pub,priv := rsa.GenPrivateKey()
	t.Logf("pub:%v------priv:%v",pub,priv)


	c := rsa.Encrypt(4,pub)

	t.Logf("c:%d",c)

	m := rsa.Decrypt(c,priv)
	t.Logf("m:%d",m)
}
