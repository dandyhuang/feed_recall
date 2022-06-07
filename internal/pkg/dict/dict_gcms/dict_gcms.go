package dict_gcms

import (
	"data_proxy/internal/pkg/dict"
	"errors"
)

var (
	// ErrNotFound is file not found.
	ErrNotFound = errors.New("file not found")
)

func init() {
	dict.GetRegister().Set("gcms", new(DictGcms))
}

type DictGcms struct {
	dict *dict.DictBase
	dictData map[string] interface{}
}

func (d DictGcms) Init(conf string) error {
	d.dictData = make(map[string] interface{})
	d.dict = dict.NewDict()
	d.dict.Init(conf)
	d.dictData["hha"] = "dfds"
	return nil
}

func (d DictGcms) Load() bool {
	panic("implement me")
}

func (d DictGcms) Close() {
	panic("implement me")
}

func NewDict(conf string, opts ...dict.Option) dict.Dict {
	dict:=dict.NewDict(opts...)
	dict.Init(conf)

	return &DictGcms{
		dict: dict,
		dictData: make(map[string] interface{}),
	}
}