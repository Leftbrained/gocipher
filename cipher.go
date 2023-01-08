package gocipher

type Cipher interface {
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}

type CharacterCipher interface {
	Cipher
	EncryptByte(from byte) byte
	DecryptByte(from byte) byte
}
