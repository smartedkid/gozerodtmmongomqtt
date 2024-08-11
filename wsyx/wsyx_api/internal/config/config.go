package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	GoodsSrv zrpc.RpcClientConf
	OrderSrv zrpc.RpcClientConf
	ScoreSrv zrpc.RpcClientConf
}
