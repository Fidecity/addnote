
// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: withdrawpermissionupdateoperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *WithdrawPermissionUpdateOperation) MarshalJSON() ([]byte, error) {
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
func (j *WithdrawPermissionUpdateOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "withdraw_from_account":`)

	{

		obj, err = j.WithdrawFromAccount.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"authorized_account":`)

	{

		obj, err = j.AuthorizedAccount.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"permission_to_update":`)

	{

		obj, err = j.PermissionToUpdate.MarshalJSON()
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
	ffjtWithdrawPermissionUpdateOperationbase = iota
	ffjtWithdrawPermissionUpdateOperationnosuchkey

	ffjtWithdrawPermissionUpdateOperationWithdrawFromAccount

	ffjtWithdrawPermissionUpdateOperationAuthorizedAccount

	ffjtWithdrawPermissionUpdateOperationPermissionToUpdate

	ffjtWithdrawPermissionUpdateOperationWithdrawalLimit

	ffjtWithdrawPermissionUpdateOperationWithdrawalPeriodSec

	ffjtWithdrawPermissionUpdateOperationPeriodStartTime

	ffjtWithdrawPermissionUpdateOperationPeriodsUntilExpiration

	ffjtWithdrawPermissionUpdateOperationFee
)

var ffjKeyWithdrawPermissionUpdateOperationWithdrawFromAccount = []byte("withdraw_from_account")

var ffjKeyWithdrawPermissionUpdateOperationAuthorizedAccount = []byte("authorized_account")

var ffjKeyWithdrawPermissionUpdateOperationPermissionToUpdate = []byte("permission_to_update")

var ffjKeyWithdrawPermissionUpdateOperationWithdrawalLimit = []byte("withdrawal_limit")

var ffjKeyWithdrawPermissionUpdateOperationWithdrawalPeriodSec = []byte("withdrawal_period_sec")

var ffjKeyWithdrawPermissionUpdateOperationPeriodStartTime = []byte("period_start_time")

var ffjKeyWithdrawPermissionUpdateOperationPeriodsUntilExpiration = []byte("periods_until_expiration")

var ffjKeyWithdrawPermissionUpdateOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *WithdrawPermissionUpdateOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *WithdrawPermissionUpdateOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtWithdrawPermissionUpdateOperationbase
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
				currentKey = ffjtWithdrawPermissionUpdateOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyWithdrawPermissionUpdateOperationAuthorizedAccount, kn) {
						currentKey = ffjtWithdrawPermissionUpdateOperationAuthorizedAccount
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyWithdrawPermissionUpdateOperationFee, kn) {
						currentKey = ffjtWithdrawPermissionUpdateOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffjKeyWithdrawPermissionUpdateOperationPermissionToUpdate, kn) {
						currentKey = ffjtWithdrawPermissionUpdateOperationPermissionToUpdate
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyWithdrawPermissionUpdateOperationPeriodStartTime, kn) {
						currentKey = ffjtWithdrawPermissionUpdateOperationPeriodStartTime
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyWithdrawPermissionUpdateOperationPeriodsUntilExpiration, kn) {
						currentKey = ffjtWithdrawPermissionUpdateOperationPeriodsUntilExpiration
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffjKeyWithdrawPermissionUpdateOperationWithdrawFromAccount, kn) {
						currentKey = ffjtWithdrawPermissionUpdateOperationWithdrawFromAccount
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyWithdrawPermissionUpdateOperationWithdrawalLimit, kn) {
						currentKey = ffjtWithdrawPermissionUpdateOperationWithdrawalLimit
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyWithdrawPermissionUpdateOperationWithdrawalPeriodSec, kn) {
						currentKey = ffjtWithdrawPermissionUpdateOperationWithdrawalPeriodSec
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyWithdrawPermissionUpdateOperationFee, kn) {
					currentKey = ffjtWithdrawPermissionUpdateOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWithdrawPermissionUpdateOperationPeriodsUntilExpiration, kn) {
					currentKey = ffjtWithdrawPermissionUpdateOperationPeriodsUntilExpiration
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWithdrawPermissionUpdateOperationPeriodStartTime, kn) {
					currentKey = ffjtWithdrawPermissionUpdateOperationPeriodStartTime
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWithdrawPermissionUpdateOperationWithdrawalPeriodSec, kn) {
					currentKey = ffjtWithdrawPermissionUpdateOperationWithdrawalPeriodSec
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyWithdrawPermissionUpdateOperationWithdrawalLimit, kn) {
					currentKey = ffjtWithdrawPermissionUpdateOperationWithdrawalLimit
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWithdrawPermissionUpdateOperationPermissionToUpdate, kn) {
					currentKey = ffjtWithdrawPermissionUpdateOperationPermissionToUpdate
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyWithdrawPermissionUpdateOperationAuthorizedAccount, kn) {
					currentKey = ffjtWithdrawPermissionUpdateOperationAuthorizedAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyWithdrawPermissionUpdateOperationWithdrawFromAccount, kn) {
					currentKey = ffjtWithdrawPermissionUpdateOperationWithdrawFromAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtWithdrawPermissionUpdateOperationnosuchkey
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

				case ffjtWithdrawPermissionUpdateOperationWithdrawFromAccount:
					goto handle_WithdrawFromAccount

				case ffjtWithdrawPermissionUpdateOperationAuthorizedAccount:
					goto handle_AuthorizedAccount

				case ffjtWithdrawPermissionUpdateOperationPermissionToUpdate:
					goto handle_PermissionToUpdate

				case ffjtWithdrawPermissionUpdateOperationWithdrawalLimit:
					goto handle_WithdrawalLimit

				case ffjtWithdrawPermissionUpdateOperationWithdrawalPeriodSec:
					goto handle_WithdrawalPeriodSec

				case ffjtWithdrawPermissionUpdateOperationPeriodStartTime:
					goto handle_PeriodStartTime

				case ffjtWithdrawPermissionUpdateOperationPeriodsUntilExpiration:
					goto handle_PeriodsUntilExpiration

				case ffjtWithdrawPermissionUpdateOperationFee:
					goto handle_Fee

				case ffjtWithdrawPermissionUpdateOperationnosuchkey:
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

handle_PermissionToUpdate:

	/* handler: j.PermissionToUpdate type=types.WithdrawPermissionID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.PermissionToUpdate.UnmarshalJSON(tbuf)
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
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.WithdrawalLimit)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_WithdrawalPeriodSec:

	/* handler: j.WithdrawalPeriodSec type=types.UInt32 kind=uint32 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.WithdrawalPeriodSec.UnmarshalJSON(tbuf)
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

handle_Fee:

	/* handler: j.Fee type=types.AssetAmount kind=struct quoted=false*/

	{
		/* Falling back. type=types.AssetAmount kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Fee)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}