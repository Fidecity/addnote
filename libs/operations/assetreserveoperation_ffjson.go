
// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: assetreserveoperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *AssetReserveOperation) MarshalJSON() ([]byte, error) {
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
func (j *AssetReserveOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "payer":`)

	{

		obj, err = j.Payer.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	/* Struct fall back. type=types.AssetAmount kind=struct */
	buf.WriteString(`,"amount_to_reserve":`)
	err = buf.Encode(&j.AmountToReserve)
	if err != nil {
		return err
	}
	buf.WriteString(`,"extensions":`)

	{

		obj, err = j.Extensions.MarshalJSON()
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
	ffjtAssetReserveOperationbase = iota
	ffjtAssetReserveOperationnosuchkey

	ffjtAssetReserveOperationPayer

	ffjtAssetReserveOperationAmountToReserve

	ffjtAssetReserveOperationExtensions

	ffjtAssetReserveOperationFee
)

var ffjKeyAssetReserveOperationPayer = []byte("payer")

var ffjKeyAssetReserveOperationAmountToReserve = []byte("amount_to_reserve")

var ffjKeyAssetReserveOperationExtensions = []byte("extensions")

var ffjKeyAssetReserveOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *AssetReserveOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *AssetReserveOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtAssetReserveOperationbase
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
				currentKey = ffjtAssetReserveOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyAssetReserveOperationAmountToReserve, kn) {
						currentKey = ffjtAssetReserveOperationAmountToReserve
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeyAssetReserveOperationExtensions, kn) {
						currentKey = ffjtAssetReserveOperationExtensions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyAssetReserveOperationFee, kn) {
						currentKey = ffjtAssetReserveOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffjKeyAssetReserveOperationPayer, kn) {
						currentKey = ffjtAssetReserveOperationPayer
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyAssetReserveOperationFee, kn) {
					currentKey = ffjtAssetReserveOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetReserveOperationExtensions, kn) {
					currentKey = ffjtAssetReserveOperationExtensions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetReserveOperationAmountToReserve, kn) {
					currentKey = ffjtAssetReserveOperationAmountToReserve
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAssetReserveOperationPayer, kn) {
					currentKey = ffjtAssetReserveOperationPayer
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtAssetReserveOperationnosuchkey
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

				case ffjtAssetReserveOperationPayer:
					goto handle_Payer

				case ffjtAssetReserveOperationAmountToReserve:
					goto handle_AmountToReserve

				case ffjtAssetReserveOperationExtensions:
					goto handle_Extensions

				case ffjtAssetReserveOperationFee:
					goto handle_Fee

				case ffjtAssetReserveOperationnosuchkey:
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

handle_Payer:

	/* handler: j.Payer type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Payer.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_AmountToReserve:

	/* handler: j.AmountToReserve type=types.AssetAmount kind=struct quoted=false*/

	{
		/* Falling back. type=types.AssetAmount kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.AmountToReserve)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Extensions:

	/* handler: j.Extensions type=types.Extensions kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Extensions.UnmarshalJSON(tbuf)
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