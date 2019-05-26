package xmrrpc

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTestServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/json_rpc" {
			w.Header().Set("Content-Type", "application/json")

			req := &jsonRPCRequest{}
			err := json.NewDecoder(r.Body).Decode(&req)
			if assert.NoError(t, err) {
				switch req.Method {
				case "get_block_count":
					io.WriteString(w, `{
						"id": 0,
						"jsonrpc": "2.0",
						"result": {
							"count": 993163,
							"status": "OK"
						}
					}`)
					break
				case "on_get_block_hash":
					io.WriteString(w, `{
						"id": 0,
						"jsonrpc": "2.0",
						"result": "e22cf75f39ae720e8b71b3d120a5ac03f0db50bba6379e2850975b4859190bc6"
					}`)
					break
				case "get_block_template":
					io.WriteString(w, `{
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
					break
				case "submit_block":
					io.WriteString(w, `{
						"id": 0,
						"jsonrpc": "2.0",
						"error": {
							"code": -7,
							"message": "Block not accepted"
						}
					}`)
					break
				case "get_last_block_header", "get_block_header_by_hash", "get_block_header_by_height":
					io.WriteString(w, `{
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
					break
				case "get_block_headers_range":
					io.WriteString(w, `{
						"id": 0,
						"jsonrpc": "2.0",
						"result": {
							"headers": [{
								"block_size": 301413,
								"depth": 16085,
								"difficulty": 134636057921,
								"hash": "86d1d20a40cefcf3dd410ff6967e0491613b77bf73ea8f1bf2e335cf9cf7d57a",
								"height": 1545999,
								"major_version": 6,
								"minor_version": 6,
								"nonce": 3246403956,
								"num_txes": 20,
								"orphan_status": false,
								"prev_hash": "0ef6e948f77b8f8806621003f5de24b1bcbea150bc0e376835aea099674a5db5",
								"reward": 5025593029981,
								"timestamp": 1523002893
							},
							{
								"block_size": 13322,
								"depth": 16084,
								"difficulty": 134716086238,
								"hash": "b408bf4cfcd7de13e7e370c84b8314c85b24f0ba4093ca1d6eeb30b35e34e91a",
								"height": 1546000,
								"major_version": 7,
								"minor_version": 7,
								"nonce": 3737164176,
								"num_txes": 1,
								"orphan_status": false,
								"prev_hash": "86d1d20a40cefcf3dd410ff6967e0491613b77bf73ea8f1bf2e335cf9cf7d57a",
								"reward": 4851952181070,
								"timestamp": 1523002931
							}],
							"status": "OK",
							"untrusted": false
						}
					}`)
					break
				case "get_block":
					io.WriteString(w, `{
						"id": 0,
						"jsonrpc": "2.0",
						"result": {
							"blob": "0102f4bedfb405b61c58b2e0be53fad5ef9d9731a55e8a81d972b8d90ed07c04fd37ca6403ff786e0600000195d83701ffd9d73704ee84ddb42102378b043c1724c92c69d923d266fe86477d3a5ddd21145062e148c64c5767700880c0fc82aa020273733cbd6e6218bda671596462a4b062f95cfe5e1dbb5b990dacb30e827d02f280f092cbdd080247a5dab669770da69a860acde21616a119818e1a489bb3c4b1b6b3c50547bc0c80e08d84ddcb01021f7e4762b8b755e3e3c72b8610cc87b9bc25d1f0a87c0c816ebb952e4f8aff3d2b01fd0a778957f4f3103a838afda488c3cdadf2697b3d34ad71234282b2fad9100e02080000000bdfc2c16c00",
							"block_header": {
								"block_size": 210,
								"depth": 649772,
								"difficulty": 815625611,
								"hash": "e22cf75f39ae720e8b71b3d120a5ac03f0db50bba6379e2850975b4859190bc6",
								"height": 912345,
								"major_version": 1,
								"minor_version": 2,
								"nonce": 1646,
								"num_txes": 0,
								"orphan_status": false,
								"prev_hash": "b61c58b2e0be53fad5ef9d9731a55e8a81d972b8d90ed07c04fd37ca6403ff78",
								"reward": 7388968946286,
								"timestamp": 1452793716
							},
							"json": "{\n  \"major_version\": 1, \n  \"minor_version\": 2, \n  \"timestamp\": 1452793716, \n  \"prev_id\": \"b61c58b2e0be53fad5ef9d9731a55e8a81d972b8d90ed07c04fd37ca6403ff78\", \n  \"nonce\": 1646, \n  \"miner_tx\": {\n    \"version\": 1, \n    \"unlock_time\": 912405, \n    \"vin\": [ {\n        \"gen\": {\n          \"height\": 912345\n        }\n      }\n    ], \n    \"vout\": [ {\n        \"amount\": 8968946286, \n        \"target\": {\n          \"key\": \"378b043c1724c92c69d923d266fe86477d3a5ddd21145062e148c64c57677008\"\n        }\n      }, {\n        \"amount\": 80000000000, \n        \"target\": {\n          \"key\": \"73733cbd6e6218bda671596462a4b062f95cfe5e1dbb5b990dacb30e827d02f2\"\n        }\n      }, {\n        \"amount\": 300000000000, \n        \"target\": {\n          \"key\": \"47a5dab669770da69a860acde21616a119818e1a489bb3c4b1b6b3c50547bc0c\"\n        }\n      }, {\n        \"amount\": 7000000000000, \n        \"target\": {\n          \"key\": \"1f7e4762b8b755e3e3c72b8610cc87b9bc25d1f0a87c0c816ebb952e4f8aff3d\"\n        }\n      }\n    ], \n    \"extra\": [ 1, 253, 10, 119, 137, 87, 244, 243, 16, 58, 131, 138, 253, 164, 136, 195, 205, 173, 242, 105, 123, 61, 52, 173, 113, 35, 66, 130, 178, 250, 217, 16, 14, 2, 8, 0, 0, 0, 11, 223, 194, 193, 108\n    ], \n    \"signatures\": [ ]\n  }, \n  \"tx_hashes\": [ ]\n}",
							"miner_tx_hash": "c7da3965f25c19b8eb7dd8db48dcd4e7c885e2491db77e289f0609bf8e08ec30",
							"status": "OK",
							"untrusted": false
						}
					}`)
					break
				case "get_connections":
					io.WriteString(w, `{
						"id": 0,
						"jsonrpc": "2.0",
						"result": {
							"connections": [{
								"address": "173.90.69.136:62950",
								"avg_download": 0,
								"avg_upload": 2,
								"connection_id": "083c301a3030329a487adb12ad981d2c",
								"current_download": 0,
								"current_upload": 2,
								"height": 1562127,
								"host": "173.90.69.136",
								"incoming": true,
								"ip": "173.90.69.136",
								"live_time": 8,
								"local_ip": false,
								"localhost": false,
								"peer_id": "c959fbfbed9e44fb",
								"port": "62950",
								"recv_count": 259,
								"recv_idle_time": 8,
								"send_count": 24342,
								"send_idle_time": 8,
								"state": "state_normal",
								"support_flags": 0
							}],
							"status": "OK"
						}
					}`)
					break
				}
			}
		}
	}))
}

