// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: committeememberupdateglobalparametersoperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *ChainParameters) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *ChainParameters) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	if j.AllowNonMemberWhitelists {
		buf.WriteString(`{"allow_non_member_whitelists":true`)
	} else {
		buf.WriteString(`{"allow_non_member_whitelists":false`)
	}
	if j.CountNonMemberVotes {
		buf.WriteString(`,"count_non_member_votes":true`)
	} else {
		buf.WriteString(`,"count_non_member_votes":false`)
	}
	buf.WriteString(`,"extensions":`)

	{

		obj, err = j.Extensions.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	/* Struct fall back. type=types.FeeSchedule kind=struct */
	buf.WriteString(`,"current_fees":`)
	err = buf.Encode(&j.CurrentFees)
	if err != nil {
		return err
	}
	buf.WriteString(`,"account_fee_scale_bitshifts":`)
	fflib.FormatBits2(buf, uint64(j.AccountFeeScaleBitshifts), 10, false)
	buf.WriteString(`,"block_interval":`)
	fflib.FormatBits2(buf, uint64(j.BlockInterval), 10, false)
	buf.WriteString(`,"maintenance_skip_slots":`)
	fflib.FormatBits2(buf, uint64(j.MaintenanceSkipSlots), 10, false)
	buf.WriteString(`,"max_authority_depth":`)
	fflib.FormatBits2(buf, uint64(j.MaxAuthorityDepth), 10, false)
	buf.WriteString(`,"maximum_asset_feed_publishers":`)
	fflib.FormatBits2(buf, uint64(j.MaximumAssetFeedPublishers), 10, false)
	buf.WriteString(`,"maximum_asset_whitelist_authorities":`)
	fflib.FormatBits2(buf, uint64(j.MaximumAssetWhitelistAuthorities), 10, false)
	buf.WriteString(`,"accounts_per_fee_scale":`)
	fflib.FormatBits2(buf, uint64(j.AccountsPerFeeScale), 10, false)
	buf.WriteString(`,"lifetime_referrer_percent_of_fee":`)
	fflib.FormatBits2(buf, uint64(j.LifetimeReferrerPercentOfFee), 10, false)
	buf.WriteString(`,"max_predicate_opcode":`)
	fflib.FormatBits2(buf, uint64