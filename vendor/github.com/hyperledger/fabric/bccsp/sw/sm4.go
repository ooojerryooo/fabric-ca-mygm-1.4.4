package sw

import (
	"github.com/hyperledger/fabric/bccsp"
	"github.com/tjfoc/gmsm/sm4"
)

// AESCBCPKCS7Encrypt combines CBC encryption and PKCS7 padding
func SM4Encrypt(key, src []byte) ([]byte, error) {
	dst := make([]byte, len(src))
	sm4.EncryptBlock(key, dst, src)
	return dst, nil
}

// AESCBCPKCS7Decrypt combines CBC decryption and PKCS7 unpadding
func SM4Decrypt(key, src []byte) ([]byte, error) {
	dst := make([]byte, len(src))
	sm4.DecryptBlock(key, dst, src)
	return dst, nil
}

type gmsm4Encryptor struct{}

//实现 Encryptor 接口
func (*gmsm4Encryptor) Encrypt(k bccsp.Key, plaintext []byte, opts bccsp.EncrypterOpts) (ciphertext []byte, err error) {
	return SM4Encrypt(k.(*gmsm4PrivateKey).privKey, plaintext)
}

type gmsm4Decryptor struct{}

//实现 Decryptor 接口
func (*gmsm4Decryptor) Decrypt(k bccsp.Key, ciphertext []byte, opts bccsp.DecrypterOpts) (plaintext []byte, err error) {
	return SM4Decrypt(k.(*gmsm4PrivateKey).privKey, ciphertext)
}
