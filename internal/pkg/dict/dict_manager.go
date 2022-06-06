package dict

import (
	"data_proxy/internal/conf"
	"flag"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	dictConf string
)

type Info struct {
	reloadTimestamp int64
	dictIdx int8
	dict []*Dict
}

type Manager struct {
	DictHash map[string]Info
}

func Init( log *log.Helper)  {
	flag.StringVar(&dictConf, "dict", "./configs/dict.yaml", "dict eg: -dict dict.yaml")
	log.Info("dict path:", dictConf)
	c := config.New(
		config.WithSource(
			file.NewSource(dictConf),
		),
	)
	defer c.Close()
	if err := c.Load(); err != nil {
		panic(err)
	}

	var dict conf.Dict
	if err := c.Scan(&dict); err != nil {
		panic(err)
	}

	log.Info("dict:", dict)
}