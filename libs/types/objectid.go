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
	return ffjson.Marshal(p.ID())
}

func (p *ObjectID) UnmarshalJSON(s []byte) error {
	var val string
	if err := ffjson.Unmarshal(s, &val); err != nil {
		return errors.Annotate(err, "Unmarshal")
	}

	if err := p.Parse(val); err != nil {
		return errors.Annotate(err, "Parse")
	}

	return nil
}

func (p *ObjectID) FromObject(ob GrapheneObject) error {
	return p.Parse(ob.ID())
}

func (p *ObjectID) MustFromObject(ob GrapheneObject) {
	if err := p.FromObject(ob); err != nil {
		panic(errors.Annotate(err, "FromObject"))
	}
}

func ObjectIDFromObject(ob GrapheneObject) ObjectID {
	id, ok := ob.(*ObjectID)
	if ok {
		return *id
	}

	p := ObjectID{}
	p.MustFromObject(ob)
	return p
}

func (p ObjectID) Equals(o GrapheneObject) bool {
	return p.ID() == o.ID()
}

func (p ObjectID) Valid() bool {
	return p.number != 0
}

//String, ObjectID implements Stringer
func (p ObjectID) String() string {
	return p.ID()
}

//ID returns the objects string representation
func (p ObjectID) ID() string {
	return fmt.Sprintf("%d.%d.%d",
		p.SpaceType(),
		p.ObjectType(),
		p.Instance(),
	)
}

//ObjectType returns the objects ObjectType
func (p ObjectID) ObjectType() ObjectType {
	return ObjectType(p.number >> 48 & 0x00ff)
}

//SpaceType returns the objects SpaceType
func (p ObjectID) SpaceType() SpaceType {
	return SpaceType(p.number >> 56)
}