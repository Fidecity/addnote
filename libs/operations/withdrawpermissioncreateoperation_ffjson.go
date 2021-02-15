// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: withdrawpermissioncreateoperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *WithdrawPermissionCreateOperation) MarshalJSON() ([]byte, error) {
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
func (j *WithdrawPermissionCreateOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "authorized_account":`)

	{

		obj, err = j.AuthorizedAccount.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"period_start_time":`)

	{

		obj, err = j.PeriodStartTime.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"periods_until_expiration":`)
	fflib.FormatBits2(buf, uint64(j.PeriodsUntilExpiration), 10, false)
	buf.WriteString(`,"withdraw_from_account":`)

	{

		obj, err = j.WithdrawFromAccount.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	/* Struct fall back. type=types.AssetAmount kind=struct */
	buf.WriteString(`,"withdrawal_limit":`)
	err = buf.Encode(&j.WithdrawalLimit)
	if err != nil {
		return err
	}
	buf.WriteString(`,"withdrawal_period_sec":`)
	fflib.FormatBits2(buf, uint64(j.WithdrawalPeriodSec), 10, false)
	buf.WriteByte(',')
	if j.Fee != nil {
		if true {
			/* Struct fall back. type=types.AssetAmount kind=struct */
			buf.WriteString(`"fee":`)
			err = buf.Encode(j.Fee)
			if err !