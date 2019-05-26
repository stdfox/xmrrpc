package xmrrpc

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomKey(t *testing.T) {
	assert.NotPanics(t, func() { randomKey() })

	k := randomKey()
	assert.Len(t, k, 12, "Key length is incorrect.")

	d, err := base64.StdEncoding.DecodeString(k)
	if assert.NoError(t, err, "Key encode is incorrect.") {
		assert.Len(t, d, 8, "Key bytes length is incorrect.")
	}
}

func TestH(t *testing.T) {
	assert.NotPanics(t, func() { h("") })
	assert.Equal(t, h("1234567890"), "e807f1fcf82d132f9bb018ca6738a19f", "MD5 hash is incorrect.")
}

func TestRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("WWW-authenticate", `Digest qop="auth",algorithm=MD5,realm="monero-rpc",nonce="jmk3hzH2xpPmUBSD0uy+uQ==",stale=false`)
		w.Header().Add("WWW-authenticate", `Digest qop="auth",algorithm=MD5-sess,realm="monero-rpc",nonce="jmk3hzH2xpPmUBSD0uy+uQ==",stale=false`)

		if r.Header.Get("Authorization") == "" {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer ts.Close()

	res, err := request(http.MethodPost, ts.URL, nil, "username", "password")
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, res.StatusCode)
	}
}
