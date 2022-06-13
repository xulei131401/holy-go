package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"holy-go/service/api/admin/internal/logic/user"
	"holy-go/service/api/admin/internal/svc"
	"holy-go/service/api/admin/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}