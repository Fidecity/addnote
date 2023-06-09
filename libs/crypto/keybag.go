
package crypto

import (
	"bufio"
	"os"
	"strings"

	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

// KeyBag is a PrivateKey collection for signing and verifying purposes.
type KeyBag struct {
	keys []*types.PrivateKey
}

func NewKeyBag() *KeyBag {
	bag := KeyBag{
		keys: make([]*types.PrivateKey, 0),
	}

	return &bag
}

func (p KeyBag) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p.keys))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, k := range p.keys {
		if err := enc.Encode(k); err != nil {
			return errors.Annotate(err, "encode Key")
		}
	}

	return nil
}

func (p *KeyBag) Unmarshal(dec *util.TypeDecoder) error {
	var len uint64
	if err := dec.DecodeUVarint(&len); err != nil {
		return errors.Annotate(err, "decode length")
	}

	for i := len; i > 0; i-- {
		key := types.PrivateKey{}
		if err := dec.Decode(&key); err != nil {
			return errors.Annotate(err, "decode key")
		}

		p.keys = append(p.keys, &key)
	}

	return nil
}

func (b *KeyBag) Add(wifKey string) error {
	privKey, err := types.NewPrivateKeyFromWif(wifKey)
	if err != nil {
		return errors.Annotate(err, "NewPrivateKeyFromWif")
	}

	b.keys = append(b.keys, privKey)
	return nil
}

func (b *KeyBag) Remove(pub string) bool {
	for _, p := range b.Publics() {
		if p.String() == pub {
			for idx, k := range b.keys {
				if k.PublicKey().Equal(&p) {
					b.keys = append(b.keys[:idx], b.keys[idx+1:]...)
					return true
				}
			}
		}
	}

	return false
}

func (b KeyBag) EncryptMemo(memo *types.Memo, msg string) error {
	priv := b.Private(&memo.From)
	if priv == nil {
		return errors.Errorf(
			"private key related to %q not found in KeyBag",
			memo.From,
		)
	}

	if err := memo.Encrypt(priv, msg); err != nil {
		return errors.Annotate(err, "Encrypt")
	}

	return nil
}

func (b *KeyBag) ImportFromFile(path string) error {
	inFile, err := os.Open(path)
	if err != nil {
		return errors.Errorf("import keys from file [%s], %s", path, err)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		key := strings.TrimSpace(strings.Split(scanner.Text(), " ")[0])

		if strings.Contains(key, "/") || strings.Contains(key, "#") || strings.Contains(key, ";") {
			return errors.Errorf("lines should consist of a private key on each line, with an optional whitespace and comment")
		}

		if err := b.Add(key); err != nil {
			return err
		}
	}

	return nil
}

// Public returns a collection of public keys in bag.
func (b KeyBag) Publics() (out types.PublicKeys) {
	for _, k := range b.keys {
		pub := k.PublicKey()
		out = append(out, *pub)
	}
	return
}

// Privates returns a collection of private keys in bag.
func (b KeyBag) Privates() (out types.PrivateKeys) {
	for _, k := range b.keys {
		priv := k
		out = append(out, *priv)
	}
	return
}

// Present checks if a private key associated with the given public key is present
func (b KeyBag) Present(pub *types.PublicKey) bool {
	return b.Private(pub) != nil
}

//Private returns the private key associated with the given public key
func (b KeyBag) Private(pub *types.PublicKey) *types.PrivateKey {
	for _, k := range b.keys {
		if k.PublicKey().Equal(pub) {
			priv := *k
			return &priv
		}
	}
	return nil
}

func (b KeyBag) PrivatesByPublics(pubKeys types.PublicKeys) (out types.PrivateKeys) {
	for _, pub := range pubKeys {
		for _, k := range b.keys {
			if pub.Equal(k.PublicKey()) {
				priv := k
				out = append(out, *priv)
			}
		}
	}

	return
}