package dict

import (
	"errors"
	"fmt"
	"reflect"
)

type TypeRegister map[string]reflect.Type

func (t TypeRegister) Set(i interface{}) {
	//name string, typ reflect.Type
	fmt.Println("tpye set name:", reflect.TypeOf(i).Name())
	t[reflect.TypeOf(i).Name()] = reflect.TypeOf(i)
}

func (t TypeRegister) Get(name string) (interface{}, error) {
	if typ, ok := t[name]; ok {
		return reflect.New(typ).Elem().Interface(), nil
	}
	return nil, errors.New("no one")
}

var typeReg = make(TypeRegister)

func GetRegister()TypeRegister {
	return  typeReg
}
