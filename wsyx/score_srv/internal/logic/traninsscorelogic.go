package logic

import (
	"2112a-6/month/wsyx/model"
	"2112a-6/month/wsyx/score_srv/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"2112a-6/month/wsyx/score_srv/internal/svc"
	"2112a-6/month/wsyx/score_srv/score_srv"

	"github.com/zeromicro/go-zero/core/logx"
)

type TranInsScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTranInsScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TranInsScoreLogic {
	return &TranInsScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TranInsScoreLogic) TranInsScore(in *score_srv.AddUserScoreRequest) (*score_srv.AddUserScoreResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println("添加积分》》》》》》》》》》》》》》")
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, err
	}
	if err := barrier.CallWithDB(l.svcCtx.DB, func(tx *sql.Tx) error {
		Findscore := model.Score{}
		score := model.Db.Table("score").Where("user_id=?", in.Userid).First(&Findscore)

		if score.RowsAffected == 0 {
			score := models.Score{
				UserId: in.Userid,
				Score:  uint64(in.Score),
			}
			err := model.Db.Table("score").Create(&score).Error
			if err != nil {
				return errors.New("添加积分失败")
			}
		}
		fmt.Println(">>>>>>>>>>", in.Score, in.Score)
		l.svcCtx.DataSource.TranInsUpdate(tx, int(in.Score), int(in.Userid))

		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &score_srv.AddUserScoreResponse{}, nil
}
