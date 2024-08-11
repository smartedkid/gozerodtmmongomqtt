package svc

import (
	"2112a-6/month/wsyx/score_srv/internal/config"
	"2112a-6/month/wsyx/score_srv/internal/models"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	DB         *sql.DB
	DataSource models.ScoreModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, _ := sqlx.NewMysql(c.DB.DataSource).RawDB()
	return &ServiceContext{
		Config:     c,
		DB:         db,
		DataSource: models.NewScoreModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
