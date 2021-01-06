package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeLimitOrderCreate] = func() types.Operation {
		op := &LimitOrderCreateOperation{}
		return op
	}
}

//LimitOrderCreateOperation instructs the blockchain to attempt to sell one asset for another.
//The blockchain will attempt to sell amount_to_sell.asset_id for as much min_to_receive.asset_id as possible.
//The fee will be paid by the seller’s account. Market fees will apply as specified by the issuer of both the selling asset and the receiving asset as a percentage of the amount exchanged.
//If either the selling asset or the receiving asset is white list restricted, the order will only be created if the seller is on the white list of the restricted asset type.
//Market orders are matched in the order they are included in the block