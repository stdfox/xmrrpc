package xmrrpc

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type requestTestSuite struct {
	suite.Suite
}

func TestRequestTestSuite(t *testing.T) {
	suite.Run(t, new(requestTestSuite))
}

func (s *requestTestSuite) TestDigestAuthParams() {
	res1 := &httptest.ResponseRecorder{}
	res1.Header().Add("WWW-authenticate", `NotDigest qop="auth",algorithm=MD5`)
	if assert.Empty(s.T(), digestAuthParams(res1.Result())) {
		res2 := &httptest.ResponseRecorder{}
		res2.Header().Add("WWW-authenticate", `Digest qop="auth",algorithm=MD5,realm="monero-rpc",nonce,stale=false`)
		d := digestAuthParams(res2.Result())
		if assert.Empty(s.T(), d["nonce"]) {
			res3 := &httptest.ResponseRecorder{}
			res3.Header().Add("WWW-authenticate", `Digest qop="auth",algorithm=MD5,realm="monero-rpc",nonce="jmk3hzH2xpPmUBSD0uy+uQ==",stale=false`)
			d = digestAuthParams(res3.Result())
			if assert.NotEmpty(s.T(), d) {
				assert.Equal(s.T(), "jmk3hzH2xpPmUBSD0uy+uQ==", d["nonce"])
			}
		}
	}
}

func (s *requestTestSuite) TestRandomKey() {
	k := randomKey()
	if assert.NotPanics(s.T(), func() { randomKey() }) {
		if assert.Len(s.T(), k, 12, "Key length is incorrect.") {
			d, err := base64.StdEncoding.DecodeString(k)
			if assert.NoError(s.T(), err, "Key encode is incorrect.") {
				assert.Len(s.T(), d, 8, "Key bytes length is incorrect.")
			}
		}
	}
}

func (s *requestTestSuite) TestH() {
	if assert.NotPanics(s.T(), func() { h("1234567890") }) {
		assert.Equal(s.T(), h("1234567890"), "e807f1fcf82d132f9bb018ca6738a19f", "MD5 hash is incorrect.")
	}
}

func (s *requestTestSuite) TestRequest() {
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

	_, err := request(http.MethodPost, "", nil, "username", "password")
	assert.Error(s.T(), err)

	res, err := request(http.MethodPost, ts.URL, nil, "username", "password")
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), http.StatusOK, res.StatusCode)
	}
}
