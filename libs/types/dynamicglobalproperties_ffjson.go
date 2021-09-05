// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: dynamicglobalproperties.go

package types

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *DynamicGlobalProperties) MarshalJSON() ([]byte, error) {
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
func (j *DynamicGlobalProperties) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
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
	buf.WriteString(`,"current_witness":`)

	{

		obj, err = j.CurrentWitness.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"last_budget_time":`)

	{

		obj, err = j.LastBudgetTime.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"time":`)

	{

		obj, err = j.Time.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"next_maintenance_time":`)

	{

		obj, err = j.NextMaintenanceTime.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"accounts_registered_this_interval":`)
	fflib.FormatBits2(buf, uint64(j.AccountsRegisteredThisInterval), 10, j.AccountsRegisteredThisInterval < 0)
	buf.WriteString(`,"dynamic_flags":`)
	fflib.FormatBits2(buf, uint64(j.DynamicFlags), 10, j.DynamicFlags < 0)
	buf.WriteString(`,"head_block_id":`)

	{

		obj, err = j.HeadBlockID.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"recent_slots_filled":`)

	{

		obj, err = j.RecentSlotsFilled.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"head_block_number":`)
	fflib.FormatBits2(buf, uint64(j.HeadBlockNumber), 10, false)
	buf.WriteString(`,"last_irreversible_block_num":`)
	fflib.FormatBits2(buf, uint64(j.LastIrreversibleBlockNum), 10, false)
	buf.WriteString(`,"current_aslot":`)
	fflib.FormatBits2(buf, uint64(j.CurrentAslot), 10, j.CurrentAslot < 0)
	buf.WriteString(`,"witness_budget":`)
	fflib.FormatBits2(buf, uint64(j.WitnessBudget), 10, j.WitnessBudget < 0)
	buf.WriteString(`,"recently_missed_count":`)
	fflib.FormatBits2(buf, uint64(j.RecentlyMissedCount), 10, j.RecentlyMissedCount < 0)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtDynamicGlobalPropertiesbase = iota
	ffjtDynamicGlobalPropertiesnosuchkey

	ffjtDynamicGlobalPropertiesID

	ffjtDynamicGlobalPropertiesCurrentWitness

	ffjtDynamicGlobalPropertiesLastBudgetTime

	ffjtDynamicGlobalPropertiesTime

	ffjtDynamicGlobalPropertiesNextMaintenanceTime

	ffjtDynamicGlobalPropertiesAccountsRegisteredThisInterval

	ffjtDynamicGlobalPropertiesDynamicFlags

	ffjtDynamicGlobalPropertiesHeadBlockID

	ffjtDynamicGlobalPropertiesRecentSlotsFilled

	ffjtDynamicGlobalPropertiesHeadBlockNumber

	ffjtDynamicGlobalPropertiesLastIrreversibleBlockNum

	ffjtDynamicGlobalPropertiesCurrentAslot

	ffjtDynamicGlobalPropertiesWitnessBudget

	ffjtDynamicGlobalPropertiesRecentlyMissedCount
)

var ffjKeyDynamicGlobalPropertiesID = []byte("id")

var ffjKeyDynamicGlobalPropertiesCurrentWitness = []byte("current_witness")

var ffjKeyDynamicGlobalPropertiesLastBudgetTime = []byte("last_budget_time")

var ffjKeyDynamicGlobalPropertiesTime = []byte("time")

var ffjKeyDynamicGlobalPropertiesNextMaintenanceTime = []byte("next_maintenance_time")

var ffjKeyDynamicGlobalPropertiesAccountsRegisteredThisInterval = []byte("accounts_registered_this_interval")

var ffjKeyDynamicGlobalPropertiesDynamicFlags = []byte("dynamic_flags")

var ffjKeyDynamicGlobalPropertiesHeadBlockID = []byte("head_block_id")

var ffjKeyDynamicGlobalPropertiesRecentSlotsFilled = []byte("recent_slots_filled")

var ffjKeyDynamicGlobalPropertiesHeadBlockNumber = []byte("head_block_number")

var ffjKeyDynamicGlobalPropertiesLastIrreversibleBlockNum = []byte("last_irreversible_block_num")

var ffjKeyDynamicGlobalPropertiesCurrentAslot = []byte("current_aslot")

var ffjKeyDynamicGlobalPropertiesWitnessBudget = []byte("witness_budget")

var ffjKeyDynamicGlobalPropertiesRecentlyMissedCount = []byte("recently_missed_count")

// UnmarshalJSON umarshall json - template of ffjson
func (j *DynamicGlobalProperties) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *DynamicGlobalProperties) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtDynamicGlobalPropertiesbase
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
				currentKey = ffjtDynamicGlobalPropertiesnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesAccountsRegisteredThisInterval, kn) {
						currentKey = ffjtDynamicGlobalPropertiesAccountsRegisteredThisInterval
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesCurrentWitness, kn) {
						currentKey = ffjtDynamicGlobalPropertiesCurrentWitness
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyDynamicGlobalPropertiesCurrentAslot, kn) {
						currentKey = ffjtDynamicGlobalPropertiesCurrentAslot
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesDynamicFlags, kn) {
						currentKey = ffjtDynamicGlobalPropertiesDynamicFlags
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'h':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesHeadBlockID, kn) {
						currentKey = ffjtDynamicGlobalPropertiesHeadBlockID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyDynamicGlobalPropertiesHeadBlockNumber, kn) {
						currentKey = ffjtDynamicGlobalPropertiesHeadBlockNumber
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesID, kn) {
						currentKey = ffjtDynamicGlobalPropertiesID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesLastBudgetTime, kn) {
						currentKey = ffjtDynamicGlobalPropertiesLastBudgetTime
						state = fflib.FFParse_want_colon
	