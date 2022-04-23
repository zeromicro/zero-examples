package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"monolithic/internal/logic"
	"monolithic/internal/svc"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
