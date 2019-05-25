package xmrrpc

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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
	Version string          `json:"jsonrpc"`
	ID      uint64          `json:"id"`
	Result  json.RawMessage `json:"result"`
	Error   jsonRPCError    `json:"error"`
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
	JSON        string      `json:"json"`
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

type CoinbaseTxSumResponse struct {
	EmissionAmount uint   `json:"emission_amount"`
	FeeAmount      uint   `json:"fee_amount"`
	Status         string `json:"status"`
}

type FeeEstimateResponse struct {
	Fee       uint   `json:"fee"`
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

type Chain struct {
	BlockHash  string `json:"block_hash"`
	Difficulty uint   `json:"difficulty"`
	Height     uint   `json:"height"`
	Length     uint   `json:"length"`
}

type AlternateChainsResponse struct {
	Chains []Chain `json:"chains"`
	Status string  `json:"status"`
}

type Peers struct {
	Info Connection `json:"info"`
}

type Span struct {
	ConnectionID     string `json:"connection_id"`
	NBlocks          uint   `json:"nblocks"`
	Rate             uint   `json:"rate"`
	RemoteAddress    string `json:"remote_address"`
	Size             uint   `json:"size"`
	Speed            uint   `json:"speed"`
	StartBlockHeight uint   `json:"start_block_height"`
}

type SyncInfoResponse struct {
	Height       uint    `json:"height"`
	Peers        []Peers `json:"peers"`
	Spans        []Span  `json:"spans"`
	Status       string  `json:"status"`
	TargetHeight uint    `json:"target_height"`
}

type TxpoolBacklogResponse struct {
	Backlog   []byte `json:"backlog"`
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

type Distribution struct {
	Amount       uint   `json:"amount"`
	Base         uint   `json:"base"`
	Binary       bool   `json:"binary"`
	Distribution []uint `json:"distribution"`
	StartHeight  uint   `json:"start_height"`
}

type OutputDistributionResponse struct {
	Distributions []Distribution `json:"distributions"`
	Status        string         `json:"status"`
	Untrusted     bool           `json:"untrusted"`
}

type HeightResponse struct {
	Height    uint   `json:"height"`
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

type TransactionEntry struct {
	AsHex           string `json:"as_hex"`
	AsJSON          string `json:"as_json"`
	BlockHeight     uint   `json:"block_height"`
	BlockTimestamp  uint   `json:"block_timestamp"`
	DoubleSpendSeen bool   `json:"double_spend_seen"`
	InPool          bool   `json:"in_pool"`
	OutputIndices   []uint `json:"output_indices"`
	TxHash          string `json:"tx_hash"`
}

type TransactionsResponse struct {
	MissedTx  []string           `json:"missed_tx"`
	Status    string             `json:"status"`
	Txs       []TransactionEntry `json:"txs"`
	TxsAsHex  []string           `json:"txs_as_hex"`
	TxsAsJSON []string           `json:"txs_as_json"`
}

type AltBlocksHashesResponse struct {
	BlksHashes []string `json:"blks_hashes"`
	Status     string   `json:"status"`
	Untrusted  bool     `json:"untrusted"`
}

type IsKeyImageSpentResponse struct {
	SpentStatus []uint `json:"spent_status"`
	Status      string `json:"status"`
	Untrusted   bool   `json:"untrusted"`
}

type SendRawTransactionResponse struct {
	DoubleSpend   bool   `json:"double_spend"`
	FeeTooLow     bool   `json:"fee_too_low"`
	InvalidInput  bool   `json:"invalid_input"`
	InvalidOutput bool   `json:"invalid_output"`
	LowMixin      bool   `json:"low_mixin"`
	NotRct        bool   `json:"not_rct"`
	NotRelayed    bool   `json:"not_relayed"`
	Overspend     bool   `json:"overspend"`
	Reason        string `json:"reason"`
	Status        string `json:"status"`
	TooBig        bool   `json:"too_big"`
	Untrusted     bool   `json:"untrusted"`
}

type MiningStatusResponse struct {
	Active                    bool   `json:"active"`
	Address                   string `json:"address"`
	IsBackgroundMiningEnabled bool   `json:"is_background_mining_enabled"`
	Speed                     uint   `json:"speed"`
	Status                    string `json:"status"`
	ThreadsCount              uint   `json:"threads_count"`
}

type Peer struct {
	Host     string `json:"host"`
	ID       uint   `json:"id"`
	IP       uint   `json:"ip"`
	LastSeen uint   `json:"last_seen"`
	Port     uint   `json:"port"`
}

type PeerListResponse struct {
	GrayList  []Peer `json:"gray_list"`
	Status    string `json:"status"`
	WhiteList []Peer `json:"white_list"`
}

type LogCategoriesResponse struct {
	Categories string `json:"categories"`
	Status     string `json:"status"`
}

type TxPoolHisto struct {
	Txs   uint `json:"txs"`
	Bytes uint `json:"bytes"`
}

type PoolStats struct {
	BytesMax        uint        `json:"bytes_max"`
	BytesMed        uint        `json:"bytes_med"`
	BytesMin        uint        `json:"bytes_min"`
	BytesTotal      uint        `json:"bytes_total"`
	Histo           TxPoolHisto `json:"histo"`
	Histo98pc       uint        `json:"histo_98pc"`
	Num10m          uint        `json:"num_10m"`
	NumDoubleSpends uint        `json:"num_double_spends"`
	NumFailing      uint        `json:"num_failing"`
	NumNotRelayed   uint        `json:"num_not_relayed"`
	Oldest          uint        `json:"oldest"`
	TxsTotal        uint        `json:"txs_total"`
}

type SpentKeyImages struct {
	IDHash    string   `json:"id_hash"`
	TxsHashes []string `json:"txs_hashes"`
}

type Transactions struct {
	BlobSize           uint   `json:"blob_size"`
	DoubleSpendSeen    bool   `json:"double_spend_seen"`
	DoNotRelay         bool   `json:"do_not_relay"`
	Fee                uint   `json:"fee"`
	IDHash             string `json:"id_hash"`
	KeptByBlock        bool   `json:"kept_by_block"`
	LastFailedHeight   uint   `json:"last_failed_height"`
	LastFailedIDHash   string `json:"last_failed_id_hash"`
	LastRelayedTime    uint   `json:"last_relayed_time"`
	MaxUsedBlockHeight uint   `json:"max_used_block_height"`
	MaxUsedBlockHash   string `json:"max_used_block_hash"`
	ReceiveTime        uint   `json:"receive_time"`
	Relayed            bool   `json:"relayed"`
	TxBlob             string `json:"tx_blob"`
	TxJSON             string `json:"tx_json"`
}

type TransactionPoolResponse struct {
	SpentKeyImages []SpentKeyImages `json:"spent_key_images"`
	Status         string           `json:"status"`
	Transactions   []Transactions   `json:"transactions"`
}

type TransactionPoolStatsResponse struct {
	PoolStats PoolStats `json:"pool_stats"`
	Status    string    `json:"status"`
	Untrusted bool      `json:"untrusted"`
}

type LimitResponse struct {
	LimitDown uint   `json:"limit_down"`
	LimitUp   uint   `json:"limit_up"`
	Status    string `json:"status"`
	Untrusted bool   `json:"untrusted"`
}

type UpdateResponse struct {
	AutoURI string `json:"auto_uri"`
	Hash    string `json:"hash"`
	Path    string `json:"path"`
	Status  string `json:"status"`
	Update  bool   `json:"update"`
	UserURI string `json:"user_uri"`
	Version string `json:"version"`
}

func NewDaemonClient(endpoint string, username string, password string) *DaemonClient {
	return &DaemonClient{endpoint: endpoint, username: username, password: password}
}

func (dc *DaemonClient) jsonRequest(method string, args interface{}, reply interface{}) error {
	params := &jsonRPCRequest{
		Version: "2.0",
		ID:      rand.Uint64(),
		Method:  method,
		Params:  args,
	}

	res := &jsonRPCResponse{}
	if err := dc.rpcRequest("/json_rpc", params, res); err != nil {
		return err
	}

	if res.Error.Code < 0 {
		return errors.New(res.Error.Message)
	}

	if res.Result == nil {
		return errors.New("Unexpected null result")
	}

	return json.Unmarshal(res.Result, reply)
}

func (dc *DaemonClient) rpcRequest(method string, args interface{}, reply interface{}) error {
	body, err := json.Marshal(args)
	if err != nil {
		return err
	}

	res, err := request(http.MethodPost, dc.endpoint+method, body, dc.username, dc.password)
	if err != nil {
		return err
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	res.Body.Close()

	return json.Unmarshal(body, reply)
}

func (dc *DaemonClient) GetBlockCount() (response BlockCountResponse, err error) {
	return response, dc.jsonRequest("get_block_count", nil, &response)
}

func (dc *DaemonClient) OnGetBlockHash(blockHeight int) (response string, err error) {
	return response, dc.jsonRequest("on_get_block_hash", []int{blockHeight}, &response)
}

func (dc *DaemonClient) GetBlockTemplate(walletAddress string, reserveSize uint) (response BlockTemplateResponse, err error) {
	type Params struct {
		WalletAddress string `json:"wallet_address"`
		ReserveSize   uint   `json:"reserve_size"`
	}

	params := Params{WalletAddress: walletAddress, ReserveSize: reserveSize}
	return response, dc.jsonRequest("get_block_template", params, &response)
}

func (dc *DaemonClient) SubmitBlock(blockBlobData string) (response string, err error) {
	return response, dc.jsonRequest("submit_block", []string{blockBlobData}, &response)
}

func (dc *DaemonClient) GetLastBlockHeader() (response BlockHeaderResponse, err error) {
	return response, dc.jsonRequest("get_last_block_header", nil, &response)
}

func (dc *DaemonClient) GetBlockHeaderByHash(hash string) (response BlockHeaderResponse, err error) {
	type Params struct {
		Hash string `json:"hash"`
	}

	params := Params{Hash: hash}
	return response, dc.jsonRequest("get_block_header_by_hash", params, &response)
}

func (dc *DaemonClient) GetBlockHeaderByHeight(height uint) (response BlockHeaderResponse, err error) {
	type Params struct {
		Height uint `json:"height"`
	}

	params := Params{Height: height}
	return response, dc.jsonRequest("get_block_header_by_height", params, &response)
}

func (dc *DaemonClient) GetBlockHeadersRange(startHeight uint, endHeight uint) (response BlockHeadersResponse, err error) {
	type Params struct {
		StartHeight uint `json:"start_height"`
		EndHeight   uint `json:"end_height"`
	}

	params := Params{StartHeight: startHeight, EndHeight: endHeight}
	return response, dc.jsonRequest("get_block_headers_range", params, &response)
}

func (dc *DaemonClient) GetBlock(height uint, hash string) (response BlockResponse, err error) {
	type Params struct {
		Height uint   `json:"height"`
		Hash   string `json:"hash"`
	}

	params := Params{Height: height, Hash: hash}
	return response, dc.jsonRequest("get_block", params, &response)
}

func (dc *DaemonClient) GetConnections() (response ConnectionsResponse, err error) {
	return response, dc.jsonRequest("get_connections", nil, &response)
}

func (dc *DaemonClient) GetInfo() (response InfoResponse, err error) {
	return response, dc.jsonRequest("get_info", nil, &response)
}

func (dc *DaemonClient) HardForkInfo() (response HardForkInfoResponse, err error) {
	return response, dc.jsonRequest("hard_fork_info", nil, &response)
}

func (dc *DaemonClient) SetBans(bans []Ban) (response StatusResponse, err error) {
	type Params struct {
		Bans []Ban `json:"bans"`
	}

	params := Params{Bans: bans}
	return response, dc.jsonRequest("set_bans", params, &response)
}

func (dc *DaemonClient) GetBans() (response BansResponse, err error) {
	return response, dc.jsonRequest("get_bans", nil, &response)
}

func (dc *DaemonClient) FlushTxpool(txids []string) (response StatusResponse, err error) {
	type Params struct {
		TxIDs []string `json:"txids"`
	}

	params := Params{TxIDs: txids}
	return response, dc.jsonRequest("flush_txpool", params, &response)
}

func (dc *DaemonClient) GetOutputHistogram(amounts []uint, minCount uint, maxCount uint, unlocked bool, recentCutoff uint) (response OutputHistogramResponse, err error) {
	type Params struct {
		Amounts      []uint `json:"amounts"`
		MinCount     uint   `json:"min_count"`
		MaxCount     uint   `json:"max_count"`
		Unlocked     bool   `json:"unlocked"`
		RecentCutoff uint   `json:"recent_cutoff"`
	}

	params := Params{Amounts: amounts, MinCount: minCount, MaxCount: maxCount, Unlocked: unlocked, RecentCutoff: recentCutoff}
	return response, dc.jsonRequest("get_output_histogram", params, &response)
}

func (dc *DaemonClient) GetVersion() (response VersionResponse, err error) {
	return response, dc.jsonRequest("get_version", nil, &response)
}

func (dc *DaemonClient) GetCoinbaseTxSum(height uint, count uint) (response CoinbaseTxSumResponse, err error) {
	type Params struct {
		Height uint `json:"height"`
		Count  uint `json:"count"`
	}

	params := Params{Height: height, Count: count}
	return response, dc.jsonRequest("get_coinbase_tx_sum", params, &response)
}

func (dc *DaemonClient) GetFeeEstimate(graceBlocks uint) (response FeeEstimateResponse, err error) {
	type Params struct {
		GraceBlocks uint `json:"grace_blocks"`
	}

	params := Params{GraceBlocks: graceBlocks}
	return response, dc.jsonRequest("get_fee_estimate", params, &response)
}

func (dc *DaemonClient) GetAlternateChains() (response AlternateChainsResponse, err error) {
	return response, dc.jsonRequest("get_alternate_chains", nil, &response)
}

func (dc *DaemonClient) RelayTx(txids []string) (response StatusResponse, err error) {
	type Params struct {
		TxIDs []string `json:"txids"`
	}

	params := Params{TxIDs: txids}
	return response, dc.jsonRequest("relay_tx", params, &response)
}

func (dc *DaemonClient) SyncInfo() (response SyncInfoResponse, err error) {
	return response, dc.jsonRequest("sync_info", nil, &response)
}

func (dc *DaemonClient) GetTxpoolBacklog() (response TxpoolBacklogResponse, err error) {
	return response, dc.jsonRequest("get_txpool_backlog", nil, &response)
}

func (dc *DaemonClient) GetOutputDistribution(amounts []uint, cumulative bool, fromHeight uint, toHeight uint) (response OutputDistributionResponse, err error) {
	type Params struct {
		Amounts    []uint `json:"amounts"`
		Cumulative bool   `json:"cumulative"`
		FromHeight uint   `json:"from_height"`
		ToHeight   uint   `json:"to_height"`
	}

	params := Params{Amounts: amounts, Cumulative: cumulative, FromHeight: fromHeight, ToHeight: toHeight}
	return response, dc.jsonRequest("get_output_distribution", params, &response)
}

func (dc *DaemonClient) GetHeight() (response HeightResponse, err error) {
	type Params struct{}

	params := Params{}
	return response, dc.rpcRequest("/get_height", params, &response)
}

func (dc *DaemonClient) GetTransactions(txs_hashes []string, decode_as_json bool, prune bool) (response TransactionsResponse, err error) {
	type Params struct {
		TxsHashes    []string `json:"txs_hashes"`
		DecodeAsJSON bool     `json:"decode_as_json"`
		Prune        bool     `json:"prune"`
	}

	params := Params{TxsHashes: txs_hashes, DecodeAsJSON: decode_as_json, Prune: prune}
	return response, dc.rpcRequest("/get_transactions", params, &response)
}

func (dc *DaemonClient) GetAltBlocksHashes() (response AltBlocksHashesResponse, err error) {
	type Params struct{}

	params := Params{}
	return response, dc.rpcRequest("/get_alt_blocks_hashes", params, &response)
}

func (dc *DaemonClient) IsKeyImageSpent(keyImages []string) (response IsKeyImageSpentResponse, err error) {
	type Params struct {
		KeyImages []string `json:"key_images"`
	}

	params := Params{KeyImages: keyImages}
	return response, dc.rpcRequest("/is_key_image_spent", params, &response)
}

func (dc *DaemonClient) SendRawTransaction(txAsHex string, doNotRelay bool) (response SendRawTransactionResponse, err error) {
	type Params struct {
		TxAsHex    string `json:"txAsHex"`
		DoNotRelay bool   `json:"doNotRelay"`
	}

	params := Params{TxAsHex: txAsHex, DoNotRelay: doNotRelay}
	return response, dc.rpcRequest("/send_raw_transaction", params, &response)
}

func (dc *DaemonClient) StartMining(doBackgroundMining bool, ignoreBattery bool, minerAddress string, threadsCount uint) (response StatusResponse, err error) {
	type Params struct {
		DoBackgroundMining bool   `json:"do_background_mining"`
		IgnoreBattery      bool   `json:"ignore_battery"`
		MinerAddress       string `json:"miner_address"`
		ThreadsCount       uint   `json:"threads_count"`
	}

	params := Params{DoBackgroundMining: doBackgroundMining, IgnoreBattery: ignoreBattery, MinerAddress: minerAddress, ThreadsCount: threadsCount}
	return response, dc.rpcRequest("/start_mining", params, &response)
}

func (dc *DaemonClient) StopMining() (response StatusResponse, err error) {
	type Params struct{}

	params := Params{}
	return response, dc.rpcRequest("/stop_mining", params, &response)
}

func (dc *DaemonClient) MiningStatus() (response MiningStatusResponse, err error) {
	type Params struct{}

	params := Params{}
	return response, dc.rpcRequest("/mining_status", params, &response)
}

func (dc *DaemonClient) SaveBC() (response StatusResponse, err error) {
	type Params struct{}

	params := Params{}
	return response, dc.rpcRequest("/save_bc", params, &response)
}

func (dc *DaemonClient) GetPeerList() (response PeerListResponse, err error) {
	type Params struct{}

	params := Params{}
	return response, dc.rpcRequest("/get_peer_list", params, &response)
}

func (dc *DaemonClient) SetLogHashRate(visible bool) (response StatusResponse, err error) {
	type Params struct {
		Visible bool `json:"visible"`
	}

	params := Params{Visible: visible}
	return response, dc.rpcRequest("/set_log_hash_rate", params, &response)
}

func (dc *DaemonClient) SetLogLevel(level uint) (response StatusResponse, err error) {
	type Params struct {
		Level uint `json:"level"`
	}

	params := Params{Level: level}
	return response, dc.rpcRequest("/set_log_level", params, &response)
}

func (dc *DaemonClient) SetLogCategories(categories string) (response LogCategoriesResponse, err error) {
	type Params struct {
		Categories string `json:"categories"`
	}

	params := Params{Categories: categories}
	return response, dc.rpcRequest("/set_log_categories", params, &response)
}

func (dc *DaemonClient) GetTransactionPool() (response TransactionPoolResponse, err error) {
	type Params struct{}

	params := Params{}
	return response, dc.rpcRequest("/get_transaction_pool", params, &response)
}

func (dc *DaemonClient) GetTransactionPoolStats() (response TransactionPoolStatsResponse, err error) {
	type Params struct{}

	params := Params{}
	return response, dc.rpcRequest("/get_transaction_pool_stats", params, &response)
}

func (dc *DaemonClient) StopDaemon() (response StatusResponse, err error) {
	type Params struct{}

	params := Params{}
	return response, dc.rpcRequest("/stop_daemon", params, &response)
}

func (dc *DaemonClient) GetLimit() (response LimitResponse, err error) {
	type Params struct{}

	params := Params{}
	return response, dc.rpcRequest("/get_limit", params, &response)
}

func (dc *DaemonClient) SetLimit(limitDown int, limitUp int) (response LimitResponse, err error) {
	type Params struct {
		LimitDown int `json:"limit_down"`
		LimitUp   int `json:"limit_up"`
	}

	params := Params{LimitDown: limitDown, LimitUp: limitUp}
	return response, dc.rpcRequest("/set_limit", params, &response)
}

func (dc *DaemonClient) OutPeers(outPeers uint) (response StatusResponse, err error) {
	type Params struct {
		OutPeers uint `json:"out_peers"`
	}

	params := Params{OutPeers: outPeers}
	return response, dc.rpcRequest("/out_peers", params, &response)
}

func (dc *DaemonClient) InPeers(inPeers uint) (response StatusResponse, err error) {
	type Params struct {
		InPeers uint `json:"in_peers"`
	}

	params := Params{InPeers: inPeers}
	return response, dc.rpcRequest("/out_peers", params, &response)
}

func (dc *DaemonClient) StartSaveGraph() (response StatusResponse, err error) {
	type Params struct{}

	params := Params{}
	return response, dc.rpcRequest("/start_save_graph", params, &response)
}

func (dc *DaemonClient) StopSaveGraph() (response StatusResponse, err error) {
	type Params struct{}

	params := Params{}
	return response, dc.rpcRequest("/stop_save_graph", params, &response)
}

func (dc *DaemonClient) Update(command string, path string) (response UpdateResponse, err error) {
	type Params struct {
		Command string `json:"command"`
		Path    string `json:"path"`
	}

	params := Params{Command: command, Path: path}
	return response, dc.rpcRequest("/update", params, &response)
}
