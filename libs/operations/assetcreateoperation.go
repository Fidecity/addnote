package operations

//go:generate ffjson $GOFILE

import (
	"strconv"

	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAssetCreate] = func() types.Operation {
		op := &AssetCreateOperation{}
		return op
	}
}

type AssetCreateOperation struct {
	types.OperationFee
	BitassetOptions    *types.BitassetOptions `json:"bitasset_opts"`
	CommonOptions      types.AssetOptions     `json:"common_options"`
	Extensions         types.Extensions       `json:"extensions"`
	IsPredictionMarket bool                   `json:"is_prediction_market"`
	Issuer             types.AccountID        `json:"issuer"`
	Precision          types.UInt8            `json:"precision"`
	Symbol             types.String           `json:"symbol"`
}

func (p AssetCreateOperation) Type() types.OperationType {
	return types.OperationTypeAssetCreate
}

func (p AssetCreateOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if s3, ok := params["symbol3"]; ok {
		s3Val, 