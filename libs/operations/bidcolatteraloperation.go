package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeBidCollateral] = func() types.Operation {
		op := &BidCollateralOperation{}
		return op
	}
}

type BidCollateralOperation struct {
	types.OperationFee
	AdditionalCollateral types.AssetAmount `json:"additional_collateral"`
	Bidder               types.AccountID   `json:"bidder"`
	DebtCovered          types.AssetAmount `json:"debt_covered"`
	Extensions           types.Extensions  `json:"extensions"`
}

func (p BidCollateralOperation) Type() types.OperationType {
	return types.OperationTypeBidCollateral
}

func (p BidCollateralOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.En