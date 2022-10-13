package model

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func TestDefaultShorturlModel_Insert(t *testing.T) {
	ast := assert.New(t)

	// build model, mock db and mock redis
	model, mock, _ := newMockShorturlModel(t)

	// mock test case
	mock.ExpectExec(fmt.Sprintf("insert into %s", model.table)).
		WithArgs("123", "123").
		WillReturnError(errors.New("exec error"))
	mock.ExpectExec(fmt.Sprintf("insert into %s", model.table)).
		WithArgs("345", "345").
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()

	_, err := model.Insert(ctx, &Shorturl{
		Shorten: "123",
		Url:     "123",
	})
	ast.NotNil(err)

	_, err = model.Insert(ctx, &Shorturl{
		Shorten: "345",
		Url:     "345",
	})
	ast.Nil(err)
}

func TestDefaultShorturlModel_Update(t *testing.T) {
	ast := assert.New(t)

	// build model, mock db and mock redis
	model, mock, _ := newMockShorturlModel(t)

	ctx := context.Background()

	// mock test fail case
	mock.ExpectExec(fmt.Sprintf("update %s", model.table)).
		WillReturnError(errors.New("exec error"))
	err := model.Update(ctx, &Shorturl{
		Shorten: "123",
		Url:     "123",
	})
	ast.NotNil(err)

	// mock test success case
	mock.ExpectExec(fmt.Sprintf("update %s", model.table)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err = model.Update(ctx, &Shorturl{
		Shorten: "123",
		Url:     "123",
	})
	ast.Nil(err)
}

func TestDefaultShorturlModel_Delete(t *testing.T) {
	ast := assert.New(t)

	// build model, mock db and mock redis
	model, mock, _ := newMockShorturlModel(t)

	ctx := context.Background()

	// mock test fail case
	mock.ExpectExec(fmt.Sprintf("delete from  %s", model.table)).
		WillReturnError(errors.New("exec error"))
	err := model.Delete(ctx, "123")
	ast.NotNil(err)

	// mock test success case
	mock.ExpectExec(fmt.Sprintf("delete from  %s", model.table)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err = model.Delete(ctx, "123")
	ast.Nil(err)
}

func TestDefaultShorturlModel_FindOne(t *testing.T) {
	ast := assert.New(t)

	// build model, mock db and mock redis
	model, mock, rds := newMockShorturlModel(t)

	ctx := context.Background()

	// mock db query error
	mock.ExpectQuery(fmt.Sprintf("select (.+) from %s", model.table)).
		WillReturnError(errors.New("query error"))

	_, err := model.FindOne(ctx, "123")
	ast.NotNil(err)

	// mock db query success
	rows := sqlmock.NewRows(
		[]string{"shorten", "url"},
	).AddRow([]driver.Value{"111", "222"}...)

	mock.ExpectQuery(fmt.Sprintf("select (.+) from %s", model.table)).
		WillReturnRows(rows)

	ret, err := model.FindOne(ctx, "111")
	ast.Nil(err)
	ast.Equal(ret, &Shorturl{
		Shorten: "111",
		Url:     "222",
	})

	// mock cache data
	su := &Shorturl{
		Shorten: "123",
		Url:     "234",
	}
	data, _ := jsonx.Marshal(su)
	rds.Set(fmt.Sprintf("%s%v", cacheShorturlShortenPrefix, su.Shorten), string(data))

	ret, err = model.FindOne(ctx, su.Shorten)
	ast.Nil(err)
	ast.Equal(ret, su)

}

func newMockShorturlModel(t *testing.T) (*defaultShorturlModel, sqlmock.Sqlmock, *miniredis.Miniredis) {
	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	rds := miniredis.RunT(t)
	return &defaultShorturlModel{
		CachedConn: sqlc.NewNodeConn(sqlx.NewSqlConnFromDB(db), redis.New(rds.Addr())),
		table:      "`shorturl`",
	}, mockDb, rds
}
