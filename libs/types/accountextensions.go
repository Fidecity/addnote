package types

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

type NullExtension struct{}

func (p NullExtension) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(uint8(AccountCreateExtensionsNullExt)); err != nil {
		return errors.Annotate(err, "encode AccountCreateExtensionsNullExt")
	}

	return nil
}

type BuybackOptions struct {
	AssetToBuy       AssetID   `json:"asset_to_buy"`
	AssetToBuyIssuer Accoun