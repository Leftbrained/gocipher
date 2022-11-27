package cipher

import (
	"bytes"
	"testing"
)

func TestPolybiusNew(t *testing.T) {
	c, err := NewPolybius(
		[]byte("CRYPTOGRAPHY"),
	)

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate: %q`, err.Error())
	}
}

func TestPolybiusNewSix(t *testing.T) {
	c, err := NewPolybius(
		[]byte("CRYPTOGRAPHY"),
		PolybiusWithAlphabet([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"), map[byte]byte{}),
	)

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate: %q`, err.Error())
	}
}

func TestPolybiusNewErrorNotSquare(t *testing.T) {
	c, err := NewPolybius(
		[]byte("CRYPTOGRAPHY"),
		PolybiusWithAlphabet([]byte("ABCDEFGHIKLMNOPQRSTUVWXY"), map[byte]byte{
			'J': 'I',
		}),
	)

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestPolybiusNewErrorCoordsTooSmall(t *testing.T) {
	c, err := NewPolybius(
		[]byte("CRYPTOGRAPHY"),
		PolybiusWithCoords([]byte("12")),
	)

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestPolybiusBasicCrypt(t *testing.T) {
	c, err := NewPolybius(
		[]byte("CRYPTOGRAPHY"),
	)
	if err != nil {
		t.Fatalf(`unexpected: could not instantiate: %q`, err.Error())
	}

	testCipherCrypt(c, t,
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
		[]byte("15243244513411352512215343332154514214452152321215243241235513312122"),
		[]byte("THEQUICKBROWNFOXUMPSOVERTHELAZYDOG"),
	)
}

func TestPolybiusExtraLetterDecrypt(t *testing.T) {
	c, err := NewPolybius(
		[]byte("CRYPTOGRAPHY"),
	)
	if err != nil {
		t.Fatalf(`unexpected: could not instantiate: %q`, err.Error())
	}

	plain := c.Decrypt([]byte("312"))

	if !bytes.Equal(plain, []byte("D")) {
		t.Fatalf("invalid decryption: %s", plain)
	}
}

func BenchmarkPolybiusEncrypt(b *testing.B) {
	c, _ := NewPolybius(
		[]byte("CRYPTOGRAPHY"),
		PolybiusWithAlphabet([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"), map[byte]byte{}),
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Encrypt([]byte("INCRYPTOGRAPHYASUBSTITUTIONCIPHERISAMETHODOFENCRYPTINGINWHICHUNITSOFPLAINTEXTAREREPLACEDWITHTHECIPHERTEXTINADEFINEDMANNERWITHTHEHELPOFAKEY"))
	}
}

func BenchmarkPolybiusDecrypt(b *testing.B) {
	c, _ := NewPolybius(
		[]byte("CRYPTOGRAPHY"),
		PolybiusWithAlphabet([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"), map[byte]byte{}),
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Decrypt([]byte("233213365134423321361134225111414312414223424342233332132334221536234111311542223314331615321336513442233221233245222313224332234241331634261123324215464211361536153426111315144523422242221513233422153642154642233211141516233215143111323215364523422242221522152634331611251551"))
	}
}
