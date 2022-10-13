package logic

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"
	"shorturl/rpc/transform/transformer"
)

func TestExpandLogic_Expand(t *testing.T) {
	ast := assert.New(t)

	// Build mock and svc context
	ctl := gomock.NewController(t)
	mockTransformer := transformer.NewMockTransformer(ctl)

	l := NewExpandLogic(context.Background(), &svc.ServiceContext{
		Transformer: mockTransformer,
	})

	// Failed to simulate transformer Expand
	mockTransformer.EXPECT().Expand(gomock.Any(), gomock.Any()).
		Return(nil, errors.New("call err")).
		Times(1)

	_, err := l.Expand(types.ExpandReq{})
	ast.NotNil(err)

	// Simulate transformer Expand success
	mockTransformer.EXPECT().Expand(gomock.Any(), gomock.Any()).
		Return(&transformer.ExpandResp{Url: "url"}, nil).
		Times(1)

	resp, err := l.Expand(types.ExpandReq{})
	ast.Nil(err)
	ast.Equal(resp, &types.ExpandResp{Url: "url"})
}
