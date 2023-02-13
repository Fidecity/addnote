package types

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/Assetsadapter/whitecoin-adapter/encoding"
	"github.com/pkg/errors"
)

// On the BitShares blockchains there are no addresses, but objects identified by a unique id,
// an type and a space in the form: space.type.id
type ObjectID struct {
	Space uint64
	Type  uint64
	ID    uint64
}

func (o ObjectID) String() string {
	return fmt.Sprintf("%d.%d.%d", o.Space, o.Type, o.ID)
}

func (o *ObjectID) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.String())
}

func (o *ObjectID) UnmarshalJSON(s []byte) error {
	str, err := unquote(string(s))
	if err != nil {
		return errors.Errorf("unable to parse ObjectID from %s", s)
	}

	objectID, err := ParseObje