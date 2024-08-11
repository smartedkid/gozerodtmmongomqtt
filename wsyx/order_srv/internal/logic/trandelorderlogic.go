package logic

import (
	"2112a-6/month/wsyx/model"
	"2112a-6/month/wsyx/order_srv/internal/svc"
	"2112a-6/month/wsyx/order_srv/order_srv"
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type TranDelOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTranDelOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TranDelOrderLogic {
	return &TranDelOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TranDelOrderLogic) TranDelOrder(in *order_srv.AddOrderRequest) (*order_srv.AddOrderResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println("删除订单")
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, err
	}
	if err := barrier.CallWithDB(l.svcCtx.DB, func(tx *sql.Tx) error {
		l.svcCtx.DataSource.TranDelUpdate(tx, int(in.Userid))
		err = model.Db.Exec("update  `goods` set goods_num=goods_num+? where id=?", in.Count, in.GoodsId).Error
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &order_srv.AddOrderResponse{}, nil
}
