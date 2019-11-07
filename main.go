package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/papandadj/nezha-chat-backend/pkg/log"
	"github.com/papandadj/nezha-chat-backend/proto/user"
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
		micro.RegisterTTL(time.Minute),
		micro.RegisterInterval(time.Second*30),
	)

	rpcuser := user.NewUserService("nezha.chat.user.srv.service", srv.Client())
	resp, err := rpcuser.Post(context.Background(), &user.PostReq{
		Username: "nezha",
		Password: "888",
	})
	fmt.Println(resp, err)
}
