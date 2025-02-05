// Code generated by goctl. DO NOT EDIT.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	goodsFieldNames          = builder.RawFieldNames(&Goods{})
	goodsRows                = strings.Join(goodsFieldNames, ",")
	goodsRowsExpectAutoSet   = strings.Join(stringx.Remove(goodsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	goodsRowsWithPlaceHolder = strings.Join(stringx.Remove(goodsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	goodsModel interface {
		Insert(ctx context.Context, data *Goods) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Goods, error)
		Update(ctx context.Context, data *Goods) error
		Delete(ctx context.Context, id int64) error
	}

	defaultGoodsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Goods struct {
		Id           int64          `db:"id"`
		GoodsTitle   sql.NullString `db:"goods_title"`   // 商品标题
		GoodsImages  sql.NullString `db:"goods_images"`  // 商品封面
		GoodsAddress sql.NullString `db:"goods_address"` // 商品产地
		GoodsNum     sql.NullInt64  `db:"goods_num"`     // 商品库存
		CreatedAt    sql.NullTime   `db:"created_at"`    // 添加时间
		UpdatedAt    sql.NullTime   `db:"updated_at"`    // 修改时间
		DeletedAt    sql.NullTime   `db:"deleted_at"`    // 删除时间
		IsHot        sql.NullInt64  `db:"is_hot"`        // 是否热门
		IsDel        sql.NullInt64  `db:"is_del"`        // 是否下架
		GoodsPrice   float64        `db:"goods_price"`   // 商品价格
	}
)

func newGoodsModel(conn sqlx.SqlConn) *defaultGoodsModel {
	return &defaultGoodsModel{
		conn:  conn,
		table: "`goods`",
	}
}

func (m *defaultGoodsModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultGoodsModel) FindOne(ctx context.Context, id int64) (*Goods, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", goodsRows, m.table)
	var resp Goods
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultGoodsModel) Insert(ctx context.Context, data *Goods) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, goodsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.GoodsTitle, data.GoodsImages, data.GoodsAddress, data.GoodsNum, data.DeletedAt, data.IsHot, data.IsDel, data.GoodsPrice)
	return ret, err
}

func (m *defaultGoodsModel) Update(ctx context.Context, data *Goods) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, goodsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.GoodsTitle, data.GoodsImages, data.GoodsAddress, data.GoodsNum, data.DeletedAt, data.IsHot, data.IsDel, data.GoodsPrice, data.Id)
	return err
}

func (m *defaultGoodsModel) tableName() string {
	return m.table
}
