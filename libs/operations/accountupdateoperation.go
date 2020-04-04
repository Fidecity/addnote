package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAccountUpdate] = func() types.Operation {
		op := &AccountUpdateOperation{}
		return op
	}
}

type AccountUpdateOperation struct {
	types.OperationFee
	Account    types.AccountID               `json:"account"`
	Active     *types.Authority           