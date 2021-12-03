package handler

import (
	"net/http"

	"upload/internal/logic"
	"upload/internal/svc"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func UploadHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadLogic(r, ctx)
		resp, err := l.Upload()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
