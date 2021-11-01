package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/adhoc/shorturl/api/internal/logic"
	"github.com/tal-tech/go-zero/adhoc/shorturl/api/internal/svc"
	"github.com/tal-tech/go-zero/adhoc/shorturl/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func shortenHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShortenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewShortenLogic(r.Context(), ctx)
		resp, err := l.Shorten(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.WriteJson(w, http.StatusOK, resp)
		}
	}
}
