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
