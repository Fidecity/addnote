// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: account.go

package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *Account) MarshalJSON() ([]byte, error) {
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
func (j *Account) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"id":`)

	{

		obj, err = j.ID.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"name":`)

	{

		obj, err = j.Name.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"statistics":`)

	{

		obj, err = j.Statistics.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"membership_expiration_date":`)

	{

		obj, err = j.MembershipExpirationDate.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"network_fee_percentage":`)
	fflib.FormatBits2(buf, uint64(j.NetworkFeePercentage), 10, false)
	buf.WriteString(`,"lifetime_referrer_fee_percentage":`)
	fflib.FormatBits2(buf, uint64(j.LifetimeReferrerFeePercentage), 10, false)
	buf.WriteString(`,"referrer_rewards_percentage":`)
	fflib.FormatBits2(buf, uint64(j.ReferrerRewardsPercentage), 10, false)
	buf.WriteString(`,"top_n_control_flags":`)
	fflib.FormatBits2(buf, uint64(j.TopNControlFlags), 10, false)
	buf.WriteString(`,"whitelisting_accounts":`)
	if j.WhitelistingAccounts != nil {
		buf.WriteString(`[`)
		for i, v := range j.WhitelistingAccounts {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				obj, err = v.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"blacklisting_accounts":`)
	if j.BlacklistingAccounts != nil {
		buf.WriteString(`[`)
		for i, v := range j.BlacklistingAccounts {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				obj, err = v.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"whitelisted_accounts":`)
	if j.WhitelistedAccounts != nil {
		buf.WriteString(`[`)
		for i, v := range j.WhitelistedAccounts {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				obj, err = v.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"blacklisted_accounts":`)
	if j.BlacklistedAccounts != nil {
		buf.WriteString(`[`)
		for i, v := range j.BlacklistedAccounts {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				obj, err = v.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	/* Struct fall back. type=types.AccountOptions kind=struct */
	buf.WriteString(`,"options":`)
	err = buf.Encode(&j.Options)
	if err != nil {
		return err
	}
	buf.WriteString(`,"registrar":`)

	{

		obj, err = j.Registrar.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"referrer":`)

	{

		obj, err = j.Referrer.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"lifetime_referrer":`)

	{

		obj, err = j.LifetimeReferrer.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"cashback_vb":`)

	{

		obj, err = j.CashbackVB.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	/* Struct fall back. type=types.Authority kind=struct */
	buf.WriteString(`,"owner":`)
	err = buf.Encode(&j.Owner)
	if err != nil {
		return err
	}
	/* Struct fall back. type=types.Authority kind=struct */
	buf.WriteString(`,"active":`)
	err = buf.Encode(&j.Active)
	if err != nil {
		return err
	}
	buf.WriteString(`,"owner_special_authority":`)

	{

		obj, err = j.OwnerSpecialAuthority.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"active_special_authority":`)

	{

		obj, err = j.ActiveSpecialAuthority.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtAccountbase = iota
	ffjtAccountnosuchkey

	ffjtAccountID

	ffjtAccountName

	ffjtAccountStatistics

	ffjtAccountMembershipExpirationDate

	ffjtAccountNetworkFeePercentage

	ffjtAccountLifetimeReferrerFeePercentage

	ffjtAccountReferrerRewardsPercentage

	ffjtAccountTopNControlFlags

	ffjtAccountWhitelistingAccounts

	ffjtAccountBlacklistingAccounts

	ffjtAccountWhitelistedAccounts

	ffjtAccountBlacklistedAccounts

	ffjtAccountOptions

	ffjtAccountRegistrar

	ffjtAccountReferrer

	ffjtAccountLifetimeReferrer

	ffjtAccountCashbackVB

	ffjtAccountOwner

	ffjtAccountActive

	ffjtAccountOwnerSpecialAuthority

	ffjtAccountActiveSpecialAuthority
)

var ffjKeyAccountID = []byte("id")

var ffjKeyAccountName = []byte("name")

var ffjKeyAccountStatistics = []byte("statistics")

var ffjKeyAccountMembershipExpirationDate = []byte("membership_expiration_date")

var ffjKeyAccountNetworkFeePercentage = []byte("network_fee_percentage")

var ffjKeyAccountLifetimeReferrerFeePercentage = []byte("lifetime_referrer_fee_percentage")

var ffjKeyAccountReferrerRewardsPercentage = []byte("referrer_rewards_percentage")

var ffjKeyAccountTopNControlFlags = []byte("top_n_control_flags")

var ffjKeyAccountWhitelistingAccounts = []byte("whitelisting_accounts")

var ffjKeyAccountBlacklistingAccounts = []byte("blacklisting_accounts")

var ffjKeyAccountWhitelistedAccounts = []byte("whitelisted_accounts")

var ffjKeyAccountBlacklistedAccounts = []byte("blacklisted_accounts")

var ffjKeyAccountOptions = []byte("options")

var ffjKeyAccountRegistrar = []byte("registrar")

var ffjKeyAccountReferrer = []byte("referrer")

var ffjKeyAccountLifetimeReferrer = []byte("lifetime_referrer")

var ffjKeyAccountCashbackVB = []byte("cashback_vb")

var ffjKeyAccountOwner = []byte("owner")

var ffjKeyAccountActive = []byte("active")

var ffjKeyAccountOwnerSpecialAuthority = []byte("owner_special_authority")

var ffjKeyAccountActiveSpecialAuthority = []byte("active_special_authority")

// UnmarshalJSON umarshall json - template of ffjson
func (j *Account) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *Account) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtAccountbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtAccountnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyAccountActive, kn) {
						currentKey = ffjtAccountActive
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountActiveSpecialAuthority, kn) {
						currentKey = ffjtAcco