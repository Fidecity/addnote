package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeWithdrawPermissionCreate] = func() types.Operation {
		op := &WithdrawPermissionCreateOperation{}
		return op
	}
}

type WithdrawPermissionCreateOperation struct {
	types.OperationFee
	AuthorizedAccount      types.AccountID   `json:"authorized_account"`
	PeriodStartTime        types.Time        `json:"period_start_time"`
	PeriodsUntilExpiration types.UInt32      `json:"periods_until_expiration"`
	WithdrawFromAccount    types.AccountID   `json:"withdraw_from_account"`
	WithdrawalLimit        types.AssetAmount `json:"withdrawal_limit"`
	WithdrawalPeriodSec    types.UInt32      `json:"withdrawal_period_sec"`
}

func (p WithdrawPermissionCreateOperation) Type() types.OperationType {
	return types.OperationTypeWithdrawPermissionCreate
}

func (p WithdrawPermissionCreateOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode Operation