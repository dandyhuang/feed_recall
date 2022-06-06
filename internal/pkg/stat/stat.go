package stat

import (
	"context"
	"data_proxy/internal/cli"
	"data_proxy/internal/conf"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/peer"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var standardStat = NewStandardStat()

type Stat interface {
	Timing(name string, time int64, extra ...string)
	Incr(ctx context.Context, metric map[string]interface{}) // name,ext...,code
	Polymeric(log *log.Helper)
}

type StandardStat struct {
	count int64
	c     chan map[string]interface{}
	localIp string
}

func GetStandardStat() *StandardStat{
	return standardStat
}

func (s *StandardStat) Timing(name string, time int64, extra ...string) {
	panic("implement me")
}

func (s *StandardStat) Incr(ctx context.Context, metric map[string]interface{}) {
	metric["rule_id"] = 153248
	metric["timestamp"] = strconv.FormatInt(time.Now().UnixNano() / 1e6 , 10)
	clientIp, _ := GetClientIP(ctx)
	metric["cli_ip"] = clientIp
	metric["local_ip"] = s.localIp
	s.c <- metric
}

func NewStandardStat() *StandardStat {
	return &StandardStat{c: make(chan map[string]interface{}, 100000)}
}

func GetClientIP(ctx context.Context) (string, error) {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		return "", fmt.Errorf("[getClinetIP] invoke FromContext() failed")
	}
	if pr.Addr == net.Addr(nil) {
		return "", fmt.Errorf("[getClientIP] peer.Addr is nil")
	}

	addSlice := strings.Split(pr.Addr.String(), ":")
	if addSlice[0] == "[" {
		//本机地址
		return "localhost", nil
	}
	return addSlice[0], nil
}

func (s *StandardStat)LocalIP(log *log.Helper) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Error("get local ip err:", err)
		os.Exit(1)
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				s.localIp = ipNet.IP.String()
				break
			}
		}
	}
}
func SendData(stat []map[string]interface{}, log *log.Helper) int16{
	if len(stat) == 0  {
		return 1
	}
	body, err := json.Marshal(stat)
	if err != nil {
		log.Error(err)
		return -1
	}
	url := "http://st-moni-business-shou.vivo.lan:8080"
	// res, err:=cli.PostHttpData(url, body)
	_, err = cli.PostHttpData(url, 1, body, log)
	log.Info("post body:", string(body))
	stat = stat[:0]
	return 0
}

func DefaultHostname() string {
	hostname, _ := os.Hostname()
	return hostname
}

func (s *StandardStat) Polymeric(conf *conf.Server_STAT, log *log.Helper) {
	go func() {
		t1 := time.Tick(conf.TickTime.AsDuration())
		stat := make([]map[string]interface{}, 0)
		s.LocalIP(log)
		for {
			select {
			case <-t1:
				if len(stat) == 0  {
					continue
				}
				body, err := json.Marshal(stat)
				if err != nil {
					log.Error(err)
					continue
				}
				url := "http://st-moni-business-shou.vivo.lan:8080"
				_, err = cli.PostHttpData(url, conf.ClientTimeout.AsDuration(), body, log)
				log.Info("post body:", string(body))
				stat = stat[:0]
			case metric, ok := <-s.c:
				if ok {
					stat = append(stat, metric)
				}
			}
		}
	}()
}
