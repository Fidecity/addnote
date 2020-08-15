package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAssetUpdate] = func() types.Operation {
		op := &AssetUpdateOperation{}
		return op
	}
}

type AssetUpdateOperation struct {
	types.OperationFee
	AssetToUpdate types.AssetID      `json:"asset_to_update"`
	Issuer        types.AccountID    `json:"issuer"`
	Extensions    types.Extensions   `json:"extensions"`
	NewIssuer     *types.AccountID   `json:"new_issuer"`
	NewOptions    types.AssetOptions `json:"new_options"`
}

func (p AssetUpdateOperation) Type() types.OperationType {
	return types.OperationTypeAssetUpdate
}

func (p AssetUpdateOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if fee, ok := params["fee"]; ok {
		if err := enc.Encode(types.UInt64(fee.(float64))); err != nil {
			return errors.Annotate(err, "encode Fee")
		}
	}

	if ppk, ok := params["price_per_kbyte"]; ok {
		if err := enc.Encode(types.UInt32(ppk.(float64))); err != nil {
			return errors.Annotate(err, "encode PricePerKByte")
		}
	}

	return nil
}

func (p AssetUpdateOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.