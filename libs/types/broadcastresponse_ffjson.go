
// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: broadcastresponse.go

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *BroadcastResponse) MarshalJSON() ([]byte, error) {
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
func (j *BroadcastResponse) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"id":`)
	fflib.WriteJsonString(buf, string(j.ID))
	buf.WriteString(`,"block_num":`)
	fflib.FormatBits2(buf, uint64(j.BlockNum), 10, false)
	buf.WriteString(`,"trx_num":`)
	fflib.FormatBits2(buf, uint64(j.TrxNum), 10, false)
	if j.Expired {
		buf.WriteString(`,"expired":true`)
	} else {
		buf.WriteString(`,"expired":false`)
	}
	/* Struct fall back. type=types.SignedTransaction kind=struct */
	buf.WriteString(`,"trx":`)
	err = buf.Encode(&j.Trx)
	if err != nil {
		return err
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtBroadcastResponsebase = iota
	ffjtBroadcastResponsenosuchkey

	ffjtBroadcastResponseID

	ffjtBroadcastResponseBlockNum

	ffjtBroadcastResponseTrxNum

	ffjtBroadcastResponseExpired

	ffjtBroadcastResponseTrx
)

var ffjKeyBroadcastResponseID = []byte("id")

var ffjKeyBroadcastResponseBlockNum = []byte("block_num")

var ffjKeyBroadcastResponseTrxNum = []byte("trx_num")

var ffjKeyBroadcastResponseExpired = []byte("expired")

var ffjKeyBroadcastResponseTrx = []byte("trx")

// UnmarshalJSON umarshall json - template of ffjson
func (j *BroadcastResponse) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *BroadcastResponse) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtBroadcastResponsebase
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
				currentKey = ffjtBroadcastResponsenosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'b':

					if bytes.Equal(ffjKeyBroadcastResponseBlockNum, kn) {
						currentKey = ffjtBroadcastResponseBlockNum
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeyBroadcastResponseExpired, kn) {
						currentKey = ffjtBroadcastResponseExpired
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyBroadcastResponseID, kn) {
						currentKey = ffjtBroadcastResponseID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffjKeyBroadcastResponseTrxNum, kn) {
						currentKey = ffjtBroadcastResponseTrxNum
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBroadcastResponseTrx, kn) {
						currentKey = ffjtBroadcastResponseTrx
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyBroadcastResponseTrx, kn) {
					currentKey = ffjtBroadcastResponseTrx
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyBroadcastResponseExpired, kn) {
					currentKey = ffjtBroadcastResponseExpired
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyBroadcastResponseTrxNum, kn) {
					currentKey = ffjtBroadcastResponseTrxNum
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBroadcastResponseBlockNum, kn) {
					currentKey = ffjtBroadcastResponseBlockNum
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyBroadcastResponseID, kn) {
					currentKey = ffjtBroadcastResponseID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtBroadcastResponsenosuchkey
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

				case ffjtBroadcastResponseID:
					goto handle_ID

				case ffjtBroadcastResponseBlockNum:
					goto handle_BlockNum

				case ffjtBroadcastResponseTrxNum:
					goto handle_TrxNum

				case ffjtBroadcastResponseExpired:
					goto handle_Expired

				case ffjtBroadcastResponseTrx:
					goto handle_Trx

				case ffjtBroadcastResponsenosuchkey:
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

	/* handler: j.ID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.ID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BlockNum:

	/* handler: j.BlockNum type=types.UInt64 kind=uint64 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.BlockNum.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_TrxNum:

	/* handler: j.TrxNum type=types.UInt32 kind=uint32 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.TrxNum.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Expired:

	/* handler: j.Expired type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				j.Expired = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				j.Expired = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Trx:

	/* handler: j.Trx type=types.SignedTransaction kind=struct quoted=false*/

	{
		/* Falling back. type=types.SignedTransaction kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Trx)
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