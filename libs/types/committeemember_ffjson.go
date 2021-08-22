// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: committeemember.go

package types

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *CommitteeMember) MarshalJSON() ([]byte, error) {
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
func (j *CommitteeMember) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
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
	buf.WriteString(`,"committee_member_account":`)

	{

		obj, err = j.CommitteeMemberAccount.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"total_votes":`)
	fflib.FormatBits2(buf, uint64(j.TotalVotes), 10, false)
	buf.WriteString(`,"url":`)

	{

		obj, err = j.URL.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"vote_id":`)

	{

		obj, err = j.VoteID.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtCommitteeMemberbase = iota
	ffjtCommitteeMembernosuchkey

	ffjtCommitteeMemberID

	ffjtCommitteeMemberCommitteeMemberAccount

	ffjtCommitteeMemberTotalVotes

	ffjtCommitteeMemberURL

	ffjtCommitteeMemberVoteID
)

var ffjKeyCommitteeMemberID = []byte("id")

var ffjKeyCommitteeMemberCommitteeMemberAccount = []byte("committee_member_account")

var ffjKeyCommitteeMemberTotalVotes = []byte("total_votes")

var ffjKeyCommitteeMemberURL = []byte("url")

var ffjKeyCommitteeMemberVoteID = []byte("vote_id")

// UnmarshalJSON umarshall json - template of ffjson
func (j *CommitteeMember) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *CommitteeMember) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtCommitteeMemberbase
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
				currentKey = ffjtCommitteeMembernosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffjKeyCommitteeMemberCommitteeMemberAccount, kn) {
						currentKey = ffjtCommitteeMemberCommitteeMemberAccount
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyCommitteeMemberID, kn) {
						currentKey = ffjtCommitteeMemberID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffjKeyCommitteeMemberTotalVotes, kn) {
						currentKey = ffjtCommitteeMemberTotalVotes
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'u':

					if bytes.Equal(ffjKeyCommitteeMemberURL, kn) {
						currentKey = ffjtCommitteeMemberURL
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'v':

					if bytes.Equal(ffjKeyCommitteeMemberVoteID, kn) {
						currentKey = ffjtCommitteeMemberVoteID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.AsciiEqualFold(ffjKeyCommitteeMemberVoteID, kn) {
					currentKey = ffjtCommitteeMemberVoteID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyCommitteeMemberURL, kn) {
					currentKey = ffjtCommitteeMemberURL
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyCommitteeMemberTotalVotes, kn) {
					currentKey = ffjtCommitteeMemberTotalVotes
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyCommitteeMemberCommitteeMemberAccount, kn) {
					currentKey = ffjtCommitteeMemberCommitteeMemberAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyCommitteeMemberID, kn) {
					currentKey = ffjtCommitteeMember