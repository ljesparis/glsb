package main

import "github.com/ljesparis/glsb/encryption"

type Configuration struct {
	Encryption *encryption.EncryptionConfig
}

var (
	defaultConfig = &Configuration{
		Encryption: &encryption.EncryptionConfig{
			Key:    "",
			Method: encryption.None,
		},
	}
)
