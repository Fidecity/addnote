package types

//go:generate ffjson $GOFILE

type LimitOrders []LimitOrder

type LimitOrder struct {
	ID          LimitOrderID `json:"id"`
	Seller      AccountID    `json:"seller"`
	Expiration  Time         `json:"expiration"`
	ForS