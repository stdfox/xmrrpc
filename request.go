package xmrrpc

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
)

func digestAuthParams(response *http.Response) map[string]string {
	s := strings.SplitN(response.Header.Get("Www-Authenticate"), " ", 2)
	if len(s) != 2 || s[0] != "Digest" {
		return nil
	}

	result := map[string]string{}
	for _, kv := range strings.Split(s[1], ",") {
		parts := strings.SplitN(kv, "=", 2)
		if len(parts) != 2 {
			continue
		}

		result[strings.Trim(parts[0], "\" ")] = strings.Trim(parts[1], "\" ")
	}

	return result
}

func randomKey() string {
	k := make([]byte, 8)
	for bytes := 0; bytes < len(k); {
		n, err := rand.Read(k[bytes:])
		if err != nil {
			panic("rand.Read() failed")
		}

		bytes += n
	}

	return base64.StdEncoding.EncodeToString(k)
}

func h(s string) string {
	digest := md5.New()
	digest.Write([]byte(s))

	return hex.EncodeToString(digest.Sum(nil))
}

func request(method string, url string, body []byte, username string, password string) (*http.Response, error) {
	req1, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req1.Header.Set("Content-Type", "application/json")

	client1 := &http.Client{}
	res1, err := client1.Do(req1)
	if err != nil {
		return nil, err
	}

	if res1.StatusCode == http.StatusUnauthorized {
		io.Copy(ioutil.Discard, res1.Body)
		res1.Body.Close()

		var authorization map[string]string = digestAuthParams(res1)
		var realmHeader = authorization["realm"]
		var qopHeader = authorization["qop"]
		var nonceHeader = authorization["nonce"]
		var algorithm = authorization["algorithm"]
		var realm = realmHeader
		var nc = "00000001"

		hash := md5.New()
		a1 := fmt.Sprintf("%s:%s:%s", username, realm, password)
		io.WriteString(hash, a1)
		ha1 := hex.EncodeToString(hash.Sum(nil))

		hash = md5.New()
		a2 := fmt.Sprintf("%s:%s", method, "/json_rpc")
		io.WriteString(hash, a2)
		ha2 := hex.EncodeToString(hash.Sum(nil))

		cnonce := randomKey()
		response := h(strings.Join([]string{ha1, nonceHeader, nc, cnonce, qopHeader, ha2}, ":"))
		authHeader := fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s", algorithm="%s", response="%s", qop=%s, nc=%s, cnonce="%s"`, username, realmHeader, nonceHeader, "/json_rpc", algorithm, response, qopHeader, nc, cnonce)

		req2, err := http.NewRequest(method, url, bytes.NewReader(body))
		if err != nil {
			return nil, err
		}
		req2.Header.Set("Content-Type", "application/json")
		req2.Header.Set("Authorization", authHeader)

		client2 := &http.Client{}
		res2, err := client2.Do(req2)
		if err != nil {
			return nil, err
		}

		return res2, nil
	}

	return res1, nil
}