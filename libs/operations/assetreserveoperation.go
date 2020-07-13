package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAssetReserve] = func() types.Operation {
		op := &AssetReserveOperation{}
		return op
	}
}

type AssetReserveOperation struct {
	types.OperationFee
	Payer    