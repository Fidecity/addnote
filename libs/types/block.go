package types

//go:generate ffjson $GOFILE

type BlockHeader struct {
	TransactionMerkleRoot Buffer     `json:"transaction_merkle_root"`
	Previous              Buffer     `json:"previous"`
	TimeStamp             Time       `json:"timestamp"