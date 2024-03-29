package v1

import (
	"context"

	"users/internal/svc"
	"users/internal/types"

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

func (l *ListLogic) List() (resp []types.UserListItem, err error) {
	resp = []types.UserListItem{
		{
			Id:   "1",
			FirstName: "John",
			LastName: "Doe",
			Email: "john@ddsio.com",
		},
		{
			Id:   "2",
			FirstName: "Jane",
			LastName: "Doe",
			Email: "jane@ddsio.com",
		},
	}

	return resp, nil
}
