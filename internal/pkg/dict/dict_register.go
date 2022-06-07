package dict

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"reflect"
)

type TypeRegister map[string]interface{}

func (t TypeRegister) Set(name string, i interface{}) {
	//name string, typ reflect.Type
	fmt.Println("tpye set name:", reflect.TypeOf(i).Name())
	t[name] = i
	fmt.Println("value：", typeReg)
}

func (t TypeRegister) Get(name string) (interface{}, error) {
	for k, v := range t {
		log.Info("k:", k , " v:", v)
	}

	if typ, ok := t[name]; ok {
		return typ, nil
	}
	return nil, ErrNotExist
}

var typeReg = make(TypeRegister)

func GetRegister() TypeRegister {
	return  typeReg
}
