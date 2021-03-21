package trivium

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestTrivium_Encrypt(t1 *testing.T) {
	want, _ := hex.DecodeString("6fa0e65512c187370979143a2dd4e5b44a4c5697a82e436d8734afd0903ac9dd1f")
	iv := []byte("security20")
	key := []byte("cryptograp")
	content := []byte("Hello world 2021, please be nice!")

	t1.Run("Test encryption", func(t1 *testing.T) {
		t := NewTrivium(iv, key)
		got := t.Encrypt(content)
		if bytes.Compare(got, want) != 0 {
			t1.Errorf("Encrypt() = %v, want %v", hex.EncodeToString(got), hex.EncodeToString(want))
		}
	})
}