func TestGetBlockCount(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	res, err := NewDaemonClient(ts.URL, "username", "password").GetBlockCount()
	if assert.NoError(t, err) {
		assert.Equal(t, "OK", res.Status)
	}
}

func TestOnGetBlockHash(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	res, err := NewDaemonClient(ts.URL, "username", "password").OnGetBlockHash(912345)
	if assert.NoError(t, err) {
		assert.Equal(t, "e22cf75f39ae720e8b71b3d120a5ac03f0db50bba6379e2850975b4859190bc6", res)
	}
}

func TestGetBlockTemplate(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	res, err := NewDaemonClient(ts.URL, "username", "password").GetBlockTemplate("44GBHzv6ZyQdJkjqZje6KLZ3xSyN1hBSFAnLP6EAqJtCRVzMzZmeXTC2AHKDS9aEDTRKmo6a6o9r9j86pYfhCWDkKjbtcns", 60)
	if assert.NoError(t, err) {
		assert.Equal(t, "OK", res.Status)
	}
}

func TestSubmitBlock(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	_, err := NewDaemonClient(ts.URL, "username", "password").SubmitBlock("0707e6bdfedc053771512f1bc27c62731ae9e8f2443db64ce742f4e57f5cf8d393de28551e441a0000000002fb830a01ffbf830a018cfe88bee283060274c0aae2ef5730e680308d9c00b6da59187ad0352efe3c71d36eeeb28782f29f2501bd56b952c3ddc3e350c2631d3a5086cac172c56893831228b17de296ff4669de020200000000")
	if assert.Error(t, err) {
		assert.EqualError(t, err, "Block not accepted")
	}
}

func TestGetLastBlockHeader(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	res, err := NewDaemonClient(ts.URL, "username", "password").GetLastBlockHeader()
	if assert.NoError(t, err) {
		assert.Equal(t, "OK", res.Status)
	}
}

func TestGetBlockHeaderByHash(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	res, err := NewDaemonClient(ts.URL, "username", "password").GetBlockHeaderByHash("3a289b8fa88b10e2163826c230b45d79f2be37d14fa3153ee58ff8a427782d14")
	if assert.NoError(t, err) {
		assert.Equal(t, "OK", res.Status)
	}
}

func TestGetBlockHeaderByHeight(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	res, err := NewDaemonClient(ts.URL, "username", "password").GetBlockHeaderByHeight(1562023)
	if assert.NoError(t, err) {
		assert.Equal(t, "OK", res.Status)
	}
}

func TestGetBlockHeadersRange(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	res, err := NewDaemonClient(ts.URL, "username", "password").GetBlockHeadersRange(1545999, 1546000)
	if assert.NoError(t, err) {
		assert.Equal(t, "OK", res.Status)
	}
}

func TestGetBlock(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	res, err := NewDaemonClient(ts.URL, "username", "password").GetBlock(1562023, "3a289b8fa88b10e2163826c230b45d79f2be37d14fa3153ee58ff8a427782d14")
	if assert.NoError(t, err) {
		assert.Equal(t, "OK", res.Status)
	}
}

func TestGetConnections(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	res, err := NewDaemonClient(ts.URL, "username", "password").GetConnections()
	if assert.NoError(t, err) {
		assert.Equal(t, "OK", res.Status)
	}
}
