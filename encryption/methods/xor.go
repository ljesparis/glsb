package methods

func encryptDecrypt(input, key string) (output string) {
	for i := range input {
		output += string(input[i] ^ key[i%len(key)])
	}

	return
}

func XorEncrypt(m, key string) string {
	return encryptDecrypt(m, key)
}

func XorDecrypt(m, key string) string {
	return encryptDecrypt(m, key)
}
