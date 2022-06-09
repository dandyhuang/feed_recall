package dict

import (
	"github.com/go-kratos/kratos/v2/log"
	"strings"
)

type TypeRegister map[string]Dict

var typeReg = make(TypeRegister)

func RegisterDict(dict Dict) TypeRegister {
	if dict == nil {
		panic("cannot register a nil Codec")
	}
	if dict.Name() == "" {
		panic("cannot register Codec with empty string result for Name()")
	}
	dictName := strings.ToLower(dict.Name())
	typeReg[dictName] = dict
	log.Infof("%v, %p", dict, dict)
	return  typeReg
}

func GetRegDict(name string) Dict {
	return typeReg[name]
}


