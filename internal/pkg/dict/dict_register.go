package dict

import (
	"reflect"
)

var dictRegistry = make(map[string]reflect.Type)

func GetRegister() map[string]reflect.Type {
	return dictRegistry
}
