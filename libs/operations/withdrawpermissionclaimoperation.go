package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeWithdrawPermissionClaim] = func() types.Operation {
		op := &WithdrawPermissionClaimOperation{}
		return op
	}
}

type WithdrawPermissionClaimOperation struct {
	types.OperationFee
	WithdrawPermission  types.WithdrawPermissionID `json:"withdraw_permission"`
	WithdrawFromAccount types.AccountID            `json:"withdraw_from_account"`
	WithdrawToAccount   types.AccountID            `json:"withdraw_to_account"`
	AmountToWithdraw    types.AssetAmount          `json:"amount_to_withdraw"`
	Memo                *types.Memo                `json:"memo,omitempty"`
}

func (p WithdrawPermissionClaimOperation) Type() types.OperationType {
	return types.OperationTypeWithdrawPermissionClaim
}

func (p WithdrawPermissionClaimOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if fee, ok := params["fee"]; ok {
		if err := enc.Encode(types.UInt64(fee.(float64))); err != nil {
			return errors.Annotate(err, "encode Fee")
		}
	}

	if ppk, ok := params["price_per_kbyte"]; ok {
		if err := enc.Encode(types.UIn