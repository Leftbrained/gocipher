package gocipher

import (
	"bytes"
	"testing"
)

func TestGetKeyedAlphabetEmptyKey(t *testing.T) {
	alphabet := GetKeyedAlphabet([]byte{}, []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))

	if !bytes.Equal([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), alphabet) {
		t.Fatalf(`unexpected alphabet: %s`, alphabet)
	}
}

func TestGetKeyedAlphabetKey(t *testing.T) {
	alphabet := GetKeyedAlphabet([]byte("NOPQRSTUVWXYZ"), []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))

	if !bytes.Equal([]byte("NOPQRSTUVWXYZABCDEFGHIJKLM"), alphabet) {
		t.Fatalf(`unexpected alphabet: %s`, alphabet)
	}
}

func TestGetKeyedAlphabetInvalidKeyChars(t *testing.T) {
	alphabet := GetKeyedAlphabet([]byte("CRYPTOGRAPHY"), []byte("ABCD"))

	if !bytes.Equal([]byte("CABD"), alphabet) {
		t.Fatalf(`unexpected alphabet: %s`, alphabet)
	}
}

func TestGetKeyedAlphabetDuplicateAlphabetChars(t *testing.T) {
	alphabet := GetKeyedAlphabet([]byte{}, []byte("ABCQEFGHIJQQMNOPQRSTUVWXYZ"))

	if !bytes.Equal([]byte("ABCQEFGHIJMNOPRSTUVWXYZ"), alphabet) {
		t.Fatalf(`unexpected alphabet: %q`, alphabet)
	}
}
