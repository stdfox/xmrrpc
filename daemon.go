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

type Block struct {
	Blob        string      `json:"blob"`
	BlockHeader BlockHeader `json:"block_header"`
	Json        string      `json:"json"`
	Status      string      `status`
	Untrusted   bool        `json:"untrusted"`
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

func (dc *DaemonClient) GetBlockHeadersRange(start_height uint, end_height uint) (BlockHeadersResponse, error) {
	var blockHeadersResponse BlockHeadersResponse

	type jsonRPCParams struct {
		StartHeight uint `json:"start_height"`
		EndHeight   uint `json:"end_height"`
	}

	params := jsonRPCParams{StartHeight: start_height, EndHeight: end_height}
	err := dc.jsonRPCRequest("get_block_headers_range", params, &blockHeadersResponse)

	return blockHeadersResponse, err
}

func (dc *DaemonClient) GetBlock(height uint, hash string) (Block, error) {
	var block Block

	type jsonRPCParams struct {
		Height uint   `json:"height"`
		Hash   string `json:"hash"`
	}

	params := jsonRPCParams{Height: height, Hash: hash}
	err := dc.jsonRPCRequest("get_block", params, &block)

	return block, err
}
