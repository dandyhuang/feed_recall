module data_proxy

go 1.17

require (
	github.com/fsnotify/fsnotify v1.5.1
	github.com/g4zhuj/go-metrics-falcon v0.0.0-20180427054158-5159ced4eafb
	github.com/go-kratos/aegis v0.1.1
	github.com/go-kratos/kratos/contrib/metrics/prometheus/v2 v2.0.0-20220506075950-eb2dcae83d7b
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20220422120629-fbf7855cf262
	github.com/go-kratos/kratos/v2 v2.2.1
	github.com/go-redis/redis/extra/redisotel v0.3.0
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db
	github.com/google/wire v0.5.0
	github.com/hashicorp/consul/api v1.12.0
	github.com/hyperledger/fabric-sdk-go v1.0.0
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/prometheus/client_golang v1.12.2
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a
	github.com/shirou/gopsutil/v3 v3.21.8
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.1.1
	github.com/stretchr/testify v1.7.1
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/exporters/jaeger v1.7.0
	go.opentelemetry.io/otel/sdk v1.7.0
	go.opentelemetry.io/otel/trace v1.7.0
	go.uber.org/zap v1.21.0
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	golang.org/x/sys v0.0.0-20220422013727-9388b58f7150 // indirect
	google.golang.org/grpc v1.44.0
	google.golang.org/protobuf v1.28.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.5
	gorm.io/plugin/prometheus v0.0.0-20220517015831-ca6bfaf20bf4
)

// replace google.golang.org/grpc v1.44.0 => gitlab.vmic.xyz/11126518/grpc-go v0.0.0-20220426114857-d0f8dd92e258
