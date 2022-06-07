package dict

import (
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"reflect"
)

type TypeRegister map[string]reflect.Type

func (t TypeRegister) Set(i interface{}) {
	//name string, typ reflect.Type
	fmt.Println("tpye set name:", reflect.TypeOf(i).Name())
	t["gcms"] = reflect.TypeOf(i)
	fmt.Println("value：", typeReg)
}

func (t TypeRegister) Get(name string) (interface{}, error) {
	for k, v := range t {
		log.Info("k:", k , " v:", v)
	}

	if typ, ok := t[name]; ok {
		return reflect.New(typ).Elem().Interface(), nil
	}
	return nil, errors.New("no one")
}

var typeReg = make(TypeRegister)

func GetRegister()TypeRegister {
	return  typeReg
}
