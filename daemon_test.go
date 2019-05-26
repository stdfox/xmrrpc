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
		fmt.Fprintln(w, `{
			"id": 0,
			"jsonrpc": "2.0",
			"result": {
				"count": 993163,
				"status": "OK"
			}
		}`)
	}))
	defer ts.Close()

	res, err := NewDaemonClient(ts.URL, "", "").GetBlockCount()
	if assert.NoError(t, err) {
		assert.Equal(t, "OK", res.Status)
	}
}

func TestOnGetBlockHash(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"id": 0,
			"jsonrpc": "2.0",
			"result": "e22cf75f39ae720e8b71b3d120a5ac03f0db50bba6379e2850975b4859190bc6"
		}`)
	}))
	defer ts.Close()

	res, err := NewDaemonClient(ts.URL, "", "").OnGetBlockHash(912345)
	if assert.NoError(t, err) {
		assert.Equal(t, "e22cf75f39ae720e8b71b3d120a5ac03f0db50bba6379e2850975b4859190bc6", res)
	}
}

func TestGetBlockTemplate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"id": 0,
			"jsonrpc": "2.0",
			"result": {
				"blockhashing_blob": "070786a498d705f8dc58791266179087907a2ff4cd883615216749b97d2f12173171c725a6f84a00000000fc751ea4a94c2f840751eaa36138eee66dda15ef554e7d6594395827994e31da10",
				"blocktemplate_blob": "070786a498d705f8dc58791266179087907a2ff4cd883615216749b97d2f12173171c725a6f84a0000000002aeab5f01fff2aa5f01e0a9d0f2f08a01028fdb3d5b5a2c363d36ea17a4add99a23a3ec7935b4c3e1e0364fcc4295c7a2ef5f01f912b15f5d17c1539d4722f79d8856d8654c5af87f54cfb3a4ff7f6b512b2a08023c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f1755090c809421d69873c161e7969b8bf33cee3b451dd4859bfc244a705f0b4900498f804b6023e13fa023a0fb759e8b7c9a39506a21442bc47077beeedc6b78d34c4ebdae91bd96097ccc9a882bc5056568b0d2f1f06559368fea4acba8e745444e883e53156d5083c1fd260edf05292934c8b40c098b81fe4e261720bdd272b209e317247a1d2c55dc4718891af0d16273c5a610f36f382a3bf50f54808aaa6a508e51d4601dd0d8fbf8b3b1685066ce121666a1409e8ac7a4d673c1cc36d10b825f764af647441f53230518e4d2efbcf8791c6060912c76e90db4982a66d51bbd96290bbb34db8080b216c2940cec407260bf5e2c3a5ee280835f15298f0801e9d98c4d414792282fbc2c28c3e20bc0fcb1829b5c3ad8f8d20847be8fdb2a949fd96f0205fbd6d271c880c5d8c83e9813606cd803a44d377fdeae45bfa67112132af601e9b3b0613ba7dff2ec3d4b935c447b47bfe39f7b950981b2f4c66c0d853e2218f1f69229a9b608c3d98be09b6d4d640a9f6ff0e920dbacf7e58b59554c0b398b1ae4b1d497104b4e4e745d850eed7eddb8aa93437427bf442ae5beb22cbf10a8fa738ea38cfa5d86dfd30675d4be11a38016e36936fd5601e52643e8b8bc433702ea7ae6149309c95b898cc854850e73fe0b95c5b8879b7325ecd4",
				"difficulty": 61043624293,
				"expected_reward": 4771949057248,
				"height": 1561970,
				"prev_hash": "f8dc58791266179087907a2ff4cd883615216749b97d2f12173171c725a6f84a",
				"reserved_offset": 129,
				"status": "OK",
				"untrusted": false
			}
		}`)
	}))
	defer ts.Close()

	res, err := NewDaemonClient(ts.URL, "", "").GetBlockTemplate("44GBHzv6ZyQdJkjqZje6KLZ3xSyN1hBSFAnLP6EAqJtCRVzMzZmeXTC2AHKDS9aEDTRKmo6a6o9r9j86pYfhCWDkKjbtcns", 60)
	if assert.NoError(t, err) {
		assert.Equal(t, "OK", res.Status)
	}
}

func TestSubmitBlock(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"id": 0,
			"jsonrpc": "2.0",
			"error": {
				"code": -7,
				"message": "Block not accepted"
			}
		}`)
	}))
	defer ts.Close()

	_, err := NewDaemonClient(ts.URL, "", "").SubmitBlock("0707e6bdfedc053771512f1bc27c62731ae9e8f2443db64ce742f4e57f5cf8d393de28551e441a0000000002fb830a01ffbf830a018cfe88bee283060274c0aae2ef5730e680308d9c00b6da59187ad0352efe3c71d36eeeb28782f29f2501bd56b952c3ddc3e350c2631d3a5086cac172c56893831228b17de296ff4669de020200000000")
	if assert.Error(t, err) {
		assert.EqualError(t, err, "Block not accepted")
	}
}

func TestGetLastBlockHeader(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"id": 0,
			"jsonrpc": "2.0",
			"result": {
				"block_header": {
					"block_size": 62774,
					"depth": 0,
					"difficulty": 60097900840,
					"hash": "3a289b8fa88b10e2163826c230b45d79f2be37d14fa3153ee58ff8a427782d14",
					"height": 1562023,
					"major_version": 7,
					"minor_version": 7,
					"nonce": 3789681204,
					"num_txes": 5,
					"orphan_status": false,
					"prev_hash": "743e5d0a26849efe27b96086f2c4ecc39a0bc744bf21473dad6710221aff6ac3",
					"reward": 4724029079703,
					"timestamp": 1525029411
				},
				"status": "OK",
				"untrusted": false
			}
		}`)
	}))
	defer ts.Close()

	res, err := NewDaemonClient(ts.URL, "", "").GetLastBlockHeader()
	if assert.NoError(t, err) {
		assert.Equal(t, "OK", res.Status)
	}
}
