package types

import (
	"testing"

	"github.com/Assetsadapter/whitecoin-adapter/libs/config"
	"github.com/stretchr/testify/assert"
)

var addresses = []string{
	"BTSFN9r6VYzBK8EKtMewfNbfiGCr56pHDBFi",
}

func TestAddress(t *testing.T) {
	config.SetCurrent(c