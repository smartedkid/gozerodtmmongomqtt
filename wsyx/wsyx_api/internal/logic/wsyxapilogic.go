package logic

import (
	"context"

	"2112a-6/month/wsyx/wsyx_api/internal/svc"
	"2112a-6/month/wsyx/wsyx_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Wsyx_apiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWsyx_apiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Wsyx_apiLogic {
	return &Wsyx_apiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Wsyx_apiLogic) Wsyx_api(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
