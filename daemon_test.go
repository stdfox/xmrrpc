package xmrrpc

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var statusOkResponse = &StatusResponse{
	Status: "OK",
}

var statusErrorResponse = &jsonRPCError{
	Code:    -7,
	Message: "Block not accepted",
}

type daemonClientTestSuite struct {
	suite.Suite
	ts *httptest.Server
}

func (s *daemonClientTestSuite) SetupTest() {
	s.ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var res []byte

		switch r.RequestURI {
		case "/json_rpc":
			req := &jsonRPCRequest{}
			err := json.NewDecoder(r.Body).Decode(&req)
			if assert.NoError(s.T(), err) {
				switch req.Method {
				case "on_get_block_hash":
					res, _ = json.Marshal("e22cf75f39ae720e8b71b3d120a5ac03f0db50bba6379e2850975b4859190bc6")
					res, _ = json.Marshal(&jsonRPCResponse{ID: req.ID, Version: "2.0", Result: res})
					break
				case "submit_block":
					res, _ = json.Marshal(&jsonRPCResponse{ID: req.ID, Version: "2.0", Error: *statusErrorResponse})
					break
				default:
					res, _ = json.Marshal(statusOkResponse)
					res, _ = json.Marshal(&jsonRPCResponse{ID: req.ID, Version: "2.0", Result: res})
					break
				}
			}
			break
		default:
			res, _ = json.Marshal(statusOkResponse)
			break
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}))
}

func TestDaemonClientTestSuite(t *testing.T) {
	suite.Run(t, new(daemonClientTestSuite))
}

func (s *daemonClientTestSuite) TestGetBlockCount() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetBlockCount()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestOnGetBlockHash() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").OnGetBlockHash(912345)
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "e22cf75f39ae720e8b71b3d120a5ac03f0db50bba6379e2850975b4859190bc6", res)
	}
}

func (s *daemonClientTestSuite) TestGetBlockTemplate() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetBlockTemplate("44GBHzv6ZyQdJ...", 60)
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestSubmitBlock() {
	_, err := NewDaemonClient(s.ts.URL, "username", "password").SubmitBlock("0707e6bdfedc0...")
	if assert.Error(s.T(), err) {
		assert.EqualError(s.T(), err, "Block not accepted")
	}
}

func (s *daemonClientTestSuite) TestGetLastBlockHeader() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetLastBlockHeader()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetBlockHeaderByHash() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetBlockHeaderByHash("3a289b8fa88b1...")
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetBlockHeaderByHeight() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetBlockHeaderByHeight(1562023)
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetBlockHeadersRange() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetBlockHeadersRange(1545999, 1546000)
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetBlock() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetBlock(1562023, "3a289b8fa88b1...")
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetConnections() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetConnections()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetInfo() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetInfo()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestHardForkInfo() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").HardForkInfo()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestSetBans() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").SetBans([]Ban{})
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetBans() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetBans()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestFlushTxpool() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").FlushTxpool([]string{})
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetOutputHistogram() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetOutputHistogram([]uint{100000000}, 0, 0, false, 0)
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetVersion() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetVersion()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetCoinbaseTxSum() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetCoinbaseTxSum(1000, 1)
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetFeeEstimate() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetFeeEstimate(1)
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetAlternateChains() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetAlternateChains()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestRelayTx() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").RelayTx([]string{""})
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestSyncInfo() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").SyncInfo()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetTxpoolBacklog() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetTxpoolBacklog()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetOutputDistribution() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetOutputDistribution([]uint{1000}, false, 1, 10)
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetHeight() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetHeight()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetTransactions() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetTransactions([]string{}, false, false)
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetAltBlocksHashes() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetAltBlocksHashes()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestIsKeyImageSpent() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").IsKeyImageSpent([]string{})
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestSendRawTransaction() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").SendRawTransaction("", false)
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestStartMining() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").StartMining(false, false, "", 0)
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestStopMining() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").StopMining()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestMiningStatus() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").MiningStatus()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestSaveBC() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").SaveBC()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetPeerList() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetPeerList()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestSetLogHashRate() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").SetLogHashRate(false)
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestSetLogLevel() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").SetLogLevel(0)
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestSetLogCategories() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").SetLogCategories("")
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetTransactionPool() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetTransactionPool()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetTransactionPoolStats() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetTransactionPoolStats()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestStopDaemon() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").StopDaemon()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestGetLimit() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").GetLimit()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestSetLimit() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").SetLimit(0, 0)
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestOutPeers() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").OutPeers(0)
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestInPeers() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").InPeers(0)
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestStartSaveGraph() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").StartSaveGraph()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestStopSaveGraph() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").StopSaveGraph()
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}

func (s *daemonClientTestSuite) TestUpdate() {
	res, err := NewDaemonClient(s.ts.URL, "username", "password").Update("", "")
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "OK", res.Status)
	}
}
