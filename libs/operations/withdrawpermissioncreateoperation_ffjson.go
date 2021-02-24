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
	ffjtWithdrawPermissionCreateOperationbase = iota
	ffjtWithdrawPermissionCreateOperationnosuchkey

	ffjtWithdrawPermissionCreateOperationAuthorizedAccount

	ffjtWithdrawPermissionCreateOperationPeriodStartTime

	ffjtWithdrawPermissionCreateOperationPeriodsUntilExpiration

	ffjtWithdrawPermissionCreateOperationWithdrawFromAccount

	ffjtWithdrawPermissionCreateOperationWithdrawalLimit

	ffjtWithdrawPermissionCreateOperationWithdrawalPeriodSec

	ffjtWithdrawPermissionCreateOperationFee
)

var ffjKeyWithdrawPermissionCreateOperationAuthorizedAccount = []byte("authorized_account")

var ffjKeyWithdrawPermissionCreateOperationPeriodStartTime = []byte("period_start_time")

var ffjKeyWithdrawPermissionCreateOperationPeriodsUntilExpiration = []byte("periods_until_expiration")

var ffjKeyWithdrawPermissionCreateOperationWithdrawFromAccount = []byte("withdraw_from_account")

var ffjKeyWithdrawPermissionCreateOperationWithdrawalLimit = []byte("withdrawal_limit")

var ffjKeyWithdrawPermissionCreateOperationWithdrawalPeriodSec = []byte("withdrawal_period_sec")

var ffjKeyWithdrawPermissionCreateOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *WithdrawPermissionCreateOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *WithdrawPermissionCreateOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtWithdrawPermissionCreateOperationbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtWithdrawPermissionCreateOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyWithdrawPermissionCreateOperationAuthorizedAccount, kn) {
						currentKey = ffjtWithdrawPermissionCreateOperationAuthorizedAccount
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyWithdrawPermissionCreateOperationFee, kn) {
						currentKey = ffjtWithdrawPermissionCreateOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffjKeyWithdrawPermissionCreateOperationPeriodStartTime, kn) {
						currentKey = ffjtWithdrawPermissionCreateOperationPeriodStartTime
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyWithdrawPermissionCreateOperationPeriodsUntilExpiration, kn) {
						currentKey = ffjtWithdrawPermissionCreateOperationPeriodsUntilExpiration
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffjKeyWithdrawPermissionCreateOperationWithdrawFromAccount, kn) {
						currentKey = ffjtWithdrawPermissionCreateOperationWithdrawFromAccount
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyWithdrawPermissionCreateOperationWithdrawalLimit, kn) {
						currentKey = ffjtWithdrawPermissionCreateOperationWithdrawalLimit
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyWithdrawPermissionCreateOperationWithdrawalPeriodSec, kn) {
						currentKey = ffjtWithdrawPermissionCreateOperationWithdrawalPeriodSec
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyWithdrawPermissionCreateOperationFee, kn) {
					currentKey = ffjtWithdrawPermissionCreateOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWithdrawPermissionCreateOperationWithdrawalPeriodSec, kn) {
					currentKey = ffjtWithdrawPermissionCreateOperationWithdrawalPeriodSec
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyWithdrawPermissionCreateOperationWithdrawalLimit, kn) {
					currentKey = ffjtWithdrawPermissionCreateOperationWithdrawalLimit
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyWithdrawPermissionCreateOperationWithdrawFromAccount, kn) {
					currentKey = ffjtWithdrawPermissionCreateOperationWithdrawFromAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWithdrawPermissionCreateOperationPeriodsUntilExpiration, kn) {
					currentKey = ffjtWithdrawPermissionCreateOperationPeriodsUntilExpiration
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWithdrawPermissionCreateOperationPeriodStartTime, kn) {
					currentKey = ffjtWithdrawPermissionCreateOperationPeriodStartTime
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyWithdrawPermissionCreateOperationAuthorizedAccount, kn) {
					currentKey = ffjtWithdrawPermissionCreateOperationAuthorizedAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtWithdrawPermissionCreateOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtWithdrawPermissionCreateOperationAuthorizedAccount:
					goto handle_AuthorizedAccount

				case ffjtWithdrawPermissionCreateOperationPeriodStartTime:
					goto handle_PeriodStartTime

				case ffjtWithdrawPermissionCreateOperationPeriodsUntilExpiration:
					goto handle_PeriodsUntilExpiration

				case ffjtWithdrawPermissionCreateOperationWithdrawFromAccount:
					goto handle_WithdrawFromAccount

				case ffjtWithdrawPermissionCreateOperationWithdrawalLimit:
					goto handle_WithdrawalLimit

				case ffjtWithdrawPermissionCreateOperationWithdrawalPeriodSec:
					goto handle_WithdrawalPeriodSec

				case ffjtWithdrawPermissionCreateOperationFee:
					goto handle_Fee

				case ffjtWithdrawPermissionCreateOperationnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_AuthorizedAccount:

	/* handler: j.AuthorizedAccount type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.AuthorizedAccount.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_PeriodStartTime:

	/* handler: j.PeriodStartTime type=types.Time kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.PeriodStartTime.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_PeriodsUntilExpiration:

	/* handler: j.PeriodsUntilExpiration type=types.UInt32 kind=uint32 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.PeriodsUntilExpiration.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_WithdrawFromAccount:

	/* handler: j.WithdrawFromAccount type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.WithdrawFromAccount.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_WithdrawalLimit:

	/* handler: j.WithdrawalLimit type=types.AssetAmount kind=struct quoted=false*/

	{
		/* Falling back. type=types.AssetAmount kind=struct */
		tbuf, err := fs.CaptureField(t