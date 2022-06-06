package main

import (
	"context"
	rec "data_proxy/api/recommend"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
	"github.com/golang/snappy"
)
var imei = flag.String("imei", "863481040330852", "bool类型参数")
var addr =  flag.String("addr", "r-t4n9729b716dd644.redis.singapore.rds.aliyuncs.com:6379", "bool类型参数")
func main() {
	str:= []string{*addr}
	r := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    str    ,
	})
	res, _:=r.Get(context.Background(), "fuv_histroy_"+ (*imei)).Result()
	sp, _:=snappy.Decode(nil, []byte(res))
	usInfo:=rec.HistoryInfo{}
	proto.Unmarshal([]byte(sp), &usInfo)
	fmt.Println("size:", len(usInfo.HistoryTermList))
	fmt.Println(usInfo)
}
