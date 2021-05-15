
// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: assetoptions.go

package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *AssetOptions) MarshalJSON() ([]byte, error) {
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
func (j *AssetOptions) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"max_supply":`)
	fflib.FormatBits2(buf, uint64(j.MaxSupply), 10, j.MaxSupply < 0)
	buf.WriteString(`,"max_market_fee":`)
	fflib.FormatBits2(buf, uint64(j.MaxMarketFee), 10, j.MaxMarketFee < 0)
	buf.WriteString(`,"market_fee_percent":`)
	fflib.FormatBits2(buf, uint64(j.MarketFeePercent), 10, false)
	buf.WriteString(`,"flags":`)
	fflib.FormatBits2(buf, uint64(j.Flags), 10, false)
	buf.WriteString(`,"description":`)

	{

		obj, err = j.Description.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	/* Struct fall back. type=types.Price kind=struct */
	buf.WriteString(`,"core_exchange_rate":`)
	err = buf.Encode(&j.CoreExchangeRate)
	if err != nil {
		return err
	}
	buf.WriteString(`,"issuer_permissions":`)
	fflib.FormatBits2(buf, uint64(j.IssuerPermissions), 10, false)
	buf.WriteString(`,"blacklist_authorities":`)
	if j.BlacklistAuthorities != nil {
		buf.WriteString(`[`)
		for i, v := range j.BlacklistAuthorities {
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
	buf.WriteString(`,"whitelist_authorities":`)
	if j.WhitelistAuthorities != nil {
		buf.WriteString(`[`)
		for i, v := range j.WhitelistAuthorities {
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
	buf.WriteString(`,"blacklist_markets":`)
	if j.BlacklistMarkets != nil {
		buf.WriteString(`[`)
		for i, v := range j.BlacklistMarkets {
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
	buf.WriteString(`,"whitelist_markets":`)
	if j.WhitelistMarkets != nil {
		buf.WriteString(`[`)
		for i, v := range j.WhitelistMarkets {
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
	buf.WriteString(`,"extensions":`)

	{

		obj, err = j.Extensions.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtAssetOptionsbase = iota
	ffjtAssetOptionsnosuchkey

	ffjtAssetOptionsMaxSupply

	ffjtAssetOptionsMaxMarketFee

	ffjtAssetOptionsMarketFeePercent

	ffjtAssetOptionsFlags

	ffjtAssetOptionsDescription

	ffjtAssetOptionsCoreExchangeRate

	ffjtAssetOptionsIssuerPermissions

	ffjtAssetOptionsBlacklistAuthorities

	ffjtAssetOptionsWhitelistAuthorities

	ffjtAssetOptionsBlacklistMarkets

	ffjtAssetOptionsWhitelistMarkets

	ffjtAssetOptionsExtensions
)

var ffjKeyAssetOptionsMaxSupply = []byte("max_supply")

var ffjKeyAssetOptionsMaxMarketFee = []byte("max_market_fee")

var ffjKeyAssetOptionsMarketFeePercent = []byte("market_fee_percent")

var ffjKeyAssetOptionsFlags = []byte("flags")

var ffjKeyAssetOptionsDescription = []byte("description")

var ffjKeyAssetOptionsCoreExchangeRate = []byte("core_exchange_rate")

var ffjKeyAssetOptionsIssuerPermissions = []byte("issuer_permissions")

var ffjKeyAssetOptionsBlacklistAuthorities = []byte("blacklist_authorities")

var ffjKeyAssetOptionsWhitelistAuthorities = []byte("whitelist_authorities")

var ffjKeyAssetOptionsBlacklistMarkets = []byte("blacklist_markets")

var ffjKeyAssetOptionsWhitelistMarkets = []byte("whitelist_markets")

var ffjKeyAssetOptionsExtensions = []byte("extensions")

// UnmarshalJSON umarshall json - template of ffjson
func (j *AssetOptions) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *AssetOptions) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtAssetOptionsbase
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
				currentKey = ffjtAssetOptionsnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'b':

					if bytes.Equal(ffjKeyAssetOptionsBlacklistAuthorities, kn) {
						currentKey = ffjtAssetOptionsBlacklistAuthorities
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAssetOptionsBlacklistMarkets, kn) {
						currentKey = ffjtAssetOptionsBlacklistMarkets
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffjKeyAssetOptionsCoreExchangeRate, kn) {
						currentKey = ffjtAssetOptionsCoreExchangeRate
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffjKeyAssetOptionsDescription, kn) {
						currentKey = ffjtAssetOptionsDescription
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeyAssetOptionsExtensions, kn) {
						currentKey = ffjtAssetOptionsExtensions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyAssetOptionsFlags, kn) {
						currentKey = ffjtAssetOptionsFlags
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyAssetOptionsIssuerPermissions, kn) {
						currentKey = ffjtAssetOptionsIssuerPermissions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffjKeyAssetOptionsMaxSupply, kn) {
						currentKey = ffjtAssetOptionsMaxSupply
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAssetOptionsMaxMarketFee, kn) {
						currentKey = ffjtAssetOptionsMaxMarketFee
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAssetOptionsMarketFeePercent, kn) {
						currentKey = ffjtAssetOptionsMarketFeePercent
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffjKeyAssetOptionsWhitelistAuthorities, kn) {
						currentKey = ffjtAssetOptionsWhitelistAuthorities
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAssetOptionsWhitelistMarkets, kn) {
						currentKey = ffjtAssetOptionsWhitelistMarkets
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyAssetOptionsExtensions, kn) {
					currentKey = ffjtAssetOptionsExtensions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetOptionsWhitelistMarkets, kn) {
					currentKey = ffjtAssetOptionsWhitelistMarkets
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetOptionsBlacklistMarkets, kn) {
					currentKey = ffjtAssetOptionsBlacklistMarkets
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetOptionsWhitelistAuthorities, kn) {
					currentKey = ffjtAssetOptionsWhitelistAuthorities
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetOptionsBlacklistAuthorities, kn) {
					currentKey = ffjtAssetOptionsBlacklistAuthorities
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetOptionsIssuerPermissions, kn) {
					currentKey = ffjtAssetOptionsIssuerPermissions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyAssetOptionsCoreExchangeRate, kn) {
					currentKey = ffjtAssetOptionsCoreExchangeRate
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetOptionsDescription, kn) {
					currentKey = ffjtAssetOptionsDescription
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetOptionsFlags, kn) {
					currentKey = ffjtAssetOptionsFlags
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetOptionsMarketFeePercent, kn) {
					currentKey = ffjtAssetOptionsMarketFeePercent
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetOptionsMaxMarketFee, kn) {
					currentKey = ffjtAssetOptionsMaxMarketFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetOptionsMaxSupply, kn) {
					currentKey = ffjtAssetOptionsMaxSupply
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtAssetOptionsnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtAssetOptionsMaxSupply:
					goto handle_MaxSupply

				case ffjtAssetOptionsMaxMarketFee:
					goto handle_MaxMarketFee

				case ffjtAssetOptionsMarketFeePercent:
					goto handle_MarketFeePercent

				case ffjtAssetOptionsFlags:
					goto handle_Flags

				case ffjtAssetOptionsDescription:
					goto handle_Description

				case ffjtAssetOptionsCoreExchangeRate:
					goto handle_CoreExchangeRate

				case ffjtAssetOptionsIssuerPermissions:
					goto handle_IssuerPermissions

				case ffjtAssetOptionsBlacklistAuthorities:
					goto handle_BlacklistAuthorities

				case ffjtAssetOptionsWhitelistAuthorities:
					goto handle_WhitelistAuthorities

				case ffjtAssetOptionsBlacklistMarkets:
					goto handle_BlacklistMarkets

				case ffjtAssetOptionsWhitelistMarkets:
					goto handle_WhitelistMarkets

				case ffjtAssetOptionsExtensions:
					goto handle_Extensions

				case ffjtAssetOptionsnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_MaxSupply:

	/* handler: j.MaxSupply type=types.Int64 kind=int64 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.MaxSupply.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_MaxMarketFee:

	/* handler: j.MaxMarketFee type=types.Int64 kind=int64 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.MaxMarketFee.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}