// Code generated by goctl. DO NOT EDIT.
// Source: order_srv.proto

package server

import (
	"context"

	"2112a-6/month/wsyx/order_srv/internal/logic"
	"2112a-6/month/wsyx/order_srv/internal/svc"
	"2112a-6/month/wsyx/order_srv/order_srv"
)

type OrderSrvServer struct {
	svcCtx *svc.ServiceContext
	order_srv.UnimplementedOrderSrvServer
}

func NewOrderSrvServer(svcCtx *svc.ServiceContext) *OrderSrvServer {
	return &OrderSrvServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderSrvServer) Ping(ctx context.Context, in *order_srv.Request) (*order_srv.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *OrderSrvServer) Order(ctx context.Context, in *order_srv.AddOrderRequest) (*order_srv.AddOrderResponse, error) {
	l := logic.NewOrderLogic(ctx, s.svcCtx)
	return l.Order(in)
}

func (s *OrderSrvServer) TranInsOrder(ctx context.Context, in *order_srv.AddOrderRequest) (*order_srv.AddOrderResponse, error) {
	l := logic.NewTranInsOrderLogic(ctx, s.svcCtx)
	return l.TranInsOrder(in)
}

func (s *OrderSrvServer) TranDelOrder(ctx context.Context, in *order_srv.AddOrderRequest) (*order_srv.AddOrderResponse, error) {
	l := logic.NewTranDelOrderLogic(ctx, s.svcCtx)
	return l.TranDelOrder(in)
}
