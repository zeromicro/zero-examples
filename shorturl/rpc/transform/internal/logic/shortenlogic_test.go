package logic

import (
	"context"
	"errors"
	"testing"

	"shorturl/rpc/transform/internal/svc"
	"shorturl/rpc/transform/model"
	transform "shorturl/rpc/transform/pb"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestShortenLogic_Shorten(t *testing.T) {
	ast := assert.New(t)

	// Build mock models and svc context
	ctl := gomock.NewController(t)
	shortModel := model.NewMockshorturlModel(ctl)
	svcCtx := &svc.ServiceContext{
		Model: shortModel,
	}

	// build shorturl logic
	logic := NewShortenLogic(context.Background(), svcCtx)

	// Failed to simulate model insert
	shortModel.EXPECT().Insert(gomock.Any(), gomock.Any()).
		Return(nil, errors.New("insert error")).
		Times(1)
	_, err := logic.Shorten(&transform.ShortenReq{Url: "testUrl"})
	ast.NotNil(err)

	// Simulate model insert success
	shortModel.EXPECT().Insert(gomock.Any(), gomock.Any()).
		Return(nil, nil).
		Times(1)
	resp, err := logic.Shorten(&transform.ShortenReq{Url: "testUrl"})
	ast.Nil(err)
	ast.True(resp.Shorten != "")
}
