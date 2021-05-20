package types

//go:generate ffjson $GOFILE

import (
	"encoding/json"

	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	sort "github.com/emirpasic/gods/utils"
	"github.com/juju/errors"
	"github.com/pquerna/ffjson/ffjson"
)

type Authority struct {
	WeightThreshold UInt32          `json:"weight_threshold"`
	AccountAuths    AccountAuthsMap `json:"account_auths"`
	KeyAuths        KeyAuthsMap     `json:"key_auths"`
	AddressAuths    AddressAuthsMap `json:"address_auths"`
}

func (p Authority) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(p.WeightThreshold); err != nil {
		return errors.Annotate(err, "encode WeightThreshold")
	}
	if err := enc.Encode(p.AccountAuths); err != nil {
		return errors.Annotate(err, "encode AccountAuths")
	}
	if err := enc.Encode(p.KeyAuths); err != nil {
		return errors.Annotate(err, "encode KeyAuths")
	}
	if err := enc.Encode(p.AddressAuths); err != nil {
		return errors.Annotate(err, "encode AddressAuths")
	}

	return nil
}

type KeyAuthsMap map[*PublicKey]UInt16

func (p *KeyAuthsMap) UnmarshalJSON(data []byte) error {
	var auths [][]interface{}
	if err := ffjson.Unmarshal(data, &auths); err != nil {
		return errors.Annotate(err, "unmarshal Auths")
	}

	(*p) = make(map[*PublicKey]UInt16)
	for _, tk := range auths {
		key, ok := tk[0].(string)
		if !ok {
			return ErrInvalidInputType
		}

		weight, ok := tk[1].(float64)
		if !ok {
			return ErrInvalidInputType
		}

		pub, err := NewPublicKeyFromString(key)
		if err != nil {
			return errors.Annotate(err, "NewPublicKeyFromString")
		}

		(*p)[pub] = UInt16(weight