package logic

import (
	"2112a-6/month/wsyx/model"
	"2112a-6/month/wsyx/score_srv/internal/models"
	"context"
	"errors"
	"fmt"

	"2112a-6/month/wsyx/score_srv/internal/svc"
	"2112a-6/month/wsyx/score_srv/score_srv"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserScoreLogic {
	return &AddUserScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserScoreLogic) AddUserScore(in *score_srv.AddUserScoreRequest) (*score_srv.AddUserScoreResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	Findscore := model.Score{}
	tx := model.Db.Table("score").Where("user_id=?", in.Userid).First(&Findscore)

	if tx.RowsAffected == 0 {
		score := models.Score{
			UserId: in.Userid,
			Score:  uint64(in.Score),
		}
		err := model.Db.Table("score").Create(&score).Error
		if err != nil {
			return &score_srv.AddUserScoreResponse{Success: false}, errors.New("添加积分失败")
		}
	}
	err := model.Db.Exec("update `score` set score=score+? where user_id=?", in.Score, in.Userid).Error
	if err != nil {
		return &score_srv.AddUserScoreResponse{Success: false}, errors.New("添加积分失败")
	}

	return &score_srv.AddUserScoreResponse{Success: true}, nil
}
