package logic

import (
	"context"
	"time"

	"timeout/internal/svc"
	"timeout/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FastLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFastLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FastLogic {
	return &FastLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FastLogic) Fast(req *types.FastRequest) error {
	// should timeout
	time.Sleep(time.Second * 4)

	return nil
}
