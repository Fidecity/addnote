// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: accountbalance.go

package types

import (
	"bytes"
	"errors"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *AccountBalance) MarshalJSON() ([]byte, error) {
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
func (j *AccountBalance) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"id":`)

	{

		obj, err = j.ID.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"owner":`)

	{

		obj, err = j.Owner.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"asset_type":`)

	{

		obj, err = j.AssetType.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"balance":`)
	fflib.FormatBits2(buf, uint64(j.Balance), 10, false)
	if j.MaintenanceFlag {
		buf.WriteString(`,"maintenance_flag":true`)
	} else {
		buf.WriteString(`,"maintenance_flag":false`)
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtAccountBalancebase = iota
	ffjtAccountBalancenosuchkey

	ffjtAccountBalanceID

	ffjtAccountBalanceOwner

	ffjtAccountBalanceAssetType

	ffjtAccountBalanceBalance

	ffjtAccountBalanceMaintenanceFlag
)

var ffjKeyAccountBalanceID = []byte("id")

var ffjKeyAccountBalanceOwner = []byte("owner")

var ffjKeyAccountBalanceAssetType = []byte("asset_type")

var ffjKeyAccountBalanceBalance = []byte("balance")

var ffjKeyAccountBalanceMaintenanceFlag = []byte("maintenance_flag")

// UnmarshalJSON umarshall json - template of ffjson
func (j *AccountBalance) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *AccountBalance) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtAccountBalancebase
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
				currentKey = ffjtAccountBalancenosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyAccountBalanceAssetType, kn) {
						currentKey = ffjtAccountBalanceAssetType
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffjKeyAccountBalanceBalance, kn) {
						currentKey = ffjtAccountBalanceBalance
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyAccountBalanceID, kn) {
						currentKey = ffjtAccountBalanceID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffjKeyAccountBalanceMaintenanceFlag, kn) {
						currentKey = ffjtAccountBalanceMaintenanceFlag
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffjKeyAccountBalanceOwner, kn) {
						currentKey = ffjtAccountBalanceOwner
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.AsciiEqualFold(ffjKeyAccountBalanceMaintenanceFlag, kn) {
					currentKey = ffjtAccountBalanceMaintenanceFlag
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountBalanceBalance, kn) {
					currentKey = ffjtAccountBalanceBalance
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountBalanceAssetType, kn) {
					currentKey = ffjtAccountBalanceAssetType
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountBalanceOwner, kn) {
					currentKey = ffjtAccountBalanceOwner
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountBalanceID, kn) {
					currentKey = ffjtAccountBalanceID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtAccountBalancenosuchkey
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

				case ffjtAccountBalanceID:
					goto handle_ID

				case ffjtAccountBalanceOwner:
					goto handle_Owner

				case ffjtAccountBalanceAssetType:
					goto handle_AssetType

				case ffjtAccountBalanceBalance:
					goto handle_Balance

				case ffjtAccountBalanceMaintenanceFlag:
					goto handle_MaintenanceFlag

				case ffjtAccountBalancenosuchkey:
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

handle_ID:

	/* handler: j.ID type=types.AccountBalanceID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.ID.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Owner:

	/* handler: j.Owner type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Owner.U