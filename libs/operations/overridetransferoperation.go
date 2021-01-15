package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeOverrideTransfer] = func() types.Operation {
		op := &OverrideTransferOperation{}
		return op
	}
}

type OverrideTransferOperation struct {
	types.OperationFee
	Amount     types.AssetAmount `json:"amount"`
	Extensions types.Extensions  `json:"extensions"`
	From       types.AccountID   `json:"from"`
	Issuer     types.AccountID   `json:"issuer"`
	Memo       *types.Memo       `json:"memo,omitempty"`
	To         types.AccountID   `json:"to"`
}

func (p OverrideTransferOperation) Type() types.OperationType {
	return types.OperationTypeOverrideTransfer
}

func (p OverrideTransferOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if fee, ok := params["fee"]; ok {
		if err := enc.Encode(types.