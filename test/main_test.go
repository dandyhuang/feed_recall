package main

import (
	"context"
	"data_proxy/internal/pkg/stat"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"

	pb "data_proxy/api/recommend"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
)
// go test -v main_test.go -run="TestGrpc"
func TestGrpc(t *testing.T) {
	m1:=make([]map[string]interface{}, 0)
	m:=make(map[string]interface{})
	m["ip"] = "werweewr"
	m["rule_id"] = 1223423
	fmt.Println("host-name", stat.DefaultHostname())
	m1 = append(m1, m)
	body, err := json.Marshal(m1)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println("body:", string(body))
	// callHTTP()
	callGRPC(t)
}

//func callHTTP() {
//	conn, err := transhttp.NewClient(
//		context.Background(),
//		transhttp.WithMiddleware(
//			recovery.Recovery(),
//		),
//		transhttp.WithEndpoint("127.0.0.1:8000"),
//	)
//	if err != nil {
//		panic(err)
//	}
//	defer conn.Close()
//	client := pb.NewGreeterHTTPClient(conn)
//	reply, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "kratos"})
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Printf("[http] SayHello %s\n", reply.Message)
//
//	// returns error
//	_, err = client.SayHello(context.Background(), &pb.HelloRequest{Name: "error"})
//	if err != nil {
//		log.Printf("[http] SayHello error: %v\n", err)
//	}
//	if errors.IsBadRequest(err) {
//		log.Printf("[http] SayHello error is invalid argument: %v\n", err)
//	}
//}

func callGRPC(t *testing.T) {
	head:=make(map[string]interface{})
	head["key"] = "testKey"
	start:=time.Now()
	conn, err := transgrpc.DialInsecure(
		context.Background(),
		transgrpc.WithEndpoint("127.0.0.1:9000"),
		transgrpc.WithMiddleware(
			recovery.Recovery(),
			//jwt.Client(func(token *jwtv4.Token) (interface{}, error) {
			//	return []byte("testKey"), nil
			//}),
			// , jwt.WithTokenHeader(head)
		),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewCommonServiceClient(conn)
	dpReq:=&pb.DPRequest{Key: "dandy1",
		Cmd: "get"}

	r:=&pb.Param{}
	r.Key = "dandy1"
	dpReq.TableName = "da_appstore_detail_recall"
	dpReq.Param = append(dpReq.Param, r)
	log.Println("len:", len(dpReq.Param))

	req:=&pb.CommonRequest{}
	value, err := ptypes.MarshalAny(dpReq)
	req.Request= value
	//value1, err := ptypes.MarshalAny(dpReq)
	//req.Request = value1
	reply, err := client.Process(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] SayHello %+v %s, %v\n", reply, " cost:", time.Since(start))
	assert.Nil(t, err)
}
