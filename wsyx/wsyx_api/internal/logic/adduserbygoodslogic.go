package logic

import (
	"2112a-6/month/wsyx/order_srv/order_srv"
	"2112a-6/month/wsyx/score_srv/score_srv"
	"context"

	"2112a-6/month/wsyx/wsyx_api/internal/svc"
	"2112a-6/month/wsyx/wsyx_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserbyGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserbyGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserbyGoodsLogic {
	return &AddUserbyGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserbyGoodsLogic) AddUserbyGoods(req *types.UserByGoodsRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	order, err := l.svcCtx.OrderSrv.Order(l.ctx, &order_srv.AddOrderRequest{
		Userid:  int64(req.Userid),
		GoodsId: int64(req.Goodsid),
		Pay:     int64(req.Pay),
		Count:   int64(req.Count),
	})
	if err != nil {
		return &types.Response{
			Code:    400,
			Message: "下单失败",
			Data:    nil,
		}, err
	}
	_, err = l.svcCtx.ScoreSrv.AddUserScore(l.ctx, &score_srv.AddUserScoreRequest{
		Userid:  int64(req.Userid),
		GoodsId: int64(req.Goodsid),
		Score:   1,
	})
	if err != nil {
		return &types.Response{
			Code:    400,
			Message: "添加积分失败",
			Data:    nil,
		}, err
	}

	return &types.Response{
		Code:    200,
		Message: "下单成功",
		Data:    order.Success,
	}, err
}
