// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: assetfundfeepooloperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *AssetFundFeePoolOperation) MarshalJSON() ([]byte, error) {
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
func (j *AssetFundFeePoolOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "amount":`)
	fflib.FormatBits2(buf, uint64(j.Amount), 10, false)
	buf.WriteString(`,"asset_id":`)

	{

		obj, err = j.AssetID.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"extensions":`)

	{

		obj, err = j.Extensions.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"from_account":`)

	{

		obj, err = j.FromAccount.MarshalJSON()
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
	ffjtAssetFundFeePoolOperationbase = iota
	ffjtAssetFundFeePoolOperationnosuchkey

	ffjtAssetFundFeePoolOperationAmount

	ffjtAssetFundFeePoolOperationAssetID

	ffjtAssetFundFeePoolOperationExtensions

	ffjtAssetFundFeePoolOperationFromAccount

	ffjtAssetFundFeePoolOperationFee
)

var ffjKeyAssetFundFeePoolOperationAmount = []byte("amount")

var ffjKeyAssetFundFeePoolOperationAssetID = []byte("asset_id")

var ffjKeyAssetFundFeePoolOperationExtensions = []byte("extensions")

var ffjKeyAssetFundFeePoolOperationFromAccount = []byte("from_account")

var ffjKeyAssetFundFeePoolOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *AssetFundFeePoolOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *AssetFundFeePoolOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtAssetFundFeePoolOperationbase
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
				currentKey = ffjtAssetFundFeePoolOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyAssetFundFeePoolOperationAmount, kn) {
						currentKey = ffjtAssetFundFeePoolOperationAmount
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAssetFundFeePoolOperationAssetID, kn) {
						currentKey = ffjtAssetFundFeePoolOperationAssetID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeyAssetFundFeePoolOperationExtensions, kn) {
						currentKey = ffjtAssetFundFeePoolOperationExtensions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyAssetFundFeePoolOperationFromAccount, kn) {
						currentKey = ffjtAssetFundFeePoolOperationFromAccount
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAssetFundFeePoolOperationFee, kn) {
						currentKey = ffjtAssetFundFeePoolOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.Simple