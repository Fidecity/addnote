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

// NewWalletClient init a rpc 