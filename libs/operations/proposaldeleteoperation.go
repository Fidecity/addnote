package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeProposalDelete] = func() types.Operation {
		op := &ProposalDeleteOperation{}
		return op
	}
}

type ProposalDeleteOperation struct {
	types.OperationFee
	Extensions          types.Extensions `json:"extensions"`
	FeePayingAccount    types.AccountID  `json:"fee_paying_account"`
	Proposal            types.