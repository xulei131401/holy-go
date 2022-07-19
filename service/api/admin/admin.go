package main

import (
	"flag"
	"fmt"
	"github.com/xulei131401/gox/server/option"
	"github.com/xulei131401/holy-go/service/api/admin/internal/handler"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"

	"github.com/xulei131401/holy-go/service/api/admin/internal/config"
	"github.com/xulei131401/holy-go/service/api/admin/internal/svc"
)

var configFile = flag.String("f", "etc/admin-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	optionFunc()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func optionFunc() {
	// 设置可选项
	rest.WithNotFoundHandler(option.NotFoundHandler())
	rest.WithNotAllowedHandler(option.NotAllowedHandler())
	rest.WithUnauthorizedCallback(option.UnauthorizedCallback())

	// 设置全局error处理
	option.ErrorHandler()

	// 取消log统计打印
	logx.DisableStat()
}
