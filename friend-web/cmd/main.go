package main

import (
	"os"

	"github.com/papandadj/nezha-chat-backend/friend-web/handler"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-plugins/registry/etcdv3"

	openTrace "github.com/opentracing/opentracing-go"

	"github.com/papandadj/nezha-chat-backend/friend-web/conf"
	"github.com/papandadj/nezha-chat-backend/pkg/log"
	"github.com/papandadj/nezha-chat-backend/pkg/tracer"
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

	//设置opentrace
	t, io, err := tracer.NewTracer(cfg.Jaeger.ServiceName, cfg.Jaeger.URL)
	if err != nil {
		logger.Fatalln(err)
	}
	defer io.Close()
	openTrace.SetGlobalTracer(t)

	srv := grpc.NewService(
		micro.Registry(registry),
	)
	handler.Init()
	engin := handler.NewHTTPHandler(srv)

	if err := engin.Run(cfg.Web.Port); err != nil {
		logger.Fatal(err)
	}

}
