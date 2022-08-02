// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: orderbook.go

package types

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *Order) MarshalJSON() ([]byte, error) {
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
func (j *Order) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"base":`)
	fflib.AppendFloat(buf, float64(j.Base), 'g', -1, 64)
	buf.WriteString(`,"quote":`)
	fflib.AppendFloat(buf, float64(j.Quote), 'g', -1, 64)
	buf.WriteString(`,"price":`)
	fflib.AppendFloat(buf, float64(j.Price), 'g', -1, 64)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtOrderbase = iota
	ffjtOrdernosuchkey

	ffjtOrderBase

	ffjtOrderQuote

	ffjtOrderPrice
)

var ffjKeyOrderBase = []byte("base")

var ffjKeyOrderQuote = []byte("quote")

var ffjKeyOrderPrice = []byte("price")

// UnmarshalJSON umarshall json - template of ffjson
func (j *Order) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *Order) UnmarshalJSONFF