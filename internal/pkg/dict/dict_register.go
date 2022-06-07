package dict

import (
	"data_proxy/internal/pkg/dict/dict_gcms"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"reflect"
)

var dictRegistry = make(map[string]reflect.Type)

func init() {
	myTypes := []interface{}{dict_gcms.DictGcms{}}
	for _, v := range myTypes {
		// typeRegistry["DictGcms"] = reflect.TypeOf(MyString{})
		log.Info(" type :", v)
		dictRegistry[fmt.Sprintf("%T", v)] = reflect.TypeOf(v)
	}
}
