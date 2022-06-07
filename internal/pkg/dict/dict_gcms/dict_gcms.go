package dict_gcms

import (
	"data_proxy/internal/pkg/dict"
)

func init() {
	//myTypes := []interface{}{DictGcms{}}
	//for _, v := range myTypes {
	//	// typeRegistry["DictGcms"] = reflect.TypeOf(MyString{})
	//	log.Info(" ====type :", v)
	//	dict.GetRegister()[fmt.Sprintf("%T", v)] = reflect.TypeOf(v)
	//}
	dict.GetRegister().Set(new(DictGcms))
}

type DictGcms struct {
	dict *dict.DictBase
	dictData map[string] interface{}
}

func (d DictGcms) Init(conf string) error {
	d.dictData = make(map[string] interface{})
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

func NewDict(opts ...dict.Option) dict.Dict {
	dict:=dict.NewDict(opts...)

	return &DictGcms{
		dict: dict,
	}
}