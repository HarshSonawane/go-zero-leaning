package v1

import (
	"context"

	"auth/internal/svc"
	"auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PathExampleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPathExampleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PathExampleLogic {
	return &PathExampleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PathExampleLogic) PathExample(req *types.PathExampleReq) (resp *types.PathExampleResp, err error) {
	// todo: add your logic here and delete this line

	return
}
