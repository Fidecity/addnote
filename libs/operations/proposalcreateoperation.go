package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeProposalCreate] = func() types.Operation {
		op := &ProposalCreateOperation{}
		return op
	}
}

type ProposalCreateOperation struct {
	types.OperationFee
	ExpirationTime      types.Time                     `json:"expiration_time"`
	Extensions          types.Extensions               `json:"extensions"`
	FeePayingAccount    types.AccountID                `json:"fee_paying_account"`
	ReviewPeriodSeconds *types.UInt32                  `json:"review_period_seconds,omitempty"`
	ProposedOps         types.OperationEnvelopeHolders `json:"proposed_ops"`
}

func (p Proposal