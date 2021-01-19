package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeProposalUpdate] = func() types.Operation {
		op := &ProposalUpdateOperation{}
		return op
	}
}

type ProposalUpdateOperation struct {
	types.OperationFee
	ActiveApprovalsToAdd    types.AccountIDs `json:"active_approvals_to_add"`
	ActiveApprovalsToRemove types.AccountIDs `json:"active_approvals_to_remove"`
	OwnerApprovalsToAdd     types.AccountIDs `json:"owner_approvals_to_add"`
	OwnerApprovalsToRemove  types.AccountIDs `json:"owner_approvals_to_remove"`
	Extensions              types.Extensions `json:"extensions"`
	FeePayingAccount        types.AccountID  `json:"fee_paying_account"`
	KeyApprovalsToAdd       types.PublicKeys `json:"key_approvals_to_add"`
	KeyApprovalsToRemove    types.PublicKeys `json:"key_approvals_to_remove"`
	Proposal                types.ProposalID `json:"proposal"`
}

func (p ProposalUpdateOperation) Type() types.OperationType {
	return types.OperationTypeProposalUpdate
}

func (p ProposalUpdateOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if fee, ok := params["f