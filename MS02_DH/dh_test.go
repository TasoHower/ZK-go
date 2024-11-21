package ms02dh

import "testing"

func TestDH(t *testing.T) {
	p, g := GenerateParams()

	// 生成私钥
	alice_private_key := GeneratePrivateKey()
	bob_private_key := GeneratePrivateKey()

	// 生成公钥
	alice_public_key := GeneratePublicKey(alice_private_key, p, g)
	bob_public_key := GeneratePublicKey(bob_private_key, p, g)

	//计算共享密钥
	alice_shared_secret := GenerateSharedSecret(alice_private_key, bob_public_key, p)
	bob_shared_secret := GenerateSharedSecret(bob_private_key, alice_public_key, p)

	t.Logf("大素数 (p):%d", p)
	t.Logf("生成元(g):%d", g)
	t.Logf("Alias 私钥:%d", alice_private_key)
	t.Logf("bob 私钥:%d", bob_private_key)
	t.Logf("Alias 公钥:%d", alice_public_key)
	t.Logf("bob 公钥:%d", bob_public_key)
	t.Logf("Alias 计算的共享密钥：%d", alice_shared_secret)
	t.Logf("Bob 计算的共享密钥：%d", bob_shared_secret)
}
