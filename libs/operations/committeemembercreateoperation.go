package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeCommitteeMemberCreate] = func() types.Operation {
		op := &CommitteeMemberCreateOperation{}
		return op
	}
}

type CommitteeMemberCreateOperation struct {
	types.OperationFee
	CommitteeMemberAccount types.AccountID `json:"committee_member_account"`
	URL                    types.String    `json:"url"`
}

func (p CommitteeMemberCreateOperation) Type() types.OperationType {
	return types.OperationTypeCommitteeMemberCreate
}

func (p CommitteeMemberCreateOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode Fee")
	}
	if err := 