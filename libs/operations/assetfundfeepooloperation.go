package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAssetFundFeePool] = func() types.Operation {
		op := &AssetFundFeePoolOperation{}
		return op
	}
}

type AssetFundFeePoolOperation struct {
	types.OperationFee
	Amount      types.UInt64     `json:"amount"`
	AssetID     types.AssetID    `json:"asset_id"`
	Extensions  types.Extensions `json:"extensions"`
	FromAccount types.AccountID  `json:"from_account"`
}

func (p AssetFundFeePoolOperation) Type() types.OperationType {
	return types.OperationTypeAssetFundFeePool
}

func (p AssetFundFeePoolOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.FromAccount); er