package dict_gcms

import (
	"data_proxy/internal/pkg/dict"
)

type DictGcms struct {
	dict *dict.DictBase
	dictData map[string] interface{}
}

func (d DictGcms) Init(conf string) error {
	panic("implement me")
	d.dict.Init(conf)
	return nil
}

func (d DictGcms) Load() bool {
	panic("implement me")
}

func (d DictGcms) Close() {
	panic("implement me")
}

func NewDict(opts ...dict.Option) dict.Dict {
	dict:=dict.NewDict(opts...)

	return &DictGcms{
		dict: dict,
	}
}