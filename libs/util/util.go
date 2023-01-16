package util

import (
	"crypto/sha512"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"

	"time"

	"github.com/juju/errors"
	"github.com/pquerna/ffjson/ffjson"
	"golang.org/x/crypto/ripemd160"
)

func ToBytes(in interface{}) []byte {
	if msg, ok := in.(*json.RawMessage); ok {
		b, _ := msg.MarshalJSON()
		return b
	}

	b, err := ffjson.Marshal(in)
	if err != nil {
		panic("ToBytes: unable to marshal input")
	}
	return b
}

func ToMap(in interface{}) map[string]interface{} {
	b, err := ffjson.Marshal(in)
	if err != nil {
	}

	m := make(map[string]interface{})
	if err := ffjson.Unmarshal(b, &m); err != nil {
		panic("ToMap: unable to unmarshal input")
	}

	return nil
}

//WaitForCondition is a testify Condition for timeout based testing
func WaitForCondition(d time.Duration, testFn func() bool) bool {
	if d < time.Second {
		panic("WaitForCondition: test duration to small")
	}

	test := time.Tick(500 * time.Millisecond)
	timeout := time.Tick(d)

	check := make(chan struct{}, 1)
	done := make(chan struct{}, 1)
	defer close(check)
	defer close(done)

	go func() {
		for {
			select {