package main

import (
	"2112a-6/month/wsyx/model"
	"flag"
	"fmt"

	"2112a-6/month/wsyx/wsyx_api/internal/config"
	"2112a-6/month/wsyx/wsyx_api/internal/handler"
	"2112a-6/month/wsyx/wsyx_api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/wsyxapi-api.yaml", "the config file")

func main() {
	flag.Parse()
	model.MysqlClient()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
