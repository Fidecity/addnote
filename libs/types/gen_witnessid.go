
// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package types

import (
	"fmt"

	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/denkhaus/logging"
	"github.com/juju/errors"
)

type WitnessID struct {
	ObjectID
}

func (p WitnessID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}

	return nil
}

func (p *WitnessID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeWitness) << 48) | instance)
	return nil
}

type WitnessIDs []WitnessID

func (p WitnessIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode WitnessID")
		}
	}

	return nil
}

func WitnessIDFromObject(ob GrapheneObject) WitnessID {
	id, ok := ob.(*WitnessID)
	if ok {
		return *id
	}

	p := WitnessID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeWitness {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeWitness'", p.ID()))
	}

	return p
}

//NewWitnessID creates an new WitnessID object
func NewWitnessID(id string) GrapheneObject {
	gid := new(WitnessID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"WitnessID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeWitness {
		logging.Errorf(
			"WitnessID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeWitness'", id),
		)
		return nil
	}

	return gid
}