package types

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
	"github.com/pquerna/ffjson/ffjson"
)

//TODO: CoinSecondsEarned is UInt128! Since golang has no
//128 bit uint, check marshal and implement this later.
type CCDVestingPolicy s