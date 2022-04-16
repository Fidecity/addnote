package types

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/denkhaus/logging"
	"github.com/juju/errors"
	"github.com/pquerna/ffjson/ffjson"
)

var (
	GrapheneMaxInstanceID = UInt64(math.MaxUint64 >> 16)
)

type GrapheneObject interface {
	util.TypeMarshaler
	util.TypeUnmarshaler
	UnmarshalJ