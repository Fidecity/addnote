
package types

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/Assetsadapter/whitecoin-adapter/encoding"
	"github.com/pkg/errors"
)

type Operation interface {
	Type() OpType
}

type Operations []Operation

func (ops *Operations) UnmarshalJSON(b []byte) (err error) {
	// unmarshal array
	var o []json.RawMessage
	if err := json.Unmarshal(b, &o); err != nil {
		return err
	}

	// foreach operation
	for _, op := range o {
		var kv []json.RawMessage
		if err := json.Unmarshal(op, &kv); err != nil {
			return err
		}

		if len(kv) != 2 {
			return errors.New("invalid operation format: should be name, value")
		}

		var opType uint16
		if err := json.Unmarshal(kv[0], &opType); err != nil {
			return err
		}

		val, err := unmarshalOperation(OpType(opType), kv[1])
		if err != nil {
			return err
		}

		*ops = append(*ops, val)
	}

	return nil
}

type operationTuple struct {
	Type OpType
	Data Operation
}

func (op *operationTuple) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{
		op.Type,
		op.Data,
	})
}

func (ops Operations) MarshalJSON() ([]byte, error) {
	tuples := make([]*operationTuple, 0, len(ops))
	for _, op := range ops {
		tuples = append(tuples, &operationTuple{
			Type: op.Type(),
			Data: op,
		})
	}
	return json.Marshal(tuples)
}

func unmarshalOperation(opType OpType, obj json.RawMessage) (Operation, error) {
	op, ok := knownOperations[opType]
	if !ok {
		// operation is unknown wrap it as an unknown operation
		val := UnknownOperation{
			kind: opType,
			Data: obj,
		}
		return &val, nil
	} else {
		val := reflect.New(op).Interface()
		if err := json.Unmarshal(obj, val); err != nil {
			return nil, err
		}
		return val.(Operation), nil
	}
}

var knownOperations = map[OpType]reflect.Type{
	TransferOpType:         reflect.TypeOf(TransferOperation{}),
	LimitOrderCreateOpType: reflect.TypeOf(LimitOrderCreateOperation{}),
	LimitOrderCancelOpType: reflect.TypeOf(LimitOrderCancelOperation{}),
}

// UnknownOperation
type UnknownOperation struct {
	kind OpType
	Data json.RawMessage
}

func (op *UnknownOperation) Type() OpType { return op.kind }

// NewTransferOperation returns a new instance of TransferOperation
func NewTransferOperation(from, to ObjectID, amount, fee AssetAmount) *TransferOperation {
	op := &TransferOperation{
		From:       from,
		To:         to,
		Amount:     amount,
		Fee:        fee,
		Extensions: []json.RawMessage{},
	}

	return op
}

// TransferOperation
type TransferOperation struct {
	From       ObjectID          `json:"from"`
	To         ObjectID          `json:"to"`
	Amount     AssetAmount       `json:"amount"`
	Fee        AssetAmount       `json:"fee"`
	Memo       Memo              `json:"memo"`
	Extensions []json.RawMessage `json:"extensions"`
	FromAddr   string            `json:"from_addr"`
	ToAddr     string            `json:"to_addr"`
}

type Buffer []byte

func (p Buffer) String() string {