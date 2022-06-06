package data

import (
	"context"
	"data_proxy/internal/biz"
	"data_proxy/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"gorm.io/plugin/prometheus"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewGreeterRepo, NewData, NewRedis, NewDB, NewTransaction, NewUserRepo, NewCardRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	rdb map[string] *redis.ClusterClient
	db *gorm.DB
}

type contextTxKey struct{}

func (d *Data) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

// NewTransaction .
func NewTransaction(d *Data) biz.Transaction {
	return d
}


func NewRedis(c *conf.Data)  map[string] *redis.ClusterClient {
	rdb := make(map[string] *redis.ClusterClient)
	for _, v := range c.Redis {
		addrss:=strings.Split(v.Addr, ",")
		r := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        addrss,
			Password:     v.Password,
			DialTimeout:  v.DialTimeout.AsDuration(),
			WriteTimeout: v.WriteTimeout.AsDuration(),
			ReadTimeout:  v.ReadTimeout.AsDuration(),
		})
		r.AddHook(redisotel.TracingHook{})
		//if err := r.Close(); err != nil {
		//	log.Error(err)
		//}
		_, err := r.Ping(context.Background()).Result()
		if err != nil {
			log.Error("redis ping", err)
			continue
		}
		log.Debug("table name:", v.TableName)
		rdb[v.TableName] = r
	}

	return rdb
}

// NewDB gorm Connecting to a Database
func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	// not impl
	return nil
	log := log.NewHelper(log.With(logger, "module", "order-service/data/gorm"))

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	if err := db.AutoMigrate(&User{}, &Card{}); err != nil {
		log.Fatal(err)
	}
	db.Use(prometheus.New(prometheus.Config{
		DBName:          "db1", // 使用 `DBName` 作为指标 label
		RefreshInterval: 15,    // 指标刷新频率（默认为 15 秒）
		PushAddr:        "", // 如果配置了 `PushAddr`，则推送指标
		StartServer:     true,  // 启用一个 http 服务来暴露指标
		HTTPServerPort:  8080,  // 配置 http 服务监听端口，默认端口为 8080 （如果您配置了多个，只有第一个 `HTTPServerPort` 会被使用）
		MetricsCollector: []prometheus.MetricsCollector {
			&prometheus.MySQL{
				VariableNames: []string{"Threads_running"},
			},
		},  // 用户自定义指标
	}))

	RegisterCallbacks(context.Background(), db)
	return db
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, rdb map[string] *redis.ClusterClient, db *gorm.DB,) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		for _, v := range rdb {
			if err := v.Close(); err != nil {
				log.Error(err)
			}
		}
	}
	return &Data{rdb: rdb, db: db}, cleanup, nil
}
