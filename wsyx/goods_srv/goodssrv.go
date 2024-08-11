package main

import (
	"flag"
	"fmt"

	"2112a-6/month/wsyx/goods_srv/goods_srv"
	"2112a-6/month/wsyx/goods_srv/internal/config"
	"2112a-6/month/wsyx/goods_srv/internal/server"
	"2112a-6/month/wsyx/goods_srv/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/goodssrv.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		goods_srv.RegisterGoodsSrvServer(grpcServer, server.NewGoodsSrvServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
