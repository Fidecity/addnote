package types

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
	"github.com/pquerna/ffjson/ffjson"
)

//TODO: CoinSecondsEarned is UInt128! Since golang has no
//128 bit uint, check marshal and implement this later.
type CCDVestingPolicy struct {
	StartClaim                  Time   `json:"start_claim"`
	CoinSecondsEarnedLastUpdate Time   `json:"coin_seconds_earned_last_update"`
	VestingSeconds              UInt32 `json:"vesting_seconds"`
	CoinSecondsEarned           UInt64 `json:"coin_seconds_earned"` //UInt128!!
}

// TODO: check order!
func (p CCDVestingPolicy) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(p.StartClaim); err != nil {
		return errors.Annotate(err, "encode StartClaim")
	}

	if err := enc.Encode(p.VestingSeconds); err != nil {
		return errors.Annotate(err, "encode VestingSeconds")
	}

	// if err := enc.Encode(p.CoinSecondsEarnedLastUpdate); err != nil {
	// 	return errors.Annotate(err, "encode CoinSecondsEarnedLastUpdate")
	// }

	// if err := enc.Encode(p.CoinSecondsEarned); err != nil {
	// 	return errors.Annotate(err, "encode CoinSecondsEarned")
	// }

	return nil
}

type LinearVestingPolicy struct {
	BeginTimestamp         Time   `json:"begin_timestamp"`
	VestingCliffSeconds    UInt32 `json:"vesting_cliff_seconds"`
	VestingDurationSeconds UInt32 `json:"vesting_duration_seconds"`
}

func (p LinearVestingPolicy) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(p.BeginTimestamp); err != nil {
		return errors.Annotate(err, "encode BeginTimestamp")
	}

	if err := enc.Encode(p.VestingCliffSeconds); err != nil {
		return errors.Annotate(err, "encode VestingCliffSeconds")
