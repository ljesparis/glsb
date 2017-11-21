package encryption

import "github.com/ljesparis/glsb/encryption/methods"

type EncryptionConfig struct {
	Key    string
	Method EMethod
}

func (ec EncryptionConfig) Encrypt(m string) string {
	switch ec.Method {
	case Xor:
		return methods.XorEncrypt(m, ec.Key)
	}

	return m
}

func (ec EncryptionConfig) Decrypt(m string) string {
	switch ec.Method {
	case Xor:
		return methods.XorDecrypt(m, ec.Key)
	}

	return m
}

var (
	Default = &EncryptionConfig{
		Key:    "hide message with plsb!",
		Method: None,
	}
)
