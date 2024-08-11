package svc

import (
	"2112a-6/month/wsyx/goods_srv/internal/config"
	"2112a-6/month/wsyx/goods_srv/internal/models"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	DB         *sql.DB
	DataSource models.GoodsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, _ := sqlx.NewMysql(c.DB.DataSource).RawDB()

	return &ServiceContext{
		Config:     c,
		DB:         db,
		DataSource: models.NewGoodsModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
