package main

import (
	"data_proxy/internal/conf"
	logZap "data_proxy/internal/log"
	"data_proxy/internal/pkg/dict"
	"data_proxy/internal/pkg/dict/dict_gcms"
	"data_proxy/internal/pkg/stat"
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	_ "data_proxy/internal/pkg/dict/dict_gcms"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
	Env string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "./configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&Env, "env", "prd", "env eg: -env prd")
}
// 先grpc，在http给consul注册使用
func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(rr), // consul 的引入
	)
}

func main() {
	flag.Parse()
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	// consul 的引入
	var rc conf.Registry
	if err := c.Scan(&rc); err != nil {
		panic(err)
	}

	//
	var kt conf.DataProxy
	if err := c.Scan(&kt); err != nil {
		panic(err)
	}
	logLevel := zapcore.DebugLevel
	if bc.Server.Log.Level == "Info" {
		logLevel = zapcore.InfoLevel
	} else if bc.Server.Log.Level == "Error" {
		logLevel = zapcore.ErrorLevel
	} else if bc.Server.Log.Level == "Warn" {
		logLevel = zapcore.WarnLevel
	}

	// buildInstance 中Register会注册consul
	Name = kt.Kratos.Name + "_" + Env

	writeLogger:=logZap.NewZapWriteLogger(bc.Server.Log.Local,
		Env,
		zap.NewAtomicLevelAt(logLevel),
		zap.AddStacktrace(
			zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
		zap.AddCallerSkip(2),
		zap.AddCaller(),
		zap.Development(),
	)
	logger := log.With(writeLogger,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		//"service.name", Name,
		//"service.version", Version,
		//"trace.id", tracing.TraceID(),
		//"span.id", tracing.SpanID(),
	)
	dict.Init(log.NewHelper(logger))
	gcms, _ :=dict.GetRegister().Get("gcms")
	g:=gcms.(dict_gcms.DictGcms)
	g.Init("../configs")
	log.Info(gcms)
	app, cleanup, err := wireApp(bc.Server, bc.Data, &rc, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	log.Info("appease:", app.Name(), " endpoint:", app.Endpoint())
	// 监控上报
	stat.GetStandardStat().Polymeric(bc.Server.Stat , log.NewHelper(logger))


	// start and wait for stop signal
	if err := app.Run(); err != nil {
		log.NewHelper(logger).Info("run err: ", err)
		panic(err)
	}
}
