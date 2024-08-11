package svc

import (
	"2112a-6/month/wsyx/goods_srv/goodssrvclient"
	"2112a-6/month/wsyx/order_srv/ordersrvclient"
	"2112a-6/month/wsyx/score_srv/scoresrvclient"
	"2112a-6/month/wsyx/wsyx_api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	GoodsSrv goodssrvclient.GoodsSrv
	OrderSrv ordersrvclient.OrderSrv
	ScoreSrv scoresrvclient.ScoreSrv
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		GoodsSrv: goodssrvclient.NewGoodsSrv(zrpc.MustNewClient(c.GoodsSrv)),
		OrderSrv: ordersrvclient.NewOrderSrv(zrpc.MustNewClient(c.OrderSrv)),
		ScoreSrv: scoresrvclient.NewScoreSrv(zrpc.MustNewClient(c.ScoreSrv)),
	}
}
