package server

import (
	"context"
	"data_proxy/api/recommend"
	"data_proxy/internal/conf"
	"data_proxy/internal/pkg/stat"
	"data_proxy/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	logv2 "github.com/hyperledger/fabric-sdk-go/pkg/common/logging"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc/grpclog"
	"strconv"
	"time"
)

func CostSrvMiddleware(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			start:= time.Now()
			reply, err = handler(ctx, req)
			attr:=make(map[string]interface{})
			costTime:= int64(time.Since(start).Microseconds())
			attr["cost_average"] = strconv.FormatInt(costTime, 10)
			attr["cost"] = strconv.FormatInt(costTime, 10)
			attr["count"] = "1"
			stat.GetStandardStat().Incr(ctx, attr)
			return
		}
	}
}
func setTracerProvider(url string) error {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return err
	}
	tp := tracesdk.NewTracerProvider(
		// Set the sampling rate based on the parent span to 100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String("data_proxy"),
			attribute.String("env", "dev"),
		)),
	)
	otel.SetTracerProvider(tp)
	// otel.Tracer("component-bar")
	return nil
}

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, common *service.CommonService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(
				recovery.WithHandler( func() recovery.HandlerFunc {
					return func(ctx context.Context, req, err interface{}) error {
						attr:=make(map[string]interface{})
						attr["count"] = "1"
						attr["panic"] = "1"
						// stat.StadardStat.C <- attr
						log.Info("error: panic")
						stat.GetStandardStat().Incr(ctx, attr)
						return nil
					}
				}(),
				),
			),
			CostSrvMiddleware(logger),
			//ratelimit.Server(),
			//tracing.Server(),
			logging.Server(logger),
			// jwt认证
			//func() (middleware.Middleware) {
			//	if c.Grpc.JwtKey != "" {
			//		return jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
			//			return []byte(c.Grpc.JwtKey), nil
			//		})
			//	}
			//	return EmptyMiddleware()
			//}(),
			// 默认打开监控
			//metrics.Server(
			//	metrics.WithSeconds(prom.NewHistogram(_metricSeconds)),
			//	metrics.WithRequests(prom.NewCounter(_metricRequests)),
			//),
			// circuitbreaker.Client(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	// 这里会覆盖上面的拦截器
	//if c.Http.Prometheus != "" {
	//	opts = append(opts, grpc.Middleware(
	//		metrics.Server(
	//			metrics.WithSeconds(prom.NewHistogram(_metricSeconds)),
	//			metrics.WithRequests(prom.NewCounter(_metricRequests)),
	//		),
	//	))
	//}

	logger = log.With(logger, "trace_id", tracing.TraceID())
	logger = log.With(logger, "span_id", tracing.SpanID())
	if c.Grpc.GlogOpen == 1 {
		var logger1 = logv2.NewLogger("grpcLogger")
		grpclog.SetLoggerV2(NewZapLogger(logger1))
		opts = append(opts, grpc.Logger(logger))
	}


	srv := grpc.NewServer(opts...)
	rec.RegisterCommonServiceServer(srv, common)
	return srv
}
