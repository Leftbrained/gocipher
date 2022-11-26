package gocipher

import (
	"bytes"
	"testing"
)

func TestPolybiusNewFive(t *testing.T) {
	c, err := NewPolybius([]byte("ABCDEFGHIKLMNOPQRSTUVWXYZ"))

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate`)
	}
}

func TestPolybiusNewSix(t *testing.T) {
	c, err := NewPolybius([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"))

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate`)
	}
}

func TestPolybiusNewErrorNotSquare(t *testing.T) {
	c, err := NewPolybius([]byte("ABCDEFGHIKLMNOPQRSTUVWXY"))

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestPolybiusBasicCrypt(t *testing.T) {
	c, err := NewPolybius([]byte("ABCDEFGHIKLMNOPQRSTUVWXYZ"))
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	testCipherCrypt(c, t,
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
		[]byte("44231541452413251242345233213453453235433451154244231531115554143422"),
		[]byte("THEQUICKBROWNFOXUMPSOVERTHELAZYDOG"),
	)
}

func TestPolybiusExtraLetterDecrypt(t *testing.T) {
	c, err := NewPolybius([]byte("ABCDEFGHIKLMNOPQRSTUVWXYZ"))
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	plain := c.Decrypt([]byte("142"))

	if !bytes.Equal(plain, []byte("D")) {
		t.Fatalf("invalid decryption: %s", plain)
	}
}

func BenchmarkPolybiusEncrypt(b *testing.B) {
	c, _ := NewPolybius([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Encrypt([]byte("INCRYPTOGRAPHYASUBSTITUTIONCIPHERISAMETHODOFENCRYPTINGINWHICHUNITSOFPLAINTEXTAREREPLACEDWITHTHECIPHERTEXTINADEFINEDMANNERWITHTHEHELPOFAKEY"))
	}
}

func BenchmarkPolybiusDecrypt(b *testing.B) {
	c, _ := NewPolybius([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Decrypt([]byte("233213365134423321361134225111414312414223424342233332132334221536234111311542223314331615321336513442233221233245222313224332234241331634261123324215464211361536153426111315144523422242221513233422153642154642233211141516233215143111323215364523422242221522152634331611251551"))
	}
}
