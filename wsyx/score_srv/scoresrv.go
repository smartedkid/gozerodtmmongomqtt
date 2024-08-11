package main

import (
	"2112a-6/month/wsyx/model"
	"flag"
	"fmt"

	"2112a-6/month/wsyx/score_srv/internal/config"
	"2112a-6/month/wsyx/score_srv/internal/server"
	"2112a-6/month/wsyx/score_srv/internal/svc"
	"2112a-6/month/wsyx/score_srv/score_srv"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/scoresrv.yaml", "the config file")

func main() {
	flag.Parse()
	model.MysqlClient()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		score_srv.RegisterScoreSrvServer(grpcServer, server.NewScoreSrvServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
