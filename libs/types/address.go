
package types

import (
	"bytes"
	"crypto/sha512"
	"fmt"

	"github.com/Assetsadapter/whitecoin-adapter/libs/config"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/btcsuite/btcutil/base58"
	"github.com/juju/errors"
	"github.com/pquerna/ffjson/ffjson"
)

//An Address is a shortened non-reversable hash of a PublicKey.
type Address struct {
	prefix   string
	data     []byte
	checksum []byte
}

func (p *Address) UnmarshalJSON(data []byte) error {
	var add string

	if err := ffjson.Unmarshal(data, &add); err != nil {
		return errors.Annotate(err, "Unmarshal")
	}

	address, err := NewAddressFromString(add)
	if err != nil {
		return errors.Annotate(err, "NewAddressFromString")
	}

	p.data = address.data
	p.prefix = address.prefix
	p.checksum = address.checksum
	return nil
}

func (p Address) MarshalJSON() ([]byte, error) {
	return ffjson.Marshal(p.String())
}

func (p Address) Marshal(enc *util.TypeEncoder) error {
	return enc.Encode(p.data)
}

func (p Address) String() string {
	b := append(p.data, p.checksum...)
	return fmt.Sprintf("%s%s", p.prefix, base58.Encode(b))
}

func (p Address) Bytes() []byte {
	return p.data
}

func NewAddress(pub *PublicKey) (*Address, error) {
	buf512 := sha512.Sum512(pub.Bytes())
	data, err := util.Ripemd160(buf512[:])
	if err != nil {
		return nil, errors.Annotate(err, "Ripemd160")
	}

	chk1, err := util.Ripemd160Checksum(data)
	if err != nil {
		return nil, errors.Annotate(err, "Ripemd160Checksum")
	}

	ad := Address{
		prefix:   pub.prefix,
		data:     data,
		checksum: chk1,
	}

	return &ad, nil
}

// NewAddressFromString creates a new Address from string
// e.g.("BTSFN9r6VYzBK8EKtMewfNbfiGCr56pHDBFi")
func NewAddressFromString(add string) (*Address, error) {
	cnf := config.Current()
	if cnf == nil {
		return nil, ErrChainConfigIsUndefined
	}

	prefixChain := cnf.Prefix
	prefix := add[:len(prefixChain)]

	if prefix != prefixChain {
		return nil, ErrAddressChainPrefixMismatch
	}

	b58 := base58.Decode(add[len(prefixChain):])
	if len(b58) < 5 {
		return nil, ErrInvalidAddress
	}

	chk1 := b58[len(b58)-4:]
	data := b58[:len(b58)-4]

	chk2, err := util.Ripemd160Checksum(data)
	if err != nil {
		return nil, errors.Annotate(err, "Ripemd160Checksum")
	}

	if !bytes.Equal(chk1, chk2) {
		return nil, ErrInvalidAddress
	}

	a := Address{
		data:     data,
		prefix:   prefix,
		checksum: chk1,
	}

	return &a, nil
}