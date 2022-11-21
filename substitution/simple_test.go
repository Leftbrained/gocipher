package substitution

import "testing"

func TestSimpleInstatiation(t *testing.T) {
	plain, cipher := []byte("cBda"), []byte("aDcB")

	c, err := NewSimple(plain, cipher)

	if c == nil || err != nil {
		t.Fatalf("NewSimple() could not instantate")
	}
}
