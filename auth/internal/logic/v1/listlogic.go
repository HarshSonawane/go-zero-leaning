package v1

import (
	"context"

	"auth/internal/svc"
	"auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List() (resp []types.ListItem, err error) {
	// generate random data
	resp = []types.ListItem{
		{
			Id:   "1",
			Name: "name1",
		},
		{
			Id:   "2",
			Name: "name2",
		},
	}
	return resp, nil
}
