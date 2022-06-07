package dict_gcms

import (
	"data_proxy/internal/pkg/dict"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"reflect"
)

func init() {
	myTypes := []interface{}{DictGcms{}}
	for _, v := range myTypes {
		// typeRegistry["DictGcms"] = reflect.TypeOf(MyString{})
		log.Info(" ====type :", v)
		dict.GetRegister()[fmt.Sprintf("%T", v)] = reflect.TypeOf(v)
	}
}

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