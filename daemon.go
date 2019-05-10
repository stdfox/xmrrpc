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

type BlockCount struct {
	Count  uint   `json:"count"`
	Status string `json:"status"`
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

func (dc *DaemonClient) GetBlockCount() (BlockCount, error) {
	var blockCount BlockCount

	if err := dc.jsonRPCRequest("getblockcount", nil, &blockCount); err != nil {
		return blockCount, err
	}

	return blockCount, nil
}

func (dc *DaemonClient) OnGetBlockHash(blockHeight int) (string, error) {
	var blockHash string

	if err := dc.jsonRPCRequest("on_getblockhash", []int{blockHeight}, &blockHash); err != nil {
		return blockHash, err
	}

	return blockHash, nil
}
