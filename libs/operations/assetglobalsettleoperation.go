package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAssetGlobalSettle] = func() types.Operation {
		op := &AssetGlobalSettleOperation{}
		return op
	}
}

type AssetGlobalSettleOperation struct {
	types.OperationFee
	AssetToSettle types.AssetID    `json:"asset_to_settle"`
	Extensions    types.Extensions `json:"extensions"`
	Issuer        types.AccountID  `json:"issuer"`
	SettlePrice   types.Price      `json:"settle_price"`
}

func (p AssetGlobalSettleOperation) Type() types.OperationType {
	re