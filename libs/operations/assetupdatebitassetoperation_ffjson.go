
// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: assetupdatebitassetoperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *AssetUpdateBitassetOperation) MarshalJSON() ([]byte, error) {
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
func (j *AssetUpdateBitassetOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "asset_to_update":`)

	{

		obj, err = j.AssetToUpdate.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"issuer":`)

	{

		obj, err = j.Issuer.MarshalJSON()
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
	/* Struct fall back. type=types.BitassetOptions kind=struct */
	buf.WriteString(`,"new_options":`)
	err = buf.Encode(&j.NewOptions)
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
	ffjtAssetUpdateBitassetOperationbase = iota
	ffjtAssetUpdateBitassetOperationnosuchkey

	ffjtAssetUpdateBitassetOperationAssetToUpdate

	ffjtAssetUpdateBitassetOperationIssuer

	ffjtAssetUpdateBitassetOperationExtensions

	ffjtAssetUpdateBitassetOperationNewOptions

	ffjtAssetUpdateBitassetOperationFee
)

var ffjKeyAssetUpdateBitassetOperationAssetToUpdate = []byte("asset_to_update")

var ffjKeyAssetUpdateBitassetOperationIssuer = []byte("issuer")

var ffjKeyAssetUpdateBitassetOperationExtensions = []byte("extensions")

var ffjKeyAssetUpdateBitassetOperationNewOptions = []byte("new_options")

var ffjKeyAssetUpdateBitassetOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *AssetUpdateBitassetOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *AssetUpdateBitassetOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtAssetUpdateBitassetOperationbase
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
				currentKey = ffjtAssetUpdateBitassetOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyAssetUpdateBitassetOperationAssetToUpdate, kn) {
						currentKey = ffjtAssetUpdateBitassetOperationAssetToUpdate
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeyAssetUpdateBitassetOperationExtensions, kn) {
						currentKey = ffjtAssetUpdateBitassetOperationExtensions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyAssetUpdateBitassetOperationFee, kn) {
						currentKey = ffjtAssetUpdateBitassetOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyAssetUpdateBitassetOperationIssuer, kn) {
						currentKey = ffjtAssetUpdateBitassetOperationIssuer
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffjKeyAssetUpdateBitassetOperationNewOptions, kn) {
						currentKey = ffjtAssetUpdateBitassetOperationNewOptions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyAssetUpdateBitassetOperationFee, kn) {
					currentKey = ffjtAssetUpdateBitassetOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetUpdateBitassetOperationNewOptions, kn) {
					currentKey = ffjtAssetUpdateBitassetOperationNewOptions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetUpdateBitassetOperationExtensions, kn) {
					currentKey = ffjtAssetUpdateBitassetOperationExtensions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetUpdateBitassetOperationIssuer, kn) {
					currentKey = ffjtAssetUpdateBitassetOperationIssuer
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetUpdateBitassetOperationAssetToUpdate, kn) {
					currentKey = ffjtAssetUpdateBitassetOperationAssetToUpdate
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtAssetUpdateBitassetOperationnosuchkey
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

				case ffjtAssetUpdateBitassetOperationAssetToUpdate:
					goto handle_AssetToUpdate

				case ffjtAssetUpdateBitassetOperationIssuer:
					goto handle_Issuer

				case ffjtAssetUpdateBitassetOperationExtensions:
					goto handle_Extensions

				case ffjtAssetUpdateBitassetOperationNewOptions:
					goto handle_NewOptions

				case ffjtAssetUpdateBitassetOperationFee:
					goto handle_Fee

				case ffjtAssetUpdateBitassetOperationnosuchkey:
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

handle_AssetToUpdate:

	/* handler: j.AssetToUpdate type=types.AssetID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.AssetToUpdate.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Issuer:

	/* handler: j.Issuer type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Issuer.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
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

handle_NewOptions:

	/* handler: j.NewOptions type=types.BitassetOptions kind=struct quoted=false*/

	{
		/* Falling back. type=types.BitassetOptions kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.NewOptions)
		if err != nil {
			return fs.WrapErr(err)
		}
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