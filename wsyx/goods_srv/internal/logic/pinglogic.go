package logic

import (
	"context"

	"2112a-6/month/wsyx/goods_srv/goods_srv"
	"2112a-6/month/wsyx/goods_srv/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *goods_srv.Request) (*goods_srv.Response, error) {
	// todo: add your logic here and delete this line

	return &goods_srv.Response{}, nil
}
