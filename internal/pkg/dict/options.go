package dict

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/log"
	"strings"
)

type Option func(*options)

type options struct {
	decoder  Decoder
	logger   log.Logger
}

type KeyValue struct {
	Key    string
	Value  []byte
	Format string
}

type Decoder func(*KeyValue, map[string]interface{}) error

// WithLogger with config logger.
func WithLogger(l log.Logger) Option {
	return func(o *options) {
		o.logger = l
	}
}

// WithSource with config source.
func WithSource(d Decoder) Option {
	return func(o *options) {
		o.decoder = d
	}
}

func defaultDecoder(src *KeyValue, target map[string]interface{}) error {
	if src.Format == "" {
		// expand key "aaa.bbb" into map[aaa]map[bbb]interface{}
		keys := strings.Split(src.Key, ".")
		for i, k := range keys {
			if i == len(keys)-1 {
				target[k] = src.Value
			} else {
				sub := make(map[string]interface{})
				target[k] = sub
				target = sub
			}
		}
		return nil
	}
	if codec := encoding.GetCodec(src.Format); codec != nil {
		return codec.Unmarshal(src.Value, &target)
	}
	return fmt.Errorf("unsupported key: %s format: %s", src.Key, src.Format)
}