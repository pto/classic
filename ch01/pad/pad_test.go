package pad

import "testing"

func TestEncryption(t *testing.T) {
	cases := []string{
		"",
		"A first test",
		"¡Vamos, Go!",
		"🚀",
		"My 🚀 is fast!",
	}
	for _, c := range cases {
		pair := Encrypt(c)
		if c != "" && string(pair.second) == c {
			t.Errorf("Encryption left %q the same as %q",
				c, string(pair.second))
		}
		if c != Decrypt(pair) {
			t.Errorf("Decrypt %#v is %q, want %q", pair, Decrypt(pair), c)
		}
	}
}
