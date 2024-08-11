package logic

import (
	"context"

	"2112a-6/month/wsyx/score_srv/internal/svc"
	"2112a-6/month/wsyx/score_srv/score_srv"

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

func (l *PingLogic) Ping(in *score_srv.Request) (*score_srv.Response, error) {
	// todo: add your logic here and delete this line

	return &score_srv.Response{}, nil
}
