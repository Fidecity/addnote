package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeWithdrawPermissionUpdate] = func() types.Operation {
		op := &WithdrawPermissionUpdateOperation{}
		return op
	}
}

type WithdrawPermissionUpdateOperation struct {
	types.OperationFee
	WithdrawFromAccount    types.AccountID            `json:"withdraw_from_account"`
	AuthorizedAccount      types.AccountID            `json:"authorized_account"`
	PermissionToUpdate     types.WithdrawPermissionID `json:"permission_to_update"`
	WithdrawalLimit        types.AssetAmount          `json:"wit