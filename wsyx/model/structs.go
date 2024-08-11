package model

import (
	"database/sql"
	"time"
)

type Goods struct {
	Id           int64         `db:"id"`
	GoodsTitle   string        `db:"goods_title"`   // 商品标题
	GoodsImages  string        `db:"goods_images"`  // 商品封面
	GoodsAddress string        `db:"goods_address"` // 商品产地
	GoodsNum     int64         `db:"goods_num"`     // 商品库存
	CreatedAt    time.Time     `db:"created_at"`    // 添加时间
	UpdatedAt    time.Time     `db:"updated_at"`    // 修改时间
	DeletedAt    time.Time     `db:"deleted_at"`    // 删除时间
	IsHot        sql.NullInt64 `db:"is_hot"`        // 是否热门
	IsDel        sql.NullInt64 `db:"is_del"`        // 是否下架
	GoodsPrice   float64       `db:"goods_price"`   // 商品价格
}
type Score struct {
	Id     uint64 `db:"id"`
	UserId int64  `db:"user_id"`
	Score  uint64 `db:"score"`
}
type Order struct {
	Id             int64     `db:"id"`
	Userid         int64     `db:"userid"`          // 用户ID
	Goodsid        int64     `db:"goodsid"`         // 商品ID
	GoodsCount     int64     `db:"goods_count"`     // 下单商品数量
	GoodsPrice     float64   `db:"goods_price"`     // 总价格
	PayType        int64     `db:"pay_type"`        // 支付方式
	State          int64     `db:"state"`           // 支付状态
	CreatedAt      time.Time `db:"created_at"`      // 创建时间
	UpdatedAt      time.Time `db:"updated_at"`      // 修改时间
	DeletedAt      time.Time `db:"deleted_at"`      // 删除时间
	OrderNumbering string    `db:"order_numbering"` // 订单号

}
