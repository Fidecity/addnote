
package encoding

type RollingEncoder struct {
	next *Encoder
	err  error
}

func NewRollingEncoder(next *Encoder) *RollingEncoder {
	return &RollingEncoder{next, nil}
}

func (encoder *RollingEncoder) EncodeVarint(i int64) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeVarint(i)
	}
}

func (encoder *RollingEncoder) EncodeUVarint(i uint64) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeUVarint(i)
	}
}

func (encoder *RollingEncoder) EncodeNumber(v interface{}) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeNumber(v)
	}
}

func (encoder *RollingEncoder) EncodeBool(v bool) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeBool(v)
	}
}

func (encoder *RollingEncoder) Encode(v interface{}) {
	if encoder.err == nil {
		encoder.err = encoder.next.Encode(v)
	}
}

func (encoder *RollingEncoder) EncodeMoney(s string) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeMoney(s)
	}
}

func (encoder *RollingEncoder) EncodeLittleEndianUInt64(i uint64) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeLittleEndianUInt64(i)
	}
}

func (encoder *RollingEncoder) EncodeLittleEndianUInt32(i uint32) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeLittleEndianUInt32(i)
	}
}

func (encoder *RollingEncoder) Err() error {
	return encoder.err
}