package gocipher

type Cipher interface {
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}
