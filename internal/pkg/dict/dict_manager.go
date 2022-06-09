package dict

import (
	"data_proxy/internal/conf"
	"errors"
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"sync"
	"sync/atomic"
	"time"
)
var manager, mErr = NewManager()

var (
	// ErrNotFound is file not found.
	ErrNotFound = errors.New("file not found")
	ErrNotExist = errors.New("dict not exist")
)

type Dict interface {
	Load(path string) bool
	Close()
	Name() string
	Get() interface{}
	Init() error
}

type dict struct {
	reloadTimestamp int64
	dictIdx int32
	dictInfo conf.DictInfo
	opts      options
	mu sync.Mutex
}

type Manager struct {
	dictHash map[string]*dict
	fw *fsnotify.Watcher
	log       *log.Helper
}

func NewManager() (*Manager, error){
	fw, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return &Manager{
		dictHash: make(map[string]*dict),
		fw: fw,
	}, nil
}

func alignDict() {
	for _ = range time.Tick(time.Second * 30) {
		for _, hash:= range manager.dictHash {
			sec := hash.reloadTimestamp - time.Now().Unix()
			manager.log.Info("align sec:", sec)
		}
	}
}

func updateDict(name string, isFirst bool)  {
	hash, ok := manager.dictHash[name]
	if !ok {
		manager.log.Error("dict name:", name, " not exist")
	}
	// 还有校准有可能会更新
	hash.mu.Lock()
	defer hash.mu.Unlock()
	index := atomic.LoadInt32(&hash.dictIdx)
	dict := GetRegDict(hash.dictInfo.Name)
	log.Infof("%v, %p", dict, &dict)
	log.Info("1111", dict)
	dict.Init()
	log.Info("1111", dict)
	dict.Load(hash.dictInfo.Path)
	log.Info("1111", dict)
	changeIndex := 1 - index
	hash.opts.dict[changeIndex] = dict
	atomic.CompareAndSwapInt32(&hash.dictIdx, index, changeIndex)
	hash.reloadTimestamp = time.Now().Unix()
	log.Info("index:", hash.dictIdx, index, changeIndex, hash.reloadTimestamp)
}

func update(log *log.Helper) {
	defer manager.fw.Close()
	for {
		select {
		case event, ok := <-manager.fw.Events:
			if !ok {
				return
			}

			log.Info("event:", event)
			updateDict(event.Name, false)
			//if event.Op&fsnotify.Write == fsnotify.Write {
			//	log.Info("modified file:", event.Name)
			//}
		case err, ok := <-manager.fw.Errors:
			if !ok {
				return
			}
			log.Info("error:", err)
		}
	}
}

func Init(log *log.Helper)  {
	var dictConf string
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

	var conf conf.Dict
	if err := c.Scan(&conf); err != nil {
		panic(err)
	}
	log.Info("dict:", conf)
	for _, v := range conf.Dicts {
		log.Info(v.Name, v.ClassName, v.Path)
		info:= &dict{}
		info.dictInfo = *v
		info.dictIdx = 0
		// info.opts.dict[info.dictIdx] = NewDict()
		manager.dictHash[v.Path] = info
		updateDict(v.Path, true)
		err:=manager.fw.Add(v.Path)
		if err != nil {
			log.Fatal("dict path:", v.Path ," not exist")
		}
	}

	go update(log)
	go alignDict()
}

func GetDict(name string) Dict {
	hash, ok := manager.dictHash[name]
	if ok {
		manager.log.Error("dict name:", name, " not exist")
	}
	index:=atomic.LoadInt32(&hash.dictIdx)
	return hash.opts.dict[index]
}