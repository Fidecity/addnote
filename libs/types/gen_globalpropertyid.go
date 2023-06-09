
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

type GlobalPropertyID struct {
	ObjectID
}

func (p GlobalPropertyID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}

	return nil
}

func (p *GlobalPropertyID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeGlobalProperty) << 48) | instance)
	return nil
}

type GlobalPropertyIDs []GlobalPropertyID

func (p GlobalPropertyIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode GlobalPropertyID")
		}
	}

	return nil
}

func GlobalPropertyIDFromObject(ob GrapheneObject) GlobalPropertyID {
	id, ok := ob.(*GlobalPropertyID)
	if ok {
		return *id
	}

	p := GlobalPropertyID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeGlobalProperty {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeGlobalProperty'", p.ID()))
	}

	return p
}

//NewGlobalPropertyID creates an new GlobalPropertyID object
func NewGlobalPropertyID(id string) GrapheneObject {
	gid := new(GlobalPropertyID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"GlobalPropertyID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeGlobalProperty {
		logging.Errorf(
			"GlobalPropertyID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeGlobalProperty'", id),
		)
		return nil
	}

	return gid
}