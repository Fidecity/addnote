
package types

//go:generate ffjson $GOFILE

import (
	"github.com/pquerna/ffjson/ffjson"

	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

type Extensions struct {
	ext interface{}
}

//TODO refactor and test
func (p Extensions) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(0)); err != nil {
		return errors.Annotate(err, "encode length")
	}

	// for _, ex := range p {
	// 	if err := enc.Encode(ex); err != nil {
	// 		return errors.Annotate(err, "encode Extension")
	// 	}
	// }

	// if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
	// 	return errors.Annotate(err, "encode length")
	// }

	// for _, ex := range p {
	// 	if err := enc.Encode(ex); err != nil {
	// 		return errors.Annotate(err, "encode Extension")
	// 	}
	// }

	return nil
}

func (p Extensions) MarshalJSON() ([]byte, error) {
	if p.ext == nil {