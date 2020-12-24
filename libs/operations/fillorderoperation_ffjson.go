// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: fillorderoperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *FillOrderOperation) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *FillOrderOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "order_id":`)

	{

		obj, err = j.OrderID.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"account_id":`)

	{

		obj, err = j.AccountID.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	/* Struct fall back. type=types.AssetAmount kind=struct */
	buf.WriteString(`,"pays":`)
	err = buf.Encode(&j.Pays)
	if err != nil {
		return err
	}
	/* Struct fall back. type=types.AssetAmount kind=struct */
	buf.WriteString(`,"receives":`)
	err = buf.Encode(&j.Receives)
	if err != nil {
		return err
	}
	if j.IsMaker {
		buf.WriteString(`,"is_maker":true`)
	} else {
		buf.WriteString(`,"is_maker":false`)
	}
	/* Struct fall back. type=types.Price kind=struct */
	buf.WriteString(`,"fill_price":`)
	err = buf.Encode(&j.FillPrice)
	if err != nil {
		return err
	}
	buf.WriteByte(',')
	if j.Fee != nil {
		if true {
			/* Struct fall back. type=types.AssetAmount kind=struct */
			buf.WriteString(`"fee":`)
			err = buf.Encode(j.Fee)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtFillOrderOperationbase = iota
	ffjtFillOrderOperationnosuchkey

	ffjtFillOrderOperationOrderID

	ffjtFillOrderOperationAccountID

	ffjtFillOrderOperationPays

	ffjtFillOrderOperationReceives

	ffjtFillOrderOperationIsMaker

	ffjtFillOrderOperationFillPrice

	ffjtFillOrderOperationFee
)

var ffjKeyFillOrderOperationOrderID = []byte("order_id")

var ffjKeyFillOrderOperationAccountID = []byte("account_id")

var ffjKeyFillOrderOperationPays = []byte("pa