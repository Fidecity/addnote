package types

//go:generate ffjson $GOFILE

type CallOrders []CallOrder
type CallOrder struct {
	ID         CallOrder