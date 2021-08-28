package types

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"

	"github.com/juju/errors"
)

//go:generate ffjson $GOFILE

type DynamicGlobalProperties struct {
	ID                             DynamicGlobalPropertyID `json:"id"`
	CurrentWitness                 WitnessID               `json:"current_witness"`
	LastBudgetTime                 Time     