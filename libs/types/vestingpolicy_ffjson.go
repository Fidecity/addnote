
// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: vestingpolicy.go

package types

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *CCDVestingPolicy) MarshalJSON() ([]byte, error) {
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
func (j *CCDVestingPolicy) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"start_claim":`)

	{

		obj, err = j.StartClaim.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"coin_seconds_earned_last_update":`)

	{

		obj, err = j.CoinSecondsEarnedLastUpdate.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"vesting_seconds":`)
	fflib.FormatBits2(buf, uint64(j.VestingSeconds), 10, false)
	buf.WriteString(`,"coin_seconds_earned":`)
	fflib.FormatBits2(buf, uint64(j.CoinSecondsEarned), 10, false)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtCCDVestingPolicybase = iota
	ffjtCCDVestingPolicynosuchkey

	ffjtCCDVestingPolicyStartClaim

	ffjtCCDVestingPolicyCoinSecondsEarnedLastUpdate

	ffjtCCDVestingPolicyVestingSeconds

	ffjtCCDVestingPolicyCoinSecondsEarned
)

var ffjKeyCCDVestingPolicyStartClaim = []byte("start_claim")

var ffjKeyCCDVestingPolicyCoinSecondsEarnedLastUpdate = []byte("coin_seconds_earned_last_update")

var ffjKeyCCDVestingPolicyVestingSeconds = []byte("vesting_seconds")

var ffjKeyCCDVestingPolicyCoinSecondsEarned = []byte("coin_seconds_earned")

// UnmarshalJSON umarshall json - template of ffjson
func (j *CCDVestingPolicy) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *CCDVestingPolicy) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtCCDVestingPolicybase
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
				currentKey = ffjtCCDVestingPolicynosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffjKeyCCDVestingPolicyCoinSecondsEarnedLastUpdate, kn) {
						currentKey = ffjtCCDVestingPolicyCoinSecondsEarnedLastUpdate
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyCCDVestingPolicyCoinSecondsEarned, kn) {
						currentKey = ffjtCCDVestingPolicyCoinSecondsEarned
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeyCCDVestingPolicyStartClaim, kn) {
						currentKey = ffjtCCDVestingPolicyStartClaim
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'v':

					if bytes.Equal(ffjKeyCCDVestingPolicyVestingSeconds, kn) {
						currentKey = ffjtCCDVestingPolicyVestingSeconds
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyCCDVestingPolicyCoinSecondsEarned, kn) {
					currentKey = ffjtCCDVestingPolicyCoinSecondsEarned
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyCCDVestingPolicyVestingSeconds, kn) {
					currentKey = ffjtCCDVestingPolicyVestingSeconds
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyCCDVestingPolicyCoinSecondsEarnedLastUpdate, kn) {
					currentKey = ffjtCCDVestingPolicyCoinSecondsEarnedLastUpdate
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyCCDVestingPolicyStartClaim, kn) {
					currentKey = ffjtCCDVestingPolicyStartClaim
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtCCDVestingPolicynosuchkey
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

				case ffjtCCDVestingPolicyStartClaim:
					goto handle_StartClaim

				case ffjtCCDVestingPolicyCoinSecondsEarnedLastUpdate:
					goto handle_CoinSecondsEarnedLastUpdate

				case ffjtCCDVestingPolicyVestingSeconds:
					goto handle_VestingSeconds

				case ffjtCCDVestingPolicyCoinSecondsEarned:
					goto handle_CoinSecondsEarned

				case ffjtCCDVestingPolicynosuchkey:
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

handle_StartClaim:

	/* handler: j.StartClaim type=types.Time kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.StartClaim.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CoinSecondsEarnedLastUpdate:

	/* handler: j.CoinSecondsEarnedLastUpdate type=types.Time kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.CoinSecondsEarnedLastUpdate.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_VestingSeconds:

	/* handler: j.VestingSeconds type=types.UInt32 kind=uint32 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.VestingSeconds.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CoinSecondsEarned:

	/* handler: j.CoinSecondsEarned type=types.UInt64 kind=uint64 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.CoinSecondsEarned.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}

// MarshalJSON marshal bytes to json - template
func (j *LinearVestingPolicy) MarshalJSON() ([]byte, error) {
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
func (j *LinearVestingPolicy) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"begin_timestamp":`)

	{

		obj, err = j.BeginTimestamp.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"vesting_cliff_seconds":`)
	fflib.FormatBits2(buf, uint64(j.VestingCliffSeconds), 10, false)
	buf.WriteString(`,"vesting_duration_seconds":`)
	fflib.FormatBits2(buf, uint64(j.VestingDurationSeconds), 10, false)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtLinearVestingPolicybase = iota
	ffjtLinearVestingPolicynosuchkey

	ffjtLinearVestingPolicyBeginTimestamp

	ffjtLinearVestingPolicyVestingCliffSeconds

	ffjtLinearVestingPolicyVestingDurationSeconds
)

var ffjKeyLinearVestingPolicyBeginTimestamp = []byte("begin_timestamp")

var ffjKeyLinearVestingPolicyVestingCliffSeconds = []byte("vesting_cliff_seconds")

var ffjKeyLinearVestingPolicyVestingDurationSeconds = []byte("vesting_duration_seconds")

// UnmarshalJSON umarshall json - template of ffjson
func (j *LinearVestingPolicy) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *LinearVestingPolicy) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtLinearVestingPolicybase
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
				currentKey = ffjtLinearVestingPolicynosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'b':

					if bytes.Equal(ffjKeyLinearVestingPolicyBeginTimestamp, kn) {
						currentKey = ffjtLinearVestingPolicyBeginTimestamp
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'v':

					if bytes.Equal(ffjKeyLinearVestingPolicyVestingCliffSeconds, kn) {
						currentKey = ffjtLinearVestingPolicyVestingCliffSeconds
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyLinearVestingPolicyVestingDurationSeconds, kn) {
						currentKey = ffjtLinearVestingPolicyVestingDurationSeconds
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyLinearVestingPolicyVestingDurationSeconds, kn) {
					currentKey = ffjtLinearVestingPolicyVestingDurationSeconds
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyLinearVestingPolicyVestingCliffSeconds, kn) {
					currentKey = ffjtLinearVestingPolicyVestingCliffSeconds
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyLinearVestingPolicyBeginTimestamp, kn) {
					currentKey = ffjtLinearVestingPolicyBeginTimestamp
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtLinearVestingPolicynosuchkey
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

				case ffjtLinearVestingPolicyBeginTimestamp:
					goto handle_BeginTimestamp

				case ffjtLinearVestingPolicyVestingCliffSeconds:
					goto handle_VestingCliffSeconds

				case ffjtLinearVestingPolicyVestingDurationSeconds:
					goto handle_VestingDurationSeconds

				case ffjtLinearVestingPolicynosuchkey:
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

handle_BeginTimestamp:

	/* handler: j.BeginTimestamp type=types.Time kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.BeginTimestamp.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_VestingCliffSeconds:

	/* handler: j.VestingCliffSeconds type=types.UInt32 kind=uint32 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.VestingCliffSeconds.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_VestingDurationSeconds:

	/* handler: j.VestingDurationSeconds type=types.UInt32 kind=uint32 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.VestingDurationSeconds.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}