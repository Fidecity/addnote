package whitecoin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Assetsadapter/whitecoin-adapter/types"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/blocktree/openwallet/v2/log"
	bt "github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
)

// WalletClient is a Bitshares RPC client. It performs RPCs over HTTP using JSON
// request and responses. A Client must be configured with a secret token
// to authenticate with other Cores on the network.
type WalletClient struct {
	WalletAPI, ServerAPI string
	Debug                bool
	client               *req.Req
}

// NewWalletClient init a rpc client
func NewWalletClient(serverAPI, walletAPI string, debug bool) *WalletClient {

	walletAPI = strings.TrimSuffix(walletAPI, "/")
	serverAPI = strings.TrimSuffix(serverAPI, "/")
	c := WalletClient{
		WalletAPI: walletAPI,
		ServerAPI: serverAPI,
		Debug:     debug,
	}

	api := req.New()
	c.client = api

	return &c
}

// Call calls a remote procedure on another node, specified by the path.
func (c *WalletClient) call(method string, request interface{}, queryWalletAPI bool) (*gjson.Result, error) {

	var (
		body = make(map[string]interface{}, 0)
	)

	if c.client == nil {
		return nil, fmt.Errorf("API url is not setup. ")
	}

	authHeader := req.Header{
		"Accept":       "application/json",
		"Content-Type": "application/json",
		"Connection":   "close",
	}

	//json-rpc
	body["jsonrpc"] = "2.0"
	body["id"] = 1
	body["method"] = method
	body["params"] = request

	if c.Debug {
		log.Std.Info("Start Request API...")
	}

	host := c.ServerAPI
	if queryWalletAPI {
		host = c.WalletAPI
	}

	r, err := c.client.Post(host, req.BodyJSON(&body), authHeader)

	if c.Debug {
		log.Std.Info("Request API Completed")
	}

	if c.Debug {
		log.Std.Info("%+v", r)
	}

	if err != nil {
		return nil, err
	}

	resp := gjson.ParseBytes(r.Bytes())
	err = c.isError(r)
	if err != nil {
		return nil, err
	}

	result := resp.Get("result")

	return &result, nil
}

// isError 是否报错
func (c *WalletClient) isError(r *req.Resp) error {

	if r.Response().StatusCode != http.StatusOK {
		message := r.Response().Status
		status := r.Response().StatusCode
		return fmt.Errorf("[%d]%s", status, message)
	}

	result := gjson.ParseBytes(r.Bytes())

	if result.Get("error").IsObject() {

		return fmt.Errorf("[%d]%s",
			result.Get("error.code").Int(),
			result.Get("error.message").String())

	}

	return nil

}

// GetObjects return a block by the given block number
func (c *WalletClient) GetObjects(assets ...types.ObjectID) (*gjson.Result, error) {
	resp, err := c.call("get_objects", []interface{}{objectsToParams(assets)}, false)
	return resp, err
}

func objectsToParams(objs []types.ObjectID) []string {
	objsStr := make([]string, len(objs))
	for i, asset := range objs {
		objsStr[i] = asset.String()
	}
	return objsStr
}

// GetBlockchainInfo returns current blockchain data
func (c *WalletClient) GetBlockchainInfo() (*BlockchainInfo, error) {
	//get_dynamic_global_properties
	r, err := c.call("get_dynamic_global_properties", []interface{}{}, false)
	if err != nil {
		return nil, err
	}
	info := NewBlockchainInfo(r)
	return info, nil
}

// GetBlockByHeight returns a certain block
func (c *WalletClient) GetBlockByHeight(height uint32) (*Block, error) {
	r, err := c.call("get_block", []interface{}{height}, false)
	if err != nil {
		return nil, err
	}
	block := NewBlock(height, r)
	return block, nil
}

// GetTransaction returns the TX
func (c *WalletClient) GetTransaction(txid string) (*types.Transaction, error) {
	r, err := c.call("get_transaction", []interface{}{txid}, false)
	if err != nil {
		return nil, err
	}
	return NewTransaction(r, txid)
}

// GetAddrBalance Returns information about the given account.
func (c *WalletClient) GetAddrBalance(addr string, asset types.ObjectID) (*Balance, error) {

	balance := &Balance{
		AssetID: asset,
		Amount:  "0",
	}

	//var resp []*types.Account
	r, err := c.call("get_addr_balances", []interface{}{addr, []interface{}{asset.String()}}, false)
	if err != nil {
		return balance, nil
	}

	if r.IsArray() {
		for _, item := range r.Array() {
			b := NewBalance(item)
			if b.AssetID.String() == asset.String() {
				return b, nil
			}

		}
	}

	return balance, nil
}

// GetAssetsBalance Returns information about the given account.
func (c *WalletClient) GetAccountID(name string) (*types.ObjectID, error) {
	r, err := c.call("lookup_accounts", []interface{}{name, 1}, false)
	if err != nil {
		return nil, err
	}
	arr := r.Array()
	if len(arr) > 0 {
		if arr[0].Array()[0].String() == name {
			id := arr[0].Array()[1].String()
			objectID := types.MustParseObjectID(id)
			return &objectID, nil
