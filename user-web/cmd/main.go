package main

import (
	"os"
	"time"

	"github.com/papandadj/nezha-chat-backend/user-web/handler"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/papandadj/nezha-chat-backend/pkg/log"
	"github.com/papandadj/nezha-chat-backend/user-web/conf"
)

var (
	logger log.Logger
	cfg    *conf.Config
)

func init() {
	confPath := os.Getenv("confPath")
	if confPath == "" {
		panic("配置文件没有找见")
	}

	err := conf.LoadGlobalConfig(confPath)
	if err != nil {
		panic(err)
	}

	cfg = conf.GetGlobalConfig()
	//初始化日志
	logger = log.New()
	logger.SetLevel(log.Level(cfg.LogLevel))
	logger.SetWorkspace(cfg.Workspace)
	logger.SetRootPackageSlash(cfg.RootPackageSlash)
}

func main() {
	addrEtcd := func(opts *registry.Options) {
		opts.Addrs = cfg.Etcd.Addrs
	}

	registry := etcdv3.NewRegistry(addrEtcd)

	srv := grpc.NewService(
		micro.Registry(registry),
		micro.Name(cfg.Micro.Name),
		micro.Version(cfg.Micro.Version),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
	)

	engin := handler.NewHTTPHandler(srv)

	if err := engin.Run(cfg.Web.Port); err != nil {
		log.Fatal(err)
	}

}
