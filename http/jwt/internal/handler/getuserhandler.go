package handler

import (
	"net/http"

	"github.com/zeromicro/zero-examples/jwt/internal/logic"
	"github.com/zeromicro/zero-examples/jwt/internal/svc"
	"github.com/zeromicro/zero-examples/jwt/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserLogic(r.Context(), ctx)
		resp, err := l.GetUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
