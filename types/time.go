package types

import (
	"time"

	"github.com/Assetsadapter/whitecoin-adapter/encoding"
)

const Layout = `"2006-01-02T15:04:05"`

type Time struct {
	*time.Time
}

func NewTime(t time.Time) Time {
	return Time{
		&t,
	}
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(Layout)), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	parsed, err := time.ParseInLocation(Layout, string(data), time.UTC)
	if err != nil {
		return err
	}
	t.Time = &parsed
	return nil
}

func (t Time) Marshal(encoder *encoding.Encoder) error {
	return encoder.EncodeLittleEndianUInt32(uint32(t.Time.UTC().Unix()))
}