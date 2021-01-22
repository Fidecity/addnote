package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeTransfer] = func() types.Operation {
		op := &TransferOperation{}
		return op
	}
}

type TransferOperation struct {
	types.OperationFee
	From       types.AccountID   `json:"from"`
	To         types.AccountID   `json:"to"`
	FromAddr   types.Address     `json:"from_addr"`
	ToAddr     types.Address     `json:"to_addr"`
	Amount     types.AssetAmount `json:"amount"`
	Memo       *types.Memo       `json:"memo,omitempty"`
	Extensions types.Extensions  `json:"extens