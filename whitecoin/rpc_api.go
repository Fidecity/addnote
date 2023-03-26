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
	r, err := c.call("get_block", []int