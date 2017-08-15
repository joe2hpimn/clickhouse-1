package column

import (
	"fmt"

	"github.com/kshvakov/clickhouse/internal/binary"
)

type Float32 struct{ base }

func (Float32) Read(decoder *binary.Decoder) (interface{}, error) {
	v, err := decoder.Float32()
	if err != nil {
		return float32(0), err
	}
	return v, nil
}

func (Float32) Write(encoder *binary.Encoder, v interface{}) error {
	switch v := v.(type) {
	case float32:
		return encoder.Float32(v)
	case float64:
		return encoder.Float32(float32(v))
	}
	return fmt.Errorf("unexpected type %T", v)
}
