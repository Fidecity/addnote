package types

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/denkhaus/logging"
	"github.com/juju/errors"
	"github.com/pquerna/ffjson/ffjson"
)

var (
	GrapheneMaxInstanceID = UInt64(math.MaxUint64 >> 16)
)

type GrapheneObject interface {
	util.TypeMarshaler
	util.TypeUnmarshaler
	UnmarshalJSON(s []byte) error
	MarshalJSON() ([]byte, error)
	ID() string
	String() string
	ObjectType() ObjectType
	SpaceType() SpaceType
	Instance() UInt64
	Equals(id GrapheneObject) bool
	Valid() bool
}

type GrapheneObjects []GrapheneObject

func (p GrapheneObjects) ToStrings() []string {
	ids := make([]string, len(p))
	for idx, o := range p {
		ids[idx] = o.ID()
	}

	return ids
}

func (p GrapheneObjects) String() string {
	return strings.Join(p.ToStrings(), " ")
}

type ObjectIDs []ObjectID

func (p ObjectIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode ObjectID")
		}
	}

	return nil
}

type ObjectID struct {
	number UInt64
}

func (p ObjectID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.number)); err != nil {
		return errors.Annotate(err, "encode number")
	}

	return nil
}

func (p *ObjectID) Unmarshal(dec *util.TypeDecoder) error {
	var ins uint64
	if err := dec.DecodeUVarint(&ins); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64(ins)
	return nil
}

func (p ObjectID) MarshalJSON() ([]byte, error) {
	return ffjson.Marsha