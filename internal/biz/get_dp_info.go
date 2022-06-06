package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Param struct {
	Key string
	Value string
	Expire int32
	Field string
	code int32
}

type DpRequest struct {
	TableName string
	Cmd string
	Key string
	Params []*Param
	ReqId string
	MonitorTag string
}

type DPResponse struct {
	Code uint32
	Params []*Param
	ReqId string
}


// CommonRepo is a Greater repo.
type CommonRepo interface {
	Process(context.Context, *DpRequest) (*DPResponse, error)
}

// CommonServicecase is a Common usecase.
type CommonServicecase struct {
	repo CommonRepo
	log  *log.Helper
}

// NewCommonServicecase new a Common usecase.
func NewCommonServicecase(repo CommonRepo, logger log.Logger) *CommonServicecase {
	return &CommonServicecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateCommon creates a Common, and returns the new Common.
func (uc *CommonServicecase) GetDpInfo(ctx context.Context, g *DpRequest) (*DPResponse, error) {
	uc.log.WithContext(ctx).Debugf("CreateCommon: %v", g.Cmd)
	return uc.repo.Process(ctx, g)
}
