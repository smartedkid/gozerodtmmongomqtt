package logic

import (
	"2112a-6/month/wsyx/model"
	"2112a-6/month/wsyx/order_srv/internal/models"
	"2112a-6/month/wsyx/pkg"
	"2112a-6/month/wsyx/produce"
	"context"
	"database/sql"
	"errors"
	"time"

	"2112a-6/month/wsyx/order_srv/internal/svc"
	"2112a-6/month/wsyx/order_srv/order_srv"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderLogic {
	return &OrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderLogic) Order(in *order_srv.AddOrderRequest) (*order_srv.AddOrderResponse, error) {
	// todo: add your logic here and delete this line
	goods := model.Goods{}
	err := model.Db.Table("goods").Where("id=?", in.GoodsId).First(&goods).Error
	if err != nil {
		return &order_srv.AddOrderResponse{Success: false}, errors.New("商品查询失败")
	}
	if goods.GoodsNum < in.Count {
		return &order_srv.AddOrderResponse{Success: false}, errors.New("商品库存不足")
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

	//_, err = l.svcCtx.DataSource.Insert(l.ctx, &addorder)
	err = model.Db.Table("order").Create(&addorder).Error
	if err != nil {
		return &order_srv.AddOrderResponse{Success: false}, errors.New("下单添加失败")
	}

	produce.SendStock(addorder)

	//err = models.Db.Exec("update  `goods` set goods_num=goods_num-? where id=?", in.Count, in.GoodsId).Error
	//if err != nil {
	//	return nil, err
	//}

	return &order_srv.AddOrderResponse{Success: true}, nil
}
