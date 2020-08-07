package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAssetUpdateFeedProducers] = func() types.Operation {
		op := &AssetUpdateFeedProducersOperation{}
		return op
	}
}

type AssetUpdateFeedProducersOperation struct {
	types.OperationFee
	AssetToUpdate    types.AssetID    `json:"asset_to_update"`
	Extensions       types.Extensions `json:"extensions"`
	Issuer           types.AccountID  `json:"issuer"`
	NewFeedProducers types.AccountIDs `json:"new_feed_producers"`
}

func (p AssetUpdateFeedProducers