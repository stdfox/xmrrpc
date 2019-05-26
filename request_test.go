package xmrrpc

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomKey(t *testing.T) {
	k := randomKey()
	assert.Len(t, k, 12, "Key length is incorrect.")

	d, err := base64.StdEncoding.DecodeString(k)
	assert.Error(t, err, "Key encode is incorrect.")
	assert.Len(t, d, 8, "Key bytes length is incorrect.")
}

func TestH(t *testing.T) {
	assert.Equal(t, h("1234567890"), "e807f1fcf82d132f9bb018ca6738a19f1", "MD5 hash is incorrect.")
}
