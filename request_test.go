package xmrrpc

import (
	"encoding/base64"
	"testing"
)

func TestRandomKey(t *testing.T) {
	k := randomKey()

	if len(k) != 12 {
		t.Errorf("Key length was incorrect, got: %d, expected: %d.", len(k), 12)
	}

	d, err := base64.StdEncoding.DecodeString(k)
	if err != nil {
		t.Errorf("Can't decode base64 random key, got: %s.", string(d))
	}

	if len(d) != 8 {
		t.Errorf("Key random length was incorrect, got: %d, expected: %d.", len(d), 8)
	}
}

func TestH(t *testing.T) {
	h := h("1234567890")

	if h != "e807f1fcf82d132f9bb018ca6738a19f" {
		t.Errorf("MD5 hash was incorrect, got: %s, expected: %s.", h, "e807f1fcf82d132f9bb018ca6738a19f")
	}
}
