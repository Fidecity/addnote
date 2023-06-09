
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

type AccountID struct {
	ObjectID
}

func (p AccountID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}

	return nil
}

func (p *AccountID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeAccount) << 48) | instance)
	return nil
}

type AccountIDs []AccountID

func (p AccountIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode AccountID")
		}
	}

	return nil
}

func AccountIDFromObject(ob GrapheneObject) AccountID {
	id, ok := ob.(*AccountID)
	if ok {
		return *id
	}

	p := AccountID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeAccount {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeAccount'", p.ID()))
	}

	return p
}

//NewAccountID creates an new AccountID object
func NewAccountID(id string) GrapheneObject {
	gid := new(AccountID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"AccountID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeAccount {
		logging.Errorf(
			"AccountID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeAccount'", id),
		)
		return nil
	}

	return gid
}