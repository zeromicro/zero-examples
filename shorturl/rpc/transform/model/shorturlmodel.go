package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ShorturlModel = (*customShorturlModel)(nil)

type (
	// ShorturlModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShorturlModel.
	ShorturlModel interface {
		shorturlModel
	}

	customShorturlModel struct {
		*defaultShorturlModel
	}
)

// NewShorturlModel returns a model for the database table.
func NewShorturlModel(conn sqlx.SqlConn, c cache.CacheConf) ShorturlModel {
	return &customShorturlModel{
		defaultShorturlModel: newShorturlModel(conn, c),
	}
}
