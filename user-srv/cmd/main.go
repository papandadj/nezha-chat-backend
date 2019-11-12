package main

import (
	"os"
	"time"

	"github.com/micro/go-plugins/wrapper/trace/opentracing"

	openTrace "github.com/opentracing/opentracing-go"

	"github.com/papandadj/nezha-chat-backend/pkg/tracer"

	"github.com/papandadj/nezha-chat-backend/user-srv/dao"

	"github.com/papandadj/nezha-chat-backend/proto/user"

	"github.com/papandadj/nezha-chat-backend/user-srv/service"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/papandadj/nezha-chat-backend/pkg/log"
	"github.com/papandadj/nezha-chat-backend/user-srv/conf"
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

	//修改Etcd地址函数
	addrEtcd := func(opts *registry.Options) {
		opts.Addrs = cfg.Etcd.Addrs
	}

	registry := etcdv3.NewRegistry(addrEtcd)

	//设置opentrace
	t, io, err := tracer.NewTracer("user-srv", "127.0.0.1:6831")
	if err != nil {
		logger.Fatalln(err)
	}
	defer io.Close()
	openTrace.SetGlobalTracer(t)

	srv := grpc.NewService(
		micro.Registry(registry),
		micro.Name(cfg.Micro.Name),
		micro.Version(cfg.Micro.Version),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		micro.WrapHandler(opentracing.NewHandlerWrapper(openTrace.GlobalTracer())),
	)

	dao.Init()

	user.RegisterUserHandler(srv.Server(), service.New(dao.GetDao()))

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
