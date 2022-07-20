// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: operationhistory.go

package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *OperationHistory) MarshalJSON() ([]byte, error) {
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
func (j *OperationHistory) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
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
	buf.WriteString(`,"block_num":`)
	fflib.FormatBits2(buf, uint64(j.BlockNum), 10, false)
	buf.WriteString(`,"trx_in_block":`)
	fflib.FormatBits2(buf, uint64(j.TrxInBlock), 10, false)
	buf.WriteString(`,"op_in_trx":`)
	fflib.FormatBits2(buf, uint64(j.OpInTrx), 10, false)
	buf.WriteString(`,"virtual_op":`)
	fflib.FormatBits2(buf, uint64(j.VirtualOp), 10, false)
	buf.WriteString(`,"op":`)

	{

		obj, err = j.Operation.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"result":`)
	/* Interface types must use runtime reflection. type=types.OperationResult kind=interface */
	err = buf.Encode(j.Result)
	if err != nil {
		return err
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtOperationHistorybase = iota
	ffjtOperationHistorynosuchkey

	ffjtOperationHistoryID

	ffjtOperationHistoryBlockNum

	ffjtOperationHistoryTrxInBlock

	ffjtOperationHistoryOpInTrx

	ffjtOperationHistoryVirtualOp

	ffjtOperationHistoryOperation

	ffjtOperationHistoryResult
)

var ffjKeyOperationHistoryID = []byte("id")

var ffjKeyOperationHistoryBlockNum = []byte("block_num")

var ffjKeyOperationHistoryTrxInBlock = []byte("trx_in_block")

var ffjKeyOperationHistoryOpInTrx = []byte("op_in_trx")

var ffjKeyOperationHistoryVirtualOp = []byte("virtual_op")

var ffjKeyOperationHistoryOperation = []byte("op")

var ffjKeyOperationHistoryResult = []byte("result")

// UnmarshalJSON umarshall json - template of ffjson
func (j *OperationHistory) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *OperationHistory) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtOperationHistorybase
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
				currentKey = ffjtOperationHistorynosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'b':

					if bytes.Equal(ffjKeyOperationHistoryBlockNum, kn) {
						currentKey = ffjtOperationHistoryBlockNum
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyOperationHistoryID, kn) {
						currentKey = ffjtOperationHistoryID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffjKeyOperationHistoryOpInTrx, kn) {
						currentKey = ffjtOperationHistoryOpInTrx
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyOperationHistoryOperation, kn) {
						currentKey = ffjtOperationHistoryOperation
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffjKeyOperationHistoryResult, kn) {
						currentKey = ffjtOperationHistoryResult
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffjKeyOperationHistoryTrxInBlock, kn) {
						currentKey = ffjtOperationHistoryTrxInBlock
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'v':

					if bytes.Equal(ffjKeyOperationHistoryVirtualOp, kn) {
						currentKey = ffjtOperationHistoryVirtualOp
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyOperationHistoryResult, kn) {
					currentKey = ffjtOperationHistoryResult
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyOperationHistoryOperation, kn) {
					currentKey = ffjtOperationHistoryOperation
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyOperationHistoryVirtualOp, kn) {
					currentKey = ffjtOperationHistoryVirtualOp
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyOperationHistoryOpInTrx, kn) {
					currentKey = ffjtOperationHistoryOpInTrx
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyOperationHistoryTrxInBlock, kn) {
					currentKey = ffjtOperationHistoryTrxInBlock
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyOperationHistoryBlockNum, kn) {
					currentKey = ffjtOperationHistoryBlockNum
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyOperationHistoryID, kn) {
					currentKey = ffjtOperationHistoryID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtOperationHistorynosuchkey
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

				case ffjtOperationHistoryID:
					goto handle_ID

				case ffjtOperationHistoryBlockNum:
					goto handle_BlockNum

				case ffjtOperationHistoryTrxInBlock:
					goto handle_TrxInBlock

				case ffjtOperationHistoryOpInTrx:
					goto handle_OpInTrx

				case ffjtOperationHistoryVirtualOp:
					goto handle_VirtualOp

				case ffjtOperationHistoryOperation:
					goto handle_Operation

				case ffjtOperationHistoryResult:
					goto handle_Result

				case ffjtOperationHistorynosuchkey:
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

	/* handler: j.ID type=types.OperationHistoryID kind=struct quoted=false*/

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

handle_BlockNum:

	/* handler: j.BlockNum type=types.UInt32 kind=uint32 quoted=false*/

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

handle_TrxInBlock:

	/* handler: j.TrxInBlock type=types.UInt16 kind=uint16 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.TrxInBlock.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_OpInTrx:

	/* handler: j.OpInTrx type=types.UInt16 kind=uint16 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.OpInTrx.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_VirtualOp:

	/* handler: j.VirtualOp type=types.UInt16 kind=uint16 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.VirtualOp.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Operation:

	/* handler: j.Operation type=types.OperationEnvelope kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Operation.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Result:

	/* handler: j.Result type=types.OperationResult kind=interface quoted=false*/

	{
		/* Falling back. type=types.OperationResult kind=interface */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Result)
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

// MarshalJSON marshal bytes to json - template
func (j *OperationRelativeHistory) MarshalJSON() ([]byte, error) {
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
func (j *OperationRelativeHistory) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"memo":`)

	{

		obj, err = j.Memo.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"description":`)

	{

		obj, err = j.Description.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"op":`)

	{

		err = j.Op.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtOperationRelativeHistorybase = iota
	ffjtOperationRelativeHistorynosuchkey

	ffjtOperationRelativeHistoryMemo

	ffjtOperationRelativeHistoryDescription

	ffjtOperationRelativeHistoryOp
)

var ffjKeyOperationRelativeHistoryMemo = []byte("memo")

var ffjKeyOperationRelativeHistoryDescription = []byte("description")

var ffjKeyOperationRelativeHistoryOp = []byte("op")

// UnmarshalJSON umarshall json - template of ffjson
func (j *OperationRelativeHistory) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *OperationRelativeHistory) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtOperationRelativeHistorybase
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
				currentKey = ffjtOperationRelativeHistorynosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'd':

					if bytes.Equal(ffjKeyOperationRelativeHistoryDescription, kn) {
						currentKey = ffjtOperationRelativeHistoryDescription
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffjKeyOperationRelativeHistoryMemo, kn) {
						currentKey = ffjtOperationRelativeHistoryMemo
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffjKeyOperationRelativeHistoryOp, kn) {
						currentKey = ffjtOperationRelativeHistoryOp
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyOperationRelativeHistoryOp, kn) {
					currentKey = ffjtOperationRelativeHistoryOp
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyOperationRelativeHistoryDescription, kn) {
					currentKey = ffjtOperationRelativeHistoryDescription
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyOperationRelativeHistoryMemo, kn) {
					currentKey = ffjtOperationRelativeHistoryMemo
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtOperationRelativeHistorynosuchkey
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

				case ffjtOperationRelativeHistoryMemo:
					goto handle_Memo

				case ffjtOperationRelativeHistoryDescription:
					goto handle_Description

				case ffjtOperationRelativeHistoryOp:
					goto handle_Op

				case ffjtOperationRelativeHistorynosuchkey:
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

handle_Memo:

	/* handler: j.Memo type=types.Buffer kind=slice quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Memo.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Description:

	/* handler: j.Description type=types.String kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

		