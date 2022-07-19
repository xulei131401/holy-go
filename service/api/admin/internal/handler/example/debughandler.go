package example

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/xulei131401/gox/validate"
	"github.com/xulei131401/holy-go/service/api/admin/internal/logic/example"
	"github.com/xulei131401/holy-go/service/api/admin/internal/svc"
	"github.com/xulei131401/holy-go/service/api/admin/internal/types"
)

func DebugHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DebugRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		if err := validate.ValidStruct(req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := example.NewDebugLogic(r.Context(), svcCtx)
		resp, err := l.Debug(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
