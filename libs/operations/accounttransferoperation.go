package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAccountTransfer] = func() types.Operation {
		op := &AccountTransferOperation{}
		return op
	}
}

type AccountTransferOperation struct {
	types.OperationFee
	AccountID  types.AccountID  `json:"account_id"`
	NewOwner   types.AccountID  `json:"new_owner"`
	Extensions types.Extensions `json:"extensions"`
}

func (p AccountTransferOperation) Type() types.OperationType {
	return types.OperationTypeAccountTransfer
}

func (p AccountTransferOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.