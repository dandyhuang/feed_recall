package dict

import (
	"bytes"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var (
	// ErrNotFound is file not found.
	ErrNotFound = errors.New("file not found")
	_ Dict = (*DictBase)(nil)
)

type Dict interface {
	Init(conf string) error
	Load() bool
	Close()
}

type DictBase struct {
	dictFile string
	dictName string
	opts      options
	log       *log.Helper
}

func NewDict(opts ...Option) *DictBase {
	o := options{
		logger:   log.GetLogger(),
		decoder:  defaultDecoder,
	}
	for _, opt := range opts {
		opt(&o)
	}

	return &DictBase{
		opts:   o,
		log:    log.NewHelper(o.logger),
	}
}

func (d DictBase) Init(fileName string) error {
	_, err := os.Stat(fileName)
	if err == nil{
		d.log.Error("File exist")
		return nil
	}
	if os.IsNotExist(err){
		d.log.Error("File not exist")
		return ErrNotFound
	}
	d.dictFile = fileName
	return nil
}

func (d DictBase) Load() bool {
	b, err := ioutil.ReadFile(d.dictFile)
	if err != nil {
		return false
	}
	r := bytes.NewBuffer(b)
	for {
		id, err := r.ReadString('\n')
		if err == io.EOF || err == nil {
			id = strings.TrimSpace(id)
			if len(id) > 0 {
				// d.blackIDs[id] = 1
			}
		}

		if err != nil {
			break
		}
	}
	return true
}

func (d DictBase) Close() {
	panic("implement me")
}
