
package types

//go:generate ffjson $GOFILE

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"strconv"

	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

type Memo struct {
	From    PublicKey `json:"from"`
	To      PublicKey `json:"to"`
	Nonce   UInt64    `json:"nonce"`
	Message Buffer    `json:"message"`
}

func (p Memo) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(p.From); err != nil {
		return errors.Annotate(err, "encode from")
	}

	if err := enc.Encode(p.To); err != nil {
		return errors.Annotate(err, "encode to")
	}

	if err := enc.Encode(p.Nonce); err != nil {
		return errors.Annotate(err, "encode nonce")
	}

	if err := enc.Encode(p.Message); err != nil {
		return errors.Annotate(err, "encode Message")
	}

	return nil
}

//Encrypt calculates a shared secret by the senders private key
//and the recipients public key, then encrypts the given memo message.
func (p *Memo) Encrypt(priv *PrivateKey, msg string) error {
	sec, err := priv.SharedSecret(&p.To, 16, 16)
	if err != nil {
		return errors.Annotate(err, "SharedSecret")
	}

	iv, blk, err := p.cypherBlock(sec)
	if err != nil {
		return errors.Annotate(err, "cypherBlock")
	}

	buf := []byte(msg)
	digest := sha256.Sum256(buf)
	mode := cipher.NewCBCEncrypter(blk, iv)

	// checksum + msg
	raw := digest[:4]
	raw = append(raw, buf...)

	if len(raw)%16 != 0 {
		raw = pad(raw, 16)
	}

	dst := make([]byte, len(raw))
	mode.CryptBlocks(dst, raw)
	p.Message = dst

	return nil
}

//Decrypt calculates a shared secret by the receivers private key
//and the senders public key, then decrypts the given memo message.
func (p Memo) Decrypt(priv *PrivateKey) (string, error) {
	var counterPartyPubKey PublicKey
	myPubKey := priv.PublicKey()

	if myPubKey.Equal(&p.To) {
		counterPartyPubKey = p.From
	} else if myPubKey.Equal(&p.From) {
		counterPartyPubKey = p.To
	} else {
		return "", errors.New("Invalid counterparty public key, cannot decrypt")
	}

	sec, err := priv.SharedSecret(&counterPartyPubKey, 16, 16)
	if err != nil {
		return "", errors.Annotate(err, "SharedSecret")
	}

	iv, blk, err := p.cypherBlock(sec)
	if err != nil {
		return "", errors.Annotate(err, "cypherBlock")
	}

	mode := cipher.NewCBCDecrypter(blk, iv)
	dst := make([]byte, len(p.Message))
	mode.CryptBlocks(dst, p.Message)

	//verify checksum
	chk1 := dst[:4]
	msg := unpad(dst[4:])
	dig := sha256.Sum256(msg)
	chk2 := dig[:4]