
// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: overridetransferoperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *OverrideTransferOperation) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")