package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"2112a-6/month/wsyx/score_srv/internal/svc"
	"2112a-6/month/wsyx/score_srv/score_srv"

	"github.com/zeromicro/go-zero/core/logx"
)

type TranDelScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTranDelScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TranDelScoreLogic {
	return &TranDelScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TranDelScoreLogic) TranDelScore(in *score_srv.AddUserScoreRequest) (*score_srv.AddUserScoreResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println("删除积分》》》》》》》》》》》》》》")
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, err
	}
	if err := barrier.CallWithDB(l.svcCtx.DB, func(tx *sql.Tx) error {

		fmt.Println(">>>>>>>>>>", in.Score, in.Score)
		l.svcCtx.DataSource.TranDelUpdate(tx, int(in.Score), int(in.Userid))
		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &score_srv.AddUserScoreResponse{}, nil
}
