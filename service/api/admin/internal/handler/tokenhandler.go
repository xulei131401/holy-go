package handler

import (
	"holy-go/service/api/admin/internal/logic"
	"holy-go/service/api/admin/internal/svc"
	"holy-go/service/api/admin/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func TokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.JwtTokenRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTokenLogic(r.Context(), svcCtx)
		resp, err := l.Token(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
