package xmrrpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
)

type DaemonClient struct {
	endpoint string
	username string
	password string
}

type jsonRPCRequest struct {
	Version string      `json:"jsonrpc"`
	ID      uint64      `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type jsonRPCResponse struct {
	Version string           `json:"jsonrpc"`
	ID      uint64           `json:"id"`
	Result  *json.RawMessage `json:"result"`
	Error   jsonRPCError     `json:"error"`
}

type jsonRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

type BlockCountResponse struct {
	Count  uint   `json:"count"`
	Status string `json:"status"`
}

type BlockTemplateResponse struct {
	BlockTemplateBlob string `json:"blocktemplate_blob"`
	BlockHashingBlob  string `json:"blockhashing_blob"`
	Difficulty        uint   `json:"difficulty"`
	ExpectedReward    uint   `json:"expected_reward"`
	Height            uint   `json:"height"`
	PrevHash          string `json:"prev_hash"`
	ReservedOffset    uint   `json:"reserved_offset"`
	Status            string `json:"status"`
	Untrusted         bool   `json:"untrusted"`
}

type BlockHeader struct {
	BlockSize    uint   `json:"block_size"`
	Depth        uint   `json:"depth"`
	Difficulty   uint   `json:"difficulty"`
	Hash         string `json:"hash"`
	Height       uint   `json:"height"`
	MajorVersion uint   `json:"major_version"`
	MinorVersion uint   `json:"minor_version"`
	Nonce        uint   `json:"nonce"`
	NumTxes      uint   `json:"num_txes"`
	OrphanStatus bool   `json:"orphan_status"`
	PrevHash     string `json:"prev_hash"`
	Reward       uint   `json:"reward"`
	Timestamp    uint   `json:"timestamp"`
}

type BlockHeaderResponse struct {
	BlockHeader BlockHeader `json:"block_header"`
	Status      string      `json:"status"`
	Untrusted   bool        `json:"untrusted"`
}

type BlockHeadersResponse struct {
	BlockHeader []BlockHeader `json:"headers"`
	Status      string        `json:"status"`
	Untrusted   bool          `json:"untrusted"`
}

type BlockResponse struct {
	Blob        string      `json:"blob"`
	BlockHeader BlockHeader `json:"block_header"`
	Json        string      `json:"json"`
	Status      string      `json:"status"`
	Untrusted   bool        `json:"untrusted"`
}

type Connection struct {
	Address         string `json:"address"`
	AvgDownload     uint   `json:"avg_download"`
	AvgUpload       uint   `json:"avg_upload"`
	ConnectionID    string `json:"connection_id"`
	CurrentDownload uint   `json:"current_download"`
	CurrentUpload   uint   `json:"current_upload"`
	Height          uint   `json:"height"`
	Host            string `json:"host"`
	Incoming        bool   `json:"incoming"`
	IP              string `json:"ip"`
	LiveTime        uint   `json:"live_time"`
	LocalIP         bool   `json:"local_ip"`
	Localhost       bool   `json:"localhost"`
	PeerID          string `json:"peer_id"`
	Port            string `json:"port"`
	RecvCount       uint   `json:"recv_count"`
	RecvIdleTime    uint   `json:"recv_idle_time"`
	SendCount       uint   `json:"send_count"`
	SendIdleTime    uint   `json:"send_idle_time"`
	State           string `json:"state"`
	SupportFlags    uint   `json:"support_flags"`
}

type ConnectionsResponse struct {
	Connections []Connection `json:"connections"`
	Status      string       `json:"status"`
}

type InfoResponse struct {
	AltBlocksCount           uint   `json:"alt_blocks_count"`
	BlockSizeLimit           uint   `json:"block_size_limit"`
	BlockSizeMedian          uint   `json:"block_size_median"`
	BootstrapDaemonAddress   string `json:"bootstrap_daemon_address"`
	CumulativeDifficulty     uint   `json:"cumulative_difficulty"`
	Difficulty               uint   `json:"difficulty"`
	FreeSpace                uint   `json:"free_space"`
	GreyPeerlistSize         uint   `json:"grey_peerlist_size"`
	Height                   uint   `json:"height"`
	HeightWithoutBootstrap   uint   `json:"height_without_bootstrap"`
	IncomingConnectionsCount uint   `json:"incoming_connections_count"`
	Mainnet                  bool   `json:"mainnet"`
	Offline                  bool   `json:"offline"`
	OutgoingConnectionsCount uint   `json:"outgoing_connections_count"`
	RPCConnectionsCount      uint   `json:"rpc_connections_count"`
	Stagenet                 bool   `json:"stagenet"`
	StartTime                uint   `json:"start_time"`
	Status                   string `json:"status"`
	Target                   uint   `json:"target"`
	TargetHeight             uint   `json:"target_height"`
	Testnet                  bool   `json:"testnet"`
	TopBlockHash             string `json:"top_block_hash"`
	TxCount                  uint   `json:"tx_count"`
	TxPoolSize               uint   `json:"tx_pool_size"`
	Untrusted                bool   `json:"untrusted"`
	WasBootstrapEverUsed     bool   `json:"was_bootstrap_ever_used"`
	WhitePeerlistSize        uint   `json:"white_peerlist_size"`
}

type HardForkInfoResponse struct {
	EarliestHeight uint   `json:"earliest_height"`
	Enabled        bool   `json:"enabled"`
	State          uint   `json:"state"`
	Status         string `json:"status"`
	Threshold      uint   `json:"threshold"`
	Version        uint   `json:"version"`
	Votes          uint   `json:"votes"`
	Voting         uint   `json:"voting"`
	Window         uint   `json:"window"`
}

type Ban struct {
	Host    string `json:"host"`
	IP      uint   `json:"ip"`
	Ban     bool   `json:"ban"`
	Seconds uint   `json:"seconds"`
}

type BansResponse struct {
	Bans   []Ban  `json:"bans"`
	Status string `json:"status"`
}

type Histogram struct {
	Amount            uint `json:"amount"`
	TotalInstances    uint `json:"total_instances"`
	UnlockedInstances uint `json:"unlocked_instances"`
	RecentInstances   uint `json:"recent_instances"`
}

type OutputHistogramResponse struct {
	Histogram []Histogram `json:"histogram"`
	Status    string      `json:"status"`
	Untrusted bool        `json:"untrusted"`
}

type VersionResponse struct {
	Version   uint   `json:"version"`
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

func NewDaemonClient(endpoint string, username string, password string) *DaemonClient {
	return &DaemonClient{endpoint: endpoint, username: username, password: password}
}

func (dc *DaemonClient) jsonRPCRequest(method string, params interface{}, reply interface{}) error {
	jsonRPCRequest := &jsonRPCRequest{
		Version: "2.0",
		ID:      rand.Uint64(),
		Method:  method,
		Params:  params,
	}

	jsonRPCRequestBody, err := json.Marshal(jsonRPCRequest)
	if err != nil {
		return err
	}

	httpRequest, err := http.NewRequest(http.MethodPost, dc.endpoint, bytes.NewReader(jsonRPCRequestBody))
	if err != nil {
		return err
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	jsonRPCResponse := &jsonRPCResponse{}
	if err := json.NewDecoder(httpResponse.Body).Decode(&jsonRPCResponse); err != nil {
		return err
	}

	if jsonRPCResponse.Error.Code < 0 {
		return errors.New(jsonRPCResponse.Error.Message)
	}

	if jsonRPCResponse.Result == nil {
		return errors.New("Unexpected null result")
	}

	return json.Unmarshal(*jsonRPCResponse.Result, reply)
}

func (dc *DaemonClient) GetBlockCount() (BlockCountResponse, error) {
	var blockCountResponse BlockCountResponse
	err := dc.jsonRPCRequest("get_block_count", nil, &blockCountResponse)

	return blockCountResponse, err
}

func (dc *DaemonClient) OnGetBlockHash(blockHeight int) (string, error) {
	var blockHash string
	err := dc.jsonRPCRequest("on_get_block_hash", []int{blockHeight}, &blockHash)

	return blockHash, err
}

func (dc *DaemonClient) GetBlockTemplate(walletAddress string, reserveSize uint) (BlockTemplateResponse, error) {
	var blockTemplateResponse BlockTemplateResponse

	type jsonRPCParams struct {
		WalletAddress string `json:"wallet_address"`
		ReserveSize   uint   `json:"reserve_size"`
	}

	params := jsonRPCParams{WalletAddress: walletAddress, ReserveSize: reserveSize}
	err := dc.jsonRPCRequest("get_block_template", params, &blockTemplateResponse)

	return blockTemplateResponse, err
}

func (dc *DaemonClient) SubmitBlock(blockBlobData string) (string, error) {
	var status string
	err := dc.jsonRPCRequest("submit_block", []string{blockBlobData}, &status)

	return status, err
}

func (dc *DaemonClient) GetLastBlockHeader() (BlockHeaderResponse, error) {
	var blockHeaderResponse BlockHeaderResponse
	err := dc.jsonRPCRequest("get_last_block_header", nil, &blockHeaderResponse)

	return blockHeaderResponse, err
}

func (dc *DaemonClient) GetBlockHeaderByHash(hash string) (BlockHeaderResponse, error) {
	var blockHeaderResponse BlockHeaderResponse

	type jsonRPCParams struct {
		Hash string `json:"hash"`
	}

	params := jsonRPCParams{Hash: hash}
	err := dc.jsonRPCRequest("get_block_header_by_hash", params, &blockHeaderResponse)

	return blockHeaderResponse, err
}

func (dc *DaemonClient) GetBlockHeaderByHeight(height uint) (BlockHeaderResponse, error) {
	var blockHeaderResponse BlockHeaderResponse

	type jsonRPCParams struct {
		Height uint `json:"height"`
	}

	params := jsonRPCParams{Height: height}
	err := dc.jsonRPCRequest("get_block_header_by_height", params, &blockHeaderResponse)

	return blockHeaderResponse, err
}

func (dc *DaemonClient) GetBlockHeadersRange(startHeight uint, endHeight uint) (BlockHeadersResponse, error) {
	var blockHeadersResponse BlockHeadersResponse

	type jsonRPCParams struct {
		StartHeight uint `json:"start_height"`
		EndHeight   uint `json:"end_height"`
	}

	params := jsonRPCParams{StartHeight: startHeight, EndHeight: endHeight}
	err := dc.jsonRPCRequest("get_block_headers_range", params, &blockHeadersResponse)

	return blockHeadersResponse, err
}

func (dc *DaemonClient) GetBlock(height uint, hash string) (BlockResponse, error) {
	var blockResponse BlockResponse

	type jsonRPCParams struct {
		Height uint   `json:"height"`
		Hash   string `json:"hash"`
	}

	params := jsonRPCParams{Height: height, Hash: hash}
	err := dc.jsonRPCRequest("get_block", params, &blockResponse)

	return blockResponse, err
}

func (dc *DaemonClient) GetConnections() (ConnectionsResponse, error) {
	var connectionsResponse ConnectionsResponse
	err := dc.jsonRPCRequest("get_connections", nil, &connectionsResponse)

	return connectionsResponse, err
}

func (dc *DaemonClient) GetInfo() (InfoResponse, error) {
	var infoResponse InfoResponse
	err := dc.jsonRPCRequest("get_info", nil, &infoResponse)

	return infoResponse, err
}

func (dc *DaemonClient) HardForkInfo() (HardForkInfoResponse, error) {
	var hardForkInfoResponse HardForkInfoResponse
	err := dc.jsonRPCRequest("hard_fork_info", nil, &hardForkInfoResponse)

	return hardForkInfoResponse, err
}

func (dc *DaemonClient) SetBans(bans []Ban) (StatusResponse, error) {
	var statusResponse StatusResponse

	type jsonRPCParams struct {
		Bans []Ban `json:"bans"`
	}

	params := jsonRPCParams{Bans: bans}
	err := dc.jsonRPCRequest("set_bans", params, &statusResponse)

	return statusResponse, err
}

func (dc *DaemonClient) GetBans() (BansResponse, error) {
	var bansResponse BansResponse
	err := dc.jsonRPCRequest("get_bans", nil, &bansResponse)

	return bansResponse, err
}

func (dc *DaemonClient) FlushTxpool(txids []string) (StatusResponse, error) {
	var statusResponse StatusResponse

	type jsonRPCParams struct {
		TxIDs []string `json:"txids"`
	}

	params := jsonRPCParams{TxIDs: txids}
	err := dc.jsonRPCRequest("flush_txpool", params, &statusResponse)

	return statusResponse, err
}

func (dc *DaemonClient) GetOutputHistogram(amounts []uint, minCount uint, maxCount uint, unlocked bool, recentCutoff uint) (OutputHistogramResponse, error) {
	var outputHistogramResponse OutputHistogramResponse

	type jsonRPCParams struct {
		Amounts      []uint `json:"amounts"`
		MinCount     uint   `json:"min_count"`
		MaxCount     uint   `json:"max_count"`
		Unlocked     bool   `json:"unlocked"`
		RecentCutoff uint   `json:"recent_cutoff"`
	}

	params := jsonRPCParams{Amounts: amounts, MinCount: minCount, MaxCount: maxCount, Unlocked: unlocked, RecentCutoff: recentCutoff}
	err := dc.jsonRPCRequest("get_output_histogram", params, &outputHistogramResponse)

	return outputHistogramResponse, err
}

func (dc *DaemonClient) GetVersion() (VersionResponse, error) {
	var versionResponse VersionResponse
	err := dc.jsonRPCRequest("get_version", nil, &versionResponse)

	return versionResponse, err
}
