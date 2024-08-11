package logic

import (
	"2112a-6/month/wsyx/order_srv/order_srv"
	"2112a-6/month/wsyx/score_srv/score_srv"
	"context"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"

	"2112a-6/month/wsyx/wsyx_api/internal/svc"
	"2112a-6/month/wsyx/wsyx_api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type DtmAddUserbyGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDtmAddUserbyGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DtmAddUserbyGoodsLogic {
	return &DtmAddUserbyGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var DtmServer = "etcd://localhost:2379/dtmservice"

func (l *DtmAddUserbyGoodsLogic) DtmAddUserbyGoods(req *types.UserByGoodsRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	gid := dtmgrpc.MustGenGid(DtmServer)
	fmt.Println(">>>>>>>>>>>>>>", gid)
	sagaGrpc := dtmgrpc.NewSagaGrpc(DtmServer, gid)
	OrderSrv, _ := l.svcCtx.Config.OrderSrv.BuildTarget()
	//"/order_srv.Order_srv/TranInsOrder"
	//"/order_srv.Order_srv/TranDelOrder"
	sagaGrpc.Add(OrderSrv+"/order_srv.Order_srv/TranInsOrder", OrderSrv+"/order_srv.Order_srv/TranDelOrder", &order_srv.AddOrderRequest{
		Userid:  int64(req.Userid),
		GoodsId: int64(req.Goodsid),
		Pay:     int64(req.Pay),
		Count:   int64(req.Count),
	})
	ScoreSrv, _ := l.svcCtx.Config.ScoreSrv.BuildTarget()

	//"/score_srv.Score_srv/TranInsScore"
	//"/score_srv.Score_srv/TranDelScore"
	sagaGrpc.Add(ScoreSrv+"/score_srv.Score_srv/TranInsScore", ScoreSrv+"/score_srv.Score_srv/TranDelScore", &score_srv.AddUserScoreRequest{
		Userid:  int64(req.Userid),
		GoodsId: int64(req.Goodsid),
		Score:   1,
	})
	err = sagaGrpc.Submit()
	if err != nil {
		return nil, fmt.Errorf("submit data to  dtm-server err  : %+v \n", err)
	}
	return &types.Response{
		Code:    200,
		Message: "下单成功",
		Data:    nil,
	}, nil
}
