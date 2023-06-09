
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

type WorkerID struct {
	ObjectID
}

func (p WorkerID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}

	return nil
}

func (p *WorkerID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeWorker) << 48) | instance)
	return nil
}

type WorkerIDs []WorkerID

func (p WorkerIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode WorkerID")
		}
	}

	return nil
}

func WorkerIDFromObject(ob GrapheneObject) WorkerID {
	id, ok := ob.(*WorkerID)
	if ok {
		return *id
	}

	p := WorkerID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeWorker {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeWorker'", p.ID()))
	}

	return p
}

//NewWorkerID creates an new WorkerID object
func NewWorkerID(id string) GrapheneObject {
	gid := new(WorkerID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"WorkerID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeWorker {
		logging.Errorf(
			"WorkerID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeWorker'", id),
		)
		return nil
	}

	return gid
}