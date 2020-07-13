package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAssetPublishFeed] = func() types.Operation {
		op := &AssetPublishFeedOperation{}
		return op
	}
}

type AssetPublishFeedOperation struct {
	types.OperationFee
	Publisher  types.AccountID  `json:"publisher"`
	AssetID    types.AssetID    `json:"asset_id"`
	Feed       types.PriceFeed  `json:"feed"`
	Extensions types.Extensions `json:"extensions"`
}

func (p AssetPublishFeedOperation) Type() types.OperationType {
	return types.OperationTypeAssetPublishFeed
}

func (p AssetPublishFeedOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

