package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"timeout/internal/logic"
	"timeout/internal/svc"
	"timeout/internal/types"
)

func FastHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FastRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFastLogic(r.Context(), svcCtx)
		err := l.Fast(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
