// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: bitassetdata.go

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *BitAssetData) MarshalJSON() ([]byte, error) {
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
func (j *BitAssetData) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
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
	buf.WriteString(`,"current_feed_publication_time":`)

	{

		obj, err = j.MembershipExpirationDate.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	if j.IsPredictionMarket {
		buf.WriteString(`,"is_prediction_market":true`)
	} else {
		buf.WriteString(`,"is_prediction_market":false`)
	}
	/* Struct fall back. type=types.Price kind=struct */
	buf.WriteString(`,"settlement_price":`)
	err = buf.Encode(&j.SettlementPrice)
	if err != nil {
		return err
	}
	buf.WriteString(`,"feeds":`)
	if j.Feeds != nil {
		buf.WriteString(`[`)
		for i, v := range j.Feeds {
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
	buf.WriteString(`,"options":`)

	{

		err = j.Options.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	/* Struct fall back. type=types.PriceFeed kind=struct */
	buf.WriteString(`,"current_feed":`)
	err = buf.Encode(&j.CurrentFeed)
	if err != nil {
		return err
	}
	buf.WriteString(`,"force_settled_volume":`)
	fflib.FormatBits2(buf, uint64(j.ForcedSettledVolume), 10, false)
	buf.WriteString(`,"settlement_fund":`)
	fflib.FormatBits2(buf, uint64(j.SettlementFund), 10, false)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtBitAssetDatabase = iota
	ffjtBitAssetDatanosuchkey

	ffjtBitAssetDataID

	ffjtBitAssetDataMembershipExpirationDate

	ffjtBitAssetDataIsPredictionMarket

	ffjtBitAssetDataSettlementPrice

	ffjtBitAssetDataFeeds

	ffjtBitAssetDataOptions

	ffjtBitAssetDataCurrentFeed

	ffjtBitAssetDataForcedSettledVolume

	ffjtBitAssetDataSettlementFund
)

var ffjKeyBitAssetDataID = []byte("id")

var ffjKeyBitAssetDataMembershipExpirationDate = []byte("current_feed_publication_time")

var ffjKeyBitAssetDataIsPredictionMarket = []byte("is_prediction_market")

var ffjKeyBitAssetDataSettlementPrice = []byte("settlement_price")

var ffjKeyBitAssetDataFeeds = []byte("feeds")

var ffjKeyBitAssetDataOptions = []byte("options")

var ffjKeyBitAssetDataCurrentFeed = []byte("current_feed")

var ffjKeyBitAssetDataForcedSettledVolume = []byte("force_settled_volume")

var ffjKeyBitAssetDataSettlementFund = []byte("settlement_fund")

// UnmarshalJSON umarshall json - template of ffjson
func (j *BitAssetData) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *BitAssetData) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtBitAssetDatabase
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
				currentKey = ffjtBitAssetDatanosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffjKeyBitAssetDataMembershipExpirationDate, kn) {
						currentKey = ffjtBitAssetDataMembershipExpirationDate
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBitAssetDataCurrentFeed, kn) {
						currentKey = ffjtBitAssetDataCurrentFeed
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyBitAssetDataFeeds, kn) {
						currentKey = ffjtBitAssetDataFeeds
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBitAssetDataForcedSettledVolume, kn) {
						currentKey = ffjtBitAssetDataForcedSettledVolume
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyBitAssetDataID, kn) {
						currentKey = ffjtBitAssetDataID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBitAssetDataIsPredictionMarket, kn) {
						currentKey = ffjtBitAssetDataIsPredictionMarket
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffjKeyBitAssetDataOptions, kn) {
						currentKey = ffjtBitAssetDataOptions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeyBitAssetDataSettlementPrice, kn) {
						currentKey = ffjtBitAssetDataSettlementPrice
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBitAssetDataSettlementFund, kn) {
						currentKey = ffjtBitAssetDataSettlementFund
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyBitAssetDataSettlementFund, kn) {
					currentKey = ffjtBitAssetDataSettlementFund
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBitAssetDataForcedSettledVolume, kn) {
					currentKey = ffjtBitAssetDataForcedSettledVolume
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyBitAssetDataCurrentFeed, kn) {
					currentKey = ffjtBitAssetDataCurrentFeed
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBitAssetDataOptions, kn) {
					currentKey = ffjtBitAssetDataOptions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBitAssetDataFeeds, kn) {
					currentKey = ffjtBitAssetDataFeeds
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBitAssetDataSettlementPrice, kn) {
					currentKey = ffjtBitAssetDataSettlementPrice
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBitAssetDataIsPredictionMarket, kn) {
					currentKey = ffjtBitAssetDataIsPredictionMarket
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyBitAssetDataMembershipExpirationDate, kn) {
					currentKey = ffjtBitAssetDataMembershipExpirationDate
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyBitAssetDataID, kn) {
					currentKey = ffjtBitAssetDataID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtBitAssetDatanosuchkey
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

				case ffjtBitAssetDataID:
					goto handle_ID

				case ffjtBitAssetDataMembershipExpirationDate:
					goto handle_MembershipExpirationDate

				case ffjtBitAssetDataIsPredictionMarket:
					goto handle_IsPredictionMarket

				case ffjtBitAssetDataSettlementPrice:
					goto handle_SettlementPrice

				case ffjtBitAssetDataFeeds:
					goto handle_Feeds

				case ffjtBitAssetDataOptions:
					goto handle_Options

				case ffjtBitAssetDataCurrentFeed:
					goto handle_CurrentFeed

				case ffjtBitAssetDataForcedSettledVolume:
					goto handle_ForcedSettledVolume

				case ffjtBitAssetDataSettlementFund:
					goto handle_SettlementFund

				case ffjtBitAssetDatanosuchkey:
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

handle_ID:

	/* handler: j.ID type=types.AssetBitAssetDataID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.ID.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_MembershipExpirationDate:

	/* handler: j.MembershipExpirationDate type=types.Time kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.MembershipExpirationDate.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_IsPredictionMarket:

	/* handler: j.IsPredictionMarket type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				j.IsPredictionMarket = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				j.IsPredictionMarket = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_SettlementPrice:

	/* handler: j.SettlementPrice type=types.Price kind=struct quoted=false*/

	{
		/* Falling back. type=types.Price kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.SettlementPrice)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Feeds:

	/* handler: j.Feeds type=types.AssetFeeds kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for AssetFeeds", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Feeds = nil
		} else {

			j.Feeds = []AssetFeed{}

			wantVal := true

			for {

				var tmpJFeeds AssetFeed

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
				