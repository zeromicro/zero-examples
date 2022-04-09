package logic

import (
	"context"
	"time"

	"timeout/internal/svc"
	"timeout/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SlowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSlowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SlowLogic {
	return &SlowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SlowLogic) Slow(req *types.SlowRequest) error {
	time.Sleep(time.Second * 4)

	return nil
}
