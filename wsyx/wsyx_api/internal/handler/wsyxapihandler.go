package handler

import (
	"net/http"

	"2112a-6/month/wsyx/wsyx_api/internal/logic"
	"2112a-6/month/wsyx/wsyx_api/internal/svc"
	"2112a-6/month/wsyx/wsyx_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Wsyx_apiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewWsyx_apiLogic(r.Context(), svcCtx)
		resp, err := l.Wsyx_api(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
