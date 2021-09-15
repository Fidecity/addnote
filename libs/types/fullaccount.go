package types

//go:generate ffjson $GOFILE

import (
	"encoding/json"

	"github.com/juju/errors"
	"github.com/pquerna/ffjson/ffjson"
)

type FullAccountInfos []FullAccountInfo

type FullAccountInfo struct {
	ID          AccountID
	AccountInfo AccountInfo
}

type AccountInfo struct {
	Account              Account               `json:"account"`
	RegistrarName        String                `json:"registrar_name"`
	ReferrerName         String                `json:"referrer_name"`
	LifetimeReferrerName String                `json:"lifetime_referrer_name"`
	CashbackBalance      VestingBalance        `json:"cashback_balance"`
	Balances             AccountBalances       `json:"balances"`
	VestingBalances      VestingBalances    