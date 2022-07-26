package types

import (
	"encoding/json"
	"fmt"

	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/denkhaus/logging"
	"github.com/juju/errors"
	"github.com/pquerna/ffjson/ffjson"
)

type GetOpFunc func() Operation

var (
	OperationMap = make(map[OperationType]GetOpFunc)
)

type Operation interface {
	util.TypeMarshaler
	SetFee(fee AssetAmount)
	GetFee() AssetA