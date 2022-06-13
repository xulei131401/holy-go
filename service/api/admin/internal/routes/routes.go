package routes

import (
	"holy-go/service/api/admin/internal/handler"
	"holy-go/service/api/admin/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: handler.AdminHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/token",
				Handler: handler.TokenHandler(serverCtx),
			},
		},
	)
}
