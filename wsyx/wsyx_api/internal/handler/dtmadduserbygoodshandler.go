package handler

import (
	"net/http"

	"2112a-6/month/wsyx/wsyx_api/internal/logic"
	"2112a-6/month/wsyx/wsyx_api/internal/svc"
	"2112a-6/month/wsyx/wsyx_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DtmAddUserbyGoodsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserByGoodsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDtmAddUserbyGoodsLogic(r.Context(), svcCtx)
		resp, err := l.DtmAddUserbyGoods(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
