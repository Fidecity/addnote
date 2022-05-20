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
			} else if tok == fflib.FFTok_right_brack