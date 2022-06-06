package service

import (
	"context"
	pb "data_proxy/api/recommend"
	"data_proxy/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

type CommonService struct {
	pb.UnimplementedCommonServiceServer
	log *log.Helper
	cc *biz.CommonServicecase
	user *biz.UserUsecase
}

func NewCommonService(logger log.Logger, cc *biz.CommonServicecase, user *biz.UserUsecase) *CommonService {
	return &CommonService{log: log.NewHelper(logger), cc: cc, user: user}
}

func (s *CommonService) Process(ctx context.Context, req *pb.CommonRequest) (*pb.CommonResponse, error) {
	dpReq:=&pb.DPRequest{}
	err := proto.Unmarshal(req.Request.Value, dpReq)
	if err != nil {
		s.log.Error("Unmarshal err")
		return nil, err
	}
	dp:=biz.DpRequest{Cmd: dpReq.Cmd,
		TableName: dpReq.TableName}
	for _, v := range  dpReq.Param {
		p:=biz.Param{}
		p.Key = v.Key
		p.Expire = v.Expire
		dp.Params = append(dp.Params, &p)
	}
	dpRsp:=&pb.DPResponse{}
	g, err := s.cc.GetDpInfo(ctx, &dp)
	dpRsp.Code = g.Code
	rsp:=&pb.CommonResponse{}
	if err != nil {
		value, err := ptypes.MarshalAny(dpRsp)
		rsp.Response = value
		return rsp, err
	}
	for _, v := range g.Params {
		r:=&pb.Param{}
		r.Key = v.Key
		r.Value = []byte(v.Value)
		dpRsp.Param = append(dpRsp.Param, r)
	}
	value, err := ptypes.MarshalAny(dpRsp)
	// value, err := proto.Marshal(dpRsp)
	// rsp.Response.Value = value
	rsp.Response = value

	//id, err := s.user.CreateUser(ctx, &biz.User{
	//	Name:  "v.Key",
	//	Email: "v.value",
	//})
	//if id != 0 {
	//	log.Info("id:", id)
	//}
	return rsp, err
}

//func (s *CommonService) process(ctx context.Context, req *pb.CommonRequest) (rsp *pb.CommonResponse,errs error) {
//	dpReq:=&pb.DPRequest{}
//	err := proto.Unmarshal(req.Request.Value, dpReq)
//	if err != nil {
//		s.log.Error("Unmarshal err")
//		return nil, err
//	}
//	dp:=biz.DpRequest{Cmd: dpReq.Cmd,
//		TableName: dpReq.TableName}
//	for _, v := range  dpReq.Param {
//		p:=biz.Param{}
//		p.Key = v.Key
//		p.Expire = v.Expire
//		dp.Params = append(dp.Params, &p)
//	}
//	dpRsp:=&pb.DPResponse{}
//	g, err := s.cc.GetDpInfo(ctx, &dp)
//	dpRsp.Code = g.Code
//	if err != nil {
//		rsp.Response.Value, err = proto.Marshal(dpRsp)
//		return rsp, err
//	}
//	for _, v := range g.Params {
//		r:=&pb.Param{}
//		r.Key = v.Key
//		r.Value = []byte(v.Value)
//		dpRsp.Param = append(dpRsp.Param, r)
//	}
//
//	rsp.Response.Value, err = proto.Marshal(dpRsp)
//	return rsp, err
//}
