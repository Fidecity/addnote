package types

import (
	"github.com/Assetsadapter/whitecoin-adapter/encoding"
	"github.com/pkg/errors"
)

type Transaction struct {
	RefBlockNum    uint16     `json:"ref_block_num"`
	RefBlockPrefix uint32     `json:"ref_block_prefix"`
	Expiration     Time       `json:"expiration"`
	Operations     Operations `json:"operations"`
	Signatures     []string   `json:"