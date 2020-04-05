package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/denkhaus/logging"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAssert] = func() types.Operation {
		op := &AssertOperation{}
		return op
	}
}

// struct account_name_eq_lit_predicate
// {
//    account_id_type account_id;
//    string          name;

//    bool validate()const;
// };

// struct asset_symbol_eq_lit_predicate
// {
//    asset_id_type   asset_id;
//    string          symbol;

//    bool validate()const;

// };

// struct block_id_predicate
// {
//    block_id_ty