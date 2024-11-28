package handler

import (
	"net/http"

	"fisrt/internal/logic"
	"fisrt/internal/svc"
	"fisrt/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FisrtHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFisrtLogic(r.Context(), svcCtx)
		resp, err := l.Fisrt(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
