package types

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAssetAmount_UnmarshalJSON(t *testing.T) {
	t.Run("amount uint64", func