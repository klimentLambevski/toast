package crypto

import (
	"crypto/aes"
	"crypto/des"

	. "github.com/89hmdys/toast/cipher"
	"github.com/89hmdys/toast/rsa"
)

/*
介绍:创建默认的AES Cipher,使用ECB工作模式、pkcs57填充,算法秘钥长度128 192 256 位 , 使用秘钥作为初始向量

作者:Alex
版本:release-1.1
*/
func NewAES(key []byte) (Cipher, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return NewECBMode().Cipher(block, key[:block.BlockSize()]), err
}

/*
介绍:根据指定的工作模式，创建AESCipher,算法秘钥长度128 192 256 位 , 使用秘钥作为初始向量

作者:Alex
版本:release-1.1
*/
func NewAESWith(key []byte, mode CipherMode) (Cipher, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return mode.Cipher(block, key[:block.BlockSize()]), nil
}

/*
介绍:创建默认DESCipher,使用ECB工作模式、pkcs57填充,算法秘钥长度64位 , 使用秘钥作为初始向量

作者:Alex
版本:release-1.1
*/
func NewDES(key []byte) (Cipher, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return NewECBMode().Cipher(block, key[:block.BlockSize()]), nil
}

/*
介绍:根据指定的工作模式，创建DESCipher,算法秘钥长度64位,使用秘钥作为初始向量

作者:Alex
版本:release-1.1
*/
func NewDESWith(key []byte, mode CipherMode) (Cipher, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return mode.Cipher(block, key[:block.BlockSize()]), nil
}

/*
介绍:创建RSACipher,默认使用pkcs1 padding.

作者:Alex
版本:release-1.1
*/
func NewRSA(key rsa.Key) (rsa.Cipher, error) {
	padding := rsa.NewPKCS1Padding(key.Modulus())
	return rsa.NewCipher(key, padding), nil
}

/*
介绍:根据指定的key,和padding来创建RSACipher

作者:Alex
版本:release-1.1
*/
func NewRSAWith(key rsa.Key, padding rsa.Padding) (rsa.Cipher, error) {
	return rsa.NewCipher(key, padding), nil
}