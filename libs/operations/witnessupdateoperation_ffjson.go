// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: witnessupdateoperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *WitnessUpdateOperation) MarshalJSON() ([]byte, error) {
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
func (j *WitnessUpdateOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if j.NewSigningKey != nil {
		if true {
			buf.WriteString(`"new_signing_key":`)

			{

				obj, err = j.NewSigningKey.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
			buf.WriteByte(',')
		}
	}
	if j.NewURL != nil {
		if true {
			buf.WriteString(`"new_url":`)

			{

				obj, err = j.NewURL.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
			buf.WriteByte(',')
		}
	}
	buf.WriteString(`"witness":`)

	{

		obj, err = j.Witness.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"witness_account":`)

	{

		obj, err = j.WitnessAccount.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

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
	ffjtWitnessUpdateOperationbase = iota
	ffjtWitnessUpdateOperationnosuchkey

	ffjtWitnessUpdateOperationNewSigningKey

	ffjtWitnessUpdateOperationNewURL

	ffjtWitnessUpdateOperationWitness

	ffjtWitnessUpdateOperationWitnessAccount

	ffjtWitnessUpdateOperationFee
)

var ffjKeyWitnessUpdateOperationNewSigningKey = []byte("new_signing_key")

var ffjKeyWitnessUpdateOperationNewURL = []byte("new_url")

var ffjKeyWitnessUpdateOperationWitness = []byte("witness")

var ffjKeyWitnessUpdateOperationWitnessAccount = []byte("witness_account")

var ffjKeyWitnessUpdateOperationFee = 