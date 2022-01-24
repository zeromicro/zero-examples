package handler

import (
	"io/ioutil"
	"net/http"

	"download/internal/svc"
	"download/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func DownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		body, err := ioutil.ReadFile(req.File)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		w.Write(body)
	}
}
