
package types

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil/base58"
	"github.com/Assetsadapter/whitecoin-adapter/libs/config"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	sort "github.com/emirpasic/gods/utils"
	"github.com/juju/errors"
	"github.com/pquerna/ffjson/ffjson"
)

var (
	nullKeyBytes = bytes.Repeat([]byte{0x0}, 33)
)

type PublicKeys []PublicKey

func (p PublicKeys) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {