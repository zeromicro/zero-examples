package logic

import (
	"context"

	"github.com/zeromicro/zero-examples/jwt/internal/svc"
	"github.com/zeromicro/zero-examples/jwt/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserLogic {
	return GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req types.GetUserRequest) (*types.GetUserResponse, error) {
	// todo: add your logic here and delete this line

	return &types.GetUserResponse{Name: "frank_test"}, nil
}
