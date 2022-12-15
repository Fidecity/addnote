package util

import (
	"encoding/binary"
	"io"
	"reflect"

	"github.com/juju/errors"
)

var (
	ErrCannotDecodeNilValue = errors.New("cannot decode nil value")
)

type TypeUnmarshaler interface {
	Unmarshal(enc *TypeDecoder) error
}

type TypeDecoder struct {
	r io.Reader
}

func NewTypeDecoder(r io.Reader) *TypeDecoder {
	return &TypeDecoder{r}
}

func (p *TypeDecoder) DecodeUVarint(v interface{}) error {
	br := ByteReader{p.r}
	val, err := binary.ReadUvarint(br)
	if err != nil {
		return errors.Annotate(err, "ReadUvarint")
	}

	reflect.ValueOf(v).Elem().SetUint(val)
	return nil
}

func (p *TypeDecoder) DecodeNumber(v interface{}) error {
	if err := binary.Read(p.r, binary.LittleEndian, v); err != nil {
		return errors.Annotate(err, "Read")
	}
	return nil
}

func (p *TypeDecoder) Decode(v interface{}) error {
	if v == nil {
		return ErrCannotDecodeNilValue
	}

	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return errors.New("v must be a pointer to a valid decode target")
	}

	if m, ok := v.(TypeUnmarshaler); ok {
		return m.Unmarshal(p)
	}

	trg := val.Elem().Interface()
	switch trg.(type) {
	case int8:
		return p.DecodeNumber(v)
	case int16:
		return p.DecodeNumber(v)
	case int32:
		return p.DecodeNumber(v)
	case int64:
		return p.DecodeNumber(v)
	case uint:
		return p.DecodeNumber(v)
	case uint8:
		return p.DecodeNumber(v)
	case ui