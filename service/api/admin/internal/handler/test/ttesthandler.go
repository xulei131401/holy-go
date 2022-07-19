package test

import (
	"net/http"

	"github.com/xulei131401/holy-go/service/api/admin/internal/logic/test"
	"github.com/xulei131401/holy-go/service/api/admin/internal/svc"
	"github.com/xulei131401/holy-go/service/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/xulei131401/gox/validate"
)

func TTestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TestRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		if err := validate.ValidStruct(req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := test.NewTTestLogic(r.Context(), svcCtx)
		resp, err := l.TTest(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
