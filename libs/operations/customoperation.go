package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeCustom] = func() types.Operation {
		op := &CustomOperation{}
		return op
	}
}

type CustomOperation struct {
	types.OperationFee
	Payer         types.AccountID  `json:"payer"`
	RequiredAuths types.AccountIDs `json:"required_auths"`
	ID            types.UInt16     `json:"id"`
	Data          types.Buffer     `json:"data"`
}

func (p CustomOperation) Type() types.OperationType {
	return types.OperationTypeCustom
}

func (p CustomOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if fee, ok := params["fee"]; ok {
		if err := enc.Encode(types.UInt64(fee.(float64))); err != nil {
			return errors.Annotate(err, "encode Fee")
		}
	}

	if ppk, ok := params["price_per_kbyte"]; ok {
		if err := enc.Encode(types.UInt32(ppk.(float64))); err != nil {
			return errors.Annotate(err, "encode PricePerKByte")
	