package data

import (
	"context"
	"data_proxy/internal/biz"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

type processRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.CommonRepo {
	return &processRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// process(context.Context, *Common) (*Common, error)
func (r *processRepo) Process(ctx context.Context, g *biz.DpRequest) (*biz.DPResponse, error) {
	rsp:=&biz.DPResponse{}
	rdb, ok := r.data.rdb[g.TableName]
	if !ok {
		rsp.Code = 9005
		return rsp, errors.New(-1,"table name not find: ",g.TableName)
	}
	r.log.Debug("table:", g.TableName,len(r.data.rdb))
	pipeline := rdb.Pipeline()
	// result := make([]*redis.StringCmd, 0)
	for _, p := range  g.Params {
		pipeline.Get(ctx, p.Key)
	}
	cmds, err := pipeline.Exec(ctx)
	if err != nil {
		r.log.Error("pipe error ", err)
		rsp.Code = 9006
		return rsp, err
	}
	for _, cmd := range  cmds {
		biz:=&biz.Param{}
		biz.Value = cmd.(*redis.StringCmd).Val()
		cmd.(*redis.StringCmd).Args()
		if len(cmd.(*redis.StringCmd).Args()) == 2 {
			biz.Key = cmd.(*redis.StringCmd).Args()[1].(string)
		}
		rsp.Params = append(rsp.Params, biz)
	}
	return rsp, nil
}

