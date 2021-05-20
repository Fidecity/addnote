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
	WeightThreshol