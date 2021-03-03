package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeWorkerCreate] = func() types.Operation {
		op := &WorkerCreateOperation{}
		return op
	}
}

type WorkerCreateOperation struct {
	types.OperationFee
	DailyPay      types.UInt64            `json:"daily_pay"`
	Initializer   types.WorkerInitializer `json:"initializer"`
	Name          strin