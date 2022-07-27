package types

import (
	"encoding/json"
	"fmt"

	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/denkhaus/logging"
	"github.com/juju/errors"
	"github.com/pquerna/ffjson/ffjson"
)

type GetOpFunc func() Operation

var (
	OperationMap = make(map[OperationType]GetOpFunc)
)

type Operation interface {
	util.TypeMarshaler
	SetFee(fee AssetAmount)
	GetFee() AssetAmount
	Type() OperationType
	MarshalFeeScheduleParams(M, *util.TypeEncoder) error
}

type OperationResult interface {
}

type OperationEnvelopeHolder struct {
	Op OperationEnvelope `json:"op"`
}

type OperationEnvelopeHolders []OperationEnvelopeHolder

func (p OperationEnvelopeHolders) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, op := range p {
		if err := enc.Encode(op.Op); err != nil {
			return errors.Annotate(err, "encode Op")
		}
	}

	return nil
}

type OperationEnvelope struct {
	Type      OperationType
	Operation Operation
}

func (p OperationEnvelope) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(p.Operation); err != nil {
		return errors.Annotate(err, "encode Operation")
	}

	return nil
}

func (p OperationEnvelope) MarshalJSON() ([]byte, error) {
	return ffjson.Marshal([]interface{}{
		p.Type,
		p.Operation,
	})
}

func (p *OperationEnvelope) UnmarshalJSON(data []byte) error {
	raw := make([]json.RawMessage, 2)
	if err := ffjson.Unmarshal(data, &raw); err != nil {
		return errors.Annotate(err, "unmarshal raw object")
	}

	if err := ffjson.Unmarshal(raw[0], &p.Type); err != nil {
		return errors.Annotate(err, "unmarshal OperationType")
	}

	descr := fmt.Sprintf("Operation %s", p.Ty