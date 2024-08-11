package logic

import (
	"2112a-6/month/wsyx/model"
	"2112a-6/month/wsyx/order_srv/internal/models"
	"2112a-6/month/wsyx/pkg"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

	"2112a-6/month/wsyx/order_srv/internal/svc"
	"2112a-6/month/wsyx/order_srv/order_srv"

	"github.com/zeromicro/go-zero/core/logx"
)

type TranInsOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTranInsOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TranInsOrderLogic {
	return &TranInsOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TranInsOrderLogic) TranInsOrder(in *order_srv.AddOrderRequest) (*order_srv.AddOrderResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Printf("创建订单 in : %+v \n", in)
	barrier, _ := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err := barrier.CallWithDB(l.svcCtx.DB, func(tx *sql.Tx) error {

		goods := model.Goods{}
		err := model.Db.Table("goods").Where("id=?", in.GoodsId).First(&goods).Error
		if err != nil {
			return errors.New("商品查询失败")
		}
		if goods.GoodsNum < in.Count {
			return errors.New("商品库存不足")
		}
		addorder := models.Order{
			Userid: sql.NullInt64{
				Int64: in.Userid,
				Valid: true,
			},
			Goodsid:    in.GoodsId,
			GoodsCount: in.Count,
			GoodsPrice: sql.NullFloat64{
				Float64: goods.GoodsPrice * float64(in.Count),
				Valid:   true,
			},
			PayType: sql.NullInt64{
				Int64: in.Pay,
				Valid: true,
			},
			State: sql.NullInt64{
				Int64: 0,
				Valid: true,
			},
			CreatedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			UpdatedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			OrderNumbering: sql.NullString{
				String: pkg.OrderNumbering(),
				Valid:  true,
			},
		}
		_, err = l.svcCtx.DataSource.TranIncInsert(tx, &addorder)
		if err != nil {
			return errors.New("下单失败")
		}
		err = model.Db.Exec("update  `goods` set goods_num=goods_num-? where id=?", in.Count, in.GoodsId).Error
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &order_srv.AddOrderResponse{}, nil
}
