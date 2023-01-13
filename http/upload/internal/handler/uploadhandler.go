package handler

import (
	"net/http"

	"upload/internal/logic"
	"upload/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadLogic(r.Context(), ctx)
		resp, err := l.Upload(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
