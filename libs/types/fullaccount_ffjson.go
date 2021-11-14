// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: fullaccount.go

package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *AccountInfo) MarshalJSON() ([]byte, error) {
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
func (j *AccountInfo) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"account":`)

	{

		err = j.Account.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteString(`,"registrar_name":`)

	{

		obj, err = j.RegistrarName.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"referrer_name":`)

	{

		obj, err = j.ReferrerName.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"lifetime_referrer_name":`)

	{

		obj, err = j.LifetimeReferrerName.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	/* Struct fall back. type=types.VestingBalance kind=struct */
	buf.WriteString(`,"cashback_balance":`)
	err = buf.Encode(&j.CashbackBalance)
	if err != nil {
		return err
	}
	buf.WriteString(`,"balances":`)
	if j.Balances != nil {
		buf.WriteString(`[`)
		for i, v := range j.Balances {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				err = v.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"vesting_balances":`)
	if j.VestingBalances != nil {
		buf.WriteString(`[`)
		for i, v := range j.VestingBalances {
			if i != 0 {
				buf.WriteString(`,`)
			}
			/* Struct fall back. type=types.VestingBalance kind=struct */
			err = buf.Encode(&v)
			if err != nil {
				return err
			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"limit_orders":`)
	if j.LimitOrders != nil {
		buf.WriteString(`[`)
		for i, v := range j.LimitOrders {
			if i != 0 {
				buf.WriteString(`,`)
			}
			/* Struct fall back. type=types.LimitOrder kind=struct */
			err = buf.Encode(&v)
			if err != nil {
				return err
			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"call_orders":`)
	if j.CallOrders != nil {
		buf.WriteString(`[`)
		for i, v := range j.CallOrders {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				err = v.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"settle_orders":`)
	if j.SettleOrders != nil {
		buf.WriteString(`[`)
		for i, v := range j.SettleOrders {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				err = v.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"statistics":`)

	{

		err = j.Statistics.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteString(`,"assets":`)
	if j.Assets != nil {
		buf.WriteString(`[`)
		for i, v := range j.Assets {
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
	buf.WriteByte('}')
	return nil
}

const (
	ffjtAccountInfobase = iota
	ffjtAccountInfonosuchkey

	ffjtAccountInfoAccount

	ffjtAccountInfoRegistrarName

	ffjtAccountInfoReferrerName

	ffjtAccountInfoLifetimeReferrerName

	ffjtAccountInfoCashbackBalance

	ffjtAccountInfoBalances

	ffjtAccountInfoVestingBalances

	ffjtAccountInfoLimitOrders

	ffjtAccountInfoCallOrders

	ffjtAccountInfoSettleOrders

	ffjtAccountInfoStatistics

	ffjtAccountInfoAssets
)

var ffjKeyAccountInfoAccount = []byte("account")

var ffjKeyAccountInfoRegistrarName = []byte("registrar_name")

var ffjKeyAccountInfoReferrerName = []byte("referrer_name")

var ffjKeyAccountInfoLifetimeReferrerName = []byte("lifetime_referrer_name")

var ffjKeyAccountInfoCashbackBalance = []byte("cashback_balance")

var ffjKeyAccountInfoBalances = []byte("balances")

var ffjKeyAccountInfoVestingBalances = []byte("vesting_balances")

var ffjKeyAccountInfoLimitOrders = []byte("limit_orders")

var ffjKeyAccountInfoCallOrders = []byte("call_orders")

var ffjKeyAccountInfoSettleOrders = []byte("settle_orders")

var ffjKeyAccountInfoStatistics = []byte("statistics")

var ffjKeyAccountInfoAssets = []byte("assets")

// UnmarshalJSON umarshall json - template of ffjson
func (j *AccountInfo) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *AccountInfo) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtAccountInfobase
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
				currentKey = ffjtAccountInfonosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyAccountInfoAccount, kn) {
						currentKey = ffjtAccountInfoAccount
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountInfoAssets, kn) {
						currentKey = ffjtAccountInfoAssets
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffjKeyAccountInfoBalances, kn) {
						currentKey = ffjtAccountInfoBalances
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffjKeyAccountInfoCashbackBalance, kn) {
						currentKey = ffjtAccountInfoCashbackBalance
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountInfoCallOrders, kn) {
						currentKey = ffjtAccountInfoCallOrders
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffjKeyAccountInfoLifetimeReferrerName, kn) {
						currentKey = ffjtAccountInfoLifetimeReferrerName
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountInfoLimitOrders, kn) {
						currentKey = ffjtAccountInfoLimitOrders
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffjKeyAccountInfoRegistrarName, kn) {
						currentKey = ffjtAccountInfoRegistrarName
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountInfoReferrerName, kn) {
						currentKey = ffjtAccountInfoReferrerName
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeyAccountInfoSettleOrders, kn) {
						currentKey = ffjtAccountInfoSettleOrders
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountInfoStatistics, kn) {
						currentKey = ffjtAccountInfoStatistics
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'v':

					if bytes.Equal(ffjKeyAccountInfoVestingBalances, kn) {
						currentKey = ffjtAccountInfoVestingBalances
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyAccountInfoAssets, kn) {
					currentKey = ffjtAccountInfoAssets
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountInfoStatistics, kn) {
					currentKey = ffjtAccountInfoStatistics
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountInfoSettleOrders, kn) {
					currentKey = ffjtAccountInfoSettleOrders
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountInfoCallOrders, kn) {
					currentKey = ffjtAccountInfoCallOrders
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountInfoLimitOrders, kn) {
					currentKey = ffjtAccountInfoLimitOrders
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountInfoVestingBalances, kn) {
					currentKey = ffjtAccountInfoVestingBalances
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountInfoBalances, kn) {
					currentKey = ffjtAccountInfoBalances
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountInfoCashbackBalance, kn) {
					currentKey = ffjtAccountInfoCashbackBalance
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyAccountInfoLifetimeReferrerName, kn) {
					currentKey = ffjtAccountInfoLifetimeReferrerName
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyAccountInfoReferrerName, kn) {
					currentKey = ffjtAccountInfoReferrerName
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountInfoRegistrarName, kn) {
					currentKey = ffjtAccountInfoRegistrarName
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountInfoAccount, kn) {
					currentKey = ffjtAccountInfoAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtAccountInfonosuchkey
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

				case ffjtAccountInfoAccount:
					goto handle_Account

				case ffjtAccountInfoRegistrarName:
					goto handle_RegistrarName

				case ffjtAccountInfoReferrerName:
					goto handle_ReferrerName

				case ffjtAccountInfoLifetimeReferrerName:
					goto handle_LifetimeReferrerName

				case ffjtAccountInfoCashbackBalance:
					goto handle_CashbackBalance

				case ffjtAccountInfoBalances:
					goto handle_Balances

				case ffjtAccountInfoVestingBalances:
					goto handle_VestingBalances

				case ffjtAccountInfoLimitOrders:
					goto handle_LimitOrders

				case ffjtAccountInfoCallOrders:
					goto handle_CallOrders

				case ffjtAccountInfoSettleOrders:
					goto handle_SettleOrders

				case ffjtAccountInfoStatistics:
					goto handle_Statistics

				case ffjtAccountInfoAssets:
					goto handle_Assets

				case ffjtAccountInfonosuchkey:
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

handle_Account:

	/* handler: j.Account type=types.Account kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			err = j.Account.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_RegistrarName:

	/* handler: j.RegistrarName type=types.String kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.RegistrarName.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ReferrerName:

	/* handler: j.ReferrerName type=types.String kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.ReferrerName.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LifetimeReferrerName:

	/* handler: j.LifetimeReferrerName type=types.String kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.LifetimeReferrerName.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CashbackBalance:

	/* handler: j.CashbackBalance type=types.VestingBalance kind=struct quoted=false*/

	{
		/* Falling back. type=types.VestingBalance kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.CashbackBalance)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Balances:

	/* handler: j.Balances type=types.AccountBalances kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for AccountBalances", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Balances = nil
		} else {

			j.Balances = []AccountBalance{}

			wantVal := true

			for {

				var tmpJBalances AccountBalance

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJBalances type=types.AccountBalance kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						err = tmpJBalances.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
						if err != nil {
							return err
						}
					}
					state = fflib.FFParse_after_value
				}

				j.Balances = append(j.Balances, tmpJBalances)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_VestingBalances:

	/* handler: j.VestingBalances type=types.VestingBalances kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for VestingBalances", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.VestingBalances = nil
		} else {

			j.VestingBalances = []VestingBalance{}

			wantVal := true

			for {

				var tmpJVestingBalances VestingBalance

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJVestingBalances type=types.VestingBalance kind=struct quoted=false*/

				{
					/* Falling back. type=types.VestingBalance kind=struct */
					tbuf, err := fs.CaptureField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}

					err = json.Unmarshal(tbuf, &tmpJVestingBalances)
					if err != nil {
						return fs.WrapErr(err)
					}
				}

				j.VestingBalances = append(j.VestingBalances, tmpJVestingBalances)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LimitOrders:

	/* handler: j.LimitOrders type=types.LimitOrders kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for LimitOrders", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.LimitOrders = nil
		} else {

			j.LimitOrders = []LimitOrder{}

			wantVal := true

			for {

				var tmpJLimitOrders LimitOrder

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJLimitOrders type=types.LimitOrder kind=struct quoted=false*/

				{
					/* Falling back. type=types.LimitOrder kind=struct */
					tbuf, err := fs.CaptureField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}

					err = json.Unmarshal(tbuf, &tmpJLimitOrders)
					if err != nil {
						return fs.WrapErr(err)
					}
				}

				j.LimitOrders = append(j.LimitOrders, tmpJLimitOrders)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CallOrders:

	/* handler: j.CallOrders type=types.CallOrders kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for CallOrders", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.CallOrders = nil
		} else {

			j.CallOrders = []CallOrder{}

			wantVal := true

			for {

				var tmpJCallOrders CallOrder

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJCallOrders type=types.CallOrder kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						err = tmpJCallOrders.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
						if err != nil {
							return err
						}
					}
					state = fflib.FFParse_after_value
				}

				j.CallOrders = append(j.CallOrders, tmpJCallOrders)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_SettleOrders:

	/* handler: j.SettleOrders type=types.ForceSettlementOrders kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ForceSettlementOrders", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.SettleOrders = nil
		} else {

			j.SettleOrders = []ForceSettlementOrder{}

			wantVal := true

			for {

				var tmpJSettleOrders ForceSettlementOrder

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJSettleOrders type=types.ForceSettlementOrder kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						err = tmpJSettleOrders.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
						if err != nil {
							return err
						}
					}
					state = fflib.FFParse_after_value
				}

				j.SettleOrders = append(j.SettleOrders, tmpJSettleOrders)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Statistics:

	/* handler: j.Statistics type=types.AccountStatistics kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			err = j.Statistics.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Assets:

	/* handler: j.Assets type=types.AssetIDs kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for AssetIDs", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Assets = nil
		} else {

			j.Assets = []AssetID{}

			wantVal := true

			for {

				var tmpJAssets AssetID

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJAssets type=types.AssetID kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJAssets.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.Assets = append(j.Assets, tmpJAssets)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr