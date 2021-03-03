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

var ffjKeyWitnessUpdateOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *WitnessUpdateOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *WitnessUpdateOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtWitnessUpdateOperationbase
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
				currentKey = ffjtWitnessUpdateOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'f':

					if bytes.Equal(ffjKeyWitnessUpdateOperationFee, kn) {
						currentKey = ffjtWitnessUpdateOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffjKeyWitnessUpdateOperationNewSigningKey, kn) {
						currentKey = ffjtWitnessUpdateOperationNewSigningKey
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyWitnessUpdateOperationNewURL, kn) {
						currentKey = ffjtWitnessUpdateOperationNewURL
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffjKeyWitnessUpdateOperationWitness, kn) {
						currentKey = ffjtWitnessUpdateOperationWitness
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyWitnessUpdateOperationWitnessAccount, kn) {
						currentKey = ffjtWitnessUpdateOperationWitnessAccount
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyWitnessUpdateOperationFee, kn) {
					currentKey = ffjtWitnessUpdateOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWitnessUpdateOperationWitnessAccount, kn) {
					currentKey = ffjtWitnessUpdateOperationWitnessAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWitnessUpdateOperationWitness, kn) {
					currentKey = ffjtWitnessUpdateOperationWitness
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyWitnessUpdateOperationNewURL, kn) {
					currentKey = ffjtWitnessUpdateOperationNewURL
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWitnessUpdateOperationNewSigningKey, kn) {
					currentKey = ffjtWitnessUpdateOperationNewSigningKey
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtWitnessUpdateOperationnosuchkey
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

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_in