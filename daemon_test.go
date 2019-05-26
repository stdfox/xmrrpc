package xmrrpc

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBlockCount(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"id": 0," jsonrpc": "2.0", "result": {"count": 993163, "status": "OK"}}`)
	}))
	defer ts.Close()

	c := NewDaemonClient(ts.URL, "", "")
	r, err := c.GetBlockCount()

	assert.NoError(t, err)
	assert.Equal(t, "OK", r.Status)
}
