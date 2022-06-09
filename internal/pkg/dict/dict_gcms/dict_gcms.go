package dict_gcms

import (
	"data_proxy/internal/pkg/dict"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrNotFound is file not found.
	ErrNotFound = errors.New("file not found")
)

func init() {
	dict.RegisterDict(DictGcms{})
}
type DictGcms struct {
	dictData map[string] interface{}
}


func (d DictGcms) Name() string {
	return "gcms"
}

func (d DictGcms) Init() error {
	d.dictData = make(map[string] interface{})
	return nil
}

func (d DictGcms) Load(path string) bool {
	log.Info("path", path)
	d.dictData["dandy"] = "hello"
	log.Info("implement me")
	return true
}

func (d DictGcms) Get() interface{} {
	return d.dictData
}

func (d DictGcms) Close() {
	log.Info("implement me")
}

