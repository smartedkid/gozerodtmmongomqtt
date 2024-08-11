package models

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ScoreModel = (*customScoreModel)(nil)

type (
	// ScoreModel is an interface to be customized, add more methods here,
	// and implement the added methods in customScoreModel.
	ScoreModel interface {
		scoreModel
		withSession(session sqlx.Session) ScoreModel
	}

	customScoreModel struct {
		*defaultScoreModel
	}
)

// NewScoreModel returns a models for the database table.
func NewScoreModel(conn sqlx.SqlConn) ScoreModel {
	return &customScoreModel{
		defaultScoreModel: newScoreModel(conn),
	}
}

func (m *customScoreModel) withSession(session sqlx.Session) ScoreModel {
	return NewScoreModel(sqlx.NewSqlConnFromSession(session))
}
