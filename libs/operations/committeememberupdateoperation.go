package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeCommitteeMemberUpdate] = func() types.Operation {
		op := &CommitteeMemberUpdateOperation{}
		return op
	}
}

type CommitteeMemberUpdateOperation struct {
	types.OperationFee
	CommitteeMember        types.CommitteeMember `json:"committee_member"`
	CommitteeMemberAccount types.AccountID       `json:"committee_member_account"`
	NewURL                 *types.String         `json:"new_url,omitempty"`
}

func (p CommitteeMemberUpdateOperation) Type() types.OperationType {
	return types.OperationTypeCommitteeMemberUpdate
}

func (p CommitteeMemberUpdateOperation) Marshal(enc *util.TypeE