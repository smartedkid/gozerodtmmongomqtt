package logic

import (
	"context"

	"2112a-6/month/wsyx/order_srv/internal/svc"
	"2112a-6/month/wsyx/order_srv/order_srv"

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

func (l *PingLogic) Ping(in *order_srv.Request) (*order_srv.Response, error) {
	// todo: add your logic here and delete this line

	return &order_srv.Response{}, nil
}
